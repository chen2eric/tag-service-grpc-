package server

import (
	"context"
	"encoding/json"

	"github.com/chen2eric/tag-service/pkg/bapi"
	"github.com/chen2eric/tag-service/pkg/bapi/errcode"
	pb "github.com/chen2eric/tag-service/proto"
)

type TagServer struct{}

func NewTagServer() *TagServer {
	return &TagServer{}
}

func (t *TagServer) GetTagList(ctx context.Context, r *pb.GetTagListRequest) (*pb.GetTagListReply, error) {
	api := bapi.NewAPI("http://127.0.0.1:8000")
	body, err := api.GetTagList(ctx, r.GetName())
	if err != nil {
		return nil, errcode.TogRPCError(errcode.ErrorGetTagListFail)
	}
	tagListReply := pb.GetTagListReply{}
	err = json.Unmarshal(body, &tagListReply)
	if err != nil {
		return nil, errcode.TogRPCError(errcode.Fail)
	}
	return &tagListReply, nil
}
