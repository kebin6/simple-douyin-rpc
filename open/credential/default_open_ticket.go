package credential

import (
	"encoding/json"
	"fmt"
	"github.com/kebin6/simple-douyin-rpc/open/cache"
	"github.com/kebin6/simple-douyin-rpc/open/util"
	"sync"
	"time"
)

// 获取ticket的url
const getOpenTicketURL = "https://open.douyin.com/open/getticket?access_token=%s"

// DefaultOpenTicket 默认获取open ticket方法
type DefaultOpenTicket struct {
	appID          string
	cacheKeyPrefix string
	cache          cache.Cache
	//OpenAPITicketLock 读写锁 同一个AppID一个
	OpenAPITicketLock *sync.Mutex
}

// NewDefaultOpenTicket new
func NewDefaultOpenTicket(appID string, cacheKeyPrefix string, cache cache.Cache) OpenTicketHandle {
	return &DefaultOpenTicket{
		appID:             appID,
		cache:             cache,
		cacheKeyPrefix:    cacheKeyPrefix,
		OpenAPITicketLock: new(sync.Mutex),
	}
}

// OpenTicket 请求open_tikcet返回结果
type OpenTicket struct {
	util.DYError
	Ticket    string `json:"ticket"`
	ExpiresIn int64  `json:"expires_in"`
}

type OpenTicketResp struct {
	Extra util.DYErrorExtra `json:"extra"`
	Data  Ticket            `json:"data"`
}

// GetTicket 获取openapi_ticket
func (open *DefaultOpenTicket) GetTicket(accessToken string) (ticketStr string, err error) {
	open.OpenAPITicketLock.Lock()
	defer open.OpenAPITicketLock.Unlock()

	//先从cache中取
	openAPITicketCacheKey := fmt.Sprintf("%s_openapi_ticket_%s", open.cacheKeyPrefix, open.appID)
	val := open.cache.Get(openAPITicketCacheKey)
	if val != nil {
		ticketStr = val.(string)
		return
	}
	var ticket Ticket
	ticket, err = open.GetTicketFromServer(accessToken)
	if err != nil {
		return
	}
	expires := ticket.ExpiresIn - 1500
	_ = open.cache.Set(openAPITicketCacheKey, ticket.Ticket, time.Duration(expires)*time.Second)

	ticketStr = ticket.Ticket
	return
}

// GetTicketFromServer 从服务器中获取ticket
func (open *DefaultOpenTicket) GetTicketFromServer(accessToken string) (ticket Ticket, err error) {
	var response []byte
	url := fmt.Sprintf(getOpenTicketURL, accessToken)
	response, err = util.HTTPGet(url)
	if err != nil {
		return
	}
	var ticketResp TicketResp
	err = json.Unmarshal(response, &ticketResp)
	if err != nil {
		return
	}

	ticket = ticketResp.Data
	if ticket.ErrCode != 0 {
		err = util.NewCodeDouYinError("GetOpenTicket", ticket.DYError, &ticketResp.Extra)
		return
	}
	return
}
