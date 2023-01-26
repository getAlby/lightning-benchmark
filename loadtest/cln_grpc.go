package main

import (
	"context"
	"crypto/tls"
	"crypto/x509"
	"encoding/hex"
	"fmt"
	"io/ioutil"
	"strconv"
	"strings"

	clightning "github.com/flitz-be/cln-grpc-go/cln"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
)

func loadTLSCredentials(caCertPath, clientCertPath, clientKeyPath string) (credentials.TransportCredentials, error) {
	// Load certificate of the CA who signed server's certificate
	pemServerCA, err := ioutil.ReadFile(caCertPath)
	if err != nil {
		return nil, err
	}

	certPool := x509.NewCertPool()
	if !certPool.AppendCertsFromPEM(pemServerCA) {
		return nil, fmt.Errorf("failed to add server CA's certificate")
	}

	// Load client's certificate and private key
	clientCert, err := tls.LoadX509KeyPair(clientCertPath, clientKeyPath)
	if err != nil {
		return nil, err
	}

	// Create the credentials and return it
	config := &tls.Config{
		Certificates: []tls.Certificate{clientCert},
		RootCAs:      certPool,
	}

	return credentials.NewTLS(config), nil
}

type ClnGrpcClient struct {
	client clightning.NodeClient
}

func getClightningGrpcConnection(cfg *clightningGrpcConfig) (*ClnGrpcClient, error) {
	credentials, err := loadTLSCredentials(
		cfg.CaPath,
		cfg.ClientPemPath,
		cfg.ClientKeyPath)
	if err != nil {
		return nil, err
	}

	opts := []grpc.DialOption{
		grpc.WithTransportCredentials(credentials),
	}

	conn, err := grpc.Dial(cfg.RpcHost, opts...)
	if err != nil {
		return nil, err
	}
	return &ClnGrpcClient{
		client: clightning.NewNodeClient(conn),
	}, nil
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
	addrType := clightning.NewaddrRequest_BECH32
	resp, err := cln.client.NewAddr(context.Background(), &clightning.NewaddrRequest{
		Addresstype: &addrType,
	})
	if err != nil {
		return "", err
	}
	return *resp.Bech32, nil
}

func (cln *ClnGrpcClient) OpenChannel(peerKey string, amtSat int64) error {
	_, err := cln.client.FundChannel(context.Background(), &clightning.FundchannelRequest{
		Id: []byte(peerKey),
		Amount: &clightning.AmountOrAll{
			Value: &clightning.AmountOrAll_Amount{
				Amount: &clightning.Amount{
					Msat: uint64(1000 * amtSat),
				},
			},
		},
	})
	return err
}

func (cln *ClnGrpcClient) ActiveChannels() (int, error) {
	resp, err := cln.client.ListFunds(context.Background(), &clightning.ListfundsRequest{})
	if err != nil {
		return 0, err
	}
	return len(resp.Channels), nil
}

func (cln *ClnGrpcClient) AddInvoice(amtMsat int64) (string, error) {
	panic("not needed")
}

func (cln *ClnGrpcClient) SendPayment(invoice string) error {
	panic("not needed") // TODO: Implement
}

func (cln *ClnGrpcClient) SendKeysend(destination string, amtMsat int64) error {
	_, err := cln.client.KeySend(context.Background(), &clightning.KeysendRequest{
		Destination: []byte(destination),
		AmountMsat: &clightning.Amount{
			Msat: uint64(amtMsat),
		},
	})
	return err
}

func (cln *ClnGrpcClient) Close() {
	return
}

func (cln *ClnGrpcClient) HasFunds() error {
	resp, err := cln.client.ListFunds(context.Background(), &clightning.ListfundsRequest{})
	if err != nil {
		return err
	}
	for _, op := range resp.Outputs {
		if op.Status == clightning.ListfundsOutputs_CONFIRMED {
			return nil
		}
	}
	return fmt.Errorf("no confirmed outputs")
}
