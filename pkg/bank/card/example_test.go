package card

import (
	"bank/pkg/bank/types"
	"fmt"
)

func ExampleTotal() {
	fmt.Println(Total([]types.Card{
		{
			Balance: 1_000_00,
			Active: true,

		},
	}))
	fmt.Println(Total([]types.Card{
		{
			Balance: 1_000_00,
			Active: true,
		},
		{
			Balance: 2_000_00,
			Active: true,
		},
	}))
	fmt.Println(Total([]types.Card{
		{
			Balance: 1_000_00,
			Active: false,

		},
	}))
	fmt.Println(Total([]types.Card{
		{
			Balance: -1_000_00,
			Active: true,

		},

	}))

	//Output:
	//100000
	//300000
	//0
	//0
}

func ExamplePaymentSources(){

	paymentSource:= []types.PaymentSource{}

	paymentSource = PaymentSources(
		[]types.Card{
			{
				PAN: "5058 xxxx xxxx 9999",
				Balance: 999_99,
				Active:true,
			},
			{
				PAN: "5058 xxxx xxxx 8888",
				Balance: 888_88,
				Active:true,

			},

		},

	)	


	for _, source := range paymentSource{

		fmt.Println(source.Number)
	}
	//Output:
	//5058 xxxx xxxx 9999
	//5058 xxxx xxxx 8888
}
