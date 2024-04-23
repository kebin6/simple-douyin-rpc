package share

import (
	"encoding/json"
	"fmt"
	"github.com/kebin6/simple-douyin-rpc/open/cache"
	"github.com/kebin6/simple-douyin-rpc/open/util"
	"sync"
)

const (
	getShareIdURL = "https://open.douyin.com/share-id?access-token=%s"
)

type DefaultShareId struct {
	ClientToken string
	cache       cache.Cache
	shareIdLock *sync.Mutex
}

func NewDefaultShareId(clientToken string, cache cache.Cache) ShareIdHandle {
	if cache == nil {
		panic("cache is need")
	}

	return &DefaultShareId{
		ClientToken: clientToken,
		cache:       cache,
		shareIdLock: new(sync.Mutex),
	}
}

// ShareId struct.
type ShareId struct {
	util.DYError

	ShareId string `json:"share_id"`
}

type ShareIdResp struct {
	Extra util.DYErrorExtra `json:"extra"`
	Data  ShareId           `json:"data"`
}

func (share *DefaultShareId) GetShareId() (shareIdStr string, err error) {
	share.shareIdLock.Lock()
	defer share.shareIdLock.Unlock()

	var response []byte
	url := fmt.Sprintf(getShareIdURL, share.ClientToken)
	response, err = util.HTTPGet(url)
	if err != nil {
		return
	}
	var shareIdResp ShareIdResp
	err = json.Unmarshal(response, &shareIdResp)
	if err != nil {
		return
	}

	shareIdStr = shareIdResp.Data.ShareId
	if shareIdResp.Data.ErrCode != 0 {
		err = util.NewCodeDouYinError("GetTicket", shareIdResp.Data.DYError, &shareIdResp.Extra)
		return
	}
	return
}
