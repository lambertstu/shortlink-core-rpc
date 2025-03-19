package model

import (
	"context"
	"errors"
	"github.com/lambertstu/shortlink-core-rpc/pkg/constant"
	"github.com/zeromicro/go-zero/core/stores/mon"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo/options"
	"time"
)

const (
	DB         = "shortlink"
	Collection = "shortlink"
)

var _ ShortlinkModel = (*customShortlinkModel)(nil)

type (
	// ShortlinkModel is an interface to be customized, add more methods here,
	// and implement the added methods in customShortlinkModel.
	ShortlinkModel interface {
		shortlinkModel
		FindOneByShortUrl(ctx context.Context, shortLink string) (*Shortlink, error)
		InsertOneShortlink(ctx context.Context, data *Shortlink) error
		UpdateShortLinkInfo(ctx context.Context, data *Shortlink) error
		Pagination(ctx context.Context, page, size, sortOrder int64, filter bson.M, sortField string, v any) (int64, error)
	}

	customShortlinkModel struct {
		*defaultShortlinkModel
	}
)

func (c *customShortlinkModel) Pagination(ctx context.Context, page, size, sortOrder int64, filter bson.M, sortField string, v any) (int64, error) {
	if page < 1 {
		page = 1
	}
	if size < 1 {
		size = 10 // 默认每页10条
	} else if size > 100 {
		size = 100 // 限制最大查询数
	}

	skip := (page - 1) * size

	opts := options.Find().
		SetSkip(skip).
		SetLimit(size)

	if sortField == "" {
		sortField = "full_short_url"
	}
	opts.SetSort(bson.D{{sortField, sortOrder}})

	err := c.conn.Find(ctx, v, filter, opts)
	if err != nil {
		return -1, err
	}

	sum, err := c.conn.CountDocuments(ctx, filter)
	if err != nil {
		return -1, err
	}

	return sum, nil
}

func (c *customShortlinkModel) UpdateShortLinkInfo(ctx context.Context, data *Shortlink) error {
	if data == nil {
		return ErrInvalidRequest
	}

	filter := bson.M{
		"origin_url": data.OriginUrl,
		"short_uri":  data.ShortUri,
		"deleteFlag": constant.ENABLE_FLAG,
	}

	updateFields := bson.M{}

	// 动态添加需要更新的字段
	if data.Describe != "" {
		updateFields["describe"] = data.Describe
	}
	if data.Favicon != "" {
		updateFields["favicon"] = data.Favicon
	}
	if data.ClickNum != 0 {
		updateFields["click_num"] = data.ClickNum
	}
	if data.TotalPv != 0 {
		updateFields["total_pv"] = data.TotalPv
	}
	if data.TotalUv != 0 {
		updateFields["total_uv"] = data.TotalUv
	}
	if data.TotalUip != 0 {
		updateFields["total_uip"] = data.TotalUip
	}
	if data.TodayPv != 0 {
		updateFields["today_pv"] = data.TodayPv
	}
	if data.TodayUv != 0 {
		updateFields["today_uv"] = data.TodayUv
	}
	if data.TodayUip != 0 {
		updateFields["today_uip"] = data.TodayUip
	}
	if data.DeleteFlag != 0 {
		updateFields["deleteFlag"] = data.DeleteFlag
	}

	if len(updateFields) == 0 {
		return nil // 没有字段需要更新
	}

	update := bson.M{"$set": updateFields}
	updateOptions := options.Update().SetUpsert(false)

	result, err := c.conn.UpdateOne(ctx, filter, update, updateOptions)
	if err != nil {
		return err
	}

	if result.MatchedCount == 0 {
		return errors.New("短链接不存在")
	}

	return nil
}

func (c *customShortlinkModel) InsertOneShortlink(ctx context.Context, data *Shortlink) error {
	if data.ID.IsZero() {
		data.ID = primitive.NewObjectID()
		data.CreateAt = time.Now()
		data.UpdateAt = time.Now()
	}

	filter := bson.M{
		"short_uri":  data.ShortUri,
		"deleteFlag": constant.ENABLE_FLAG,
	}

	count, err := c.conn.CountDocuments(ctx, filter)
	if err != nil {
		return err
	}
	if count > 0 {
		return ErrShortUriExist
	}

	_, err = c.conn.InsertOne(ctx, data)
	if err != nil {
		return err
	}

	return nil
}

func (c *customShortlinkModel) FindOneByShortUrl(ctx context.Context, shortLink string) (*Shortlink, error) {
	filter := bson.M{
		"short_uri":  shortLink,
		"deleteFlag": constant.ENABLE_FLAG,
	}

	var shortlink Shortlink

	err := c.conn.FindOne(ctx, &shortlink, filter)
	if err != nil {
		return nil, err
	}

	return &shortlink, nil
}

// NewShortlinkModel returns a model for the mongo.
func NewShortlinkModel(url, db, collection string) ShortlinkModel {
	conn := mon.MustNewModel(url, db, collection)
	return &customShortlinkModel{
		defaultShortlinkModel: newDefaultShortlinkModel(conn),
	}
}
