package main

import (
	"os"

	"github.com/cosmos/osmosis-sdk/server"
	svrcmd "github.com/cosmos/osmosis-sdk/server/cmd"
	"github.com/cosmos/osmosis-sdk/simapp"
	"github.com/cosmos/osmosis-sdk/simapp/simd/cmd"
)

func main() {
	rootCmd, _ := cmd.NewRootCmd()

	if err := svrcmd.Execute(rootCmd, simapp.DefaultNodeHome); err != nil {
		switch e := err.(type) {
		case server.ErrorCode:
			os.Exit(e.Code)

		default:
			os.Exit(1)
		}
	}
}
