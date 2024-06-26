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
const getTicketURL = "https://open.douyin.com/js/getticket?access_token=%s"

// DefaultJsTicket 默认获取js ticket方法
type DefaultJsTicket struct {
	appID          string
	cacheKeyPrefix string
	cache          cache.Cache
	//jsAPITicket 读写锁 同一个AppID一个
	jsAPITicketLock *sync.Mutex
}

// NewDefaultJsTicket new
func NewDefaultJsTicket(appID string, cacheKeyPrefix string, cache cache.Cache) JsTicketHandle {
	return &DefaultJsTicket{
		appID:           appID,
		cache:           cache,
		cacheKeyPrefix:  cacheKeyPrefix,
		jsAPITicketLock: new(sync.Mutex),
	}
}

// Ticket 请求jsapi_tikcet返回结果
type Ticket struct {
	util.DYError
	Ticket    string `json:"ticket"`
	ExpiresIn int64  `json:"expires_in"`
}

type TicketResp struct {
	Extra util.DYErrorExtra `json:"extra"`
	Data  Ticket            `json:"data"`
}

// GetTicket 获取jsapi_ticket
func (js *DefaultJsTicket) GetTicket(accessToken string) (ticketStr string, err error) {
	js.jsAPITicketLock.Lock()
	defer js.jsAPITicketLock.Unlock()

	//先从cache中取
	jsAPITicketCacheKey := fmt.Sprintf("%s_jsapi_ticket_%s", js.cacheKeyPrefix, js.appID)
	val := js.cache.Get(jsAPITicketCacheKey)
	if val != nil {
		ticketStr = val.(string)
		return
	}
	var ticket Ticket
	ticket, err = GetTicketFromServer(accessToken)
	if err != nil {
		return
	}
	expires := ticket.ExpiresIn - 1500
	_ = js.cache.Set(jsAPITicketCacheKey, ticket.Ticket, time.Duration(expires)*time.Second)

	ticketStr = ticket.Ticket
	return
}

// GetTicketFromServer 从服务器中获取ticket
func GetTicketFromServer(accessToken string) (ticket Ticket, err error) {
	var response []byte
	url := fmt.Sprintf(getTicketURL, accessToken)
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
		err = util.NewCodeDouYinError("GetTicket", ticket.DYError, &ticketResp.Extra)
		return
	}
	return
}
