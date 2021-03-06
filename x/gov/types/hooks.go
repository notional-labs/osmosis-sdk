package types

import (
	sdk "github.com/cosmos/osmosis-sdk/types"
)

var _ GovHooks = MultiGovHooks{}

// combine multiple governance hooks, all hook functions are run in array sequence
type MultiGovHooks []GovHooks

func NewMultiGovHooks(hooks ...GovHooks) MultiGovHooks {
	return hooks
}

func (h MultiGovHooks) AfterProposalSubmission(ctx sdk.Context, proposalID uint64) {
	for i := range h {
		h[i].AfterProposalSubmission(ctx, proposalID)
	}
}

func (h MultiGovHooks) AfterProposalDeposit(ctx sdk.Context, proposalID uint64, depositorAddr sdk.AccAddress) {
	for i := range h {
		h[i].AfterProposalDeposit(ctx, proposalID, depositorAddr)
	}
}

func (h MultiGovHooks) AfterProposalVote(ctx sdk.Context, proposalID uint64, voterAddr sdk.AccAddress) {
	for i := range h {
		h[i].AfterProposalVote(ctx, proposalID, voterAddr)
	}
}
func (h MultiGovHooks) AfterProposalInactive(ctx sdk.Context, proposalID uint64) {
	for i := range h {
		h[i].AfterProposalInactive(ctx, proposalID)
	}
}
func (h MultiGovHooks) AfterProposalActive(ctx sdk.Context, proposalID uint64) {
	for i := range h {
		h[i].AfterProposalActive(ctx, proposalID)
	}
}
