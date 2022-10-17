package main

import (
	"github.com/bloc-chain/bc-faucet/cmd"
)

//go:generate npm run build
func main() {
	cmd.Execute()
}
