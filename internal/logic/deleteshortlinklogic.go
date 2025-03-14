package logic

import (
	"context"
	model "github.com/lambertstu/shortlink-core-rpc/mongo/shortlink"
	"github.com/lambertstu/shortlink-core-rpc/pkg/constant"
	"time"

	"github.com/lambertstu/shortlink-core-rpc/internal/svc"
	"github.com/lambertstu/shortlink-core-rpc/pb/shortlink"

	"github.com/zeromicro/go-zero/core/logx"
)

type DeleteShortLinkLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewDeleteShortLinkLogic(ctx context.Context, svcCtx *svc.ServiceContext) *DeleteShortLinkLogic {
	return &DeleteShortLinkLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *DeleteShortLinkLogic) DeleteShortLink(in *shortlink.ShortLinkDeleteRequest) (*shortlink.ShortLinkDeleteResponse, error) {
	err := l.svcCtx.ShortlinkModel.UpdateShortLinkInfo(l.ctx, &model.Shortlink{
		OriginUrl:  in.GetOriginUrl(),
		ShortUri:   in.GetShortUri(),
		DeleteFlag: constant.DELETE_FLAG,
		UpdateAt:   time.Now(),
	})
	if err != nil {
		return &shortlink.ShortLinkDeleteResponse{
			Success: false,
		}, err
	}

	return &shortlink.ShortLinkDeleteResponse{
		Success: true,
	}, nil
}
