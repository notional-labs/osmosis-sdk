package client

import (
	govclient "github.com/cosmos/osmosis-sdk/x/gov/client"
	"github.com/cosmos/osmosis-sdk/x/upgrade/client/cli"
	"github.com/cosmos/osmosis-sdk/x/upgrade/client/rest"
)

var ProposalHandler = govclient.NewProposalHandler(cli.NewCmdSubmitUpgradeProposal, rest.ProposalRESTHandler)
var CancelProposalHandler = govclient.NewProposalHandler(cli.NewCmdSubmitCancelUpgradeProposal, rest.ProposalCancelRESTHandler)
