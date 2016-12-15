package main

import (
	"log"
	"os"

	"github.com/micro/discovery-srv/discovery"
	"github.com/micro/discovery-srv/handler"
	"github.com/micro/go-micro"

	proto "github.com/micro/discovery-srv/proto/discovery"
	proto2 "github.com/micro/discovery-srv/proto/registry"
	_ "github.com/micro/go-os/discovery"
)

func init() {
	os.Setenv("MICRO_REGISTRY", "os")
}

func main() {
	service := micro.NewService(
		micro.Name("go.micro.srv.discovery"),
		micro.BeforeStart(discovery.Start),
		micro.AfterStop(discovery.Stop),
	)

	service.Init()

	service.Server().Subscribe(
		service.Server().NewSubscriber(
			discovery.HeartbeatTopic,
			discovery.Default.ProcessHeartbeat,
		),
	)

	service.Server().Subscribe(
		service.Server().NewSubscriber(
			discovery.WatchTopic,
			discovery.Default.ProcessResult,
		),
	)

	proto.RegisterDiscoveryHandler(service.Server(), new(handler.Discovery))
	proto2.RegisterRegistryHandler(service.Server(), new(handler.Registry))

	if err := service.Run(); err != nil {
		log.Fatal(err)
	}
}
