package tapjoy

import (
	"crypto/md5"
	"encoding/hex"
	"fmt"
)

type RewardNotify struct {
	ID          string `form:"id"`
	SNUid       string `form:"snuid"`
	Currency    string `form:"currency"`
	MacAddress  string `form:"mac_address"`
	Verifier    string `form:"verifier"`
	Application string `form:"application"`
	Rev         string `form:"rev"`
	StoreId     string `form:"storeid"`
}

func (rn *RewardNotify) IsValid(token string) bool {
	verifier := fmt.Sprintf("%v:%v:%v:%v", rn.ID, rn.SNUid, rn.Currency, token)
	md5Result := md5.Sum([]byte(verifier))
	return hex.EncodeToString(md5Result[:]) == rn.Verifier
}
