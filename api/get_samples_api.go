package api

import (
	"context"
	"fmt"
	"io"

	"github.com/goboilerplates/core"

	"github.com/goboilerplates/micro-rpc/proto"
	"github.com/goboilerplates/micro-rpc/util"
)

// GetSamplesAPI .
type GetSamplesAPI struct {
	Interactor core.GetSamplesInteractor
}

// All gets all samples via unary RPC.
func (api *GetSamplesAPI) All(ctx context.Context, commonRequest *proto.CommonRequest) (*proto.GetSamplesReply, error) {
	fmt.Println(commonRequest.Name, " with TimeStamp: ", commonRequest.Timestamp)
	result, err := api.Interactor.All()
	samples := util.ConvertCoreSampleArrayToReplySampleArray(result)
	return &proto.GetSamplesReply{Samples: samples}, err
}

// AllByName All gets all samples by name keyword via unary RPC.
func (api *GetSamplesAPI) AllByName(ctx context.Context, keywordRequest *proto.KeywordRequest) (*proto.GetSamplesReply, error) {
	fmt.Println(keywordRequest.Keyword, " with TimeStamp: ", keywordRequest.Timestamp)
	result, err := api.Interactor.AllByName(keywordRequest.Keyword)
	samples := util.ConvertCoreSampleArrayToReplySampleArray(result)
	return &proto.GetSamplesReply{Samples: samples}, err
}

// ServerStreamAll gets all samples via server-side stream RPC.
func (api *GetSamplesAPI) ServerStreamAll(commonRequest *proto.CommonRequest, stream proto.GetSamples_ServerStreamAllServer) error {
	fmt.Println(commonRequest.Name, " with TimeStamp: ", commonRequest.Timestamp)
	result, err := api.Interactor.All()
	if err != nil {
		return err
	}

	samples := util.ConvertCoreSampleArrayToReplySampleArray(result)
	err = stream.Send(&proto.GetSamplesReply{Samples: samples})
	return err
}

// ServerStreamAllByName gets all samples by name keyword via server-side stream RPC.
func (api *GetSamplesAPI) ServerStreamAllByName(keywordRequest *proto.KeywordRequest, stream proto.GetSamples_ServerStreamAllByNameServer) error {
	fmt.Println(keywordRequest.Keyword, " with TimeStamp: ", keywordRequest.Timestamp)
	result, err := api.Interactor.AllByName(keywordRequest.Keyword)
	if err != nil {
		return err
	}

	samples := util.ConvertCoreSampleArrayToReplySampleArray(result)
	err = stream.Send(&proto.GetSamplesReply{Samples: samples})
	return err
}

// ClientStreamAll gets all samples via client-side stream RPC.
func (api *GetSamplesAPI) ClientStreamAll(stream proto.GetSamples_ClientStreamAllServer) error {
	for {
		commonRequest, err := stream.Recv()
		if err == io.EOF {
			fmt.Println(commonRequest.Name, " with TimeStamp: ", commonRequest.Timestamp)
			result, err := api.Interactor.All()
			samples := util.ConvertCoreSampleArrayToReplySampleArray(result)
			err = stream.SendAndClose(&proto.GetSamplesReply{Samples: samples})
			return err
		}
		if err != nil {
			return err
		}
	}
}

// ClientStreamAllByName gets all samples by name keyword via client-side stream RPC.
func (api *GetSamplesAPI) ClientStreamAllByName(stream proto.GetSamples_ClientStreamAllByNameServer) error {
	for {
		keywordRequest, err := stream.Recv()
		if err == io.EOF {
			fmt.Println(keywordRequest.Keyword, " with TimeStamp: ", keywordRequest.Timestamp)
			result, err := api.Interactor.AllByName(keywordRequest.Keyword)
			samples := util.ConvertCoreSampleArrayToReplySampleArray(result)
			err = stream.SendAndClose(&proto.GetSamplesReply{Samples: samples})
			return err
		}
		if err != nil {
			return err
		}
	}
}

// StreamAll gets all samples via bidirectional stream RPC.
func (api *GetSamplesAPI) StreamAll(stream proto.GetSamples_StreamAllServer) error {
	for {
		commonRequest, err := stream.Recv()
		if err == io.EOF {
			return nil
		}
		if err != nil {
			return err
		}
		fmt.Println(commonRequest.Name, " with TimeStamp: ", commonRequest.Timestamp)
		result, err := api.Interactor.All()
		samples := util.ConvertCoreSampleArrayToReplySampleArray(result)
		err = stream.Send(&proto.GetSamplesReply{Samples: samples})
	}
}

// StreamAllByName gets all samples by name keyword via bidirectional stream RPC.
func (api *GetSamplesAPI) StreamAllByName(stream proto.GetSamples_StreamAllByNameServer) error {
	for {
		keywordRequest, err := stream.Recv()
		if err == io.EOF {
			return nil
		}
		if err != nil {
			return err
		}
		fmt.Println(keywordRequest.Keyword, " with TimeStamp: ", keywordRequest.Timestamp)
		result, err := api.Interactor.AllByName(keywordRequest.Keyword)
		samples := util.ConvertCoreSampleArrayToReplySampleArray(result)
		stream.Send(&proto.GetSamplesReply{Samples: samples})
	}
}
