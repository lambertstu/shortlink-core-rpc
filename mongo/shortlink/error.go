package model

import (
	"errors"

	"github.com/zeromicro/go-zero/core/stores/mon"
)

var (
	ErrNotFound        = mon.ErrNotFound
	ErrInvalidObjectId = errors.New("invalid objectId")
	ErrShortUriExist   = errors.New("短链接重复")
	ErrInvalidRequest  = errors.New("更新参数错误")
)
