package query

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/umee-network/umee/v2/x/leverage/types"
)

// GetBorrow wraps the leverage GetBorrow query.
type GetBorrow struct {
	BorrowerAddr sdk.AccAddress `json:"borrower_addr"`
	Denom        string         `json:"denom"`
}

// GetBorrowResponse wraps the response of GetBorrow query.
type GetBorrowResponse struct {
	BorrowedAmount sdk.Coin `json:"borrowed_amount"`
}

// GetAllRegisteredTokens wraps the leverage RegisteredTokens query.
type GetAllRegisteredTokens struct {
}

// GetAllRegisteredTokensResponse wraps the response of RegisteredTokens query.
type GetAllRegisteredTokensResponse struct {
	Registry []types.Token `json:"registry"`
}