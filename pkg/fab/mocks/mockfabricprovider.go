/*
Copyright SecureKey Technologies Inc. All Rights Reserved.

SPDX-License-Identifier: Apache-2.0
*/

package mocks

import (
	"fmt"

	"github.com/hyperledger/fabric-sdk-go/pkg/common/context"
	"github.com/hyperledger/fabric-sdk-go/pkg/context/api/core"
	"github.com/hyperledger/fabric-sdk-go/pkg/context/api/fab"
	"github.com/hyperledger/fabric-sdk-go/pkg/context/api/msp"
)

// MockInfraProvider represents the default implementation of Fabric objects.
type MockInfraProvider struct {
	providerContext context.Providers
	customOrderer   fab.Orderer
}

// CreateChannelLedger returns a new client initialized for the current instance of the SDK.
func (f *MockInfraProvider) CreateChannelLedger(ic msp.Identity, channelName string) (fab.ChannelLedger, error) {
	return nil, nil
}

// CreateEventService creates the event service.
func (f *MockInfraProvider) CreateEventService(ic fab.ClientContext, chConfig fab.ChannelCfg) (fab.EventService, error) {
	panic("not implemented")
}

// CreateChannelConfig initializes the channel config
func (f *MockInfraProvider) CreateChannelConfig(ic msp.Identity, channelID string) (fab.ChannelConfig, error) {
	if ic == nil {
		return &MockChannelConfig{channelID: channelID}, nil
	}
	return &MockChannelConfig{channelID: channelID, ctx: NewMockContext(ic)}, nil
}

// CreateChannelMembership returns a channel member identifier
func (f *MockInfraProvider) CreateChannelMembership(cfg fab.ChannelCfg) (fab.ChannelMembership, error) {
	return nil, fmt.Errorf("Not implemented")
}

// CreateChannelTransactor initializes the transactor
func (f *MockInfraProvider) CreateChannelTransactor(ic msp.Identity, cfg fab.ChannelCfg) (fab.Transactor, error) {
	if cfg == nil {
		return &MockTransactor{}, nil
	} else if ic == nil {
		return &MockTransactor{ChannelID: cfg.ID()}, nil
	}
	return &MockTransactor{ChannelID: cfg.ID(), Ctx: NewMockContext(ic)}, nil
}

// CreatePeerFromConfig returns a new default implementation of Peer based configuration
func (f *MockInfraProvider) CreatePeerFromConfig(peerCfg *core.NetworkPeer) (fab.Peer, error) {
	if peerCfg != nil {
		return NewMockPeer(peerCfg.MspID, peerCfg.URL), nil
	}
	return &MockPeer{}, nil
}

// CreateOrdererFromConfig creates a default implementation of Orderer based on configuration.
func (f *MockInfraProvider) CreateOrdererFromConfig(cfg *core.OrdererConfig) (fab.Orderer, error) {
	if f.customOrderer != nil {
		return f.customOrderer, nil
	}

	return &MockOrderer{}, nil
}

//CommManager returns comm provider
func (f *MockInfraProvider) CommManager() fab.CommManager {
	return nil
}

// SetCustomOrderer creates a default implementation of Orderer based on configuration.
func (f *MockInfraProvider) SetCustomOrderer(customOrderer fab.Orderer) {
	f.customOrderer = customOrderer
}

//Close mock close function
func (f *MockInfraProvider) Close() {
}
