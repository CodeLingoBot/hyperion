package server

import (
	"context"
	"os"
	"os/signal"
	"sync"

	"github.com/nickysemenza/gola"
	"github.com/nickysemenza/hyperion/util/clock"

	log "github.com/sirupsen/logrus"

	"github.com/nickysemenza/hyperion/util/tracing"

	"github.com/nickysemenza/hyperion/api"
	"github.com/nickysemenza/hyperion/control/homekit"
	"github.com/nickysemenza/hyperion/core/config"
	"github.com/nickysemenza/hyperion/core/cue"
	"github.com/nickysemenza/hyperion/core/light"
)

//Run starts the server
func Run(ctx context.Context) {

	ctx, cancel := context.WithCancel(ctx)
	wg := sync.WaitGroup{}

	master := cue.InitializeMaster(clock.RealClock{})

	go tracing.InitTracer(ctx)
	light.Initialize(ctx)
	//Set up Homekit Server
	wg.Add(1)
	go homekit.Start(ctx, &wg)

	//Set up RPC server
	wg.Add(1)
	go api.ServeRPC(ctx, &wg)

	//Setup API server
	wg.Add(1)
	go api.ServeHTTP(ctx, &wg)

	//proceess cues forever
	master.ProcessForever(ctx)

	olaConfig := config.GetServerConfig(ctx).Outputs.OLA
	if !olaConfig.Enabled {
		log.Info("ola output is not enabled")
	} else {
		client, err := gola.New(olaConfig.Address)
		if err != nil {
			log.Errorf("could not start DMX worker: could not connect to ola: %v", err)
		} else {
			wg.Add(1)
			go light.SendDMXWorker(ctx, client, olaConfig.Tick, &wg)
		}
	}

	//handle CTRL+C interrupt
	quit := make(chan os.Signal)
	signal.Notify(quit, os.Interrupt)

	<-quit
	log.Println("shutting down hyperion server")
	cancel()
	wg.Wait()
}
