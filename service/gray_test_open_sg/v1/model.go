// Package gray_test_open_sg code generated by oapi sdk gen
package larkgray_test_open_sg

import (
	"fmt"

	"context"
	"errors"

	"github.com/larksuite/oapi-sdk-go/v3/core"
)

// 生成枚举值

// 生成数据类型

type Level struct {
	Level *string `json:"level,omitempty"`
	Body  *string `json:"body,omitempty"`
	Type  *string `json:"type,omitempty"`
}

// builder开始
type LevelBuilder struct {
	level     string
	levelFlag bool
	body      string
	bodyFlag  bool
	type_     string
	typeFlag  bool
}

func NewLevelBuilder() *LevelBuilder {
	builder := &LevelBuilder{}
	return builder
}

func (builder *LevelBuilder) Level(level string) *LevelBuilder {
	builder.level = level
	builder.levelFlag = true
	return builder
}
func (builder *LevelBuilder) Body(body string) *LevelBuilder {
	builder.body = body
	builder.bodyFlag = true
	return builder
}
func (builder *LevelBuilder) Type(type_ string) *LevelBuilder {
	builder.type_ = type_
	builder.typeFlag = true
	return builder
}

func (builder *LevelBuilder) Build() *Level {
	req := &Level{}
	if builder.levelFlag {
		req.Level = &builder.level

	}
	if builder.bodyFlag {
		req.Body = &builder.body

	}
	if builder.typeFlag {
		req.Type = &builder.type_

	}
	return req
}

// builder结束

type Moto struct {
	MotoId   *string `json:"moto_id,omitempty"`
	Id       *string `json:"id,omitempty"`
	UserName *string `json:"user_name,omitempty"`
	Type     *string `json:"type,omitempty"`
}

// builder开始
type MotoBuilder struct {
	motoId       string
	motoIdFlag   bool
	id           string
	idFlag       bool
	userName     string
	userNameFlag bool
	type_        string
	typeFlag     bool
}

func NewMotoBuilder() *MotoBuilder {
	builder := &MotoBuilder{}
	return builder
}

func (builder *MotoBuilder) MotoId(motoId string) *MotoBuilder {
	builder.motoId = motoId
	builder.motoIdFlag = true
	return builder
}
func (builder *MotoBuilder) Id(id string) *MotoBuilder {
	builder.id = id
	builder.idFlag = true
	return builder
}
func (builder *MotoBuilder) UserName(userName string) *MotoBuilder {
	builder.userName = userName
	builder.userNameFlag = true
	return builder
}
func (builder *MotoBuilder) Type(type_ string) *MotoBuilder {
	builder.type_ = type_
	builder.typeFlag = true
	return builder
}

func (builder *MotoBuilder) Build() *Moto {
	req := &Moto{}
	if builder.motoIdFlag {
		req.MotoId = &builder.motoId

	}
	if builder.idFlag {
		req.Id = &builder.id

	}
	if builder.userNameFlag {
		req.UserName = &builder.userName

	}
	if builder.typeFlag {
		req.Type = &builder.type_

	}
	return req
}

// builder结束

// 生成请求和响应结果类型，以及请求对象的Builder构造器

// 1.4 生成请求的builder结构体
type CreateMotoReqBuilder struct {
	*larkcore.HttpReq
	level *Level
}

// 生成请求的New构造器
func NewCreateMotoReqBuilder() *CreateMotoReqBuilder {
	builder := &CreateMotoReqBuilder{}
	builder.HttpReq = &larkcore.HttpReq{
		PathParams:  larkcore.PathParams{},
		QueryParams: larkcore.QueryParams{},
	}
	return builder
}

// 1.5 生成请求的builder属性方法
func (builder *CreateMotoReqBuilder) DepartmentIdType(departmentIdType string) *CreateMotoReqBuilder {
	builder.QueryParams.Set("department_id_type", fmt.Sprint(departmentIdType))
	return builder
}
func (builder *CreateMotoReqBuilder) Level(level *Level) *CreateMotoReqBuilder {
	builder.level = level
	return builder
}

// 1.5 生成请求的builder的build方法
func (builder *CreateMotoReqBuilder) Build() *CreateMotoReq {
	req := &CreateMotoReq{}
	req.HttpReq = &larkcore.HttpReq{}
	req.HttpReq.QueryParams = builder.QueryParams
	req.HttpReq.Body = builder.level
	return req
}

type CreateMotoReq struct {
	*larkcore.HttpReq
	Level *Level `body:""`
}

type CreateMotoRespData struct {
	Moto *Moto `json:"moto,omitempty"`
}

type CreateMotoResp struct {
	*larkcore.RawResponse `json:"-"`
	larkcore.CodeError
	Data *CreateMotoRespData `json:"data"`
}

func (resp *CreateMotoResp) Success() bool {
	return resp.Code == 0
}

// 1.4 生成请求的builder结构体
type GetMotoReqBuilder struct {
	*larkcore.HttpReq
}

// 生成请求的New构造器
func NewGetMotoReqBuilder() *GetMotoReqBuilder {
	builder := &GetMotoReqBuilder{}
	builder.HttpReq = &larkcore.HttpReq{
		PathParams:  larkcore.PathParams{},
		QueryParams: larkcore.QueryParams{},
	}
	return builder
}

// 1.5 生成请求的builder属性方法
func (builder *GetMotoReqBuilder) MotoId(motoId string) *GetMotoReqBuilder {
	builder.PathParams.Set("moto_id", fmt.Sprint(motoId))
	return builder
}
func (builder *GetMotoReqBuilder) BodyLevel(bodyLevel string) *GetMotoReqBuilder {
	builder.QueryParams.Set("body_level", fmt.Sprint(bodyLevel))
	return builder
}

// 1.5 生成请求的builder的build方法
func (builder *GetMotoReqBuilder) Build() *GetMotoReq {
	req := &GetMotoReq{}
	req.HttpReq = &larkcore.HttpReq{}
	req.HttpReq.PathParams = builder.PathParams
	req.HttpReq.QueryParams = builder.QueryParams
	return req
}

type GetMotoReq struct {
	*larkcore.HttpReq
}

type GetMotoRespData struct {
	Moto *Moto `json:"moto,omitempty"`
}

type GetMotoResp struct {
	*larkcore.RawResponse `json:"-"`
	larkcore.CodeError
	Data *GetMotoRespData `json:"data"`
}

func (resp *GetMotoResp) Success() bool {
	return resp.Code == 0
}

// 1.4 生成请求的builder结构体
type ListMotoReqBuilder struct {
	*larkcore.HttpReq
	limit int
}

// 生成请求的New构造器
func NewListMotoReqBuilder() *ListMotoReqBuilder {
	builder := &ListMotoReqBuilder{}
	builder.HttpReq = &larkcore.HttpReq{
		PathParams:  larkcore.PathParams{},
		QueryParams: larkcore.QueryParams{},
	}
	return builder
}

// 1.5 生成请求的builder属性方法
func (builder *ListMotoReqBuilder) Limit(limit int) *ListMotoReqBuilder {
	builder.limit = limit
	return builder
}
func (builder *ListMotoReqBuilder) PageSize(pageSize int) *ListMotoReqBuilder {
	builder.QueryParams.Set("page_size", fmt.Sprint(pageSize))
	return builder
}
func (builder *ListMotoReqBuilder) PageToken(pageToken string) *ListMotoReqBuilder {
	builder.QueryParams.Set("page_token", fmt.Sprint(pageToken))
	return builder
}
func (builder *ListMotoReqBuilder) Level(level int) *ListMotoReqBuilder {
	builder.QueryParams.Set("level", fmt.Sprint(level))
	return builder
}

// 1.5 生成请求的builder的build方法
func (builder *ListMotoReqBuilder) Build() *ListMotoReq {
	req := &ListMotoReq{}
	req.HttpReq = &larkcore.HttpReq{}
	req.Limit = builder.limit
	req.HttpReq.QueryParams = builder.QueryParams
	return req
}

type ListMotoReq struct {
	*larkcore.HttpReq
	Limit int
}

type ListMotoRespData struct {
	Items     []string `json:"items,omitempty"`
	PageToken *string  `json:"page_token,omitempty"`
	HasMore   *bool    `json:"has_more,omitempty"`
}

type ListMotoResp struct {
	*larkcore.RawResponse `json:"-"`
	larkcore.CodeError
	Data *ListMotoRespData `json:"data"`
}

func (resp *ListMotoResp) Success() bool {
	return resp.Code == 0
}

// 生成消息事件结构体

// 生成请求的builder构造器
// 1.1 生成body的builder结构体
type ListMotoIterator struct {
	nextPageToken *string
	items         []string
	index         int
	limit         int
	ctx           context.Context
	req           *ListMotoReq
	listFunc      func(ctx context.Context, req *ListMotoReq, options ...larkcore.RequestOptionFunc) (*ListMotoResp, error)
	options       []larkcore.RequestOptionFunc
	curlNum       int
}

func (iterator *ListMotoIterator) Next() (bool, string, error) {
	// 达到最大量，则返回
	if iterator.limit > 0 && iterator.curlNum >= iterator.limit {
		return false, "", nil
	}

	// 为0则拉取数据
	if iterator.index == 0 || iterator.index >= len(iterator.items) {
		if iterator.index != 0 && iterator.nextPageToken == nil {
			return false, "", nil
		}
		if iterator.nextPageToken != nil {
			iterator.req.QueryParams.Set("page_token", *iterator.nextPageToken)
		}
		resp, err := iterator.listFunc(iterator.ctx, iterator.req, iterator.options...)
		if err != nil {
			return false, "", err
		}

		if resp.Code != 0 {
			return false, "", errors.New(fmt.Sprintf("Code:%d,Msg:%s", resp.Code, resp.Msg))
		}

		if len(resp.Data.Items) == 0 {
			return false, "", nil
		}

		iterator.nextPageToken = resp.Data.PageToken
		iterator.items = resp.Data.Items
		iterator.index = 0
	}

	block := iterator.items[iterator.index]
	iterator.index++
	iterator.curlNum++
	return true, block, nil
}

func (iterator *ListMotoIterator) NextPageToken() *string {
	return iterator.nextPageToken
}
