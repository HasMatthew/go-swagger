// Copyright 2015 go-swagger maintainers
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//    http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package httpkit

import (
	"bytes"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
)

// Gob byte slice representation of {Name: "Somebody", ID: 1}
var consProdGob = []byte{28, 255, 129, 3, 1, 2, 255, 130, 0, 1, 2, 1, 4, 78, 97, 109, 101, 1, 12, 0, 1, 2, 73, 68, 1, 4, 0, 0, 0, 15, 255, 130, 1, 8, 83, 111, 109, 101, 98, 111, 100, 121, 1, 2, 0}

func TestGobConsumer(t *testing.T) {
	cons := GobConsumer()
	var data struct {
		Name string
		ID   int
	}
	err := cons.Consume(bytes.NewBuffer(consProdGob), &data)
	assert.NoError(t, err)
	assert.Equal(t, "Somebody", data.Name)
	assert.Equal(t, 1, data.ID)
}

func TestGobProducer(t *testing.T) {
	prod := GobProducer()
	data := struct {
		Name string `json:"name"`
		ID   int    `json:"id"`
	}{Name: "Somebody", ID: 1}

	rw := httptest.NewRecorder()
	err := prod.Produce(rw, data)
	assert.NoError(t, err)
	assert.Equal(t, consProdGob, rw.Body.Bytes())
}
