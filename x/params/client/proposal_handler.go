package client

import (
	govclient "github.com/cosmos/osmosis-sdk/x/gov/client"
	"github.com/cosmos/osmosis-sdk/x/params/client/cli"
	"github.com/cosmos/osmosis-sdk/x/params/client/rest"
)

// ProposalHandler is the param change proposal handler.
var ProposalHandler = govclient.NewProposalHandler(cli.NewSubmitParamChangeProposalTxCmd, rest.ProposalRESTHandler)
