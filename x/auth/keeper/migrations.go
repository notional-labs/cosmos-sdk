package keeper

import (
	"github.com/gogo/protobuf/grpc"

	v043 "github.com/cosmos/cosmos-sdk/x/auth/legacy/v043"
	v045 "github.com/cosmos/cosmos-sdk/x/auth/migrations/v045"
	"github.com/cosmos/cosmos-sdk/x/auth/types"

	sdk "github.com/cosmos/cosmos-sdk/types"
)

// Migrator is a struct for handling in-place store migrations.
type Migrator struct {
	keeper      AccountKeeper
	queryServer grpc.Server
}

// NewMigrator returns a new Migrator.
func NewMigrator(keeper AccountKeeper, queryServer grpc.Server) Migrator {
	return Migrator{keeper: keeper, queryServer: queryServer}
}

// Migrate1to2 migrates from version 1 to 2.
func (m Migrator) Migrate1to2(ctx sdk.Context) error {
	var iterErr error

	m.keeper.IterateAccounts(ctx, func(account types.AccountI) (stop bool) {
		wb, err := v043.MigrateAccount(ctx, account, m.queryServer)
		if err != nil {
			iterErr = err
			return true
		}

		if wb == nil {
			return false
		}

		m.keeper.SetAccount(ctx, wb)
		return false
	})

	return iterErr
}

// Migrate2to3 migrates x/auth state from consensus version 2 to 3.
func (m Migrator) Migrate2to3(ctx sdk.Context) error {
	return v045.MigrateStore(ctx, m.keeper.key, m.keeper.cdc, m.keeper.paramSubspace)
}
