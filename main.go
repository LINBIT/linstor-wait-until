package main

import (
	"context"
	"fmt"
	"os"
	"os/signal"
	"time"

	"github.com/LINBIT/golinstor/client"
	"github.com/sirupsen/logrus"
)

var (
	Version = "unknown"
)

type WaitFunc func(ctx context.Context, lc *client.Client) error

func main() {
	ctx, cancel := signal.NotifyContext(context.Background(), os.Interrupt)
	defer cancel()

	log := logrus.WithField("version", Version)

	if len(os.Args) < 2 {
		log.Fatal("expected at least 1 argument")
	}

	var err error
	var waitCheck WaitFunc
	switch os.Args[1] {
	case "api-online":
		waitCheck, err = waitApiOnline(os.Args[2:]...)
	case "satellite-online":
		waitCheck, err = waitSatelliteOnline(os.Args[2:]...)
	default:
		log.Fatalf("unknown argument %s", os.Args[1])
	}

	if err != nil {
		log.WithError(err).Fatal("error parsing arguments")
	}

	lc, err := client.NewClient(client.Log(log))
	if err != nil {
		log.WithError(err).Fatal("failed to create LINSTOR client")
	}

	for {
		err := waitCheck(ctx, lc)
		if err == nil {
			return
		}

		if ctx.Err() != nil {
			log.WithError(err).Fatal("context cancelled")
		}

		log.WithError(err).Info("not ready")

		time.Sleep(10 * time.Second)
	}
}

func waitApiOnline(args ...string) (WaitFunc, error) {
	if len(args) != 0 {
		return nil, fmt.Errorf("expected 0 arguments, got %d", len(args))
	}

	return func(ctx context.Context, lc *client.Client) error {
		_, err := lc.Controller.GetVersion(ctx)
		return err
	}, nil
}

func waitSatelliteOnline(args ...string) (WaitFunc, error) {
	if len(args) != 1 {
		return nil, fmt.Errorf("expected exactly 1 argument (satellite-name), got %d", len(args))
	}

	return func(ctx context.Context, lc *client.Client) error {
		node, err := lc.Nodes.Get(ctx, args[0])
		if err != nil {
			return err
		}

		if node.ConnectionStatus != "ONLINE" {
			return fmt.Errorf("satellite %s is not ONLINE: %s", node.Name, node.ConnectionStatus)
		}

		return nil
	}, nil
}
