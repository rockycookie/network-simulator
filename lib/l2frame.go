package lib

type L2Frame struct {
	SrcMac    string
	DstMac    string
	Name      string // content to be updated later
	NeedReply bool   // simple flag to indicate if this frame needs a reply
}
