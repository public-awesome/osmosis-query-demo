package main

import (
	"context"
	"crypto/tls"
	"log"
	"time"

	poolmanagerqueryproto "github.com/osmosis-labs/osmosis/v25/x/poolmanager/client/queryproto"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
)

func main() {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	conn, err := grpc.DialContext(ctx, "osmosis-grpc.lavenderfive.com:443",
		grpc.WithTransportCredentials(credentials.NewTLS(&tls.Config{})),
	)
	if err != nil {
		log.Fatalf("did not connect: %v", err)
		return
	}
	defer conn.Close()

	client := poolmanagerqueryproto.NewQueryClient(conn)
	// tokenIn := "1000000ibc/5DD1F95ED336014D00CE2520977EC71566D282F9749170ADC83A392E0EA7426A"

	resp, err := client.AllPools(context.Background(), &poolmanagerqueryproto.AllPoolsRequest{})
	if err != nil {
		log.Fatalf("could not get all pools: %v", err)
		return
	}
	log.Printf("All pools: %v", resp.Pools)

	// resp, err := client.EstimateSwapExactAmountIn(context.Background(), &poolmanagerqueryproto.EstimateSwapExactAmountInRequest{
	// 	TokenIn: tokenIn,
	// 	Routes: []poolmanagertypes.SwapAmountInRoute{
	// 		{
	// 			PoolId:        1397,
	// 			TokenOutDenom: "ibc/987C17B11ABC2B20019178ACE62929FE9840202CE79498E29FE8E5CB02B7C0A4",
	// 		},
	// 	},
	// })
	// if err != nil {
	// 	log.Fatalf("could not estimate swap exact amount in: %v", err)
	// }

	// log.Printf("Estimated swap exact amount in: %s", resp.TokenOutAmount)
}
