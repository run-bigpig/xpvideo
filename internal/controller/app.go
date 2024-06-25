package controller

import (
	"context"
	"errors"
	"github.com/run-bigpig/xpvideo/internal/config"
	"github.com/run-bigpig/xpvideo/internal/core"
	"github.com/run-bigpig/xpvideo/internal/types"
	"github.com/run-bigpig/xpvideo/internal/utils"
	"github.com/spf13/viper"
	"strconv"
)

type AppController struct {
	ctx context.Context
}

func NewAppController(ctx context.Context) *AppController {
	return &AppController{
		ctx: ctx,
	}
}

func (ac *AppController) Class() *types.Response {
	var response []*types.ClassResponse
	s := core.NewSpider(config.GetDefaultSourceUrl())
	class, err := s.Class()
	if err != nil {
		return utils.Fail(400, err)
	}
	for _, item := range class.Class {
		response = append(response, &types.ClassResponse{
			Id:   item.TypeId,
			Pid:  item.TypePid,
			Name: item.TypeName,
		})
	}
	return utils.Success(response)
}

func (ac *AppController) List(req types.ListRequest) *types.Response {
	var (
		listItems []*types.ListItem
		response  types.ListResponse
	)
	s := core.NewSpider(config.GetDefaultSourceUrl())
	list, err := s.DetailList(&types.VodListRequest{
		Page:     req.Page,
		PageSize: req.PageSize,
		TypeId:   req.Class,
	})
	if err != nil {
		return utils.Fail(400, err)
	}
	for _, item := range list.List {
		listItems = append(listItems, &types.ListItem{
			Id:       item.VodId,
			Name:     item.VodName,
			Actor:    item.VodActor,
			Director: item.VodDirector,
			Img:      item.VodPic,
		})
	}
	response = types.ListResponse{
		List:  listItems,
		Total: list.Total,
	}
	return utils.Success(response)
}

func (ac *AppController) Detail(req types.DetailRequest) *types.Response {
	var (
		response types.DetailResponse
	)
	s := core.NewSpider(config.GetDefaultSourceUrl())
	detail, err := s.Detail(&types.VodDetailRequest{
		Ids: strconv.Itoa(req.Id),
	})
	if err != nil {
		return utils.Fail(400, err)
	}
	if len(detail.List) == 1 {
		response = types.DetailResponse{
			Id:       detail.List[0].VodId,
			Name:     detail.List[0].VodName,
			Actor:    detail.List[0].VodActor,
			Director: detail.List[0].VodDirector,
			Img:      detail.List[0].VodPic,
			Desc:     utils.RemoveHTMLTags(detail.List[0].VodContent),
			Play:     utils.DealPlayUrl(detail.List[0].VodPlayUrl),
		}
	}
	return utils.Success(response)
}

func (ac *AppController) Search(req types.SearchRequest) *types.Response {
	var (
		listItems []*types.ListItem
		response  types.ListResponse
	)
	s := core.NewSpider(config.GetDefaultSourceUrl())
	list, err := s.Search(&types.VodSearchRequest{
		Keyword:  req.Keyword,
		Page:     req.Page,
		PageSize: req.PageSize,
	})
	if err != nil {
		return utils.Fail(400, err)
	}
	for _, item := range list.List {
		listItems = append(listItems, &types.ListItem{
			Id:       item.VodId,
			Name:     item.VodName,
			Actor:    item.VodActor,
			Director: item.VodDirector,
			Img:      item.VodPic,
		})
	}
	response = types.ListResponse{
		List:  listItems,
		Total: list.Total,
	}
	return utils.Success(response)
}

func (ac *AppController) GetSourceList() *types.Response {
	var (
		list []*types.Source
	)
	err := viper.UnmarshalKey("source.list", &list)
	if err != nil {
		return utils.Fail(400, err)
	}
	return utils.Success(&types.SourceResponse{
		List: list,
	})
}

func (ac *AppController) Setting(req types.SettingRequest) *types.Response {
	if len(req.Sources) <= 0 {
		return utils.Fail(400, errors.New("sources is empty"))
	}
	for _, item := range req.Sources {
		if item.Url == "" || item.Name == "" {
			return utils.Fail(400, errors.New("source is empty"))
		}
	}
	viper.Set("source.list", req.Sources)
	err := viper.WriteConfig()
	if err != nil {
		return utils.Fail(400, err)
	}
	return utils.Success("ok")
}
