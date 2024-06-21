package core

import (
	"bytes"
	"fmt"
	"github.com/run-bigpig/xpvideo/internal/constant"
	"github.com/run-bigpig/xpvideo/internal/types"
	"io"
	"net/http"
)

type Spider struct {
	api    string
	header map[string]string
}

type Unmarshaler interface {
	Unmarshal([]byte) error
}

func NewSpider(api string) *Spider {
	return &Spider{
		api:    api,
		header: map[string]string{"User-Agent": constant.UserAgent, "Content-Type": constant.ContentType},
	}
}

// Class 获取分类
func (s *Spider) Class() (*types.VodListResponse, error) {
	var (
		response types.VodListResponse
	)
	err := s.send(http.MethodGet, s.api, nil, &response)
	if err != nil {
		return nil, err
	}
	return &response, err
}

// List 获取列表
func (s *Spider) List(params *types.VodListRequest) (*types.VodListResponse, error) {
	var (
		response types.VodListResponse
	)
	err := s.send(http.MethodGet, s.api, []byte(fmt.Sprintf("ac=%s&pg=%d&pagesize=%d&t=%d", "list", params.Page, params.PageSize, params.TypeId)), &response)
	if err != nil {
		return nil, err
	}
	return &response, err
}

// DetailList 获取详情列表
func (s *Spider) DetailList(params *types.VodListRequest) (*types.VodDetailResponse, error) {
	var (
		response types.VodDetailResponse
	)
	if params.Page == 0 {
		params.Page = 1
	}
	if params.PageSize == 0 {
		params.PageSize = 24
	}
	err := s.send(http.MethodGet, s.api, []byte(fmt.Sprintf("ac=%s&pg=%d&pagesize=%d&t=%d", "videolist", params.Page, params.PageSize, params.TypeId)), &response)
	if err != nil {
		return nil, err
	}
	return &response, err
}

// Detail 获取详情
func (s *Spider) Detail(params *types.VodDetailRequest) (*types.VodDetailResponse, error) {
	var (
		response types.VodDetailResponse
	)
	err := s.send(http.MethodGet, s.api, []byte(fmt.Sprintf("ac=%s&ids=%s", "videolist", params.Ids)), &response)
	if err != nil {
		return nil, err
	}
	return &response, err
}

// Search 搜索
func (s *Spider) Search(params *types.VodSearchRequest) (*types.VodDetailResponse, error) {
	var (
		response types.VodDetailResponse
	)
	err := s.send(http.MethodGet, s.api, []byte(fmt.Sprintf("ac=%s&wd=%s&pg=%d&pagesize=%d", "videolist", params.Keyword, params.Page, params.PageSize)), &response)
	if err != nil {
		return nil, err
	}
	return &response, err
}

func (s *Spider) send(method string, url string, params []byte, response Unmarshaler) error {
	var (
		client = &http.Client{}
		buf    = new(bytes.Buffer)
	)
	if method == http.MethodGet {
		url = fmt.Sprintf("%s?%s", url, params)
	}
	if method == http.MethodPost {
		buf = bytes.NewBuffer(params)
	}
	req, err := http.NewRequest(method, url, buf)
	if err != nil {
		return err
	}
	for k, v := range s.header {
		req.Header.Set(k, v)
	}
	res, err := client.Do(req)
	if err != nil {
		return err
	}
	defer res.Body.Close()
	body, err := io.ReadAll(res.Body)
	if err != nil {
		return err
	}
	return response.Unmarshal(body)
}
