package credential

// OpenTicketHandle ticket获取, 用于 h5 链接拉起抖音发布器分享视频时对开发者身份进行验签；本接口适用于抖音
type OpenTicketHandle interface {
	//GetTicket 获取ticket
	GetTicket(accessToken string) (ticket string, err error)
}
