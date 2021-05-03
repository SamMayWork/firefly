// Copyright © 2021 Kaleido, Inc.
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

package broadcast

import (
	"context"

	"github.com/kaleido-io/firefly/internal/batching"
	"github.com/kaleido-io/firefly/internal/blockchain"
	"github.com/kaleido-io/firefly/internal/fftypes"
	"github.com/kaleido-io/firefly/internal/i18n"
	"github.com/kaleido-io/firefly/internal/persistence"
)

type Broadcast interface {
	BroadcastMessage(ctx context.Context, identity string, msg *fftypes.MessageRefsOnly) error
	Close()
}

type broadcast struct {
	ctx         context.Context
	persistence persistence.Plugin
	blockchain  blockchain.Plugin
	batch       batching.BatchManager
}

func NewBroadcast(ctx context.Context, persistence persistence.Plugin, blockchain blockchain.Plugin, batch batching.BatchManager) (Broadcast, error) {
	if persistence == nil || batch == nil {
		return nil, i18n.NewError(ctx, i18n.MsgInitializationNilDepError)
	}
	b := &broadcast{
		ctx:         ctx,
		persistence: persistence,
		blockchain:  blockchain,
		batch:       batch,
	}
	return b, nil
}

func (b *broadcast) BroadcastMessage(ctx context.Context, identity string, msg *fftypes.MessageRefsOnly) error {

	return nil
}

func (b *broadcast) Close() {}
