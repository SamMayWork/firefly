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

package orchestrator

import (
	"context"
	"fmt"
	"testing"

	"github.com/kaleido-io/firefly/internal/config"
	"github.com/kaleido-io/firefly/mocks/batchmocks"
	"github.com/kaleido-io/firefly/mocks/blockchainmocks"
	"github.com/kaleido-io/firefly/mocks/broadcastmocks"
	"github.com/kaleido-io/firefly/mocks/databasemocks"
	"github.com/kaleido-io/firefly/mocks/eventmocks"
	"github.com/kaleido-io/firefly/mocks/publicstoragemocks"
	"github.com/kaleido-io/firefly/pkg/fftypes"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

type testOrchestrator struct {
	orchestrator

	mdi *databasemocks.Plugin
	mbm *broadcastmocks.Manager
	mba *batchmocks.Manager
	mei *eventmocks.EventManager
	mps *publicstoragemocks.Plugin
	mbi *blockchainmocks.Plugin
}

func newTestOrchestrator() *testOrchestrator {
	tor := &testOrchestrator{
		orchestrator: orchestrator{
			ctx: context.Background(),
		},
		mdi: &databasemocks.Plugin{},
		mbm: &broadcastmocks.Manager{},
		mba: &batchmocks.Manager{},
		mei: &eventmocks.EventManager{},
		mps: &publicstoragemocks.Plugin{},
		mbi: &blockchainmocks.Plugin{},
	}
	tor.orchestrator.database = tor.mdi
	tor.orchestrator.batch = tor.mba
	tor.orchestrator.broadcast = tor.mbm
	tor.orchestrator.events = tor.mei
	tor.orchestrator.publicstorage = tor.mps
	tor.orchestrator.blockchain = tor.mbi
	tor.mdi.On("Name").Return("mock-di").Maybe()
	tor.mbm.On("Name").Return("mock-bi").Maybe()
	tor.mba.On("Name").Return("mock-ba").Maybe()
	tor.mei.On("Name").Return("mock-ei").Maybe()
	tor.mps.On("Name").Return("mock-ps").Maybe()
	tor.mbi.On("Name").Return("mock-bi").Maybe()
	return tor
}

func TestNewOrchestrator(t *testing.T) {
	or := NewOrchestrator()
	assert.NotNil(t, or)
}

func TestBadDatabasePlugin(t *testing.T) {
	or := newTestOrchestrator()
	config.Set(config.DatabaseType, "wrong")
	or.database = nil
	err := or.Init(context.Background())
	assert.Regexp(t, "FF10122.*wrong", err.Error())
}

func TestBadDatabaseInitFail(t *testing.T) {
	or := newTestOrchestrator()
	config.Set(config.DatabaseType, "wrong")
	or.mdi.On("Init", mock.Anything, mock.Anything, mock.Anything).Return(fmt.Errorf("pop"))
	err := or.Init(context.Background())
	assert.EqualError(t, err, "pop")
}

func TestBadBlockchainPlugin(t *testing.T) {
	or := newTestOrchestrator()
	config.Set(config.BlockchainType, "wrong")
	or.blockchain = nil
	or.mdi.On("Init", mock.Anything, mock.Anything, mock.Anything).Return(nil)
	err := or.Init(context.Background())
	assert.Regexp(t, "FF10110.*wrong", err.Error())
}

func TestBlockchaiInitFail(t *testing.T) {
	or := newTestOrchestrator()
	or.mdi.On("Init", mock.Anything, mock.Anything, mock.Anything).Return(nil)
	or.mbi.On("Init", mock.Anything, mock.Anything, mock.Anything).Return(fmt.Errorf("pop"))
	err := or.Init(context.Background())
	assert.EqualError(t, err, "pop")
}

func TestBlockchainVerifyIDFail(t *testing.T) {
	or := newTestOrchestrator()
	or.mdi.On("Init", mock.Anything, mock.Anything, mock.Anything).Return(nil)
	or.mbi.On("Init", mock.Anything, mock.Anything, mock.Anything).Return(nil)
	or.mbi.On("VerifyIdentitySyntax", mock.Anything, mock.Anything, mock.Anything).Return("", fmt.Errorf("pop"))
	err := or.Init(context.Background())
	assert.EqualError(t, err, "pop")
}

func TestBadPublicStoragePlugin(t *testing.T) {
	or := newTestOrchestrator()
	config.Set(config.PublicStorageType, "wrong")
	or.publicstorage = nil
	or.mdi.On("Init", mock.Anything, mock.Anything, mock.Anything).Return(nil)
	or.mbi.On("Init", mock.Anything, mock.Anything, mock.Anything).Return(nil)
	or.mbi.On("VerifyIdentitySyntax", mock.Anything, mock.Anything, mock.Anything).Return("", nil)
	err := or.Init(context.Background())
	assert.Regexp(t, "FF10134.*wrong", err.Error())
}

func TestBadPublicStorageInitFail(t *testing.T) {
	or := newTestOrchestrator()
	or.mdi.On("Init", mock.Anything, mock.Anything, mock.Anything).Return(nil)
	or.mbi.On("Init", mock.Anything, mock.Anything, mock.Anything).Return(nil)
	or.mbi.On("VerifyIdentitySyntax", mock.Anything, mock.Anything, mock.Anything).Return("", nil)
	or.mps.On("Init", mock.Anything, mock.Anything, mock.Anything).Return(fmt.Errorf("pop"))
	err := or.Init(context.Background())
	assert.EqualError(t, err, "pop")
}

func TestInitEventsComponentFail(t *testing.T) {
	or := &orchestrator{}
	err := or.initComponents(context.Background())
	assert.Regexp(t, "FF10128", err.Error())
}

func TestInitBatchComponentFail(t *testing.T) {
	or := newTestOrchestrator()
	or.database = nil
	or.batch = nil
	err := or.initComponents(context.Background())
	assert.Regexp(t, "FF10128", err.Error())
}

func TestInitBroadcastComponentFail(t *testing.T) {
	or := newTestOrchestrator()
	or.database = nil
	or.broadcast = nil
	err := or.initComponents(context.Background())
	assert.Regexp(t, "FF10128", err.Error())
}

func TestStartBatchFail(t *testing.T) {
	config.Reset()
	or := newTestOrchestrator()
	or.mba.On("Start").Return(fmt.Errorf("pop"))
	or.mbi.On("Start").Return(nil)
	err := or.Start()
	assert.Regexp(t, "pop", err.Error())
}

func TestStartStopOk(t *testing.T) {
	config.Reset()
	or := newTestOrchestrator()
	or.mbi.On("Start").Return(nil)
	or.mba.On("Start").Return(nil)
	or.mei.On("Start").Return(nil)
	or.mbm.On("Start").Return(nil)
	or.mbi.On("WaitStop").Return(nil)
	or.mba.On("WaitStop").Return(nil)
	or.mei.On("WaitStop").Return(nil)
	or.mbm.On("WaitStop").Return(nil)
	err := or.Start()
	assert.NoError(t, err)
	or.WaitStop()
	or.WaitStop() // swallows dups
}

func TestInitNamespacesBadName(t *testing.T) {
	config.Reset()
	config.Set(config.NamespacesPredefined, fftypes.JSONObjectArray{
		{"name": "!Badness"},
	})
	or := newTestOrchestrator()
	err := or.initNamespaces(context.Background())
	assert.Regexp(t, "FF10131", err.Error())
}

func TestInitNamespacesGetFail(t *testing.T) {
	config.Reset()
	or := newTestOrchestrator()
	or.mdi.On("GetNamespace", mock.Anything, mock.Anything).Return(nil, fmt.Errorf("pop"))
	err := or.initNamespaces(context.Background())
	assert.Regexp(t, "pop", err.Error())
}

func TestInitNamespacesUpsertFail(t *testing.T) {
	config.Reset()
	or := newTestOrchestrator()
	or.mdi.On("GetNamespace", mock.Anything, mock.Anything).Return(nil, nil)
	or.mdi.On("UpsertNamespace", mock.Anything, mock.Anything, true).Return(fmt.Errorf("pop"))
	err := or.initNamespaces(context.Background())
	assert.Regexp(t, "pop", err.Error())
}

func TestInitNamespacesUpsertNotNeeded(t *testing.T) {
	config.Reset()
	or := newTestOrchestrator()
	or.mdi.On("GetNamespace", mock.Anything, mock.Anything).Return(&fftypes.Namespace{
		Type: fftypes.NamespaceTypeBroadcast, // any broadcasted NS will not be updated
	}, nil)
	err := or.initNamespaces(context.Background())
	assert.NoError(t, err)
}

func TestInitNamespacesDefaultMissing(t *testing.T) {
	config.Reset()
	or := newTestOrchestrator()
	config.Set(config.NamespacesPredefined, fftypes.JSONObjectArray{})
	err := or.initNamespaces(context.Background())
	assert.Regexp(t, "FF10166", err.Error())
}

func TestInitOK(t *testing.T) {
	or := newTestOrchestrator()
	or.mdi.On("Init", mock.Anything, mock.Anything, mock.Anything).Return(nil)
	or.mbi.On("Init", mock.Anything, mock.Anything, mock.Anything).Return(nil)
	or.mbi.On("VerifyIdentitySyntax", mock.Anything, mock.Anything, mock.Anything).Return("", nil)
	or.mps.On("Init", mock.Anything, mock.Anything, mock.Anything).Return(nil)
	or.mdi.On("GetNamespace", mock.Anything, mock.Anything).Return(nil, nil)
	or.mdi.On("UpsertNamespace", mock.Anything, mock.Anything, true).Return(nil)
	err := config.ReadConfig("../../test/config/firefly.core.yaml")
	assert.NoError(t, err)
	err = or.Init(context.Background())
	assert.NoError(t, err)
}