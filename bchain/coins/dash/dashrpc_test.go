// +build integration

package dash

import (
	"blockbook/bchain"
	"blockbook/bchain/tests/rpc"
	"encoding/json"
	"testing"
)

func getRPCClient(cfg json.RawMessage) (bchain.BlockChain, error) {
	c, err := NewDashRPC(cfg, nil)
	if err != nil {
		return nil, err
	}
	cli := c.(*DashRPC)
	cli.Parser = NewDashParser(GetChainParams("test"), cli.ChainConfig)
	cli.Mempool = bchain.NewUTXOMempool(cli, cli.ChainConfig.MempoolWorkers, cli.ChainConfig.MempoolSubWorkers)
	return cli, nil
}

var rpcTest *rpc.Test

func init() {
	t, err := rpc.NewTest("Dash Testnet", getRPCClient)
	if err != nil {
		panic(err)
	}
	rpcTest = t
}

func TestDashRPC_GetBlockHash(t *testing.T) {
	rpcTest.TestGetBlockHash(t)
}

func TestDashRPC_GetBlock(t *testing.T) {
	rpcTest.TestGetBlock(t)
}

func TestDashRPC_GetTransaction(t *testing.T) {
	rpcTest.TestGetTransaction(t)
}

func TestDashRPC_TestGetTransactionForMempool(t *testing.T) {
	rpcTest.TestGetTransactionForMempool(t)
}

func TestDashRPC_TestMempoolSync(t *testing.T) {
	rpcTest.TestMempoolSync(t)
}

func TestDashRPC_SendRawTransaction(t *testing.T) {
	rpcTest.TestSendRawTransaction(t)
}

func TestDashRPC_EstimateSmartFee(t *testing.T) {
	rpcTest.TestEstimateSmartFee(t)
}

func TestDashRPC_EstimateFee(t *testing.T) {
	rpcTest.TestEstimateFee(t)
}
