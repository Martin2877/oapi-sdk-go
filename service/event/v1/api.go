// Package event code generated by oapi sdk gen
package larkevent

import (
	"context"
	"net/http"

	"github.com/larksuite/oapi-sdk-go/v3/core"
)

// 构建业务域服务实例
func NewService(config *larkcore.Config) *EventService {
	e := &EventService{config: config}
	e.OutboundIp = &outboundIp{service: e}
	return e
}

// 业务域服务定义
type EventService struct {
	config     *larkcore.Config
	OutboundIp *outboundIp
}

// 资源服务定义
type outboundIp struct {
	service *EventService
}

// 资源服务方法定义
func (o *outboundIp) List(ctx context.Context, req *ListOutboundIpReq, options ...larkcore.RequestOptionFunc) (*ListOutboundIpResp, error) {
	// 发起请求
	httpReq := req.HttpReq
	httpReq.ApiPath = "/open-apis/event/v1/outbound_ip"
	httpReq.HttpMethod = http.MethodGet
	httpReq.SupportedAccessTokenTypes = []larkcore.AccessTokenType{larkcore.AccessTokenTypeTenant}
	rawResp, err := larkcore.Request(ctx, req.HttpReq, o.service.config, options...)
	if err != nil {
		return nil, err
	}
	// 反序列响应结果
	resp := &ListOutboundIpResp{RawResponse: rawResp}
	err = rawResp.JSONUnmarshalBody(resp)
	if err != nil {
		return nil, err
	}
	return resp, err
}
func (o *outboundIp) ListByIterator(ctx context.Context, req *ListOutboundIpReq, options ...larkcore.RequestOptionFunc) (*ListOutboundIpIterator, error) {
	return &ListOutboundIpIterator{
		ctx:      ctx,
		req:      req,
		listFunc: o.List,
		options:  options,
		limit:    req.Limit}, nil
}
