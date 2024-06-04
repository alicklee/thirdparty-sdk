package login

import (
	"github.com/CloudcadeSF/thirdparty-sdk/utils/log"
	"testing"
)

func TestVerifyString(t *testing.T) {

	lt := NewLeiTing("eov3fmox", "44", "310001", "8fc5fd52ebfdb6dc5625aca5707d0f75", "uuidxxx")
	isSuccess, uid, err := lt.Verity()

	if err != nil {
		log.Infoln(isSuccess, uid, err)
		t.Fail()
	}
}

func TestLeiTing_IsSkipVerify(t *testing.T) {
	lt := NewLeiTing("", "44", "310001", "8fc5fd52ebfdb6dc5625aca5707d0f75", "eov3fmo2x")

	if !lt.IsSkipVerify() {
		log.Infof("IsSkipVerify false")
		t.Fail()
	}
}
