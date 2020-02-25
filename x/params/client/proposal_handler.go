package client

import (
	govclient "github.com/serjplus/cosmos-sdk/x/gov/client"
	"github.com/serjplus/cosmos-sdk/x/params/client/cli"
	"github.com/serjplus/cosmos-sdk/x/params/client/rest"
)

// param change proposal handler
var ProposalHandler = govclient.NewProposalHandler(cli.GetCmdSubmitProposal, rest.ProposalRESTHandler)
