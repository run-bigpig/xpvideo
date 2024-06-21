package utils

import (
	"github.com/run-bigpig/xpvideo/internal/constant"
	"github.com/run-bigpig/xpvideo/internal/types"
)

func Success(data interface{}) *types.Response {
	return &types.Response{
		Code: 0,
		Data: data,
		Msg:  constant.Success,
	}
}

func Fail(code int, error error) *types.Response {
	return &types.Response{
		Code: code,
		Data: nil,
		Msg:  error.Error(),
	}
}
