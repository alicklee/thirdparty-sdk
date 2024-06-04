package firebase

import (
	"firebase.google.com/go/messaging"
	"fmt"
	"github.com/stretchr/testify/assert"
	"testing"
)

var iosToken = "fzAzKB8QEEBKnhgsoWhWcq:APA91bH8_1v3jrqRWmoLRIhScjyIxDiI6lESrMG2_0iIc3XaKPi9yIt9U_NsK6MLhVikYhWmeRRs-rUfc-ysv7hjdo0eX524xAXPj4j0bI0R-lWgdh1zPDNskA67JiAYE_a4TBdyY18M"
var android = "e7QZnAdFSwm9f9vrq3btPd:APA91bE82Pv-_UoHvwUEc8QcVWkTuR7iaHFheZyiuB8ZL1Gv15x9oP3cudefI-tow1KsUOglRql7bsL6HdMI90Y-WvsSOk6s3oWGJDp7ys-7j2Qa9yrFOH1PeuzZ_f-u2R3beLCxRvRF"
var topic = "en."

/*
;
*
键值内容推送测试
*/
func TestKeyArgsNotification_Push(t *testing.T) {
	ts := make([]string, 0)
	ts = append(ts, iosToken)
	//topics := []string{"en", "US", "IOS"}
	message := NewNotification("", "", []string{}, ts)
	//AddUserToTopic(topics, ts)

	message.TitleLocKey = KeyArgs{
		Key: "UI_crafting_title",
		//Args: []string{"123", "456"},
	}

	message.BodyLocKey = KeyArgs{
		Key: "UI_crafting_slots",
		//Args: []string{"999", "fufufufufu"},
	}

	message.Push()
}

/*
*
Android Token的Notification的测试
*/
func TestTokenNotification_Push(t *testing.T) {
	ts := make([]string, 0)
	ts = append(ts, iosToken)
	message := NewNotification("This is IOS title", "This Ios 789", nil, ts)
	message.Push()
}

/*
*
Android Token的Notification的推送测试
*/
func TestAndroidNotification_Push(t *testing.T) {
	ts := make([]string, 0)
	ts = append(ts, android)
	message := NewNotification("This is Android title", "This Android fda", nil, ts)
	message.Push()
}

func TestMessage(t *testing.T) {
	msg := "{\"notification\":{},\"android\":{\"ttl\":\"3600s\",\"notification\":{\"icon\":\"stock_ticker_update\",\"color\":\"#f45342\",\"sound\":\"s\",\"body_loc_key\":\"shl_ntf_craft_msg\",\"title_loc_key\":\"shl_ntf_craft_title\"}},\"apns\":{\"payload\":{\"aps\":{\"alert\":{\"loc-key\":\"shl_ntf_craft_msg\",\"title-loc-key\":\"shl_ntf_craft_title\"},\"badge\":1,\"sound\":\"s\"}}},\"token\":\"fiTSqICCKEL3lUp7m2pj8J:APA91bHPKxxqYbsEc316PW_JH9YnYsfIjcNPn0Eiz5cKyiMD2ZO60zDUcxso0ayqrZR65T5db4uPGYfPH0zzG86tMU2KAwkm-kCLrwaoED-EWGhY2388kAW8bT5s0QtwovncaTZKekDL\"}"
	m := messaging.Message{}
	if err := m.UnmarshalJSON([]byte(msg)); err != nil {
		assert.NoError(t, err)
		return
	}

	client, err := firebaseApp.Messaging(cxt)
	if err != nil {
		assert.NoError(t, err)
	}

	send, err := client.Send(cxt, &m)
	if err != nil {
		assert.NoError(t, err)
	} else {
		fmt.Println(send)
	}
}

///*
//*
//Topic 的Notification的推送测试
//*/
//func TestTopicNotification_Push(t *testing.T) {
//	ts := make([]string, 0)
//	ts = append(ts, android)
//	en := NewNotification(topic, "this is topic title", "this is topic body", nil)
//	en.AddUserToTopic()
//	en.Push()
//}
//
//func TestMPush(t *testing.T) {
//	ts := make([]string, 0)
//	ts = append(ts, android)
//	message := NewNotification(topic, "This is topic title", "This topic body6666666", ts)
//	message.PushMulticast()
//}
