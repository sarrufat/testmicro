package main

import (
	"github.com/sarrufat/testmicro/hello/handler"
	pb "github.com/sarrufat/testmicro/hello/proto"

	"github.com/micro/micro/v3/service"
	"github.com/micro/micro/v3/service/logger"
)

func main() {
	// Create service
	srv := service.New(
		service.Name("github.com/sarrufat/testmicro/hello"),
		service.Version("latest"),
	)

	// Register handler
	pb.RegisterGithub.Com/Sarrufat/Testmicro/HelloHandler(srv.Server(), new(handler.Github.Com/Sarrufat/Testmicro/Hello))

	// Run service
	if err := srv.Run(); err != nil {
		logger.Fatal(err)
	}
}
