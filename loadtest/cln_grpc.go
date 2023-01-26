package main

import (
	"context"
	"encoding/hex"
	"fmt"
	"strconv"
	"strings"

	clightning "github.com/flitz-be/cln-grpc-go/cln"
)

type ClnGrpcClient struct {
	client clightning.NodeClient
}

func (cln *ClnGrpcClient) GetInfo() (*info, error) {
	resp, err := cln.client.Getinfo(context.Background(), &clightning.GetinfoRequest{})
	if err != nil {
		return nil, err
	}
	return &info{
		key:    hex.EncodeToString(resp.Id),
		synced: resp.WarningLightningdSync == nil,
	}, nil
}

func (cln *ClnGrpcClient) Connect(key string, host string) error {
	parts := strings.Split(host, ":")
	if len(parts) != 2 {
		return fmt.Errorf("Host is malformed: %s. should be host:port format", host)
	}
	hostPart := parts[0]
	port, err := strconv.Atoi(parts[1])
	if err != nil {
		return fmt.Errorf("Host is malformed: %s. should be host:port format", host)
	}
	portPart := uint32(port)
	_, err = cln.client.ConnectPeer(context.Background(), &clightning.ConnectRequest{
		Id:   key,
		Host: &hostPart,
		Port: &portPart,
	})
	return err
}

func (cln *ClnGrpcClient) NewAddress() (string, error) {
	panic("not implemented") // TODO: Implement
}

func (cln *ClnGrpcClient) OpenChannel(peerKey string, amtSat int64) error {
	panic("not implemented") // TODO: Implement
}

func (cln *ClnGrpcClient) ActiveChannels() (int, error) {
	panic("not implemented") // TODO: Implement
}

func (cln *ClnGrpcClient) AddInvoice(amtMsat int64) (string, error) {
	panic("not implemented") // TODO: Implement
}

func (cln *ClnGrpcClient) SendPayment(invoice string) error {
	panic("not implemented") // TODO: Implement
}

func (cln *ClnGrpcClient) SendKeysend(destination string, amtMsat int64) error {
	panic("not implemented") // TODO: Implement
}

func (cln *ClnGrpcClient) Close() {
	panic("not implemented") // TODO: Implement
}

func (cln *ClnGrpcClient) HasFunds() error {
	panic("not implemented") // TODO: Implement
}
