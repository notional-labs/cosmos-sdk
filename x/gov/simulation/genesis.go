package simulation

// DONTCOVER

import (
	"encoding/json"
	"fmt"
	"math/rand"
	"time"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/module"
	"github.com/cosmos/cosmos-sdk/types/simulation"
	"github.com/cosmos/cosmos-sdk/x/gov/types"
)

// Simulation parameter constants
const (
	DepositParamsMinDeposit           = "deposit_params_min_deposit"
	DepositParamsMinExpeditedDeposit  = "deposit_params_min_expedited_deposit"
	DepositParamsDepositPeriod        = "deposit_params_deposit_period"
	VotingParamsVotingPeriod          = "voting_params_voting_period"
	ExpeditedVotingParamsVotingPeriod = "expedited_voting_params_voting_period"
	TallyParamsQuorum                 = "tally_params_quorum"
	TallyParamsExpeditedQuorum        = "tally_params_expedited_quorum"
	TallyParamsThreshold              = "tally_params_threshold"
	TallyParamsExpeditedThreshold     = "tally_params_expedited_threshold"
	TallyParamsVeto                   = "tally_params_veto"
)

// GenDepositParamsDepositPeriod randomized DepositParamsDepositPeriod
func GenDepositParamsDepositPeriod(r *rand.Rand) time.Duration {
	return time.Duration(simulation.RandIntBetween(r, 1, 60*60*24*2)) * time.Second
}

// GenDepositParamsMinDeposit randomized DepositParamsMinDeposit
func GenDepositParamsMinDeposit(r *rand.Rand) sdk.Coins {
	return sdk.NewCoins(sdk.NewInt64Coin(sdk.DefaultBondDenom, int64(simulation.RandIntBetween(r, 1, 1e3))))
}

// GenDepositParamsMinExpeditedDeposit randomized DepositParamsMinExpeditedDeposit
func GenDepositParamsMinExpeditedDeposit(r *rand.Rand) sdk.Coins {
	return sdk.NewCoins(sdk.NewInt64Coin(sdk.DefaultBondDenom, int64(simulation.RandIntBetween(r, 2e3, 3e3))))
}

// GenVotingParamsVotingPeriod randomized VotingParamsVotingPeriod
func GenVotingParamsVotingPeriod(r *rand.Rand) time.Duration {
	return time.Duration(simulation.RandIntBetween(r, 60*60*24, 60*60*24*2)) * time.Second
}

// GenVotingParamsExpeditedVotingPeriod randomized VotingParamsExpeditedVotingPeriod
func GenVotingParamsExpeditedVotingPeriod(r *rand.Rand) time.Duration {
	return time.Duration(simulation.RandIntBetween(r, 1, 60*60*24)) * time.Second
}

// GenTallyParamsQuorum randomized TallyParamsQuorum
func GenTallyParamsQuorum(r *rand.Rand) sdk.Dec {
	return sdk.NewDecWithPrec(int64(simulation.RandIntBetween(r, 334, 500)), 3)
}

// GenTallyExpeditedParamsQuorum randomized TallyParamsExpeditedQuorum
func GenTallyExpeditedParamsQuorum(r *rand.Rand) sdk.Dec {
	return sdk.NewDecWithPrec(int64(simulation.RandIntBetween(r, 500, 550)), 3)
}

// GenTallyParamsThreshold randomized TallyParamsThreshold
func GenTallyParamsThreshold(r *rand.Rand) sdk.Dec {
	return sdk.NewDecWithPrec(int64(simulation.RandIntBetween(r, 450, 500)), 3)
}

// GenTallyParamsExpeditedThreshold randomized TallyParamsExpeditedThreshold
func GenTallyParamsExpeditedThreshold(r *rand.Rand) sdk.Dec {
	return sdk.NewDecWithPrec(int64(simulation.RandIntBetween(r, 500, 550)), 3)
}

// GenTallyParamsVeto randomized TallyParamsVeto
func GenTallyParamsVeto(r *rand.Rand) sdk.Dec {
	return sdk.NewDecWithPrec(int64(simulation.RandIntBetween(r, 250, 334)), 3)
}

// RandomizedGenState generates a random GenesisState for gov
func RandomizedGenState(simState *module.SimulationState) {
	startingProposalID := uint64(simState.Rand.Intn(100))

	var minDeposit sdk.Coins
	simState.AppParams.GetOrGenerate(
		simState.Cdc, DepositParamsMinDeposit, &minDeposit, simState.Rand,
		func(r *rand.Rand) { minDeposit = GenDepositParamsMinDeposit(r) },
	)

	var minExpeditedDeposit sdk.Coins
	simState.AppParams.GetOrGenerate(
		simState.Cdc, DepositParamsMinExpeditedDeposit, &minDeposit, simState.Rand,
		func(r *rand.Rand) { minDeposit = GenDepositParamsMinExpeditedDeposit(r) },
	)

	var depositPeriod time.Duration
	simState.AppParams.GetOrGenerate(
		simState.Cdc, DepositParamsDepositPeriod, &depositPeriod, simState.Rand,
		func(r *rand.Rand) { depositPeriod = GenDepositParamsDepositPeriod(r) },
	)

	var votingPeriod time.Duration
	simState.AppParams.GetOrGenerate(
		simState.Cdc, VotingParamsVotingPeriod, &votingPeriod, simState.Rand,
		func(r *rand.Rand) { votingPeriod = GenVotingParamsVotingPeriod(r) },
	)

	var expeditedVotingPeriod time.Duration
	simState.AppParams.GetOrGenerate(
		simState.Cdc, ExpeditedVotingParamsVotingPeriod, &expeditedVotingPeriod, simState.Rand,
		func(r *rand.Rand) { expeditedVotingPeriod = GenVotingParamsExpeditedVotingPeriod(r) },
	)

	var quorum sdk.Dec
	simState.AppParams.GetOrGenerate(
		simState.Cdc, TallyParamsQuorum, &quorum, simState.Rand,
		func(r *rand.Rand) { quorum = GenTallyParamsQuorum(r) },
	)

	var expeditedQuorum sdk.Dec
	simState.AppParams.GetOrGenerate(
		simState.Cdc, TallyParamsExpeditedQuorum, &expeditedQuorum, simState.Rand,
		func(r *rand.Rand) { expeditedQuorum = GenTallyExpeditedParamsQuorum(r) },
	)

	var threshold sdk.Dec
	simState.AppParams.GetOrGenerate(
		simState.Cdc, TallyParamsThreshold, &threshold, simState.Rand,
		func(r *rand.Rand) { threshold = GenTallyParamsThreshold(r) },
	)

	var expeditedThreshold sdk.Dec
	simState.AppParams.GetOrGenerate(
		simState.Cdc, TallyParamsExpeditedThreshold, &expeditedThreshold, simState.Rand,
		func(r *rand.Rand) { expeditedThreshold = GenTallyParamsExpeditedThreshold(r) },
	)

	var veto sdk.Dec
	simState.AppParams.GetOrGenerate(
		simState.Cdc, TallyParamsVeto, &veto, simState.Rand,
		func(r *rand.Rand) { veto = GenTallyParamsVeto(r) },
	)

	govGenesis := types.NewGenesisState(
		startingProposalID,
		types.NewDepositParams(minDeposit, minExpeditedDeposit, depositPeriod),
		types.NewVotingParams(votingPeriod, expeditedVotingPeriod),
		types.NewTallyParams(quorum, expeditedQuorum, threshold, expeditedThreshold, veto),
	)

	bz, err := json.MarshalIndent(&govGenesis, "", " ")
	if err != nil {
		panic(err)
	}
	fmt.Printf("Selected randomly generated governance parameters:\n%s\n", bz)
	simState.GenState[types.ModuleName] = simState.Cdc.MustMarshalJSON(govGenesis)
}
