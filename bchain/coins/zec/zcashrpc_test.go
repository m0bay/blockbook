// +build integration

package zec

import (
	"blockbook/bchain"
	"blockbook/bchain/tests/rpc"
	"encoding/json"
	"testing"
)

func getRPCClient(cfg json.RawMessage) (bchain.BlockChain, error) {
	c, err := NewZCashRPC(cfg, nil)
	if err != nil {
		return nil, err
	}
	cli := c.(*ZCashRPC)
	cli.Parser = NewZCashParser(cli.ChainConfig)
	cli.Mempool = bchain.NewUTXOMempool(cli, cli.ChainConfig.MempoolWorkers, cli.ChainConfig.MempoolSubWorkers)
	return cli, nil
}

var rpcTest *rpc.Test

func init() {
	t, err := rpc.NewTest("Zcash Testnet", getRPCClient)
	if err != nil {
		panic(err)
	}
	rpcTest = t
}

func TestZCashRPC_GetBlockHash(t *testing.T) {
	rpcTest.TestGetBlockHash(t)
}

func TestZCashRPC_GetBlock(t *testing.T) {
	rpcTest.TestGetBlock(t)
}

func TestZCashRPC_GetTransaction(t *testing.T) {
	rpcTest.TestGetTransaction(t)
}

func TestZCashRPC_TestGetTransactionForMempool(t *testing.T) {
	rpcTest.TestGetTransactionForMempool(t)
}

func TestZCashRPC_TestMempoolSync(t *testing.T) {
	rpcTest.TestMempoolSync(t)
}

func TestZCashRPC_SendRawTransaction(t *testing.T) {
	rpcTest.TestSendRawTransaction(t)
}

func TestZCashRPC_EstimateSmartFee(t *testing.T) {
	rpcTest.TestEstimateSmartFee(t)
}

func TestZCashRPC_EstimateFee(t *testing.T) {
	rpcTest.TestEstimateFee(t)
}
