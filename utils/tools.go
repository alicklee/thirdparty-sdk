package utils

import (
	"fmt"
	"os"
	"path"

	"github.com/joho/godotenv"
)

// 获取环境变量
func GetenvDefault(envKey, defaultValue string) string {
	env := defaultValue
	if val, ok := os.LookupEnv(envKey); ok {
		env = val
	}
	fmt.Println(envKey, env, defaultValue)
	return env
}

//loadEnv 初始化环境变量
func LoadEnv() {
	if _, ok := os.LookupEnv("ENV_DOCKER"); ok {
		return
	}
	gopath := os.ExpandEnv("$GOPATH")
	godotenv.Load(path.Join(gopath, "src", "github.com/CloudcadeSF/thirdparty-sdk", ".env"))
	// ENV=test   测试时使用
	////if _, ok := os.LookupEnv("ENV_TEST"); ok {
	////	gopath := os.ExpandEnv("$GOPATH")
	////	godotenv.Load(path.Join(gopath, "src", "shop-heroes-legends-user-center", ".env"))
	////	return
	////}
	//if _, ok := os.LookupEnv("ENV_DOCKER"); ok {
	//	return
	//}
	//// APP_ENV=
	//if err := godotenv.Load(); err != nil {
	//	log.Panic("Error loading .env file")
	//}
}
