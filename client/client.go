package main

import (
	"context"
	"fmt"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"

	"github.com/asadbek/app/dict"
)

func main() {

	conn, err := grpc.Dial("localhost:9101",
		grpc.WithTransportCredentials(insecure.NewCredentials()),
		grpc.WithBlock(),
	)

	if err != nil {
		panic(err)
	}
	defer conn.Close()

	c := dict.NewTranslateClient(conn)

	// resp, err := c.Dictionary(
	// 	context.Background(),
	// 	&dict.DictionaryRequest{
	// 		Key: "apple",
	// 	},
	// )
	// resp2, err := c.Add(
	// 	context.Background(),
	// 	&dict.AddRequest{
	// 		Key:  1,
	// 		Key1: 2,
	// 	},
	// )
	// resp3, err := c.GetCurrency(
	// 	context.Background(),
	// 	&dict.CurrencyRequest{
	// 		Ccy: "US",
	// 	},
	// )

	// resp4, err := c.Converter(
	// 	context.Background(),
	// 	&dict.MoneyConverterRequest{
	// 		Money: "115000000",
	// 		Ccy:   "EUR",
	// 	},
	// )
	// resp5, err := c.Square(
	// 	context.Background(),
	// 	&dict.NumberRequest{
	// 		Number: 2,
	// 		Degree: 4,
	// 	},
	// )
	resp6, err := c.MaxNum(
		context.Background(),
		&dict.MaxRequest{
			Nums: []int64{9, 5, 2, 5, 3, 67},
		},
	)
	if err != nil {
		fmt.Println("translate -> Dictionary ->> ", err.Error())
		return
	}

	// fmt.Println(resp.Value)
	// fmt.Println(resp2.Result)
	// fmt.Println(resp3)
	// fmt.Println(resp4)
	// fmt.Println(resp5)
	fmt.Println(resp6)

}
