package client

import (
	"github.com/cosmos/osmosis-sdk/x/distribution/client/cli"
	"github.com/cosmos/osmosis-sdk/x/distribution/client/rest"
	govclient "github.com/cosmos/osmosis-sdk/x/gov/client"
)

// ProposalHandler is the community spend proposal handler.
var (
	ProposalHandler = govclient.NewProposalHandler(cli.GetCmdSubmitProposal, rest.ProposalRESTHandler)
)
