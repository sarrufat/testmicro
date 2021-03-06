package handler

import (
	"context"

	log "github.com/micro/micro/v3/service/logger"

	github.com/sarrufat/testmicro/hello "github.com/sarrufat/testmicro/hello/proto"
)

type Github.Com/Sarrufat/Testmicro/Hello struct{}

// Call is a single request handler called via client.Call or the generated client code
func (e *Github.Com/Sarrufat/Testmicro/Hello) Call(ctx context.Context, req *github.com/sarrufat/testmicro/hello.Request, rsp *github.com/sarrufat/testmicro/hello.Response) error {
	log.Info("Received Github.Com/Sarrufat/Testmicro/Hello.Call request")
	rsp.Msg = "Hello " + req.Name
	return nil
}

// Stream is a server side stream handler called via client.Stream or the generated client code
func (e *Github.Com/Sarrufat/Testmicro/Hello) Stream(ctx context.Context, req *github.com/sarrufat/testmicro/hello.StreamingRequest, stream github.com/sarrufat/testmicro/hello.Github.Com/Sarrufat/Testmicro/Hello_StreamStream) error {
	log.Infof("Received Github.Com/Sarrufat/Testmicro/Hello.Stream request with count: %d", req.Count)

	for i := 0; i < int(req.Count); i++ {
		log.Infof("Responding: %d", i)
		if err := stream.Send(&github.com/sarrufat/testmicro/hello.StreamingResponse{
			Count: int64(i),
		}); err != nil {
			return err
		}
	}

	return nil
}

// PingPong is a bidirectional stream handler called via client.Stream or the generated client code
func (e *Github.Com/Sarrufat/Testmicro/Hello) PingPong(ctx context.Context, stream github.com/sarrufat/testmicro/hello.Github.Com/Sarrufat/Testmicro/Hello_PingPongStream) error {
	for {
		req, err := stream.Recv()
		if err != nil {
			return err
		}
		log.Infof("Got ping %v", req.Stroke)
		if err := stream.Send(&github.com/sarrufat/testmicro/hello.Pong{Stroke: req.Stroke}); err != nil {
			return err
		}
	}
}
