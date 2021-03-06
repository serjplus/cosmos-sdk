package client

import (
	govclient "github.com/serjplus/cosmos-sdk/x/gov/client"
	"github.com/serjplus/cosmos-sdk/x/upgrade/client/cli"
	"github.com/serjplus/cosmos-sdk/x/upgrade/client/rest"
)

var ProposalHandler = govclient.NewProposalHandler(cli.GetCmdSubmitUpgradeProposal, rest.ProposalRESTHandler)
