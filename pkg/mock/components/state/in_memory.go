/*
 * Copyright 2021 Layotto Authors
 *
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 *
 *     http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 */

package mock_state

import (
	"context"
	"fmt"
	"sync"

	"github.com/dapr/components-contrib/state"
	"github.com/dapr/kit/logger"
	jsoniter "github.com/json-iterator/go"
)

type inMemStateStoreItem struct {
	data []byte
	etag *string
}

type inMemoryStore struct {
	items map[string]*inMemStateStoreItem
	lock  *sync.RWMutex
	log   logger.Logger
}

func New(logger logger.Logger) state.Store {
	return &inMemoryStore{
		items: map[string]*inMemStateStoreItem{},
		lock:  &sync.RWMutex{},
		log:   logger,
	}
}

func (store *inMemoryStore) newItem(data []byte, etagString *string) *inMemStateStoreItem {
	return &inMemStateStoreItem{
		data: data,
		etag: etagString,
	}
}

func (store *inMemoryStore) Init(ctx context.Context, metadata state.Metadata) error {
	return nil
}

func (store *inMemoryStore) Ping() error {
	return nil
}

func (store *inMemoryStore) Features() []state.Feature {
	return []state.Feature{state.FeatureETag, state.FeatureTransactional}
}

func (store *inMemoryStore) Delete(ctx context.Context, req *state.DeleteRequest) error {
	store.lock.Lock()
	defer store.lock.Unlock()
	delete(store.items, req.Key)

	return nil
}

func (store *inMemoryStore) BulkDelete(ctx context.Context, req []state.DeleteRequest, opts state.BulkStoreOpts) error {
	if len(req) == 0 {
		return nil
	}
	for _, dr := range req {
		err := store.Delete(ctx, &dr)
		if err != nil {
			store.log.Error(err)
			return err
		}
	}
	return nil
}

func (store *inMemoryStore) Get(ctx context.Context, req *state.GetRequest) (*state.GetResponse, error) {
	store.lock.RLock()
	defer store.lock.RUnlock()
	item := store.items[req.Key]

	if item == nil {
		return &state.GetResponse{Data: nil, ETag: nil}, nil
	}
	if data, err := unmarshal(item.data); err != nil {
		return nil, err
	} else {
		return &state.GetResponse{Data: data, ETag: item.etag}, nil
	}
}

func (store *inMemoryStore) BulkGet(ctx context.Context, req []state.GetRequest, opts state.BulkGetOpts) ([]state.BulkGetResponse, error) {
	res := []state.BulkGetResponse{}
	for _, oneRequest := range req {
		oneResponse, err := store.Get(ctx, &state.GetRequest{
			Key:      oneRequest.Key,
			Metadata: oneRequest.Metadata,
			Options:  oneRequest.Options,
		})
		if err != nil {
			store.log.Error(err)
			return nil, err
		}

		res = append(res, state.BulkGetResponse{
			Key:  oneRequest.Key,
			Data: oneResponse.Data,
			ETag: oneResponse.ETag,
		})
	}

	return res, nil
}

func (store *inMemoryStore) Set(ctx context.Context, req *state.SetRequest) error {
	b, _ := marshal(req.Value)
	store.lock.Lock()
	defer store.lock.Unlock()
	store.items[req.Key] = store.newItem(b, req.ETag)

	return nil
}

func (store *inMemoryStore) BulkSet(ctx context.Context, req []state.SetRequest, opts state.BulkStoreOpts) error {
	for _, r := range req {
		err := store.Set(ctx, &r)
		if err != nil {
			store.log.Error(err)
			return err
		}
	}
	return nil
}

func (store *inMemoryStore) Multi(ctx context.Context, request *state.TransactionalStateRequest) error {
	store.lock.Lock()
	defer store.lock.Unlock()
	// First we check all eTags
	for _, o := range request.Operations {
		var eTag *string
		key := ""
		if o.Operation() == state.OperationUpsert {
			key = o.(state.SetRequest).Key
			eTag = o.(state.SetRequest).ETag
		} else if o.Operation() == state.OperationDelete {
			key = o.(state.SetRequest).Key
			eTag = o.(state.SetRequest).ETag
		}
		item := store.items[key]
		if eTag != nil && item != nil {
			if *eTag != *item.etag {
				return fmt.Errorf("etag does not match for key %v", key)
			}
		}
		if eTag != nil && item == nil {
			return fmt.Errorf("etag does not match for key not found %v", key)
		}
	}

	// Now we can perform the operation.
	for _, o := range request.Operations {
		if o.Operation() == state.OperationUpsert {
			req := o.(state.SetRequest)
			b, _ := marshal(req.Value)
			store.items[req.Key] = store.newItem(b, req.ETag)
		} else if o.Operation() == state.OperationDelete {
			req := o.(state.DeleteRequest)
			delete(store.items, req.Key)
		}
	}

	return nil
}

func marshal(value interface{}) ([]byte, error) {
	if v, err := jsoniter.MarshalToString(value); err != nil {
		return nil, err
	} else {
		return []byte(v), nil
	}
}

func unmarshal(val interface{}) ([]byte, error) {
	var output string

	if err := jsoniter.UnmarshalFromString(string(val.([]byte)), &output); err != nil {
		return nil, err
	}
	return []byte(output), nil
}
