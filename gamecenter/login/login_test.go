package login

import (
	"testing"

	"github.com/CloudcadeSF/thirdparty-sdk/utils/log"
)

func TestVerifyString(t *testing.T) {
	publicUrl := "https://static.gc.apple.com/public-key/gc-prod-5.cer"
	data := "m5Tq/YfFp2fc0gj2pOtMeqS2miK3xZ5rZ5Tjh7tA6K0="
	sig := "p0UGDX4UZvZeLeZVs64Lfo8lJJjzegj11mPYQKUusG2guzGBhWXEIjBtjb6v2lukLkHvFNUQB1dxmjdmjdUS4hVabIPDKkQ6UgHMX2HZBD32mPBSE8MdcmAn4nzsNwt6HiOjiEj0/ElLjpT9ucHFYwfkfXNnobVIw8DdSRNLjTFdtkc4wtlThS5s/833fV18XGLDFW4yNyhrO0ZPBwC7fIQfhmYVtbCPSHsLa0SzDJGo64qsgCwSLEwNr1ejRmD2h+tO+i3zmCxVLBFdvaKcqr+vDTAmI1BfhK2kzn6qv9XVMr5QmSIJqtLM1QdzLFEyuPlS0Qh/CKUGu+9nDHnMng=="

	err := VerifyString(publicUrl, data, sig)

	if err != nil {
		log.Infoln(err)
		t.Fail()
	}
}

func TestVerifySig(t *testing.T) {
	publicUrl := "https://static.gc.apple.com/public-key/gc-prod-5.cer"
	cid := "G:594478901"
	sig := "YKbfmtB1wdRqgIPmfRFSZ7fdL6nqkIBxI2/896pqqf2UykLfrgT90rerKnEv8VXanVcSNQFnin6vmf+rTNdMvC/7DN41SNgC4AbwSPixCSHoSgEyIPSfD3KoHzp1TW0I/xVBLmnrWjHAyjOsytS6lg/dO4c0Yxu4Yv2QKBlOOhn5ui2gwj27ynXlqAuqqK7/sJraXQlBMK3WEDZuAxyc9yFx62obJFM66XBkf4Gd+9V2wBCDSyHZ3pwv5iaoBvEWPCp7HJMH3+zx6NB6IQdlMZ0PiutP9u00iD8WDrZRRFIp7dAlbIKtN8XlRwHqu67pCEdVpKOT6IBlJwfosH4btA=="
	sig = "p0UGDX4UZvZeLeZVs64Lfo8lJJjzegj11mPYQKUusG2guzGBhWXEIjBtjb6v2lukLkHvFNUQB1dxmjdmjdUS4hVabIPDKkQ6UgHMX2HZBD32mPBSE8MdcmAn4nzsNwt6HiOjiEj0/ElLjpT9ucHFYwfkfXNnobVIw8DdSRNLjTFdtkc4wtlThS5s/833fV18XGLDFW4yNyhrO0ZPBwC7fIQfhmYVtbCPSHsLa0SzDJGo64qsgCwSLEwNr1ejRmD2h+tO+i3zmCxVLBFdvaKcqr+vDTAmI1BfhK2kzn6qv9XVMr5QmSIJqtLM1QdzLFEyuPlS0Qh/CKUGu+9nDHnMng=="
	salt := "jlP8xw=="
	timestamp := "1607425879142"
	verifySig, id, err := Verify(sig, cid, salt, timestamp, publicUrl)
	if err != nil {
		t.Fail()
		log.Info(err)
	}
	log.Info(verifySig)
	log.Info(id)
}
