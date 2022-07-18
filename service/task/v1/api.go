// Package task code generated by oapi sdk gen
package larktask

import (
	"context"
	"net/http"

	"github.com/larksuite/oapi-sdk-go/v3/core"
)

// 构建业务域服务实例
func NewService(config *larkcore.Config) *TaskService {
	t := &TaskService{config: config}
	t.Task = &task{service: t}
	t.TaskCollaborator = &taskCollaborator{service: t}
	t.TaskComment = &taskComment{service: t}
	t.TaskFollower = &taskFollower{service: t}
	t.TaskReminder = &taskReminder{service: t}
	return t
}

// 业务域服务定义
type TaskService struct {
	config           *larkcore.Config
	Task             *task
	TaskCollaborator *taskCollaborator
	TaskComment      *taskComment
	TaskFollower     *taskFollower
	TaskReminder     *taskReminder
}

// 资源服务定义
type task struct {
	service *TaskService
}
type taskCollaborator struct {
	service *TaskService
}
type taskComment struct {
	service *TaskService
}
type taskFollower struct {
	service *TaskService
}
type taskReminder struct {
	service *TaskService
}

// 资源服务方法定义
func (t *task) Complete(ctx context.Context, req *CompleteTaskReq, options ...larkcore.RequestOptionFunc) (*CompleteTaskResp, error) {
	// 发起请求
	httpReq := req.HttpReq
	httpReq.ApiPath = "/open-apis/task/v1/tasks/:task_id/complete"
	httpReq.HttpMethod = http.MethodPost
	httpReq.SupportedAccessTokenTypes = []larkcore.AccessTokenType{larkcore.AccessTokenTypeTenant, larkcore.AccessTokenTypeUser}
	rawResp, err := larkcore.Request(ctx, req.HttpReq, t.service.config, options...)
	if err != nil {
		return nil, err
	}
	// 反序列响应结果
	resp := &CompleteTaskResp{RawResponse: rawResp}
	err = rawResp.JSONUnmarshalBody(resp)
	if err != nil {
		return nil, err
	}
	return resp, err
}
func (t *task) Create(ctx context.Context, req *CreateTaskReq, options ...larkcore.RequestOptionFunc) (*CreateTaskResp, error) {
	// 发起请求
	httpReq := req.HttpReq
	httpReq.ApiPath = "/open-apis/task/v1/tasks"
	httpReq.HttpMethod = http.MethodPost
	httpReq.SupportedAccessTokenTypes = []larkcore.AccessTokenType{larkcore.AccessTokenTypeTenant, larkcore.AccessTokenTypeUser}
	rawResp, err := larkcore.Request(ctx, req.HttpReq, t.service.config, options...)
	if err != nil {
		return nil, err
	}
	// 反序列响应结果
	resp := &CreateTaskResp{RawResponse: rawResp}
	err = rawResp.JSONUnmarshalBody(resp)
	if err != nil {
		return nil, err
	}
	return resp, err
}
func (t *task) Delete(ctx context.Context, req *DeleteTaskReq, options ...larkcore.RequestOptionFunc) (*DeleteTaskResp, error) {
	// 发起请求
	httpReq := req.HttpReq
	httpReq.ApiPath = "/open-apis/task/v1/tasks/:task_id"
	httpReq.HttpMethod = http.MethodDelete
	httpReq.SupportedAccessTokenTypes = []larkcore.AccessTokenType{larkcore.AccessTokenTypeTenant, larkcore.AccessTokenTypeUser}
	rawResp, err := larkcore.Request(ctx, req.HttpReq, t.service.config, options...)
	if err != nil {
		return nil, err
	}
	// 反序列响应结果
	resp := &DeleteTaskResp{RawResponse: rawResp}
	err = rawResp.JSONUnmarshalBody(resp)
	if err != nil {
		return nil, err
	}
	return resp, err
}
func (t *task) Get(ctx context.Context, req *GetTaskReq, options ...larkcore.RequestOptionFunc) (*GetTaskResp, error) {
	// 发起请求
	httpReq := req.HttpReq
	httpReq.ApiPath = "/open-apis/task/v1/tasks/:task_id"
	httpReq.HttpMethod = http.MethodGet
	httpReq.SupportedAccessTokenTypes = []larkcore.AccessTokenType{larkcore.AccessTokenTypeTenant, larkcore.AccessTokenTypeUser}
	rawResp, err := larkcore.Request(ctx, req.HttpReq, t.service.config, options...)
	if err != nil {
		return nil, err
	}
	// 反序列响应结果
	resp := &GetTaskResp{RawResponse: rawResp}
	err = rawResp.JSONUnmarshalBody(resp)
	if err != nil {
		return nil, err
	}
	return resp, err
}
func (t *task) List(ctx context.Context, req *ListTaskReq, options ...larkcore.RequestOptionFunc) (*ListTaskResp, error) {
	// 发起请求
	httpReq := req.HttpReq
	httpReq.ApiPath = "/open-apis/task/v1/tasks"
	httpReq.HttpMethod = http.MethodGet
	httpReq.SupportedAccessTokenTypes = []larkcore.AccessTokenType{larkcore.AccessTokenTypeUser, larkcore.AccessTokenTypeTenant}
	rawResp, err := larkcore.Request(ctx, req.HttpReq, t.service.config, options...)
	if err != nil {
		return nil, err
	}
	// 反序列响应结果
	resp := &ListTaskResp{RawResponse: rawResp}
	err = rawResp.JSONUnmarshalBody(resp)
	if err != nil {
		return nil, err
	}
	return resp, err
}
func (t *task) ListByIterator(ctx context.Context, req *ListTaskReq, options ...larkcore.RequestOptionFunc) (*ListTaskIterator, error) {
	return &ListTaskIterator{
		ctx:      ctx,
		req:      req,
		listFunc: t.List,
		options:  options,
		limit:    req.Limit}, nil
}
func (t *task) Patch(ctx context.Context, req *PatchTaskReq, options ...larkcore.RequestOptionFunc) (*PatchTaskResp, error) {
	// 发起请求
	httpReq := req.HttpReq
	httpReq.ApiPath = "/open-apis/task/v1/tasks/:task_id"
	httpReq.HttpMethod = http.MethodPatch
	httpReq.SupportedAccessTokenTypes = []larkcore.AccessTokenType{larkcore.AccessTokenTypeTenant, larkcore.AccessTokenTypeUser}
	rawResp, err := larkcore.Request(ctx, req.HttpReq, t.service.config, options...)
	if err != nil {
		return nil, err
	}
	// 反序列响应结果
	resp := &PatchTaskResp{RawResponse: rawResp}
	err = rawResp.JSONUnmarshalBody(resp)
	if err != nil {
		return nil, err
	}
	return resp, err
}
func (t *task) Uncomplete(ctx context.Context, req *UncompleteTaskReq, options ...larkcore.RequestOptionFunc) (*UncompleteTaskResp, error) {
	// 发起请求
	httpReq := req.HttpReq
	httpReq.ApiPath = "/open-apis/task/v1/tasks/:task_id/uncomplete"
	httpReq.HttpMethod = http.MethodPost
	httpReq.SupportedAccessTokenTypes = []larkcore.AccessTokenType{larkcore.AccessTokenTypeTenant, larkcore.AccessTokenTypeUser}
	rawResp, err := larkcore.Request(ctx, req.HttpReq, t.service.config, options...)
	if err != nil {
		return nil, err
	}
	// 反序列响应结果
	resp := &UncompleteTaskResp{RawResponse: rawResp}
	err = rawResp.JSONUnmarshalBody(resp)
	if err != nil {
		return nil, err
	}
	return resp, err
}
func (t *taskCollaborator) Create(ctx context.Context, req *CreateTaskCollaboratorReq, options ...larkcore.RequestOptionFunc) (*CreateTaskCollaboratorResp, error) {
	// 发起请求
	httpReq := req.HttpReq
	httpReq.ApiPath = "/open-apis/task/v1/tasks/:task_id/collaborators"
	httpReq.HttpMethod = http.MethodPost
	httpReq.SupportedAccessTokenTypes = []larkcore.AccessTokenType{larkcore.AccessTokenTypeTenant, larkcore.AccessTokenTypeUser}
	rawResp, err := larkcore.Request(ctx, req.HttpReq, t.service.config, options...)
	if err != nil {
		return nil, err
	}
	// 反序列响应结果
	resp := &CreateTaskCollaboratorResp{RawResponse: rawResp}
	err = rawResp.JSONUnmarshalBody(resp)
	if err != nil {
		return nil, err
	}
	return resp, err
}
func (t *taskCollaborator) Delete(ctx context.Context, req *DeleteTaskCollaboratorReq, options ...larkcore.RequestOptionFunc) (*DeleteTaskCollaboratorResp, error) {
	// 发起请求
	httpReq := req.HttpReq
	httpReq.ApiPath = "/open-apis/task/v1/tasks/:task_id/collaborators/:collaborator_id"
	httpReq.HttpMethod = http.MethodDelete
	httpReq.SupportedAccessTokenTypes = []larkcore.AccessTokenType{larkcore.AccessTokenTypeTenant, larkcore.AccessTokenTypeUser}
	rawResp, err := larkcore.Request(ctx, req.HttpReq, t.service.config, options...)
	if err != nil {
		return nil, err
	}
	// 反序列响应结果
	resp := &DeleteTaskCollaboratorResp{RawResponse: rawResp}
	err = rawResp.JSONUnmarshalBody(resp)
	if err != nil {
		return nil, err
	}
	return resp, err
}
func (t *taskCollaborator) List(ctx context.Context, req *ListTaskCollaboratorReq, options ...larkcore.RequestOptionFunc) (*ListTaskCollaboratorResp, error) {
	// 发起请求
	httpReq := req.HttpReq
	httpReq.ApiPath = "/open-apis/task/v1/tasks/:task_id/collaborators"
	httpReq.HttpMethod = http.MethodGet
	httpReq.SupportedAccessTokenTypes = []larkcore.AccessTokenType{larkcore.AccessTokenTypeTenant, larkcore.AccessTokenTypeUser}
	rawResp, err := larkcore.Request(ctx, req.HttpReq, t.service.config, options...)
	if err != nil {
		return nil, err
	}
	// 反序列响应结果
	resp := &ListTaskCollaboratorResp{RawResponse: rawResp}
	err = rawResp.JSONUnmarshalBody(resp)
	if err != nil {
		return nil, err
	}
	return resp, err
}
func (t *taskCollaborator) ListByIterator(ctx context.Context, req *ListTaskCollaboratorReq, options ...larkcore.RequestOptionFunc) (*ListTaskCollaboratorIterator, error) {
	return &ListTaskCollaboratorIterator{
		ctx:      ctx,
		req:      req,
		listFunc: t.List,
		options:  options,
		limit:    req.Limit}, nil
}
func (t *taskComment) Create(ctx context.Context, req *CreateTaskCommentReq, options ...larkcore.RequestOptionFunc) (*CreateTaskCommentResp, error) {
	// 发起请求
	httpReq := req.HttpReq
	httpReq.ApiPath = "/open-apis/task/v1/tasks/:task_id/comments"
	httpReq.HttpMethod = http.MethodPost
	httpReq.SupportedAccessTokenTypes = []larkcore.AccessTokenType{larkcore.AccessTokenTypeTenant, larkcore.AccessTokenTypeUser}
	rawResp, err := larkcore.Request(ctx, req.HttpReq, t.service.config, options...)
	if err != nil {
		return nil, err
	}
	// 反序列响应结果
	resp := &CreateTaskCommentResp{RawResponse: rawResp}
	err = rawResp.JSONUnmarshalBody(resp)
	if err != nil {
		return nil, err
	}
	return resp, err
}
func (t *taskComment) Delete(ctx context.Context, req *DeleteTaskCommentReq, options ...larkcore.RequestOptionFunc) (*DeleteTaskCommentResp, error) {
	// 发起请求
	httpReq := req.HttpReq
	httpReq.ApiPath = "/open-apis/task/v1/tasks/:task_id/comments/:comment_id"
	httpReq.HttpMethod = http.MethodDelete
	httpReq.SupportedAccessTokenTypes = []larkcore.AccessTokenType{larkcore.AccessTokenTypeTenant, larkcore.AccessTokenTypeUser}
	rawResp, err := larkcore.Request(ctx, req.HttpReq, t.service.config, options...)
	if err != nil {
		return nil, err
	}
	// 反序列响应结果
	resp := &DeleteTaskCommentResp{RawResponse: rawResp}
	err = rawResp.JSONUnmarshalBody(resp)
	if err != nil {
		return nil, err
	}
	return resp, err
}
func (t *taskComment) Get(ctx context.Context, req *GetTaskCommentReq, options ...larkcore.RequestOptionFunc) (*GetTaskCommentResp, error) {
	// 发起请求
	httpReq := req.HttpReq
	httpReq.ApiPath = "/open-apis/task/v1/tasks/:task_id/comments/:comment_id"
	httpReq.HttpMethod = http.MethodGet
	httpReq.SupportedAccessTokenTypes = []larkcore.AccessTokenType{larkcore.AccessTokenTypeTenant, larkcore.AccessTokenTypeUser}
	rawResp, err := larkcore.Request(ctx, req.HttpReq, t.service.config, options...)
	if err != nil {
		return nil, err
	}
	// 反序列响应结果
	resp := &GetTaskCommentResp{RawResponse: rawResp}
	err = rawResp.JSONUnmarshalBody(resp)
	if err != nil {
		return nil, err
	}
	return resp, err
}
func (t *taskComment) List(ctx context.Context, req *ListTaskCommentReq, options ...larkcore.RequestOptionFunc) (*ListTaskCommentResp, error) {
	// 发起请求
	httpReq := req.HttpReq
	httpReq.ApiPath = "/open-apis/task/v1/tasks/:task_id/comments"
	httpReq.HttpMethod = http.MethodGet
	httpReq.SupportedAccessTokenTypes = []larkcore.AccessTokenType{larkcore.AccessTokenTypeTenant, larkcore.AccessTokenTypeUser}
	rawResp, err := larkcore.Request(ctx, req.HttpReq, t.service.config, options...)
	if err != nil {
		return nil, err
	}
	// 反序列响应结果
	resp := &ListTaskCommentResp{RawResponse: rawResp}
	err = rawResp.JSONUnmarshalBody(resp)
	if err != nil {
		return nil, err
	}
	return resp, err
}
func (t *taskComment) ListByIterator(ctx context.Context, req *ListTaskCommentReq, options ...larkcore.RequestOptionFunc) (*ListTaskCommentIterator, error) {
	return &ListTaskCommentIterator{
		ctx:      ctx,
		req:      req,
		listFunc: t.List,
		options:  options,
		limit:    req.Limit}, nil
}
func (t *taskComment) Update(ctx context.Context, req *UpdateTaskCommentReq, options ...larkcore.RequestOptionFunc) (*UpdateTaskCommentResp, error) {
	// 发起请求
	httpReq := req.HttpReq
	httpReq.ApiPath = "/open-apis/task/v1/tasks/:task_id/comments/:comment_id"
	httpReq.HttpMethod = http.MethodPut
	httpReq.SupportedAccessTokenTypes = []larkcore.AccessTokenType{larkcore.AccessTokenTypeTenant, larkcore.AccessTokenTypeUser}
	rawResp, err := larkcore.Request(ctx, req.HttpReq, t.service.config, options...)
	if err != nil {
		return nil, err
	}
	// 反序列响应结果
	resp := &UpdateTaskCommentResp{RawResponse: rawResp}
	err = rawResp.JSONUnmarshalBody(resp)
	if err != nil {
		return nil, err
	}
	return resp, err
}
func (t *taskFollower) Create(ctx context.Context, req *CreateTaskFollowerReq, options ...larkcore.RequestOptionFunc) (*CreateTaskFollowerResp, error) {
	// 发起请求
	httpReq := req.HttpReq
	httpReq.ApiPath = "/open-apis/task/v1/tasks/:task_id/followers"
	httpReq.HttpMethod = http.MethodPost
	httpReq.SupportedAccessTokenTypes = []larkcore.AccessTokenType{larkcore.AccessTokenTypeTenant, larkcore.AccessTokenTypeUser}
	rawResp, err := larkcore.Request(ctx, req.HttpReq, t.service.config, options...)
	if err != nil {
		return nil, err
	}
	// 反序列响应结果
	resp := &CreateTaskFollowerResp{RawResponse: rawResp}
	err = rawResp.JSONUnmarshalBody(resp)
	if err != nil {
		return nil, err
	}
	return resp, err
}
func (t *taskFollower) Delete(ctx context.Context, req *DeleteTaskFollowerReq, options ...larkcore.RequestOptionFunc) (*DeleteTaskFollowerResp, error) {
	// 发起请求
	httpReq := req.HttpReq
	httpReq.ApiPath = "/open-apis/task/v1/tasks/:task_id/followers/:follower_id"
	httpReq.HttpMethod = http.MethodDelete
	httpReq.SupportedAccessTokenTypes = []larkcore.AccessTokenType{larkcore.AccessTokenTypeTenant, larkcore.AccessTokenTypeUser}
	rawResp, err := larkcore.Request(ctx, req.HttpReq, t.service.config, options...)
	if err != nil {
		return nil, err
	}
	// 反序列响应结果
	resp := &DeleteTaskFollowerResp{RawResponse: rawResp}
	err = rawResp.JSONUnmarshalBody(resp)
	if err != nil {
		return nil, err
	}
	return resp, err
}
func (t *taskFollower) List(ctx context.Context, req *ListTaskFollowerReq, options ...larkcore.RequestOptionFunc) (*ListTaskFollowerResp, error) {
	// 发起请求
	httpReq := req.HttpReq
	httpReq.ApiPath = "/open-apis/task/v1/tasks/:task_id/followers"
	httpReq.HttpMethod = http.MethodGet
	httpReq.SupportedAccessTokenTypes = []larkcore.AccessTokenType{larkcore.AccessTokenTypeTenant, larkcore.AccessTokenTypeUser}
	rawResp, err := larkcore.Request(ctx, req.HttpReq, t.service.config, options...)
	if err != nil {
		return nil, err
	}
	// 反序列响应结果
	resp := &ListTaskFollowerResp{RawResponse: rawResp}
	err = rawResp.JSONUnmarshalBody(resp)
	if err != nil {
		return nil, err
	}
	return resp, err
}
func (t *taskFollower) ListByIterator(ctx context.Context, req *ListTaskFollowerReq, options ...larkcore.RequestOptionFunc) (*ListTaskFollowerIterator, error) {
	return &ListTaskFollowerIterator{
		ctx:      ctx,
		req:      req,
		listFunc: t.List,
		options:  options,
		limit:    req.Limit}, nil
}
func (t *taskReminder) Create(ctx context.Context, req *CreateTaskReminderReq, options ...larkcore.RequestOptionFunc) (*CreateTaskReminderResp, error) {
	// 发起请求
	httpReq := req.HttpReq
	httpReq.ApiPath = "/open-apis/task/v1/tasks/:task_id/reminders"
	httpReq.HttpMethod = http.MethodPost
	httpReq.SupportedAccessTokenTypes = []larkcore.AccessTokenType{larkcore.AccessTokenTypeTenant, larkcore.AccessTokenTypeUser}
	rawResp, err := larkcore.Request(ctx, req.HttpReq, t.service.config, options...)
	if err != nil {
		return nil, err
	}
	// 反序列响应结果
	resp := &CreateTaskReminderResp{RawResponse: rawResp}
	err = rawResp.JSONUnmarshalBody(resp)
	if err != nil {
		return nil, err
	}
	return resp, err
}
func (t *taskReminder) Delete(ctx context.Context, req *DeleteTaskReminderReq, options ...larkcore.RequestOptionFunc) (*DeleteTaskReminderResp, error) {
	// 发起请求
	httpReq := req.HttpReq
	httpReq.ApiPath = "/open-apis/task/v1/tasks/:task_id/reminders/:reminder_id"
	httpReq.HttpMethod = http.MethodDelete
	httpReq.SupportedAccessTokenTypes = []larkcore.AccessTokenType{larkcore.AccessTokenTypeTenant, larkcore.AccessTokenTypeUser}
	rawResp, err := larkcore.Request(ctx, req.HttpReq, t.service.config, options...)
	if err != nil {
		return nil, err
	}
	// 反序列响应结果
	resp := &DeleteTaskReminderResp{RawResponse: rawResp}
	err = rawResp.JSONUnmarshalBody(resp)
	if err != nil {
		return nil, err
	}
	return resp, err
}
func (t *taskReminder) List(ctx context.Context, req *ListTaskReminderReq, options ...larkcore.RequestOptionFunc) (*ListTaskReminderResp, error) {
	// 发起请求
	httpReq := req.HttpReq
	httpReq.ApiPath = "/open-apis/task/v1/tasks/:task_id/reminders"
	httpReq.HttpMethod = http.MethodGet
	httpReq.SupportedAccessTokenTypes = []larkcore.AccessTokenType{larkcore.AccessTokenTypeTenant, larkcore.AccessTokenTypeUser}
	rawResp, err := larkcore.Request(ctx, req.HttpReq, t.service.config, options...)
	if err != nil {
		return nil, err
	}
	// 反序列响应结果
	resp := &ListTaskReminderResp{RawResponse: rawResp}
	err = rawResp.JSONUnmarshalBody(resp)
	if err != nil {
		return nil, err
	}
	return resp, err
}
func (t *taskReminder) ListByIterator(ctx context.Context, req *ListTaskReminderReq, options ...larkcore.RequestOptionFunc) (*ListTaskReminderIterator, error) {
	return &ListTaskReminderIterator{
		ctx:      ctx,
		req:      req,
		listFunc: t.List,
		options:  options,
		limit:    req.Limit}, nil
}
