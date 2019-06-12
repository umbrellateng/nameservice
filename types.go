package nameservice

import (
	"fmt"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"strings"
)


type Whois struct {
	Value string 			`json:"value"`
	Owner sdk.AccAddress	`json:"owner"`
	Price sdk.Coins			`json:"price"`
}


var MinNamePrice = sdk.Coins{sdk.NewInt64Coin("nametoken", 1)}

func NewWhois() Whois {

	return Whois{
		Price: MinNamePrice,
	}
}

func (w Whois) String() string  {

	return strings.TrimSpace(
		fmt.Sprint(
			"Whois {\n\t\t",
					"Owner: ", w.Owner, "\n\t\t",
					"Value: ", w.Value, "\n\t\t",
					"Price: ", w.Price, "\n\t}",
			))
}



