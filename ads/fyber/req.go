package fyber

import (
	"crypto/sha1"
	"fmt"
	"reflect"
	"sort"

	"github.com/CloudcadeSF/thirdparty-sdk/utils/log"
)

type RewardNotify struct {
	UID         string `form:"uid"`         //玩家ID
	SID         string `form:"sid"`         //校验字段
	Amount      string `form:"amount"`      //指定价
	PlacementId string `form:"placementId"` //广告位信息

	//custom_param
	RewardType string `form:"rewardType"`
	//-- for eventSpeedUP
	EventType string `form:"evenType"`
	TargetID  string `form:"targetID"`
}

// IsValid 判断是否为有效的数据
func (rn *RewardNotify) IsValid(token string) bool {
	type param struct {
		n string
		v string
	}

	var params []param

	fVal := reflect.ValueOf(*rn)
	fTyp := fVal.Type()

	for i := 0; i < fVal.NumField(); i++ {
		filedTyp := fTyp.Field(i)
		filedVal := fVal.Field(i)

		name := filedTyp.Tag.Get("form")

		if name == "sid" {
			continue
		}

		value := filedVal.String()

		if len(value) == 0 {
			continue
		}

		params = append(params, param{
			n: name,
			v: value,
		})
	}

	sort.Slice(params, func(i, j int) bool {
		return params[i].n[0] < params[j].n[0]
	})

	log.Infof("params %+v", params)

	var sum string

	for _, param := range params {
		sum += param.v
	}

	sum += token

	calSha1 := sha1.New()
	calSha1.Write([]byte(sum))

	log.Infof("cal %v get %v", fmt.Sprintf("%x", calSha1.Sum(nil)), rn.SID)

	return fmt.Sprintf("%x", calSha1.Sum(nil)) == rn.SID
}
