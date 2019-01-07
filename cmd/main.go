package main

import (
	"context"
	"flag"

	log "github.com/sirupsen/logrus"

	"github.com/kyeett/gameserver/types"
	"github.com/onrik/logrus/filename"

	"github.com/kyeett/gameserver/grpc"
)

func main() {
	var port string
	flag.StringVar(&port, "port", "10001", "port to serve on")
	flag.Parse()

	//Enable tracing
	log.SetLevel(log.DebugLevel)
	filenameHook := filename.NewHook()
	filenameHook.Field = "source"
	log.AddHook(filenameHook)

	ss, err := grpc.NewServer(types.FirstWorld)
	if err != nil {
		log.Fatal(err)
	}

	ss.RunWeb(context.Background(), port)
}
