package btp

import (
	"context"
	"github.com/nnicora/sap-sdk-go/sap/oauth2"
	"github.com/nnicora/sap-sdk-go/service/btp/btpentitlements"
	"github.com/nnicora/sap-sdk-go/service/btp/btpglobalaccounts"
	"github.com/nnicora/sap-sdk-go/service/btp/btpsubaccounts"
)

type Client struct {
	config *Config

	GlobalAccounts *btpglobalaccounts.GlobalAccounts
	SubAccounts    *btpsubaccounts.SubAccounts
	Entitlements   *btpentitlements.Entitlements
}

type Config struct {
	AccountsHost     string
	EntitlementsHost string

	Auth2Config *oauth2.Config
}

func NewBtpClient(ctx context.Context, cfg *Config) (*Client, error) {
	http, err := oauth2.NewOAuth2ClientWithContext(ctx, cfg.Auth2Config)
	if err != nil {
		return nil, err
	}
	return &Client{
		config: cfg,

		GlobalAccounts: btpglobalaccounts.New(cfg.AccountsHost, http),
		SubAccounts:    btpsubaccounts.New(cfg.AccountsHost, http),
		Entitlements:   btpentitlements.New(cfg.EntitlementsHost, http),
	}, nil
}
