package msg

type ChatContext struct {
	Role    string
	Content string
}

type MsgTalkReq struct {
	Req
	Context []ChatContext
}

type MsgTalkResp struct {
	Resp
	Index int32
}
