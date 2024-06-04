package iap

import (
	"context"

	"github.com/awa/go-iap/appstore"
)

func VerifyAppStoreSubscription(receiptData string, password string) (*appstore.IAPResponse, error) {
	client := appstore.New()
	req := appstore.IAPRequest{
		ReceiptData: receiptData,
		Password:    password,
	}
	resp := &appstore.IAPResponse{}
	ctx := context.Background()
	err := client.Verify(ctx, req, resp)
	return resp, err
}

func VerifyAppStoreProduct(receiptData string) (*appstore.IAPResponse, error) {
	client := appstore.New()
	req := appstore.IAPRequest{
		ReceiptData: receiptData,
	}
	resp := &appstore.IAPResponse{}
	ctx := context.Background()
	err := client.Verify(ctx, req, resp)
	return resp, err
}
