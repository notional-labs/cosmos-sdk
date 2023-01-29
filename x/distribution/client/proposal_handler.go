package client

import (
	"github.com/cheqd/cosmos-sdk/x/distribution/client/cli"
	"github.com/cheqd/cosmos-sdk/x/distribution/client/rest"
	govclient "github.com/cheqd/cosmos-sdk/x/gov/client"
)

// ProposalHandler is the community spend proposal handler.
var (
	ProposalHandler = govclient.NewProposalHandler(cli.GetCmdSubmitProposal, rest.ProposalRESTHandler)
)
