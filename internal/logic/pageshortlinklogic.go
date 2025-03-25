package logic

import (
	"context"
	"strconv"

	model "github.com/lambertstu/shortlink-core-rpc/mongo/shortlink"
	"github.com/lambertstu/shortlink-core-rpc/pkg/constant"
	"go.mongodb.org/mongo-driver/bson"

	"github.com/lambertstu/shortlink-core-rpc/internal/svc"
	"github.com/lambertstu/shortlink-core-rpc/pb/shortlink"

	"github.com/zeromicro/go-zero/core/logx"
)

type PageShortLinkLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewPageShortLinkLogic(ctx context.Context, svcCtx *svc.ServiceContext) *PageShortLinkLogic {
	return &PageShortLinkLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *PageShortLinkLogic) PageShortLink(in *shortlink.ShortLinkPageRequest) (*shortlink.ShortLinkPageResponse, error) {
	var list []model.Shortlink
	filter := bson.M{
		"gid":        in.GetGid(),
		"deleteFlag": constant.ENABLE_FLAG,
	}
	total, err := l.svcCtx.ShortlinkModel.Pagination(l.ctx, in.GetPage(), in.GetSize(), in.GetOrderTag(), filter, "full_short_url", &list)
	if err != nil {
		return nil, err
	}

	var responseList []*shortlink.ShortLinkPageData
	for _, item := range list {
		responseList = append(responseList, &shortlink.ShortLinkPageData{
			ShortUri:     item.ShortUri,
			FullShortUrl: item.FullShortUrl,
			OriginUrl:    item.OriginUrl,
			Gid:          item.Gid,
			Describe:     item.Describe,
			Favicon:      item.Favicon,
			ClickNum:     strconv.Itoa(item.ClickNum),
			TotalPv:      int32(item.TotalPv),
			TodayPv:      int32(item.TodayPv),
			TotalUv:      int32(item.TotalUv),
			TodayUv:      int32(item.TodayUv),
			TotalUip:     int32(item.TotalUip),
			TodayUip:     int32(item.TodayUip),
			CreateTime:   item.CreateAt.Format("2006-01-02 15:04:05"),
			UpdateTime:   item.UpdateAt.Format("2006-01-02 15:04:05"),
		})
	}

	return &shortlink.ShortLinkPageResponse{
		List:    responseList,
		Page:    in.GetPage(),
		MaxPage: (total + in.GetSize() - 1) / in.GetSize(),
		Total:   total,
	}, nil
}
