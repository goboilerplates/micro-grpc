package api

import (
	"context"
	"fmt"

	"github.com/goboilerplates/core"

	"github.com/goboilerplates/micro-rpc/proto"
	"github.com/goboilerplates/micro-rpc/util"
)

// GetSamplesAPI .
type GetSamplesAPI struct {
	Interactor core.GetSamplesInteractor
}

// All .
func (api *GetSamplesAPI) All(ctx context.Context, commonRequest *proto.CommonRequest) (*proto.GetSamplesReply, error) {
	fmt.Println(commonRequest.Name, " with TimeStamp: ", commonRequest.Timestamp)
	result, err := api.Interactor.All()
	samples := util.ConvertCoreSampleArrayToReplySampleArray(result)
	return &proto.GetSamplesReply{Samples: samples}, err
}

// AllByName .
func (api *GetSamplesAPI) AllByName(ctx context.Context, keywordRequest *proto.KeywordRequest) (*proto.GetSamplesReply, error) {
	fmt.Println(keywordRequest.Keyword, " with TimeStamp: ", keywordRequest.Timestamp)
	result, err := api.Interactor.AllByName(keywordRequest.Keyword)
	samples := util.ConvertCoreSampleArrayToReplySampleArray(result)
	return &proto.GetSamplesReply{Samples: samples}, err
}
