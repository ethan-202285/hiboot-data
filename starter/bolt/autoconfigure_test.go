// Copyright 2018 John Deng (hi.devops.io@gmail.com).
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

package bolt

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestNewBolt(t *testing.T) {
	configuration := new(boltConfiguration)
	configuration.Properties = &Properties{
		Database: "test.db",
		Mode:     0600,
		Timeout:  1,
	}

	repository := configuration.Repository()
	assert.NotEqual(t, nil, repository)
	repository.DataSource().(DataSource).Close()
}

func TestNewBoltWithError(t *testing.T) {
	configuration := new(boltConfiguration)

	repository := configuration.Repository()
	assert.NotEqual(t, nil, repository)

	configuration.Properties = Properties{
		Timeout: 1,
	}

	repository = configuration.Repository()
	assert.NotEqual(t, nil, repository)
}
