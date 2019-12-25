package tx

import (
	"fmt"
	"gitlab.bianjie.ai/irita/irita-sdk-go/client/lcd"
	"gitlab.bianjie.ai/irita/irita-sdk-go/client/rpc"
	"gitlab.bianjie.ai/irita/irita-sdk-go/client/types"
	"gitlab.bianjie.ai/irita/irita-sdk-go/keys"
	commontypes "gitlab.bianjie.ai/irita/irita-sdk-go/types"
	"gitlab.bianjie.ai/irita/irita-sdk-go/util/constant"
	iritaConfig "gitlab.bianjie.ai/irita/irita/config"
)

type TxClient interface {
	SendToken(receiver string, coins []types.Coin, memo string, commit bool) (types.BroadcastTxResult, error)
	PostServiceRequest(request ServiceRequest, memo string, commit bool) (types.BroadcastTxResult, error)
	PostServiceResponse(response ServiceResponse, memo string, commit bool) (types.BroadcastTxResult, error)
}

type client struct {
	chainId    string
	keyManager keys.KeyManager
	liteClient lcd.LiteClient
	rpcClient  rpc.RPCClient
}

func NewClient(chainId string, networkType commontypes.NetworkType, keyManager keys.KeyManager,
	liteClient lcd.LiteClient, rpcClient rpc.RPCClient) (TxClient, error) {
	var (
		network string
	)
	switch networkType {
	case commontypes.Mainnet:
		network = constant.NetworkTypeMainnet
	case commontypes.Testnet:
		network = constant.NetworkTypeTestnet
	default:
		return &client{}, fmt.Errorf("invalid networktype, %d", networkType)
	}
	iritaConfig.SetNetworkType(network)

	return &client{
		chainId:    chainId,
		keyManager: keyManager,
		liteClient: liteClient,
		rpcClient:  rpcClient,
	}, nil
}
