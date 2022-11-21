# Currency

multi currency support for concurrent pricing

## Usage

```go

import (
    "github.com/badappsdev/currency"
)

func main() {

    // create new usd 
    amount := currency.NewUSD(1, 99)

    // retrieve the values 
    dollars, cents := amount.Value()

    // prep for adding two values
    tip := currency.NewUSD(0, 1)

    // add the tip to the total amount
    amount.Add(tip)

    // print total amount
    // expected: total amount $2.00
    fmt.Sprintf("total amount $%d.%d", amount.dollars, m.cents)
}

```

## Supported Currencies

supported currencies will change based on project requirements.

currently supported

    - United States Dollar (usd) 