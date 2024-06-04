package login

import (
	"bytes"
	"crypto"
	"crypto/rsa"
	"crypto/sha256"
	"crypto/x509"
	"encoding/base64"
	"encoding/binary"
	"encoding/hex"
	"io/ioutil"
	"log"
	"net/http"
	"strconv"
	"sync"

	"github.com/CloudcadeSF/thirdparty-sdk/gamecenter/config"
)

var (
	keyLock      sync.Mutex
	publicKeyMap map[string][]byte
)

type GameCenter struct {
	PublicKeyURL           string
	Signature              string
	Timestamp              string
	PlayerId               string
	TeamPlayerId           string
	GamePlayerId           string
	Sha256HashPlayerID     string
	Sha256HashTeamPlayerID string
}

func NewGameCenterWithSha256String(pub, sig, sha256HashPlayerID, sha256HashTeamPlayerID string) *GameCenter {
	result := &GameCenter{
		PublicKeyURL:           pub,
		Signature:              sig,
		Sha256HashPlayerID:     sha256HashPlayerID,
		Sha256HashTeamPlayerID: sha256HashTeamPlayerID,
	}
	return result
}

func getPublicKeyFromRemote(url string) []byte {
	res, _ := http.Get(url)
	buff, _ := ioutil.ReadAll(res.Body)
	res.Body.Close()
	return buff
}

func getPublicKey(url string) []byte {
	keyLock.Lock()
	defer keyLock.Unlock()
	if v, ok := publicKeyMap[url]; ok {
		return v
	}
	tmpBuff := getPublicKeyFromRemote(url)
	publicKeyMap[url] = tmpBuff
	return tmpBuff
}

/**
客户端直接编码的验证
*/
func VerifyString(puk, data, sig string) error {
	realData, err1 := base64.StdEncoding.DecodeString(data)
	if err1 != nil {
		return err1
	}
	realSig, err2 := base64.StdEncoding.DecodeString(sig)
	if err2 != nil {
		return err2
	}
	publicKey := getPublicKey(puk)
	certificate, err3 := x509.ParseCertificate(publicKey)
	pub := certificate.PublicKey.(*rsa.PublicKey)
	err5 := rsa.VerifyPKCS1v15(pub, crypto.SHA256, realData, realSig)
	if err5 != nil {
		return err5
	} else {
		return nil
	}
	if err3 != nil {
		return err3
	}
	if err4 := certificate.CheckSignature(certificate.SignatureAlgorithm, realData, realSig); err4 == nil {
		return nil
	} else {
		return err4
	}
	return nil
}

/**
验证签名
*/
func Verify(sSig, sGcId, sSalt, sTimeStamp, publicUrl string) (status bool, userId string, err error) {
	sig, err := base64.StdEncoding.DecodeString(sSig)
	if err != nil {
		return
	}
	salt, err := base64.StdEncoding.DecodeString(sSalt)
	//salt = []byte(sSalt)
	if err != nil {
		return
	}
	timeStamp, err := strconv.ParseUint(sTimeStamp, 10, 64)
	if err != nil {
		return
	}
	payload := new(bytes.Buffer)
	payload.WriteString(sGcId)
	payload.WriteString(config.IOSBundleId)
	binary.Write(payload, binary.BigEndian, timeStamp)
	payload.Write(salt)
	log.Println("===========")
	cert := getPublicKey(publicUrl)
	if err := verifyRsa(cert, sig, payload.Bytes()); err != nil {
		return false, "", err
	}
	return true, sGcId, nil
}

/**
验证Rsa
*/
func verifyRsa(key, sig, content []byte) error {
	log.Println(base64.StdEncoding.EncodeToString(content))
	cert, err := x509.ParseCertificate(key)
	if err != nil {
		log.Printf("parse cert error %s", err)
		return err
	}
	pub := cert.PublicKey.(*rsa.PublicKey)

	h := sha256.New()
	h.Write(content)
	digest := h.Sum(nil)
	log.Println(base64.StdEncoding.EncodeToString(digest))
	hashCode := hex.EncodeToString(digest)
	log.Println(hashCode)
	err = rsa.VerifyPKCS1v15(pub, crypto.SHA256, digest, sig)
	return err
}

func init() {
	publicKeyMap = make(map[string][]byte)
}
