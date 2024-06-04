package login

import (
	"errors"
	"fmt"
	"github.com/CloudcadeSF/thirdparty-sdk/googleplay/config"
	"net/http"
	"os"

	"golang.org/x/net/context"

	"google.golang.org/api/oauth2/v2"
	"google.golang.org/api/option"
)

type GooglePlay struct {
	IdToken string
}

var httpClient = &http.Client{}

func NewGooglePlay(idToken string) *GooglePlay {
	result := &GooglePlay{IdToken: idToken}
	return result
}

func (g *GooglePlay) Verity() (*oauth2.Tokeninfo, error) {
	ctx := context.Background()
	opt := option.WithCredentialsFile(config.GooglePlayKeyPath)
	service, _ := oauth2.NewService(ctx, opt)
	tokenInfoCall := service.Tokeninfo()
	tokenInfoCall.IdToken(g.IdToken)
	tokenInfo, err := tokenInfoCall.Context(ctx).Do()
	if err != nil {
		return nil, err
	}
	return tokenInfo, nil
}
func (g *GooglePlay) Verity2() (bool, string, error) {
	ctx := context.Background()
	token := g.IdToken
	svc, err := oauth2.New(http.DefaultClient)
	fmt.Println(svc.Userinfo)
	ti, err := svc.Tokeninfo().IdToken(token).Context(ctx).Do()
	if err != nil {

		return false, "", err
	}
	googlePlayClientId := os.Getenv("GooglePlayClientId")
	if ti.Audience != googlePlayClientId {
		return false, "", errors.New("not from our app. ClientId:" + googlePlayClientId)
	}
	if !ti.VerifiedEmail {
		return false, "", errors.New("tokeninfo: email address not verified")
	}
	return true, ti.UserId, nil
}
