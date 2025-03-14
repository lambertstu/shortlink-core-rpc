package model

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Shortlink struct {
	ID           primitive.ObjectID `bson:"_id,omitempty" json:"id,omitempty"`
	Domain       string             `bson:"domain" json:"domain"`
	OriginUrl    string             `bson:"origin_url" json:"originUrl"`
	Gid          string             `bson:"gid" json:"gid"`
	Describe     string             `bson:"describe" json:"describe"`
	FullShortUrl string             `bson:"full_short_url" json:"fullShortUrl"`
	ShortUri     string             `bson:"short_uri" json:"shortUri"`
	Favicon      string             `bson:"favicon" json:"favicon"`
	ClickNum     int                `bson:"click_num" json:"clickNum"`
	TotalPv      int                `bson:"total_pv" json:"totalPv"`
	TotalUv      int                `bson:"total_uv" json:"totalUv"`
	TotalUip     int                `bson:"total_uip" json:"totalUip"`
	TodayPv      int                `bson:"today_pv" json:"todayPv,omitempty"`
	TodayUv      int                `bson:"today_uv" json:"todayUv,omitempty"`
	TodayUip     int                `bson:"today_uip" json:"todayUip,omitempty"`
	DeleteFlag   int                `bson:"deleteFlag" json:"deleteFlag"`
	UpdateAt     time.Time          `bson:"updateAt,omitempty" json:"updateAt,omitempty"`
	CreateAt     time.Time          `bson:"createAt,omitempty" json:"createAt,omitempty"`
}
