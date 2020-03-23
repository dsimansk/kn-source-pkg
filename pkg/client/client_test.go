// Copyright © 2019 The Knative Authors
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

package client

import (
	"testing"

	"gotest.tools/assert"

	"github.com/maximilien/kn-source-pkg/pkg/types"
)

func TestNewKnSourceClient(t *testing.T) {
	knSourceClient := NewKnSourceClient(&types.KnSourceParams{}, "fake-namespace")
	assert.Assert(t, knSourceClient != nil)
}

func TestKnSourceParams(t *testing.T) {
	fakeKnSourceParams := &types.KnSourceParams{}
	knSourceClient := NewKnSourceClient(fakeKnSourceParams, "fake-namespace")
	assert.Equal(t, knSourceClient.KnSourceParams(), fakeKnSourceParams)
}

func TestNamespace(t *testing.T) {
	knSourceClient := NewKnSourceClient(&types.KnSourceParams{}, "fake-namespace")
	assert.Equal(t, knSourceClient.Namespace(), "fake-namespace")
}