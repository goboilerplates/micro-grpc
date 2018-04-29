package request

import (
	"context"
	"log"

	"github.com/goboilerplates/micro-rpc/example/goclient/util"
	"github.com/goboilerplates/micro-rpc/proto"
	"google.golang.org/grpc"
)

// GetSampleRequester is the GetSampleRequester interface.
type GetSampleRequester interface {
	All() bool
	AllByName() bool

	ServerStreamAll() bool
	ServerStreamAllByName() bool

	ClientStreamAll() bool
	ClientStreamAllByName() bool

	StreamAll() bool
	StreamAllByName() bool
}

// GetSampleRequesterImpl is the implementation of GetSampleRequester interface.
type GetSampleRequesterImpl struct {
	Context context.Context
	Conn    *grpc.ClientConn
}

// All requests all samples via unary rpc.
func (requester GetSampleRequesterImpl) All() bool {
	client := proto.NewGetSamplesClient(requester.Conn)
	result, err := client.All(requester.Context, &proto.CommonRequest{
		Name:      "All",
		Timestamp: util.GetCurrentTimeStamp(),
	})
	if err != nil {
		return false
	}
	log.Println("All with result: ", result.Samples)
	return true
}

// AllByName requests all samples by name keyword via unary rpc.
func (requester GetSampleRequesterImpl) AllByName() bool {
	client := proto.NewGetSamplesClient(requester.Conn)
	result, err := client.AllByName(requester.Context, &proto.KeywordRequest{
		Name:      "AllByName",
		Keyword:   "Ka",
		Timestamp: util.GetCurrentTimeStamp(),
	})
	if err != nil {
		return false
	}
	log.Println("AllByName with result: ", result.Samples)
	return true
}

// ServerStreamAll requests all samples via server-side stream rpc.
func (requester GetSampleRequesterImpl) ServerStreamAll() bool {
	client := proto.NewGetSamplesClient(requester.Conn)
	stream, err := client.ServerStreamAll(requester.Context, &proto.CommonRequest{
		Name:      "ServerStreamAll",
		Timestamp: util.GetCurrentTimeStamp(),
	})
	if err != nil {
		return false
	}

	result, err := stream.Recv()
	if err != nil {
		return false
	}

	log.Println("ServerStreamAll with result: ", result.Samples)
	return true
}

// ServerStreamAllByName requests all samples by name keyword via server-side stream rpc.
func (requester GetSampleRequesterImpl) ServerStreamAllByName() bool {
	client := proto.NewGetSamplesClient(requester.Conn)
	stream, err := client.ServerStreamAllByName(requester.Context, &proto.KeywordRequest{
		Name:      "ServerStreamAllByName",
		Keyword:   "Ka",
		Timestamp: util.GetCurrentTimeStamp(),
	})
	if err != nil {
		return false
	}

	result, err := stream.Recv()
	if err != nil {
		return false
	}

	log.Println("ServerStreamAllByName with result: ", result.Samples)
	return true
}

// ClientStreamAll requests all samples via client-side stream rpc.
func (requester GetSampleRequesterImpl) ClientStreamAll() bool {
	client := proto.NewGetSamplesClient(requester.Conn)
	stream, err := client.ClientStreamAll(requester.Context)
	for i := 0; i < 6; i++ {
		commonRequest := &proto.CommonRequest{
			Name:      "ClientStreamAll",
			Timestamp: util.GetCurrentTimeStamp(),
		}
		if err := stream.Send(commonRequest); err != nil {
			log.Fatalf("%v.Send(%v) = %v", stream, commonRequest, err)
		}
	}
	result, err := stream.CloseAndRecv()
	if err != nil {
		return false
	}
	log.Println("ClientStreamAll with result: ", result.Samples)
	return true
}

// ClientStreamAllByName requests all samples by keyword name via client-side rpc.
func (requester GetSampleRequesterImpl) ClientStreamAllByName() bool {
	client := proto.NewGetSamplesClient(requester.Conn)
	stream, err := client.ClientStreamAllByName(requester.Context)
	for i := 0; i < 6; i++ {
		keywordRequest := &proto.KeywordRequest{
			Name:      "ClientStreamAllByName",
			Keyword:   "Ka",
			Timestamp: util.GetCurrentTimeStamp(),
		}
		if err := stream.Send(keywordRequest); err != nil {
			log.Fatalf("%v.Send(%v) = %v", stream, keywordRequest, err)
		}
	}
	result, err := stream.CloseAndRecv()
	if err != nil {
		return false
	}
	log.Println("ClientStreamAllByName with result: ", result.Samples)
	return true
}

// StreamAll requests all samples via bidirectional stream rpc.
func (requester GetSampleRequesterImpl) StreamAll() bool {
	client := proto.NewGetSamplesClient(requester.Conn)
	stream, err := client.StreamAll(requester.Context)
	for i := 0; i < 6; i++ {
		commonRequest := &proto.CommonRequest{
			Name:      "StreamAll",
			Timestamp: util.GetCurrentTimeStamp(),
		}
		if err := stream.Send(commonRequest); err != nil {
			log.Fatalf("%v.Send(%v) = %v", stream, commonRequest, err)
		}
	}
	err = stream.CloseSend()
	if err != nil {
		return false
	}

	result, err := stream.Recv()
	if err != nil {
		return false
	}
	log.Println("StreamAll with result: ", result.Samples)
	return true
}

// StreamAllByName requests all samples by name keyword via bidirectional stream rpc.
func (requester GetSampleRequesterImpl) StreamAllByName() bool {
	client := proto.NewGetSamplesClient(requester.Conn)
	stream, err := client.StreamAllByName(requester.Context)
	for i := 0; i < 6; i++ {
		keywordRequest := &proto.KeywordRequest{
			Name:      "StreamAllByName",
			Keyword:   "Ka",
			Timestamp: util.GetCurrentTimeStamp(),
		}
		if err := stream.Send(keywordRequest); err != nil {
			log.Fatalf("%v.Send(%v) = %v", stream, keywordRequest, err)
		}
	}
	err = stream.CloseSend()
	if err != nil {
		return false
	}

	result, err := stream.Recv()
	if err != nil {
		return false
	}
	log.Println("StreamAllByName with result: ", result.Samples)
	return true
}
