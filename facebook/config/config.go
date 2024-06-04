package config

import (
	"github.com/CloudcadeSF/shop-heroes-legends-common/project"
	"os"
)

const faceBookTokenSHL = "b61bfba5c8771f3ba43fa49e36bde333"
const appIdSHL = "307697273770688"
const faceBookTokenCB = "deca31a25c09aeb3be3fec3f99458bd0"
const appIdCB = "325469786382841"
const faceBookTokenGH = "1f687776835df70cd8c313ac130f93f5"
const appIdGH = "1109469862977299"
const FacebookVerifyUrl = "https://graph.facebook.com/debug_token"

func GetFBToken() string {
	env := os.Getenv("FacebookAppToken")
	if env != "" {
		return env
	}

	if project.IsProjectSHL {
		return faceBookTokenSHL
	}
	if project.IsProjectCB {
		return faceBookTokenCB
	}
	if project.IsProjectGH {
		return faceBookTokenGH
	}
	return ""
}

func GetFBAppId() string {
	env := os.Getenv("FacebookAppId")
	if env != "" {
		return env
	}

	if project.IsProjectSHL {
		return appIdSHL
	}
	if project.IsProjectCB {
		return appIdCB
	}
	if project.IsProjectGH {
		return appIdGH
	}
	return ""
}
