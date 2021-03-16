// Code generated by lark suite oapi sdk gen
package v1

import (
	"github.com/larksuite/oapi-sdk-go/api/core/tools"
	"github.com/larksuite/oapi-sdk-go/event/core/model"
)

type DocsLink struct {
	Url             string   `json:"url,omitempty"`
	ForceSendFields []string `json:"-"`
}

func (s *DocsLink) MarshalJSON() ([]byte, error) {
	type cp DocsLink
	raw := cp(*s)
	return tools.MarshalJSON(raw, s.ForceSendFields)
}

type File struct {
	FileToken       string   `json:"file_token,omitempty"`
	FileName        string   `json:"file_name,omitempty"`
	Size            int      `json:"size,omitempty"`
	MimeType        string   `json:"mime_type,omitempty"`
	ForceSendFields []string `json:"-"`
}

func (s *File) MarshalJSON() ([]byte, error) {
	type cp File
	raw := cp(*s)
	return tools.MarshalJSON(raw, s.ForceSendFields)
}

type FileComment struct {
	CommentId       string     `json:"comment_id,omitempty"`
	UserId          string     `json:"user_id,omitempty"`
	CreateTime      int        `json:"create_time,omitempty"`
	UpdateTime      int        `json:"update_time,omitempty"`
	IsSolved        bool       `json:"is_solved,omitempty"`
	SolvedTime      int        `json:"solved_time,omitempty"`
	SolverUserId    string     `json:"solver_user_id,omitempty"`
	ReplyList       *ReplyList `json:"reply_list,omitempty"`
	ForceSendFields []string   `json:"-"`
}

func (s *FileComment) MarshalJSON() ([]byte, error) {
	type cp FileComment
	raw := cp(*s)
	return tools.MarshalJSON(raw, s.ForceSendFields)
}

type FileCommentReply struct {
	ReplyId         string        `json:"reply_id,omitempty"`
	UserId          string        `json:"user_id,omitempty"`
	CreateTime      int           `json:"create_time,omitempty"`
	UpdateTime      int           `json:"update_time,omitempty"`
	Content         *ReplyContent `json:"content,omitempty"`
	ForceSendFields []string      `json:"-"`
}

func (s *FileCommentReply) MarshalJSON() ([]byte, error) {
	type cp FileCommentReply
	raw := cp(*s)
	return tools.MarshalJSON(raw, s.ForceSendFields)
}

type Media struct {
	FileToken       string   `json:"file_token,omitempty"`
	FileName        string   `json:"file_name,omitempty"`
	Size            int      `json:"size,omitempty"`
	MimeType        string   `json:"mime_type,omitempty"`
	ForceSendFields []string `json:"-"`
}

func (s *Media) MarshalJSON() ([]byte, error) {
	type cp Media
	raw := cp(*s)
	return tools.MarshalJSON(raw, s.ForceSendFields)
}

type Person struct {
	UserId          string   `json:"user_id,omitempty"`
	ForceSendFields []string `json:"-"`
}

func (s *Person) MarshalJSON() ([]byte, error) {
	type cp Person
	raw := cp(*s)
	return tools.MarshalJSON(raw, s.ForceSendFields)
}

type ReplyContent struct {
	Elements        []*ReplyElement `json:"elements,omitempty"`
	ForceSendFields []string        `json:"-"`
}

func (s *ReplyContent) MarshalJSON() ([]byte, error) {
	type cp ReplyContent
	raw := cp(*s)
	return tools.MarshalJSON(raw, s.ForceSendFields)
}

type ReplyElement struct {
	Type            string    `json:"type,omitempty"`
	TextRun         *TextRun  `json:"text_run,omitempty"`
	DocsLink        *DocsLink `json:"docs_link,omitempty"`
	Person          *Person   `json:"person,omitempty"`
	ForceSendFields []string  `json:"-"`
}

func (s *ReplyElement) MarshalJSON() ([]byte, error) {
	type cp ReplyElement
	raw := cp(*s)
	return tools.MarshalJSON(raw, s.ForceSendFields)
}

type ReplyList struct {
	Replies         []*FileCommentReply `json:"replies,omitempty"`
	ForceSendFields []string            `json:"-"`
}

func (s *ReplyList) MarshalJSON() ([]byte, error) {
	type cp ReplyList
	raw := cp(*s)
	return tools.MarshalJSON(raw, s.ForceSendFields)
}

type TextRun struct {
	Text            string   `json:"text,omitempty"`
	ForceSendFields []string `json:"-"`
}

func (s *TextRun) MarshalJSON() ([]byte, error) {
	type cp TextRun
	raw := cp(*s)
	return tools.MarshalJSON(raw, s.ForceSendFields)
}

type TmpDownloadUrl struct {
	FileToken       string   `json:"file_token,omitempty"`
	TmpDownloadUrl  string   `json:"tmp_download_url,omitempty"`
	ForceSendFields []string `json:"-"`
}

func (s *TmpDownloadUrl) MarshalJSON() ([]byte, error) {
	type cp TmpDownloadUrl
	raw := cp(*s)
	return tools.MarshalJSON(raw, s.ForceSendFields)
}

type UploadInfo struct {
	FileName        string   `json:"file_name,omitempty"`
	ParentType      string   `json:"parent_type,omitempty"`
	ParentNode      string   `json:"parent_node,omitempty"`
	Size            int      `json:"size,omitempty"`
	ForceSendFields []string `json:"-"`
}

func (s *UploadInfo) MarshalJSON() ([]byte, error) {
	type cp UploadInfo
	raw := cp(*s)
	return tools.MarshalJSON(raw, s.ForceSendFields)
}

type UserId struct {
	UserId          string   `json:"user_id,omitempty"`
	OpenId          string   `json:"open_id,omitempty"`
	UnionId         string   `json:"union_id,omitempty"`
	ForceSendFields []string `json:"-"`
}

func (s *UserId) MarshalJSON() ([]byte, error) {
	type cp UserId
	raw := cp(*s)
	return tools.MarshalJSON(raw, s.ForceSendFields)
}

type FileUploadFinishReqBody struct {
	UploadId        string   `json:"upload_id,omitempty"`
	BlockNum        int      `json:"block_num,omitempty"`
	ForceSendFields []string `json:"-"`
}

func (s *FileUploadFinishReqBody) MarshalJSON() ([]byte, error) {
	type cp FileUploadFinishReqBody
	raw := cp(*s)
	return tools.MarshalJSON(raw, s.ForceSendFields)
}

type FileUploadFinishResult struct {
	FileToken string `json:"file_token,omitempty"`
}

type FileUploadPrepareResult struct {
	UploadId  string `json:"upload_id,omitempty"`
	BlockSize int    `json:"block_size,omitempty"`
	BlockNum  int    `json:"block_num,omitempty"`
}

type MediaUploadAllResult struct {
	FileToken string `json:"file_token,omitempty"`
}

type MediaUploadFinishReqBody struct {
	UploadId        string   `json:"upload_id,omitempty"`
	BlockNum        int      `json:"block_num,omitempty"`
	ForceSendFields []string `json:"-"`
}

func (s *MediaUploadFinishReqBody) MarshalJSON() ([]byte, error) {
	type cp MediaUploadFinishReqBody
	raw := cp(*s)
	return tools.MarshalJSON(raw, s.ForceSendFields)
}

type MediaUploadFinishResult struct {
	FileToken string `json:"file_token,omitempty"`
}

type FileUploadAllResult struct {
	FileToken string `json:"file_token,omitempty"`
}

type MediaUploadPrepareResult struct {
	UploadId  string `json:"upload_id,omitempty"`
	BlockSize int    `json:"block_size,omitempty"`
	BlockNum  int    `json:"block_num,omitempty"`
}

type MediaBatchGetTmpDownloadUrlResult struct {
	TmpDownloadUrls []*TmpDownloadUrl `json:"tmp_download_urls,omitempty"`
}

type FileCommentReplyUpdateReqBody struct {
	Content         *ReplyContent `json:"content,omitempty"`
	ForceSendFields []string      `json:"-"`
}

func (s *FileCommentReplyUpdateReqBody) MarshalJSON() ([]byte, error) {
	type cp FileCommentReplyUpdateReqBody
	raw := cp(*s)
	return tools.MarshalJSON(raw, s.ForceSendFields)
}

type FileCommentListResult struct {
	HasMore   bool           `json:"has_more,omitempty"`
	PageToken string         `json:"page_token,omitempty"`
	Items     []*FileComment `json:"items,omitempty"`
}

type FileCommentPatchReqBody struct {
	IsSolved        bool     `json:"is_solved,omitempty"`
	ForceSendFields []string `json:"-"`
}

func (s *FileCommentPatchReqBody) MarshalJSON() ([]byte, error) {
	type cp FileCommentPatchReqBody
	raw := cp(*s)
	return tools.MarshalJSON(raw, s.ForceSendFields)
}

type FileDeletedEventData struct {
	FileType   string  `json:"file_type,omitempty"`
	FileToken  string  `json:"file_token,omitempty"`
	OperatorId *UserId `json:"operator_id,omitempty"`
}

type FileDeletedEvent struct {
	*model.BaseEventV2
	Event *FileDeletedEventData `json:"event"`
}

type FilePermissionMemberAddedEventData struct {
	FileType   string    `json:"file_type,omitempty"`
	FileToken  string    `json:"file_token,omitempty"`
	OperatorId *UserId   `json:"operator_id,omitempty"`
	UserList   []*UserId `json:"user_list,omitempty"`
	ChatList   []string  `json:"chat_list,omitempty"`
}

type FilePermissionMemberAddedEvent struct {
	*model.BaseEventV2
	Event *FilePermissionMemberAddedEventData `json:"event"`
}

type FilePermissionMemberRemovedEventData struct {
	FileType   string    `json:"file_type,omitempty"`
	FileToken  string    `json:"file_token,omitempty"`
	OperatorId *UserId   `json:"operator_id,omitempty"`
	UserList   []*UserId `json:"user_list,omitempty"`
	ChatList   []string  `json:"chat_list,omitempty"`
}

type FilePermissionMemberRemovedEvent struct {
	*model.BaseEventV2
	Event *FilePermissionMemberRemovedEventData `json:"event"`
}

type FileReadEventData struct {
	FileType       string    `json:"file_type,omitempty"`
	FileToken      string    `json:"file_token,omitempty"`
	OperatorIdList []*UserId `json:"operator_id_list,omitempty"`
}

type FileReadEvent struct {
	*model.BaseEventV2
	Event *FileReadEventData `json:"event"`
}

type FileTitleUpdatedEventData struct {
	FileType   string  `json:"file_type,omitempty"`
	FileToken  string  `json:"file_token,omitempty"`
	OperatorId *UserId `json:"operator_id,omitempty"`
}

type FileTitleUpdatedEvent struct {
	*model.BaseEventV2
	Event *FileTitleUpdatedEventData `json:"event"`
}

type FileTrashedEventData struct {
	FileType   string  `json:"file_type,omitempty"`
	FileToken  string  `json:"file_token,omitempty"`
	OperatorId *UserId `json:"operator_id,omitempty"`
}

type FileTrashedEvent struct {
	*model.BaseEventV2
	Event *FileTrashedEventData `json:"event"`
}
