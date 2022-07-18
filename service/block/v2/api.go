// Package block code generated by oapi sdk gen
package larkblock

import (
	"context"
	"net/http"

	"github.com/larksuite/oapi-sdk-go/v3/core"
)

// 构建业务域服务实例
func NewService(config *larkcore.Config) *BlockService {
	b := &BlockService{config: config}
	b.Entity = &entity{service: b}
	b.Message = &message{service: b}
	return b
}

// 业务域服务定义
type BlockService struct {
	config  *larkcore.Config
	Entity  *entity
	Message *message
}

// 资源服务定义
type entity struct {
	service *BlockService
}
type message struct {
	service *BlockService
}

// 资源服务方法定义
func (e *entity) Create(ctx context.Context, req *CreateEntityReq, options ...larkcore.RequestOptionFunc) (*CreateEntityResp, error) {
	// 发起请求
	httpReq := req.HttpReq
	httpReq.ApiPath = "/open-apis/block/v2/entities"
	httpReq.HttpMethod = http.MethodPost
	httpReq.SupportedAccessTokenTypes = []larkcore.AccessTokenType{larkcore.AccessTokenTypeTenant}
	rawResp, err := larkcore.Request(ctx, req.HttpReq, e.service.config, options...)
	if err != nil {
		return nil, err
	}
	// 反序列响应结果
	resp := &CreateEntityResp{RawResponse: rawResp}
	err = rawResp.JSONUnmarshalBody(resp)
	if err != nil {
		return nil, err
	}
	return resp, err
}
func (e *entity) Update(ctx context.Context, req *UpdateEntityReq, options ...larkcore.RequestOptionFunc) (*UpdateEntityResp, error) {
	// 发起请求
	httpReq := req.HttpReq
	httpReq.ApiPath = "/open-apis/block/v2/entities/:block_id"
	httpReq.HttpMethod = http.MethodPut
	httpReq.SupportedAccessTokenTypes = []larkcore.AccessTokenType{larkcore.AccessTokenTypeTenant}
	rawResp, err := larkcore.Request(ctx, req.HttpReq, e.service.config, options...)
	if err != nil {
		return nil, err
	}
	// 反序列响应结果
	resp := &UpdateEntityResp{RawResponse: rawResp}
	err = rawResp.JSONUnmarshalBody(resp)
	if err != nil {
		return nil, err
	}
	return resp, err
}
func (m *message) Create(ctx context.Context, req *CreateMessageReq, options ...larkcore.RequestOptionFunc) (*CreateMessageResp, error) {
	// 发起请求
	httpReq := req.HttpReq
	httpReq.ApiPath = "/open-apis/block/v2/message"
	httpReq.HttpMethod = http.MethodPost
	httpReq.SupportedAccessTokenTypes = []larkcore.AccessTokenType{larkcore.AccessTokenTypeTenant}
	rawResp, err := larkcore.Request(ctx, req.HttpReq, m.service.config, options...)
	if err != nil {
		return nil, err
	}
	// 反序列响应结果
	resp := &CreateMessageResp{RawResponse: rawResp}
	err = rawResp.JSONUnmarshalBody(resp)
	if err != nil {
		return nil, err
	}
	return resp, err
}
