package derive

import (
	"os"

	"github.com/ethereum-optimism/optimism/op-node/rollup"
	"github.com/ethereum/go-ethereum/log"
)

// DerivationVersionCelestia is a byte marker for celestia references submitted
// to the batch inbox address as calldata.
// Mnemonic 0xce = celestia
// version 0xce references are encoded as:
// [8]byte block height ++ [32]byte commitment
// in little-endian encoding.
// see: https://github.com/rollkit/celestia-da/blob/1f2df375fd2fcc59e425a50f7eb950daa5382ef0/celestia.go#L141-L160
const DerivationVersionCelestia = 0xce

var daClient *rollup.DAClient

func init() {
	daRpc := os.Getenv("OP_NODE_DA_RPC")
	if daRpc == "" {
		daRpc = "localhost:26650"
	}
	var err error
	daClient, err = rollup.NewDAClient(daRpc)
	if err != nil {
		log.Error("celestia: unable to create DA client", "rpc", daRpc, "err", err)
		panic(err)
	}
}
