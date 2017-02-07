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

	// Handlers
	proto.RegisterDiscoveryHandler(service.Server(), new(handler.Discovery))
	proto2.RegisterRegistryHandler(service.Server(), new(handler.Registry))

	// Subscribers
	proto.RegisterSubscriber(discovery.HeartbeatTopic, service.Server(), discovery.Default.ProcessHeartbeat)
	proto.RegisterSubscriber(discovery.WatchTopic, service.Server(), discovery.Default.ProcessResult)

	if err := service.Run(); err != nil {
		log.Fatal(err)
	}
}
