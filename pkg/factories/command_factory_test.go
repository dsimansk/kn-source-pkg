// Copyright © 2018 The Knative Authors
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

	"github.com/maximilien/kn-source-pkg/pkg/types"
)

func TestNewDefaultCommandFactory(t *testing.T) {
	commandFactory := NewDefaultCommandFactory(&types.KnSourceParams{})

	assert.Assert(t, commandFactory != nil)
}

func TestKnSourceParams(t *testing.T) {
	knSourceParams := &types.KnSourceParams{}
	commandFactory := NewDefaultCommandFactory(knSourceParams)

	assert.Equal(t, commandFactory.KnSourceParams(), knSourceParams)
}

func TestSourceCommand(t *testing.T) {
	commandFactory := NewDefaultCommandFactory(&types.KnSourceParams{})

	assert.Assert(t, commandFactory.SourceCommand() != nil)
}

func TestCreateCommand(t *testing.T) {
	commandFactory := NewDefaultCommandFactory(&types.KnSourceParams{})

	assert.Assert(t, commandFactory.CreateCommand() != nil)
}

func TestDeleteCommand(t *testing.T) {
	commandFactory := NewDefaultCommandFactory(&types.KnSourceParams{})

	assert.Assert(t, commandFactory.DeleteCommand() != nil)
}

func TestUpdateCommand(t *testing.T) {
	commandFactory := NewDefaultCommandFactory(&types.KnSourceParams{})

	assert.Assert(t, commandFactory.UpdateCommand() != nil)
}

func TestDescribeCommand(t *testing.T) {
	commandFactory := NewDefaultCommandFactory(&types.KnSourceParams{})

	assert.Assert(t, commandFactory.DescribeCommand() != nil)
}
