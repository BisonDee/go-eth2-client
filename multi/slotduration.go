// Copyright © 2021 Attestant Limited.
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

package multi

import (
	"context"
	"time"

	eth2client "github.com/attestantio/go-eth2-client"
	"github.com/pkg/errors"
)

// SlotDuration provides the duration of a slot of the chain.
func (s *Service) SlotDuration(ctx context.Context) (time.Duration, error) {
	s.clientMu.RLock()
	defer s.clientMu.RUnlock()
	if len(s.activeClients) == 0 {
		return 0, errors.New("no active Ethereum 2 clients")
	}

	// Slot duration is static so no need to worry about failover.
	return s.activeClients[0].(eth2client.SlotDurationProvider).SlotDuration(ctx)
}
