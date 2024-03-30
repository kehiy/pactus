// Code generated by protoc-gen-jrpc-gateway. DO NOT EDIT.
// source: blockchain.proto

/*
Package pactus is a reverse proxy.

It translates gRPC into JSON-RPC 2.0
*/
package pactus

import (
	"context"
	"encoding/json"

	"google.golang.org/protobuf/encoding/protojson"
)

type BlockchainJsonRpcService struct {
	client BlockchainClient
}

func NewBlockchainJsonRpcService(client BlockchainClient) BlockchainJsonRpcService {
	return BlockchainJsonRpcService{
		client: client,
	}
}

func (s *BlockchainJsonRpcService) Methods() map[string]func(ctx context.Context, message json.RawMessage) (any, error) {
	return map[string]func(ctx context.Context, params json.RawMessage) (any, error){

		"pactus.blockchain.get_block": func(ctx context.Context, data json.RawMessage) (any, error) {
			req := new(GetBlockRequest)
			err := protojson.Unmarshal(data, req)
			if err != nil {
				return nil, err
			}
			return s.client.GetBlock(ctx, req)
		},

		"pactus.blockchain.get_block_hash": func(ctx context.Context, data json.RawMessage) (any, error) {
			req := new(GetBlockHashRequest)
			err := protojson.Unmarshal(data, req)
			if err != nil {
				return nil, err
			}
			return s.client.GetBlockHash(ctx, req)
		},

		"pactus.blockchain.get_block_height": func(ctx context.Context, data json.RawMessage) (any, error) {
			req := new(GetBlockHeightRequest)
			err := protojson.Unmarshal(data, req)
			if err != nil {
				return nil, err
			}
			return s.client.GetBlockHeight(ctx, req)
		},

		"pactus.blockchain.get_blockchain_info": func(ctx context.Context, data json.RawMessage) (any, error) {
			req := new(GetBlockchainInfoRequest)
			err := protojson.Unmarshal(data, req)
			if err != nil {
				return nil, err
			}
			return s.client.GetBlockchainInfo(ctx, req)
		},

		"pactus.blockchain.get_consensus_info": func(ctx context.Context, data json.RawMessage) (any, error) {
			req := new(GetConsensusInfoRequest)
			err := protojson.Unmarshal(data, req)
			if err != nil {
				return nil, err
			}
			return s.client.GetConsensusInfo(ctx, req)
		},

		"pactus.blockchain.get_account": func(ctx context.Context, data json.RawMessage) (any, error) {
			req := new(GetAccountRequest)
			err := protojson.Unmarshal(data, req)
			if err != nil {
				return nil, err
			}
			return s.client.GetAccount(ctx, req)
		},

		"pactus.blockchain.get_validator": func(ctx context.Context, data json.RawMessage) (any, error) {
			req := new(GetValidatorRequest)
			err := protojson.Unmarshal(data, req)
			if err != nil {
				return nil, err
			}
			return s.client.GetValidator(ctx, req)
		},

		"pactus.blockchain.get_validator_by_number": func(ctx context.Context, data json.RawMessage) (any, error) {
			req := new(GetValidatorByNumberRequest)
			err := protojson.Unmarshal(data, req)
			if err != nil {
				return nil, err
			}
			return s.client.GetValidatorByNumber(ctx, req)
		},

		"pactus.blockchain.get_validator_addresses": func(ctx context.Context, data json.RawMessage) (any, error) {
			req := new(GetValidatorAddressesRequest)
			err := protojson.Unmarshal(data, req)
			if err != nil {
				return nil, err
			}
			return s.client.GetValidatorAddresses(ctx, req)
		},

		"pactus.blockchain.get_public_key": func(ctx context.Context, data json.RawMessage) (any, error) {
			req := new(GetPublicKeyRequest)
			err := protojson.Unmarshal(data, req)
			if err != nil {
				return nil, err
			}
			return s.client.GetPublicKey(ctx, req)
		},
	}
}
