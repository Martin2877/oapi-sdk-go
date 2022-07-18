// Package human_authentication code generated by oapi sdk gen
package larkhuman_authentication

import (
	"context"
	"net/http"

	"github.com/larksuite/oapi-sdk-go/v3/core"
)

// 构建业务域服务实例
func NewService(config *larkcore.Config) *HumanAuthenticationService {
	h := &HumanAuthenticationService{config: config}
	h.Identity = &identity{service: h}
	return h
}

// 业务域服务定义
type HumanAuthenticationService struct {
	config   *larkcore.Config
	Identity *identity
}

// 资源服务定义
type identity struct {
	service *HumanAuthenticationService
}

// 资源服务方法定义
func (i *identity) Create(ctx context.Context, req *CreateIdentityReq, options ...larkcore.RequestOptionFunc) (*CreateIdentityResp, error) {
	// 发起请求
	httpReq := req.HttpReq
	httpReq.ApiPath = "/open-apis/human_authentication/v1/identities"
	httpReq.HttpMethod = http.MethodPost
	httpReq.SupportedAccessTokenTypes = []larkcore.AccessTokenType{larkcore.AccessTokenTypeTenant}
	rawResp, err := larkcore.Request(ctx, req.HttpReq, i.service.config, options...)
	if err != nil {
		return nil, err
	}
	// 反序列响应结果
	resp := &CreateIdentityResp{RawResponse: rawResp}
	err = rawResp.JSONUnmarshalBody(resp)
	if err != nil {
		return nil, err
	}
	return resp, err
}
