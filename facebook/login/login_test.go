package login

import (
	"fmt"
	"testing"

	"github.com/CloudcadeSF/thirdparty-sdk/utils/log"
)

func TestFaceBook_Verity(t *testing.T) {
	//idToken := "EAAEX2VqzksABADx9BBLovudOZBmcnQr6MyweKABkT4E282jTuY6ZAnZAhczKTSGxRiYLOUKiaCpIDC1zTpHv93oeQJZBRbOAxOm4HRMOBCi3Yzr1gZC8m0cG9NXkZCExjd4JLsQcddc5lb5Yfpg9O9kM2UjxBtRdbYOCZAxdtPjIQeoqrt8wbLZBzAjgKjxHZAINtSj6iZCxGNnHllZCYjEWFhsQOENr9Dfp4vHv29KEmt0YTiB7TsOBTkhWaRqYwwLdWoZD"
	idToken := "EAAEoA1cAofkBAFKLDO3ZCuqSkZCHe88t8JqZAnaFvrO5yzSkeJd3FtgBi4MP51irP4hRXolwYkHysACZBwfKqG7ABe5ZBWo99w3skIfZCztdvcN2LsLClJIZBtHjJScr23Tmx3uneyihvCZBKrzBED20ZBC41txVUqQ6TwdK5qyNZALnqytX5IrqZALn6N5fFU9lQD2FtvkhgP49lk6BlIQi2IcsZBnwmalcPwThu5MPlbbxkSZBnSCCZC3FPbz6n29pYvZAGEZD"
	facebook := NewFaceBook(idToken)
	verity, s, err := facebook.Verity()
	if !verity {
		fmt.Println(err)
		t.Fail()
	}
	log.Info(s)
}
