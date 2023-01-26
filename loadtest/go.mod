module github.com/bottlepay/loadtest

go 1.13

require (
	github.com/btcsuite/btcd v0.22.0-beta.0.20211005184431-e3449998be39
	github.com/btcsuite/btcutil v1.0.3-0.20210527170813-e2ba6805a890
	github.com/flitz-be/cln-grpc-go v0.0.0-20230126121549-549d2f9ede89
	github.com/gorilla/websocket v1.4.2
	github.com/lightningnetwork/lnd v0.14.3-beta
	github.com/niftynei/glightning v0.8.2
	github.com/urfave/cli v1.22.4
	go.uber.org/zap v1.17.0
	golang.org/x/net v0.5.0 // indirect
	google.golang.org/genproto v0.0.0-20230125152338-dcaf20b6aeaa // indirect
	google.golang.org/grpc v1.52.3
	gopkg.in/macaroon.v2 v2.1.0
	gopkg.in/yaml.v2 v2.4.0
)

replace github.com/niftynei/glightning => github.com/joostjager/glightning v0.8.3-0.20210325135629-f1548ac8aeb8
