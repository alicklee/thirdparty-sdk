package firebase

import (
	"context"
	"errors"
	"fmt"
	"time"

	firebase "firebase.google.com/go"
	"firebase.google.com/go/messaging"
	"github.com/CloudcadeSF/thirdparty-sdk/firebase/config"
	"github.com/CloudcadeSF/thirdparty-sdk/utils/log"
	"google.golang.org/api/option"
)

var firebaseApp *firebase.App
var cxt = context.Background()

type Notification struct {
	Topics []string
	Tokens []string

	Title       string
	Body        string
	TitleLocKey KeyArgs
	BodyLocKey  KeyArgs
}
type KeyArgs struct {
	Key  string
	Args []string
}

/*
*
key的生成地址为
https://console.firebase.google.com/u/1/project/shop-heroes-legends-45168147/settings/serviceaccounts/adminsdk
*/
func init() {
	opt := option.WithCredentialsJSON([]byte(config.FirebaseKeyData2))
	app, err := firebase.NewApp(cxt, nil, opt)
	if err != nil {
		log.Errorln()
	} else {
		firebaseApp = app
	}
}

func NewNotification(title, body string, topics, tokens []string) *Notification {
	result := &Notification{
		Topics: topics,
		Tokens: tokens,
		Title:  title,
		Body:   body,
	}
	return result
}

/*
*
消息推送的统一方法
*/
func (n *Notification) Push() error {
	message := &messaging.Message{
		Notification: &messaging.Notification{
			Title: n.Title,
			Body:  n.Body,
		},
		Android: &messaging.AndroidConfig{
			TTL: &config.AndroidTTL,
			Notification: &messaging.AndroidNotification{
				Sound:        "s",
				Icon:         config.Icon,
				Color:        config.Color,
				TitleLocKey:  n.TitleLocKey.Key,
				TitleLocArgs: n.TitleLocKey.Args,
				BodyLocKey:   n.BodyLocKey.Key,
				BodyLocArgs:  n.BodyLocKey.Args,
			},
		},
		APNS: &messaging.APNSConfig{
			Payload: &messaging.APNSPayload{
				Aps: &messaging.Aps{
					Badge: &config.Badge,
					Sound: "s",
					Alert: &messaging.ApsAlert{
						TitleLocKey:  n.TitleLocKey.Key,
						TitleLocArgs: n.TitleLocKey.Args,
						LocKey:       n.BodyLocKey.Key,
						LocArgs:      n.BodyLocKey.Args,
					},
				},
			},
		},
	}

	var cond Condition
	if len(n.Tokens) > 0 {
		message.Token = n.Tokens[0]
	}
	for _, t := range n.Topics {
		cond = cond.Add(t)
	}
	message.Condition = cond.String()

	client, err := firebaseApp.Messaging(cxt)
	if err != nil {
		log.Errorln(err)
		return err
	}

	log.Infof("MSG :%+v", message)

	ctxTimeout, cancel := context.WithTimeout(cxt, 10*time.Second)
	defer cancel()

	send, err := client.Send(ctxTimeout, message)
	if err != nil {
		log.Errorln(err)
		return err
	} else {
		log.Infoln(send)
	}

	return nil
}

// https://firebase.google.com/docs/cloud-messaging/send-message#go_4
type Condition string

func (c Condition) String() string {
	return string(c)
}
func (c Condition) Add(topic string) Condition {
	f := fmt.Sprintf("'%v' in topics", topic)
	var result string
	if len(c) > 0 {
		result = fmt.Sprintf("%v && %v", c, f)
	} else {
		result = f
	}

	return Condition(result)
}

/*
*
消息多用户推送的方法入口
*/
func (n *Notification) PushMulticast() error {
	if len(n.Tokens) > 500 {
		return errors.New("too much token ")
	}
	registrationTokens := n.Tokens
	message := &messaging.MulticastMessage{
		Notification: &messaging.Notification{
			Title: n.Title,
			Body:  n.Body,
		},
		Android: &messaging.AndroidConfig{
			TTL: &config.AndroidTTL,
			Notification: &messaging.AndroidNotification{
				Icon:  config.Icon,
				Color: config.Color,
			},
		},
		APNS: &messaging.APNSConfig{
			Payload: &messaging.APNSPayload{
				Aps: &messaging.Aps{
					Badge: &config.Badge,
				},
			},
		},
		Tokens: registrationTokens,
	}
	client, _ := firebaseApp.Messaging(cxt)
	br, err := client.SendMulticast(cxt, message)
	if err != nil {
		log.Errorln(err)
		return err
	}
	fmt.Printf("%d messages were sent successfully\n", br.SuccessCount)
	return nil
}

func DelUserFromMutilTopic(token string, topics ...string) error {
	client, _ := firebaseApp.Messaging(cxt)
	for _, topic := range topics {
		response, err := client.UnsubscribeFromTopic(cxt, []string{token}, topic)
		if err != nil {
			log.Errorln(err)
			return err
		}
		fmt.Println(response.SuccessCount, "tokens were subscribed successfully")
	}
	return nil
}

func AddUserToMutilTopic(token string, topics ...string) error {
	client, _ := firebaseApp.Messaging(cxt)
	for _, topic := range topics {
		response, err := client.SubscribeToTopic(cxt, []string{token}, topic)
		if err != nil {
			log.Errorln(err)
			return err
		}
		fmt.Println(response.SuccessCount, "tokens were subscribed successfully")
	}
	return nil
}

func AddUserToTopic(topic string, token []string) error {
	client, _ := firebaseApp.Messaging(cxt)
	response, err := client.SubscribeToTopic(cxt, token, topic)
	if err != nil {
		log.Errorln(err)
		return err
	}
	fmt.Println(response.SuccessCount, "tokens were subscribed successfully")
	return nil
}

func DeluserToTopic(topic string, token []string) error {
	client, _ := firebaseApp.Messaging(cxt)
	response, err := client.UnsubscribeFromTopic(cxt, token, topic)
	if err != nil {
		log.Errorln(err)
		return err
	}
	fmt.Println(response.SuccessCount, "tokens were unsubscribed successfully")
	return nil
}
