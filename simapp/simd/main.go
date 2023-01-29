package main

import (
	"os"

	"github.com/cheqd/cosmos-sdk/server"
	svrcmd "github.com/cheqd/cosmos-sdk/server/cmd"
	"github.com/cheqd/cosmos-sdk/simapp"
	"github.com/cheqd/cosmos-sdk/simapp/simd/cmd"
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
