package context

import (
	"fmt"
	"path/filepath"

	"github.com/pkg/errors"
	"github.com/tendermint/tendermint/libs/log"
	tmlite "github.com/tendermint/tendermint/lite"
	tmliteproxy "github.com/tendermint/tendermint/lite/proxy"
	rpcclient "github.com/tendermint/tendermint/rpc/client"
)

const (
	verifierDir = ".lite_verifier"

	// DefaultVerifierCacheSize defines the default Tendermint cache size.
	DefaultVerifierCacheSize = 10
)

// CreateVerifier returns a Tendermint verifier from a CLIContext object and
// cache size. An error is returned if the CLIContext is missing required values
// or if the verifier could not be created. A CLIContext must at the very least
// have the chain ID and home directory set. If the CLIContext has TrustNode
// enabled, no verifier will be created.
func CreateVerifier(ctx CLIContext, cacheSize int) (tmlite.Verifier, error) {
	if ctx.TrustNode {
		return nil, nil
	}

	fmt.Println("CREATEVERIFIER-CHAINID", ctx.ChainID)

	switch {
	case ctx.ChainID == "":
		return nil, errors.New("must provide a valid chain ID to create verifier")

	case ctx.HomeDir == "":
		return nil, errors.New("must provide a valid home directory to create verifier")

	case ctx.Client == nil && ctx.NodeURI == "":
		return nil, errors.New("must provide a valid RPC client or RPC URI to create verifier")
	}

	// create an RPC client based off of the RPC URI if no RPC client exists
	client := ctx.Client
	if client == nil {
		client = rpcclient.NewHTTP(ctx.NodeURI, "/websocket")
	}

	return tmliteproxy.NewVerifier(
		ctx.ChainID, filepath.Join(ctx.HomeDir, ctx.ChainID, verifierDir),
		client, log.NewNopLogger(), cacheSize,
	)
}