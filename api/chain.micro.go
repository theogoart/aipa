// Code generated by protoc-gen-micro. DO NOT EDIT.
// source: github.com/aipadad/aipa/api/chain.proto

package api

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

import (
	client "github.com/micro/go-micro/client"
	server "github.com/micro/go-micro/server"
	context "context"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ client.Option
var _ server.Option

// Client API for Chain service

type ChainClient interface {
	SendTransaction(ctx context.Context, in *Transaction, opts ...client.CallOption) (*SendTransactionResponse, error)
	GetTransaction(ctx context.Context, in *GetTransactionRequest, opts ...client.CallOption) (*GetTransactionResponse, error)
	GetBlock(ctx context.Context, in *GetBlockRequest, opts ...client.CallOption) (*GetBlockResponse, error)
	GetInfo(ctx context.Context, in *GetInfoRequest, opts ...client.CallOption) (*GetInfoResponse, error)
	GetAccount(ctx context.Context, in *GetAccountRequest, opts ...client.CallOption) (*GetAccountResponse, error)
	GetKeyValue(ctx context.Context, in *GetKeyValueRequest, opts ...client.CallOption) (*GetKeyValueResponse, error)
	GetAbi(ctx context.Context, in *GetAbiRequest, opts ...client.CallOption) (*GetAbiResponse, error)
	GetTransferCredit(ctx context.Context, in *GetTransferCreditRequest, opts ...client.CallOption) (*GetTransferCreditResponse, error)
}

type chainClient struct {
	c           client.Client
	serviceName string
}

func NewChainClient(serviceName string, c client.Client) ChainClient {
	if c == nil {
		c = client.NewClient()
	}
	if len(serviceName) == 0 {
		serviceName = "api"
	}
	return &chainClient{
		c:           c,
		serviceName: serviceName,
	}
}

func (c *chainClient) SendTransaction(ctx context.Context, in *Transaction, opts ...client.CallOption) (*SendTransactionResponse, error) {
	req := c.c.NewRequest(c.serviceName, "Chain.SendTransaction", in)
	out := new(SendTransactionResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *chainClient) GetTransaction(ctx context.Context, in *GetTransactionRequest, opts ...client.CallOption) (*GetTransactionResponse, error) {
	req := c.c.NewRequest(c.serviceName, "Chain.GetTransaction", in)
	out := new(GetTransactionResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *chainClient) GetBlock(ctx context.Context, in *GetBlockRequest, opts ...client.CallOption) (*GetBlockResponse, error) {
	req := c.c.NewRequest(c.serviceName, "Chain.GetBlock", in)
	out := new(GetBlockResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *chainClient) GetInfo(ctx context.Context, in *GetInfoRequest, opts ...client.CallOption) (*GetInfoResponse, error) {
	req := c.c.NewRequest(c.serviceName, "Chain.GetInfo", in)
	out := new(GetInfoResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *chainClient) GetAccount(ctx context.Context, in *GetAccountRequest, opts ...client.CallOption) (*GetAccountResponse, error) {
	req := c.c.NewRequest(c.serviceName, "Chain.GetAccount", in)
	out := new(GetAccountResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *chainClient) GetKeyValue(ctx context.Context, in *GetKeyValueRequest, opts ...client.CallOption) (*GetKeyValueResponse, error) {
	req := c.c.NewRequest(c.serviceName, "Chain.GetKeyValue", in)
	out := new(GetKeyValueResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *chainClient) GetAbi(ctx context.Context, in *GetAbiRequest, opts ...client.CallOption) (*GetAbiResponse, error) {
	req := c.c.NewRequest(c.serviceName, "Chain.GetAbi", in)
	out := new(GetAbiResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *chainClient) GetTransferCredit(ctx context.Context, in *GetTransferCreditRequest, opts ...client.CallOption) (*GetTransferCreditResponse, error) {
	req := c.c.NewRequest(c.serviceName, "Chain.GetTransferCredit", in)
	out := new(GetTransferCreditResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for Chain service

type ChainHandler interface {
	SendTransaction(context.Context, *Transaction, *SendTransactionResponse) error
	GetTransaction(context.Context, *GetTransactionRequest, *GetTransactionResponse) error
	GetBlock(context.Context, *GetBlockRequest, *GetBlockResponse) error
	GetInfo(context.Context, *GetInfoRequest, *GetInfoResponse) error
	GetAccount(context.Context, *GetAccountRequest, *GetAccountResponse) error
	GetKeyValue(context.Context, *GetKeyValueRequest, *GetKeyValueResponse) error
	GetAbi(context.Context, *GetAbiRequest, *GetAbiResponse) error
	GetTransferCredit(context.Context, *GetTransferCreditRequest, *GetTransferCreditResponse) error
}

func RegisterChainHandler(s server.Server, hdlr ChainHandler, opts ...server.HandlerOption) {
	s.Handle(s.NewHandler(&Chain{hdlr}, opts...))
}

type Chain struct {
	ChainHandler
}

func (h *Chain) SendTransaction(ctx context.Context, in *Transaction, out *SendTransactionResponse) error {
	return h.ChainHandler.SendTransaction(ctx, in, out)
}

func (h *Chain) GetTransaction(ctx context.Context, in *GetTransactionRequest, out *GetTransactionResponse) error {
	return h.ChainHandler.GetTransaction(ctx, in, out)
}

func (h *Chain) GetBlock(ctx context.Context, in *GetBlockRequest, out *GetBlockResponse) error {
	return h.ChainHandler.GetBlock(ctx, in, out)
}

func (h *Chain) GetInfo(ctx context.Context, in *GetInfoRequest, out *GetInfoResponse) error {
	return h.ChainHandler.GetInfo(ctx, in, out)
}

func (h *Chain) GetAccount(ctx context.Context, in *GetAccountRequest, out *GetAccountResponse) error {
	return h.ChainHandler.GetAccount(ctx, in, out)
}

func (h *Chain) GetKeyValue(ctx context.Context, in *GetKeyValueRequest, out *GetKeyValueResponse) error {
	return h.ChainHandler.GetKeyValue(ctx, in, out)
}

func (h *Chain) GetAbi(ctx context.Context, in *GetAbiRequest, out *GetAbiResponse) error {
	return h.ChainHandler.GetAbi(ctx, in, out)
}

func (h *Chain) GetTransferCredit(ctx context.Context, in *GetTransferCreditRequest, out *GetTransferCreditResponse) error {
	return h.ChainHandler.GetTransferCredit(ctx, in, out)
}
