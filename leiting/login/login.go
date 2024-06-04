package login

import (
	"encoding/json"
	"errors"
	"fmt"
	"github.com/CloudcadeSF/thirdparty-sdk/leiting/config"
	"github.com/CloudcadeSF/thirdparty-sdk/utils/log"
	"io/ioutil"
	"net/http"
	"net/url"
	"os"
	"strings"
)

// "https://graph.facebook.com/debug_token"
var verifyUrl string

func getUrlHost() string {
	if verifyUrl != "" {
		return verifyUrl
	}
	verifyUrl = os.Getenv("verifyUrlLeiTing")
	if verifyUrl == "" {
		return config.LeiTingUrlHostDefault
	}
	return verifyUrl
}

/**
登陆结构体
*/
type LeiTing struct {
	UserId    string //  用户ID
	Game      string //  游戏标识
	ChannelNo string //  渠道编号
	Token     string //  登录验证token信息
	UUID      string // 用于跳过验证
}

/**
验证成功时返回的结构体
*/
type response struct {
	Message string `json:"message"`
	Status  string `json:"status"` // 状态码 1:成功 其他失败
}

/**
验证失败时候的结构体
*/
type errorResponse struct {
	Error jsonErrorMessage
}

/**
返回失败的json数据结构
@deprecated
*/
type jsonErrorMessage struct {
	Message string `json:"message"`
	Type    string `json:"type"`
	Code    int32  `json:"code"`
}

/**
返回成功的结构体
*/
type jsonResponse struct {
	Message string `json:"message"`
	Status  string `json:"status"` // 状态码 1:成功 其他失败
}

func NewLeiTing(userId, game, channelNo, token, uuid string) *LeiTing {
	result := &LeiTing{
		UserId:    userId,
		Game:      game,
		ChannelNo: channelNo,
		Token:     token,
		UUID:      uuid,
	}
	return result
}

/**
登陆验证主函数
@Return
bool 是否通过验证
error 没有通过验证的错误信息
*/
func (f *LeiTing) Verity() (bool, string, error) {
	if f.IsSkipVerify() {
		return true, f.UUID, nil
	}

	urlPath := f.getVerifyUrl()
	resp, err := http.Get(urlPath)
	//判断是否可以联通api地址
	if resp == nil {
		fmt.Println("http.Get  resp == nil")
		return false, "", errors.New("Can't  connection with the LeiTing api")
	}
	if err != nil {
		fmt.Println("http.Get  err != nil")
		return false, "", err
	}
	body, _ := ioutil.ReadAll(resp.Body)
	var r response
	_ = json.Unmarshal(body, &r)
	//如果没有数据返回则说明验证失败尝试获取失败信息
	if &r == nil {
		var er errorResponse
		_ = json.Unmarshal(body, &er)
		if &er != nil {
			return false, "", errors.New(er.Error.Message)
		}
	}
	defer resp.Body.Close()
	status, userId := f.verifyResponseData(&r)
	return status, userId, nil

}

/**
curl --location --request GET 'https://sdkoverseas.leiting.com/login/verifyLoginToken.do?userId=nnzm1jxh&game=441&channelNo=310001&token=b184ca3bee8ae5d140b49a579ec375cc' \
--header 'Cookie: JSESSIONID=19B5DBCA3E48AE7449C3D3D63597CB6F'
*/
func (f *LeiTing) getVerifyUrl() string {
	params := url.Values{}
	Url, err := url.Parse(getUrlHost())
	if err != nil {
		return ""
	}
	//var accessToken = config.AppId + "%7C" + config.FaceBookToken
	params.Set("userId", f.UserId)
	params.Set("game", f.Game)
	params.Set("channelNo", f.ChannelNo)
	params.Set("token", f.Token)
	//如果参数中有中文参数,这个方法会进行URLEncode
	Url.RawQuery = params.Encode()
	urlPath := Url.String()
	fmt.Println(urlPath)
	return urlPath
}

/**
验证返回的数据结果和用户信息是否对的上
*/
func (f *LeiTing) verifyResponseData(data *response) (bool, string) {
	//验证是否通过了验证
	if data.Status != config.LeiTingSuccessStatus {
		return false, f.UserId
	}
	return true, f.UserId
}

// IsSkipVerify 是否可跳过验证
func (f *LeiTing) IsSkipVerify() bool {
	userId := f.UUID
	uidList := strings.Split(os.Getenv("SKIP_LT_UID_LIST"), ",")
	log.Infof("uidList %+v", uidList)
	for _, uid := range uidList {
		// 匹配到需要跳过的 uid
		if uid == userId {
			log.Infof("uid == userId %+v %+v", uid, userId)
			return true
		}
	}
	return false
}
