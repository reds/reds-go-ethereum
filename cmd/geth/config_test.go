package main

import (
	"testing"
	"time"

	"github.com/ethereum/go-ethereum/log"
	"gopkg.in/urfave/cli.v1"
)

func TestConfigArgs(t *testing.T) {
	durationArg := "2m11s"

	// command line
	args := []string{"", "--http.timeout", durationArg, "--dev"}
	// app created in init()
	app.Action = action(t, durationArg, "HTTPRpcTimeout should be "+durationArg)
	err := app.Run(args)
	if err != nil {
		t.Fatal(err)
	}

	// config file
	args = []string{"", "--config", "testdata/config.toml", "--dev"}
	app.Action = action(t, "20s", "HTTPRpcTimeout should 20s")
	err = app.Run(args)

	// commandline overrides config file
	args = []string{"", "--http.timeout", durationArg, "--config", "testdata/config.toml", "--dev"}
	app.Action = action(t, durationArg, "HTTPRpcTimeout should be "+durationArg)
	err = app.Run(args)

	if err != nil {
		t.Fatal(err)
	}
}

func action(t *testing.T, da, es string) func(ctx *cli.Context) error {
	d, err := time.ParseDuration(da)
	if err != nil {
		t.Fatal(err)
	}
	return func(ctx *cli.Context) error {
		log.Root().SetHandler(log.DiscardHandler())
		stack, cfg := makeConfigNode(ctx)
		if d != cfg.Eth.HTTPRpcTimeout {
			t.Fatal(es)
		}
		stack.Close() // cleanup db
		return nil
	}
}
