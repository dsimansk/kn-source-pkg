// Copyright © 2020 The Knative Authors
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package factories

import (
	"testing"

	"gotest.tools/assert"
)

func TestNewDefaultKnSourceFactory(t *testing.T) {
	knSourceFactory := NewDefaultKnSourceFactory()

	assert.Assert(t, knSourceFactory != nil)
}

func TestKnSourceFactory_KnSourceParams(t *testing.T) {
	knSourceFactory := NewDefaultKnSourceFactory()

	knSourceParams := knSourceFactory.KnSourceParams()
	assert.Assert(t, knSourceParams != nil)

	knSourceParams = knSourceFactory.CreateKnSourceParams()
	assert.Equal(t, knSourceFactory.KnSourceParams(), knSourceParams)
}

func TestCreateKnSourceParams(t *testing.T) {
	knSourceFactory := NewDefaultKnSourceFactory()

	knSourceParams := knSourceFactory.CreateKnSourceParams()
	assert.Assert(t, knSourceParams != nil)
}

func TestCreateKnSourceClient(t *testing.T) {
	knSourceFactory := NewDefaultKnSourceFactory()
	client := knSourceFactory.CreateKnSourceClient("fake-namespace")

	assert.Assert(t, client != nil)
	assert.Equal(t, client.Namespace(), "fake-namespace")
}
