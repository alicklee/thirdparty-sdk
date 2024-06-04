package iap

import (
	"context"
	"github.com/CloudcadeSF/shop-heroes-legends-common/project"
	"log"

	"google.golang.org/api/androidpublisher/v3"

	"io/ioutil"

	"github.com/awa/go-iap/playstore"
)

var jsonKey []byte

func Init() error {
	// You need to prepare a public key for your Android app's in app billing
	// at https://console.developers.google.com.
	var err error
	jsonKey, err = ioutil.ReadFile(getGoogleKeyFilename())
	if err != nil {
		log.Fatal(err)
	}
	return nil
}

func getGoogleKeyFilename() string {
	filename := ""
	if project.IsProjectSHL {
		filename = "./key/googleplay-iap.json"
	} else if project.IsProjectCB {
		filename = "./key/googleplay-iap-cb.json"
	} else if project.IsProjectGH {
		filename = "./key/googleplay-iap-gh.json"
	}
	return filename
}

func VerifyPlayStoreSubscription(packageName string, subscriptionID string, purchaseToken string) (*androidpublisher.SubscriptionPurchase, error) {
	client, err := playstore.New(jsonKey)
	if err != nil {
		return nil, err
	}
	ctx := context.Background()
	resp, err := client.VerifySubscription(ctx, packageName, subscriptionID, purchaseToken)
	return resp, err
}

func VerifyPlayStoreProduct(packageName string, productID string, purchaseToken string) (*androidpublisher.ProductPurchase, error) {
	client, err := playstore.New(jsonKey)
	if err != nil {
		return nil, err
	}
	ctx := context.Background()
	resp, err := client.VerifyProduct(ctx, packageName, productID, purchaseToken)
	return resp, err
}

func AcknowledgePlayStoreProduct(packageName string, productID string, purchaseToken string, developerPayload string) error {
	client, err := playstore.New(jsonKey)
	if err != nil {
		return err
	}
	ctx := context.Background()
	err = client.AcknowledgeProduct(ctx, packageName, productID, purchaseToken, developerPayload)
	return err
}
