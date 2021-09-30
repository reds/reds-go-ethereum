package main

import (
	"testing"
	"time"

	"github.com/ethereum/go-ethereum/log"
	"gopkg.in/urfave/cli.v1"
)

func TestConfigArgs(t *testing.T) {
	// app created in init()
	durationArg := "2m11s"
	duration, err := time.ParseDuration(durationArg)
	if err != nil {
		t.Fatal(err)
	}

	// command line
	args := []string{"", "--http.timeout", durationArg, "--datadir", ""}
	app.Action = func(ctx *cli.Context) error {
		log.Root().SetHandler(log.DiscardHandler())
		stack, cfg := makeConfigNode(ctx)
		if duration != cfg.Eth.HTTPRpcTimeout {
			t.Fatal("HTTPRpcTimeout should be", durationArg)
		}
		stack.Close() // cleanup db
		return nil
	}
	err = app.Run(args)

	// config file
	args = []string{"", "--config", "testdata/config.toml", "--datadir", ""}
	app.Action = func(ctx *cli.Context) error {
		log.Root().SetHandler(log.DiscardHandler())
		stack, cfg := makeConfigNode(ctx)
		if cfg.Eth.HTTPRpcTimeout.Seconds() != 20 {
			t.Fatal("HTTPRpcTimeout should be 20s")
		}
		stack.Close()
		return nil
	}
	err = app.Run(args)

	// command line overrides config file
	args = []string{"", "--http.timeout", durationArg, "--config", "testdata/config.toml", "--datadir", ""}
	app.Action = func(ctx *cli.Context) error {
		log.Root().SetHandler(log.DiscardHandler())
		stack, cfg := makeConfigNode(ctx)
		if duration != cfg.Eth.HTTPRpcTimeout {
			t.Fatal("HTTPRpcTimeout should be", duration)
		}
		stack.Close()
		return nil
	}
	err = app.Run(args)

	if err != nil {
		t.Fatal(err)
	}
}
