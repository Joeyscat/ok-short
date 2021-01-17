/*
 Copyright 2020 The Qmgo Authors.
 Licensed under the Apache License, Version 2.0 (the "License");
 you may not use this file except in compliance with the License.
 You may obtain a copy of the License at
     http://www.apache.org/licenses/LICENSE-2.0
 Unless required by applicable law or agreed to in writing, software
 distributed under the License is distributed on an "AS IS" BASIS,
 WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 See the License for the specific language governing permissions and
 limitations under the License.
*/

package qmgo

import (
	"context"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

// Pipeline define the pipeline for aggregate
type Pipeline []bson.D

// Aggregate is a handle to a aggregate
type Aggregate struct {
	ctx        context.Context
	pipeline   interface{}
	collection *mongo.Collection
}

// All iterates the cursor from aggregate and decodes each document into results.
func (a *Aggregate) All(results interface{}) error {
	c, err := a.collection.Aggregate(a.ctx, a.pipeline)
	if err != nil {
		return err
	}
	return c.All(a.ctx, results)
}

// One iterates the cursor from aggregate and decodes current document into result.
func (a *Aggregate) One(result interface{}) error {
	c, err := a.collection.Aggregate(a.ctx, a.pipeline)
	if err != nil {
		return err
	}
	cr := Cursor{
		ctx:    a.ctx,
		cursor: c,
		err:    err,
	}
	defer cr.Close()
	if !cr.Next(result) {
		return ErrNoSuchDocuments
	}
	return err
}

// Iter return the cursor after aggregate
func (a *Aggregate) Iter() CursorI {
	c, err := a.collection.Aggregate(a.ctx, a.pipeline)
	return &Cursor{
		ctx:    a.ctx,
		cursor: c,
		err:    err,
	}
}
