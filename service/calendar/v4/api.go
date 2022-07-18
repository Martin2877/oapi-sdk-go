// Package calendar code generated by oapi sdk gen
package larkcalendar

import (
	"context"
	"net/http"

	"github.com/larksuite/oapi-sdk-go/v3/core"
)

// 构建业务域服务实例
func NewService(config *larkcore.Config) *CalendarService {
	c := &CalendarService{config: config}
	c.Calendar = &calendar{service: c}
	c.CalendarAcl = &calendarAcl{service: c}
	c.CalendarEvent = &calendarEvent{service: c}
	c.CalendarEventAttendee = &calendarEventAttendee{service: c}
	c.CalendarEventAttendeeChatMember = &calendarEventAttendeeChatMember{service: c}
	c.ExchangeBinding = &exchangeBinding{service: c}
	c.Freebusy = &freebusy{service: c}
	c.Setting = &setting{service: c}
	c.TimeoffEvent = &timeoffEvent{service: c}
	return c
}

// 业务域服务定义
type CalendarService struct {
	config                          *larkcore.Config
	Calendar                        *calendar
	CalendarAcl                     *calendarAcl
	CalendarEvent                   *calendarEvent
	CalendarEventAttendee           *calendarEventAttendee
	CalendarEventAttendeeChatMember *calendarEventAttendeeChatMember
	ExchangeBinding                 *exchangeBinding
	Freebusy                        *freebusy
	Setting                         *setting
	TimeoffEvent                    *timeoffEvent
}

// 资源服务定义
type calendar struct {
	service *CalendarService
}
type calendarAcl struct {
	service *CalendarService
}
type calendarEvent struct {
	service *CalendarService
}
type calendarEventAttendee struct {
	service *CalendarService
}
type calendarEventAttendeeChatMember struct {
	service *CalendarService
}
type exchangeBinding struct {
	service *CalendarService
}
type freebusy struct {
	service *CalendarService
}
type setting struct {
	service *CalendarService
}
type timeoffEvent struct {
	service *CalendarService
}

// 资源服务方法定义
func (c *calendar) Create(ctx context.Context, req *CreateCalendarReq, options ...larkcore.RequestOptionFunc) (*CreateCalendarResp, error) {
	// 发起请求
	httpReq := req.HttpReq
	httpReq.ApiPath = "/open-apis/calendar/v4/calendars"
	httpReq.HttpMethod = http.MethodPost
	httpReq.SupportedAccessTokenTypes = []larkcore.AccessTokenType{larkcore.AccessTokenTypeTenant, larkcore.AccessTokenTypeUser}
	rawResp, err := larkcore.Request(ctx, req.HttpReq, c.service.config, options...)
	if err != nil {
		return nil, err
	}
	// 反序列响应结果
	resp := &CreateCalendarResp{RawResponse: rawResp}
	err = rawResp.JSONUnmarshalBody(resp)
	if err != nil {
		return nil, err
	}
	return resp, err
}
func (c *calendar) Delete(ctx context.Context, req *DeleteCalendarReq, options ...larkcore.RequestOptionFunc) (*DeleteCalendarResp, error) {
	// 发起请求
	httpReq := req.HttpReq
	httpReq.ApiPath = "/open-apis/calendar/v4/calendars/:calendar_id"
	httpReq.HttpMethod = http.MethodDelete
	httpReq.SupportedAccessTokenTypes = []larkcore.AccessTokenType{larkcore.AccessTokenTypeTenant, larkcore.AccessTokenTypeUser}
	rawResp, err := larkcore.Request(ctx, req.HttpReq, c.service.config, options...)
	if err != nil {
		return nil, err
	}
	// 反序列响应结果
	resp := &DeleteCalendarResp{RawResponse: rawResp}
	err = rawResp.JSONUnmarshalBody(resp)
	if err != nil {
		return nil, err
	}
	return resp, err
}
func (c *calendar) Get(ctx context.Context, req *GetCalendarReq, options ...larkcore.RequestOptionFunc) (*GetCalendarResp, error) {
	// 发起请求
	httpReq := req.HttpReq
	httpReq.ApiPath = "/open-apis/calendar/v4/calendars/:calendar_id"
	httpReq.HttpMethod = http.MethodGet
	httpReq.SupportedAccessTokenTypes = []larkcore.AccessTokenType{larkcore.AccessTokenTypeTenant, larkcore.AccessTokenTypeUser}
	rawResp, err := larkcore.Request(ctx, req.HttpReq, c.service.config, options...)
	if err != nil {
		return nil, err
	}
	// 反序列响应结果
	resp := &GetCalendarResp{RawResponse: rawResp}
	err = rawResp.JSONUnmarshalBody(resp)
	if err != nil {
		return nil, err
	}
	return resp, err
}
func (c *calendar) List(ctx context.Context, req *ListCalendarReq, options ...larkcore.RequestOptionFunc) (*ListCalendarResp, error) {
	// 发起请求
	httpReq := req.HttpReq
	httpReq.ApiPath = "/open-apis/calendar/v4/calendars"
	httpReq.HttpMethod = http.MethodGet
	httpReq.SupportedAccessTokenTypes = []larkcore.AccessTokenType{larkcore.AccessTokenTypeTenant, larkcore.AccessTokenTypeUser}
	rawResp, err := larkcore.Request(ctx, req.HttpReq, c.service.config, options...)
	if err != nil {
		return nil, err
	}
	// 反序列响应结果
	resp := &ListCalendarResp{RawResponse: rawResp}
	err = rawResp.JSONUnmarshalBody(resp)
	if err != nil {
		return nil, err
	}
	return resp, err
}
func (c *calendar) Patch(ctx context.Context, req *PatchCalendarReq, options ...larkcore.RequestOptionFunc) (*PatchCalendarResp, error) {
	// 发起请求
	httpReq := req.HttpReq
	httpReq.ApiPath = "/open-apis/calendar/v4/calendars/:calendar_id"
	httpReq.HttpMethod = http.MethodPatch
	httpReq.SupportedAccessTokenTypes = []larkcore.AccessTokenType{larkcore.AccessTokenTypeTenant, larkcore.AccessTokenTypeUser}
	rawResp, err := larkcore.Request(ctx, req.HttpReq, c.service.config, options...)
	if err != nil {
		return nil, err
	}
	// 反序列响应结果
	resp := &PatchCalendarResp{RawResponse: rawResp}
	err = rawResp.JSONUnmarshalBody(resp)
	if err != nil {
		return nil, err
	}
	return resp, err
}
func (c *calendar) Primary(ctx context.Context, req *PrimaryCalendarReq, options ...larkcore.RequestOptionFunc) (*PrimaryCalendarResp, error) {
	// 发起请求
	httpReq := req.HttpReq
	httpReq.ApiPath = "/open-apis/calendar/v4/calendars/primary"
	httpReq.HttpMethod = http.MethodPost
	httpReq.SupportedAccessTokenTypes = []larkcore.AccessTokenType{larkcore.AccessTokenTypeTenant, larkcore.AccessTokenTypeUser}
	rawResp, err := larkcore.Request(ctx, req.HttpReq, c.service.config, options...)
	if err != nil {
		return nil, err
	}
	// 反序列响应结果
	resp := &PrimaryCalendarResp{RawResponse: rawResp}
	err = rawResp.JSONUnmarshalBody(resp)
	if err != nil {
		return nil, err
	}
	return resp, err
}
func (c *calendar) Search(ctx context.Context, req *SearchCalendarReq, options ...larkcore.RequestOptionFunc) (*SearchCalendarResp, error) {
	// 发起请求
	httpReq := req.HttpReq
	httpReq.ApiPath = "/open-apis/calendar/v4/calendars/search"
	httpReq.HttpMethod = http.MethodPost
	httpReq.SupportedAccessTokenTypes = []larkcore.AccessTokenType{larkcore.AccessTokenTypeTenant}
	rawResp, err := larkcore.Request(ctx, req.HttpReq, c.service.config, options...)
	if err != nil {
		return nil, err
	}
	// 反序列响应结果
	resp := &SearchCalendarResp{RawResponse: rawResp}
	err = rawResp.JSONUnmarshalBody(resp)
	if err != nil {
		return nil, err
	}
	return resp, err
}
func (c *calendar) SearchByIterator(ctx context.Context, req *SearchCalendarReq, options ...larkcore.RequestOptionFunc) (*SearchCalendarIterator, error) {
	return &SearchCalendarIterator{
		ctx:      ctx,
		req:      req,
		listFunc: c.Search,
		options:  options,
		limit:    req.Limit}, nil
}
func (c *calendar) Subscribe(ctx context.Context, req *SubscribeCalendarReq, options ...larkcore.RequestOptionFunc) (*SubscribeCalendarResp, error) {
	// 发起请求
	httpReq := req.HttpReq
	httpReq.ApiPath = "/open-apis/calendar/v4/calendars/:calendar_id/subscribe"
	httpReq.HttpMethod = http.MethodPost
	httpReq.SupportedAccessTokenTypes = []larkcore.AccessTokenType{larkcore.AccessTokenTypeUser, larkcore.AccessTokenTypeTenant}
	rawResp, err := larkcore.Request(ctx, req.HttpReq, c.service.config, options...)
	if err != nil {
		return nil, err
	}
	// 反序列响应结果
	resp := &SubscribeCalendarResp{RawResponse: rawResp}
	err = rawResp.JSONUnmarshalBody(resp)
	if err != nil {
		return nil, err
	}
	return resp, err
}
func (c *calendar) Subscription(ctx context.Context, options ...larkcore.RequestOptionFunc) (*SubscriptionCalendarResp, error) {
	// 发起请求
	httpReq := &larkcore.HttpReq{
		PathParams:  larkcore.PathParams{},
		QueryParams: larkcore.QueryParams{},
	}
	httpReq.ApiPath = "/open-apis/calendar/v4/calendars/subscription"
	httpReq.HttpMethod = http.MethodPost
	httpReq.SupportedAccessTokenTypes = []larkcore.AccessTokenType{larkcore.AccessTokenTypeUser}
	rawResp, err := larkcore.Request(ctx, nil, c.service.config, options...)
	if err != nil {
		return nil, err
	}
	// 反序列响应结果
	resp := &SubscriptionCalendarResp{RawResponse: rawResp}
	err = rawResp.JSONUnmarshalBody(resp)
	if err != nil {
		return nil, err
	}
	return resp, err
}
func (c *calendar) Unsubscribe(ctx context.Context, req *UnsubscribeCalendarReq, options ...larkcore.RequestOptionFunc) (*UnsubscribeCalendarResp, error) {
	// 发起请求
	httpReq := req.HttpReq
	httpReq.ApiPath = "/open-apis/calendar/v4/calendars/:calendar_id/unsubscribe"
	httpReq.HttpMethod = http.MethodPost
	httpReq.SupportedAccessTokenTypes = []larkcore.AccessTokenType{larkcore.AccessTokenTypeUser, larkcore.AccessTokenTypeTenant}
	rawResp, err := larkcore.Request(ctx, req.HttpReq, c.service.config, options...)
	if err != nil {
		return nil, err
	}
	// 反序列响应结果
	resp := &UnsubscribeCalendarResp{RawResponse: rawResp}
	err = rawResp.JSONUnmarshalBody(resp)
	if err != nil {
		return nil, err
	}
	return resp, err
}
func (c *calendarAcl) Create(ctx context.Context, req *CreateCalendarAclReq, options ...larkcore.RequestOptionFunc) (*CreateCalendarAclResp, error) {
	// 发起请求
	httpReq := req.HttpReq
	httpReq.ApiPath = "/open-apis/calendar/v4/calendars/:calendar_id/acls"
	httpReq.HttpMethod = http.MethodPost
	httpReq.SupportedAccessTokenTypes = []larkcore.AccessTokenType{larkcore.AccessTokenTypeTenant, larkcore.AccessTokenTypeUser}
	rawResp, err := larkcore.Request(ctx, req.HttpReq, c.service.config, options...)
	if err != nil {
		return nil, err
	}
	// 反序列响应结果
	resp := &CreateCalendarAclResp{RawResponse: rawResp}
	err = rawResp.JSONUnmarshalBody(resp)
	if err != nil {
		return nil, err
	}
	return resp, err
}
func (c *calendarAcl) Delete(ctx context.Context, req *DeleteCalendarAclReq, options ...larkcore.RequestOptionFunc) (*DeleteCalendarAclResp, error) {
	// 发起请求
	httpReq := req.HttpReq
	httpReq.ApiPath = "/open-apis/calendar/v4/calendars/:calendar_id/acls/:acl_id"
	httpReq.HttpMethod = http.MethodDelete
	httpReq.SupportedAccessTokenTypes = []larkcore.AccessTokenType{larkcore.AccessTokenTypeTenant, larkcore.AccessTokenTypeUser}
	rawResp, err := larkcore.Request(ctx, req.HttpReq, c.service.config, options...)
	if err != nil {
		return nil, err
	}
	// 反序列响应结果
	resp := &DeleteCalendarAclResp{RawResponse: rawResp}
	err = rawResp.JSONUnmarshalBody(resp)
	if err != nil {
		return nil, err
	}
	return resp, err
}
func (c *calendarAcl) List(ctx context.Context, req *ListCalendarAclReq, options ...larkcore.RequestOptionFunc) (*ListCalendarAclResp, error) {
	// 发起请求
	httpReq := req.HttpReq
	httpReq.ApiPath = "/open-apis/calendar/v4/calendars/:calendar_id/acls"
	httpReq.HttpMethod = http.MethodGet
	httpReq.SupportedAccessTokenTypes = []larkcore.AccessTokenType{larkcore.AccessTokenTypeTenant, larkcore.AccessTokenTypeUser}
	rawResp, err := larkcore.Request(ctx, req.HttpReq, c.service.config, options...)
	if err != nil {
		return nil, err
	}
	// 反序列响应结果
	resp := &ListCalendarAclResp{RawResponse: rawResp}
	err = rawResp.JSONUnmarshalBody(resp)
	if err != nil {
		return nil, err
	}
	return resp, err
}
func (c *calendarAcl) ListByIterator(ctx context.Context, req *ListCalendarAclReq, options ...larkcore.RequestOptionFunc) (*ListCalendarAclIterator, error) {
	return &ListCalendarAclIterator{
		ctx:      ctx,
		req:      req,
		listFunc: c.List,
		options:  options,
		limit:    req.Limit}, nil
}
func (c *calendarAcl) Subscription(ctx context.Context, req *SubscriptionCalendarAclReq, options ...larkcore.RequestOptionFunc) (*SubscriptionCalendarAclResp, error) {
	// 发起请求
	httpReq := req.HttpReq
	httpReq.ApiPath = "/open-apis/calendar/v4/calendars/:calendar_id/acls/subscription"
	httpReq.HttpMethod = http.MethodPost
	httpReq.SupportedAccessTokenTypes = []larkcore.AccessTokenType{larkcore.AccessTokenTypeUser}
	rawResp, err := larkcore.Request(ctx, req.HttpReq, c.service.config, options...)
	if err != nil {
		return nil, err
	}
	// 反序列响应结果
	resp := &SubscriptionCalendarAclResp{RawResponse: rawResp}
	err = rawResp.JSONUnmarshalBody(resp)
	if err != nil {
		return nil, err
	}
	return resp, err
}
func (c *calendarEvent) Create(ctx context.Context, req *CreateCalendarEventReq, options ...larkcore.RequestOptionFunc) (*CreateCalendarEventResp, error) {
	// 发起请求
	httpReq := req.HttpReq
	httpReq.ApiPath = "/open-apis/calendar/v4/calendars/:calendar_id/events"
	httpReq.HttpMethod = http.MethodPost
	httpReq.SupportedAccessTokenTypes = []larkcore.AccessTokenType{larkcore.AccessTokenTypeTenant, larkcore.AccessTokenTypeUser}
	rawResp, err := larkcore.Request(ctx, req.HttpReq, c.service.config, options...)
	if err != nil {
		return nil, err
	}
	// 反序列响应结果
	resp := &CreateCalendarEventResp{RawResponse: rawResp}
	err = rawResp.JSONUnmarshalBody(resp)
	if err != nil {
		return nil, err
	}
	return resp, err
}
func (c *calendarEvent) Delete(ctx context.Context, req *DeleteCalendarEventReq, options ...larkcore.RequestOptionFunc) (*DeleteCalendarEventResp, error) {
	// 发起请求
	httpReq := req.HttpReq
	httpReq.ApiPath = "/open-apis/calendar/v4/calendars/:calendar_id/events/:event_id"
	httpReq.HttpMethod = http.MethodDelete
	httpReq.SupportedAccessTokenTypes = []larkcore.AccessTokenType{larkcore.AccessTokenTypeTenant, larkcore.AccessTokenTypeUser}
	rawResp, err := larkcore.Request(ctx, req.HttpReq, c.service.config, options...)
	if err != nil {
		return nil, err
	}
	// 反序列响应结果
	resp := &DeleteCalendarEventResp{RawResponse: rawResp}
	err = rawResp.JSONUnmarshalBody(resp)
	if err != nil {
		return nil, err
	}
	return resp, err
}
func (c *calendarEvent) Get(ctx context.Context, req *GetCalendarEventReq, options ...larkcore.RequestOptionFunc) (*GetCalendarEventResp, error) {
	// 发起请求
	httpReq := req.HttpReq
	httpReq.ApiPath = "/open-apis/calendar/v4/calendars/:calendar_id/events/:event_id"
	httpReq.HttpMethod = http.MethodGet
	httpReq.SupportedAccessTokenTypes = []larkcore.AccessTokenType{larkcore.AccessTokenTypeTenant, larkcore.AccessTokenTypeUser}
	rawResp, err := larkcore.Request(ctx, req.HttpReq, c.service.config, options...)
	if err != nil {
		return nil, err
	}
	// 反序列响应结果
	resp := &GetCalendarEventResp{RawResponse: rawResp}
	err = rawResp.JSONUnmarshalBody(resp)
	if err != nil {
		return nil, err
	}
	return resp, err
}
func (c *calendarEvent) List(ctx context.Context, req *ListCalendarEventReq, options ...larkcore.RequestOptionFunc) (*ListCalendarEventResp, error) {
	// 发起请求
	httpReq := req.HttpReq
	httpReq.ApiPath = "/open-apis/calendar/v4/calendars/:calendar_id/events"
	httpReq.HttpMethod = http.MethodGet
	httpReq.SupportedAccessTokenTypes = []larkcore.AccessTokenType{larkcore.AccessTokenTypeTenant, larkcore.AccessTokenTypeUser}
	rawResp, err := larkcore.Request(ctx, req.HttpReq, c.service.config, options...)
	if err != nil {
		return nil, err
	}
	// 反序列响应结果
	resp := &ListCalendarEventResp{RawResponse: rawResp}
	err = rawResp.JSONUnmarshalBody(resp)
	if err != nil {
		return nil, err
	}
	return resp, err
}
func (c *calendarEvent) Patch(ctx context.Context, req *PatchCalendarEventReq, options ...larkcore.RequestOptionFunc) (*PatchCalendarEventResp, error) {
	// 发起请求
	httpReq := req.HttpReq
	httpReq.ApiPath = "/open-apis/calendar/v4/calendars/:calendar_id/events/:event_id"
	httpReq.HttpMethod = http.MethodPatch
	httpReq.SupportedAccessTokenTypes = []larkcore.AccessTokenType{larkcore.AccessTokenTypeTenant, larkcore.AccessTokenTypeUser}
	rawResp, err := larkcore.Request(ctx, req.HttpReq, c.service.config, options...)
	if err != nil {
		return nil, err
	}
	// 反序列响应结果
	resp := &PatchCalendarEventResp{RawResponse: rawResp}
	err = rawResp.JSONUnmarshalBody(resp)
	if err != nil {
		return nil, err
	}
	return resp, err
}
func (c *calendarEvent) Search(ctx context.Context, req *SearchCalendarEventReq, options ...larkcore.RequestOptionFunc) (*SearchCalendarEventResp, error) {
	// 发起请求
	httpReq := req.HttpReq
	httpReq.ApiPath = "/open-apis/calendar/v4/calendars/:calendar_id/events/search"
	httpReq.HttpMethod = http.MethodPost
	httpReq.SupportedAccessTokenTypes = []larkcore.AccessTokenType{larkcore.AccessTokenTypeUser}
	rawResp, err := larkcore.Request(ctx, req.HttpReq, c.service.config, options...)
	if err != nil {
		return nil, err
	}
	// 反序列响应结果
	resp := &SearchCalendarEventResp{RawResponse: rawResp}
	err = rawResp.JSONUnmarshalBody(resp)
	if err != nil {
		return nil, err
	}
	return resp, err
}
func (c *calendarEvent) SearchByIterator(ctx context.Context, req *SearchCalendarEventReq, options ...larkcore.RequestOptionFunc) (*SearchCalendarEventIterator, error) {
	return &SearchCalendarEventIterator{
		ctx:      ctx,
		req:      req,
		listFunc: c.Search,
		options:  options,
		limit:    req.Limit}, nil
}
func (c *calendarEvent) Subscription(ctx context.Context, req *SubscriptionCalendarEventReq, options ...larkcore.RequestOptionFunc) (*SubscriptionCalendarEventResp, error) {
	// 发起请求
	httpReq := req.HttpReq
	httpReq.ApiPath = "/open-apis/calendar/v4/calendars/:calendar_id/events/subscription"
	httpReq.HttpMethod = http.MethodPost
	httpReq.SupportedAccessTokenTypes = []larkcore.AccessTokenType{larkcore.AccessTokenTypeUser}
	rawResp, err := larkcore.Request(ctx, req.HttpReq, c.service.config, options...)
	if err != nil {
		return nil, err
	}
	// 反序列响应结果
	resp := &SubscriptionCalendarEventResp{RawResponse: rawResp}
	err = rawResp.JSONUnmarshalBody(resp)
	if err != nil {
		return nil, err
	}
	return resp, err
}
func (c *calendarEventAttendee) BatchDelete(ctx context.Context, req *BatchDeleteCalendarEventAttendeeReq, options ...larkcore.RequestOptionFunc) (*BatchDeleteCalendarEventAttendeeResp, error) {
	// 发起请求
	httpReq := req.HttpReq
	httpReq.ApiPath = "/open-apis/calendar/v4/calendars/:calendar_id/events/:event_id/attendees/batch_delete"
	httpReq.HttpMethod = http.MethodPost
	httpReq.SupportedAccessTokenTypes = []larkcore.AccessTokenType{larkcore.AccessTokenTypeTenant, larkcore.AccessTokenTypeUser}
	rawResp, err := larkcore.Request(ctx, req.HttpReq, c.service.config, options...)
	if err != nil {
		return nil, err
	}
	// 反序列响应结果
	resp := &BatchDeleteCalendarEventAttendeeResp{RawResponse: rawResp}
	err = rawResp.JSONUnmarshalBody(resp)
	if err != nil {
		return nil, err
	}
	return resp, err
}
func (c *calendarEventAttendee) Create(ctx context.Context, req *CreateCalendarEventAttendeeReq, options ...larkcore.RequestOptionFunc) (*CreateCalendarEventAttendeeResp, error) {
	// 发起请求
	httpReq := req.HttpReq
	httpReq.ApiPath = "/open-apis/calendar/v4/calendars/:calendar_id/events/:event_id/attendees"
	httpReq.HttpMethod = http.MethodPost
	httpReq.SupportedAccessTokenTypes = []larkcore.AccessTokenType{larkcore.AccessTokenTypeTenant, larkcore.AccessTokenTypeUser}
	rawResp, err := larkcore.Request(ctx, req.HttpReq, c.service.config, options...)
	if err != nil {
		return nil, err
	}
	// 反序列响应结果
	resp := &CreateCalendarEventAttendeeResp{RawResponse: rawResp}
	err = rawResp.JSONUnmarshalBody(resp)
	if err != nil {
		return nil, err
	}
	return resp, err
}
func (c *calendarEventAttendee) List(ctx context.Context, req *ListCalendarEventAttendeeReq, options ...larkcore.RequestOptionFunc) (*ListCalendarEventAttendeeResp, error) {
	// 发起请求
	httpReq := req.HttpReq
	httpReq.ApiPath = "/open-apis/calendar/v4/calendars/:calendar_id/events/:event_id/attendees"
	httpReq.HttpMethod = http.MethodGet
	httpReq.SupportedAccessTokenTypes = []larkcore.AccessTokenType{larkcore.AccessTokenTypeTenant, larkcore.AccessTokenTypeUser}
	rawResp, err := larkcore.Request(ctx, req.HttpReq, c.service.config, options...)
	if err != nil {
		return nil, err
	}
	// 反序列响应结果
	resp := &ListCalendarEventAttendeeResp{RawResponse: rawResp}
	err = rawResp.JSONUnmarshalBody(resp)
	if err != nil {
		return nil, err
	}
	return resp, err
}
func (c *calendarEventAttendee) ListByIterator(ctx context.Context, req *ListCalendarEventAttendeeReq, options ...larkcore.RequestOptionFunc) (*ListCalendarEventAttendeeIterator, error) {
	return &ListCalendarEventAttendeeIterator{
		ctx:      ctx,
		req:      req,
		listFunc: c.List,
		options:  options,
		limit:    req.Limit}, nil
}
func (c *calendarEventAttendeeChatMember) List(ctx context.Context, req *ListCalendarEventAttendeeChatMemberReq, options ...larkcore.RequestOptionFunc) (*ListCalendarEventAttendeeChatMemberResp, error) {
	// 发起请求
	httpReq := req.HttpReq
	httpReq.ApiPath = "/open-apis/calendar/v4/calendars/:calendar_id/events/:event_id/attendees/:attendee_id/chat_members"
	httpReq.HttpMethod = http.MethodGet
	httpReq.SupportedAccessTokenTypes = []larkcore.AccessTokenType{larkcore.AccessTokenTypeTenant, larkcore.AccessTokenTypeUser}
	rawResp, err := larkcore.Request(ctx, req.HttpReq, c.service.config, options...)
	if err != nil {
		return nil, err
	}
	// 反序列响应结果
	resp := &ListCalendarEventAttendeeChatMemberResp{RawResponse: rawResp}
	err = rawResp.JSONUnmarshalBody(resp)
	if err != nil {
		return nil, err
	}
	return resp, err
}
func (c *calendarEventAttendeeChatMember) ListByIterator(ctx context.Context, req *ListCalendarEventAttendeeChatMemberReq, options ...larkcore.RequestOptionFunc) (*ListCalendarEventAttendeeChatMemberIterator, error) {
	return &ListCalendarEventAttendeeChatMemberIterator{
		ctx:      ctx,
		req:      req,
		listFunc: c.List,
		options:  options,
		limit:    req.Limit}, nil
}
func (e *exchangeBinding) Create(ctx context.Context, req *CreateExchangeBindingReq, options ...larkcore.RequestOptionFunc) (*CreateExchangeBindingResp, error) {
	// 发起请求
	httpReq := req.HttpReq
	httpReq.ApiPath = "/open-apis/calendar/v4/exchange_bindings"
	httpReq.HttpMethod = http.MethodPost
	httpReq.SupportedAccessTokenTypes = []larkcore.AccessTokenType{larkcore.AccessTokenTypeUser}
	rawResp, err := larkcore.Request(ctx, req.HttpReq, e.service.config, options...)
	if err != nil {
		return nil, err
	}
	// 反序列响应结果
	resp := &CreateExchangeBindingResp{RawResponse: rawResp}
	err = rawResp.JSONUnmarshalBody(resp)
	if err != nil {
		return nil, err
	}
	return resp, err
}
func (e *exchangeBinding) Delete(ctx context.Context, req *DeleteExchangeBindingReq, options ...larkcore.RequestOptionFunc) (*DeleteExchangeBindingResp, error) {
	// 发起请求
	httpReq := req.HttpReq
	httpReq.ApiPath = "/open-apis/calendar/v4/exchange_bindings/:exchange_binding_id"
	httpReq.HttpMethod = http.MethodDelete
	httpReq.SupportedAccessTokenTypes = []larkcore.AccessTokenType{larkcore.AccessTokenTypeUser}
	rawResp, err := larkcore.Request(ctx, req.HttpReq, e.service.config, options...)
	if err != nil {
		return nil, err
	}
	// 反序列响应结果
	resp := &DeleteExchangeBindingResp{RawResponse: rawResp}
	err = rawResp.JSONUnmarshalBody(resp)
	if err != nil {
		return nil, err
	}
	return resp, err
}
func (e *exchangeBinding) Get(ctx context.Context, req *GetExchangeBindingReq, options ...larkcore.RequestOptionFunc) (*GetExchangeBindingResp, error) {
	// 发起请求
	httpReq := req.HttpReq
	httpReq.ApiPath = "/open-apis/calendar/v4/exchange_bindings/:exchange_binding_id"
	httpReq.HttpMethod = http.MethodGet
	httpReq.SupportedAccessTokenTypes = []larkcore.AccessTokenType{larkcore.AccessTokenTypeUser}
	rawResp, err := larkcore.Request(ctx, req.HttpReq, e.service.config, options...)
	if err != nil {
		return nil, err
	}
	// 反序列响应结果
	resp := &GetExchangeBindingResp{RawResponse: rawResp}
	err = rawResp.JSONUnmarshalBody(resp)
	if err != nil {
		return nil, err
	}
	return resp, err
}
func (f *freebusy) List(ctx context.Context, req *ListFreebusyReq, options ...larkcore.RequestOptionFunc) (*ListFreebusyResp, error) {
	// 发起请求
	httpReq := req.HttpReq
	httpReq.ApiPath = "/open-apis/calendar/v4/freebusy/list"
	httpReq.HttpMethod = http.MethodPost
	httpReq.SupportedAccessTokenTypes = []larkcore.AccessTokenType{larkcore.AccessTokenTypeTenant}
	rawResp, err := larkcore.Request(ctx, req.HttpReq, f.service.config, options...)
	if err != nil {
		return nil, err
	}
	// 反序列响应结果
	resp := &ListFreebusyResp{RawResponse: rawResp}
	err = rawResp.JSONUnmarshalBody(resp)
	if err != nil {
		return nil, err
	}
	return resp, err
}
func (s *setting) GenerateCaldavConf(ctx context.Context, req *GenerateCaldavConfSettingReq, options ...larkcore.RequestOptionFunc) (*GenerateCaldavConfSettingResp, error) {
	// 发起请求
	httpReq := req.HttpReq
	httpReq.ApiPath = "/open-apis/calendar/v4/settings/generate_caldav_conf"
	httpReq.HttpMethod = http.MethodPost
	httpReq.SupportedAccessTokenTypes = []larkcore.AccessTokenType{larkcore.AccessTokenTypeUser}
	rawResp, err := larkcore.Request(ctx, req.HttpReq, s.service.config, options...)
	if err != nil {
		return nil, err
	}
	// 反序列响应结果
	resp := &GenerateCaldavConfSettingResp{RawResponse: rawResp}
	err = rawResp.JSONUnmarshalBody(resp)
	if err != nil {
		return nil, err
	}
	return resp, err
}
func (t *timeoffEvent) Create(ctx context.Context, req *CreateTimeoffEventReq, options ...larkcore.RequestOptionFunc) (*CreateTimeoffEventResp, error) {
	// 发起请求
	httpReq := req.HttpReq
	httpReq.ApiPath = "/open-apis/calendar/v4/timeoff_events"
	httpReq.HttpMethod = http.MethodPost
	httpReq.SupportedAccessTokenTypes = []larkcore.AccessTokenType{larkcore.AccessTokenTypeTenant}
	rawResp, err := larkcore.Request(ctx, req.HttpReq, t.service.config, options...)
	if err != nil {
		return nil, err
	}
	// 反序列响应结果
	resp := &CreateTimeoffEventResp{RawResponse: rawResp}
	err = rawResp.JSONUnmarshalBody(resp)
	if err != nil {
		return nil, err
	}
	return resp, err
}
func (t *timeoffEvent) Delete(ctx context.Context, req *DeleteTimeoffEventReq, options ...larkcore.RequestOptionFunc) (*DeleteTimeoffEventResp, error) {
	// 发起请求
	httpReq := req.HttpReq
	httpReq.ApiPath = "/open-apis/calendar/v4/timeoff_events/:timeoff_event_id"
	httpReq.HttpMethod = http.MethodDelete
	httpReq.SupportedAccessTokenTypes = []larkcore.AccessTokenType{larkcore.AccessTokenTypeTenant}
	rawResp, err := larkcore.Request(ctx, req.HttpReq, t.service.config, options...)
	if err != nil {
		return nil, err
	}
	// 反序列响应结果
	resp := &DeleteTimeoffEventResp{RawResponse: rawResp}
	err = rawResp.JSONUnmarshalBody(resp)
	if err != nil {
		return nil, err
	}
	return resp, err
}
