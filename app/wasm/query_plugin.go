package wasm

import (
	"encoding/json"

	wasmvmtypes "github.com/CosmWasm/wasmvm/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"

	"github.com/umee-network/umee/v2/app/wasm/query"
	leveragekeeper "github.com/umee-network/umee/v2/x/leverage/keeper"
	oraclekeeper "github.com/umee-network/umee/v2/x/oracle/keeper"
)

// QueryPlugin wraps the query plugin with keepers
type QueryPlugin struct {
	leverageKeeper leveragekeeper.Keeper
	oracleKeeper   oraclekeeper.Keeper
}

// NewQueryPlugin basic constructor
func NewQueryPlugin(
	leverageKeeper leveragekeeper.Keeper,
	oracleKeeper oraclekeeper.Keeper,
) *QueryPlugin {
	return &QueryPlugin{
		leverageKeeper: leverageKeeper,
		oracleKeeper:   oracleKeeper,
	}
}

// GetBorrow wraps leverage GetBorrow
func (qp *QueryPlugin) GetBorrow(ctx sdk.Context, borrowerAddr sdk.AccAddress, denom string) sdk.Coin {
	return qp.leverageKeeper.GetBorrow(ctx, borrowerAddr, denom)
}

// GetExchangeRateBase wraps oracle GetExchangeRateBase
func (qp *QueryPlugin) GetExchangeRateBase(ctx sdk.Context, denom string) (sdk.Dec, error) {
	return qp.oracleKeeper.GetExchangeRateBase(ctx, denom)
}

// CustomQuerier implements custom querier for wasm smartcontracts acess umee native modules
func CustomQuerier(queryPlugin *QueryPlugin) func(ctx sdk.Context, request json.RawMessage) ([]byte, error) {
	return func(ctx sdk.Context, request json.RawMessage) ([]byte, error) {
		var smartcontractQuery query.UmeeQuery
		if err := json.Unmarshal(request, &smartcontractQuery); err != nil {
			return nil, sdkerrors.Wrap(err, "umee query")
		}

		switch smartcontractQuery.AssignedQuery {
		case query.AssignedQueryGetBorrow:
			return smartcontractQuery.HandleGetBorrow(ctx, queryPlugin)
		case query.AssignedQueryGetExchangeRateBase:
			return smartcontractQuery.HandleGetExchangeRateBase(ctx, queryPlugin)

		default:
			return nil, wasmvmtypes.UnsupportedRequest{Kind: "invalid assigned umee query"}
		}
	}
}