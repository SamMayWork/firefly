// Copyright © 2021 Kaleido, Inc.
//
// SPDX-License-Identifier: Apache-2.0
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

package apiserver

import (
	"bytes"
	"encoding/json"
	"net/http/httptest"
	"testing"

	"github.com/hyperledger-labs/firefly/mocks/orchestratormocks"
	"github.com/hyperledger-labs/firefly/pkg/fftypes"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func TestPutConfigRecord(t *testing.T) {
	o := &orchestratormocks.Orchestrator{}
	r := createAdminMuxRouter(o)
	input := &fftypes.ConfigRecord{
		Key:   "foo",
		Value: fftypes.Byteable(`{"foo": "bar"}`),
	}
	var buf bytes.Buffer
	json.NewEncoder(&buf).Encode(&input)
	req := httptest.NewRequest("PUT", "/admin/api/v1/config/records/foo", &buf)
	req.Header.Set("Content-Type", "application/json; charset=utf-8")
	res := httptest.NewRecorder()

	o.On("PutConfigRecord", mock.Anything, mock.Anything, mock.Anything).
		Return(input.Value, nil)
	r.ServeHTTP(res, req)

	assert.Equal(t, 200, res.Result().StatusCode)
	outputBuf := new(bytes.Buffer)
	outputBuf.ReadFrom(res.Body)
	assert.Equal(t, "{\"foo\":\"bar\"}\n", outputBuf.String())
}