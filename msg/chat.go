package msg

type MsgRoleType int32

const (
	RoleType_None MsgRoleType = iota
	RoleType_User
	RoleType_Ass
)

type ChatContext struct {
	Role    MsgRoleType `json:"role"`
	Context string      `json:"context"`
}

type MsgTalkReq struct {
	Req
	Messages []ChatContext `json:"messages"`
}

type MsgTalkResp struct {
	Resp
	Index   int32
	Context string
}
