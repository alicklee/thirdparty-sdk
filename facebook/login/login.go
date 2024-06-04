package login

import (
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"

	"github.com/CloudcadeSF/thirdparty-sdk/facebook/config"
)

const facebookVerifyUrl = "https://graph.facebook.com/debug_token"

/**
Facebook 登陆结构体
*/
type FaceBook struct {
	IdToken string
}

/**
Facebook验证成功时返回的结构体
*/
type facebookResponse struct {
	Data jsonResponse `json:"data"`
}

/**
Facebook验证失败时候的结构体
*/
type facebookErrorResponse struct {
	Error jsonErrorMessage
}

/**
返回失败的json数据结构
*/
type jsonErrorMessage struct {
	Message string `json:"message"`
	Type    string `json:"type"`
	Code    int32  `json:"code"`
	SubCode int32  `json:"error_subcode"`
	TraceId string `json:"fbtrace_id"`
}

/**
返回成功的结构体
*/
type jsonResponse struct {
	AppId     string `json:"app_id"`
	ExpiresAt int32  `json:"expires_at"`
	IsValid   bool   `json:"is_valid"`
	IssuedAt  int32  `json:"issued_at"`
	UserId    string `json:"user_id"`
}

func NewFaceBook(idToken string) *FaceBook {
	result := &FaceBook{IdToken: idToken}
	return result
}

/**
Facebook登陆验证主函数
@Return
bool 是否通过验证
error 没有通过验证的错误信息
*/
func (f *FaceBook) Verity() (bool, string, error) {
	urlPath := getVerifyUrl(f.IdToken)
	resp, err := http.Get(urlPath)
	//判断是否可以联通过facebook的api地址
	if resp == nil {
		return false, "", errors.New("Can't  connection with the facebook api")
	}
	if err != nil {
		return false, "", err
	}
	body, _ := ioutil.ReadAll(resp.Body)
	var r facebookResponse
	_ = json.Unmarshal(body, &r)
	//如果没有数据返回则说明验证失败尝试获取失败信息
	if &r == nil {
		var er facebookErrorResponse
		json.Unmarshal(body, &er)
		if &er != nil {
			return false, "", errors.New(er.Error.Message)
		}
	}
	defer resp.Body.Close()
	status, userId := verifyResponseData(&r)
	return status, userId, nil

}

/**
获取Facebook的验证URL
*/
func getVerifyUrl(inputToken string) string {
	params := url.Values{}
	Url, err := url.Parse(facebookVerifyUrl)
	if err != nil {
		return ""
	}
	var accessToken = config.GetFBAppId() + "%7C" + config.GetFBToken()
	params.Set("input_token", inputToken)
	params.Set("access_token", accessToken)
	//如果参数中有中文参数,这个方法会进行URLEncode
	Url.RawQuery = params.Encode()
	urlPath := Url.String()
	fmt.Println(urlPath)
	return urlPath
}

/**
验证返回的数据结果和用户信息是否对的上
*/
func verifyResponseData(data *facebookResponse) (bool, string) {
	//验证是否来自指定App的登陆
	if data.Data.AppId != config.GetFBAppId() {
		return false, data.Data.UserId
	}
	//验证是否通过了facebook的验证
	if data.Data.IsValid != true {
		return false, data.Data.UserId
	}
	return true, data.Data.UserId
}
