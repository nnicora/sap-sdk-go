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
	Endpoint *Endpoint

	OAuth2 *oauth2.Config
}

type Endpoint struct {
	Account     string
	Entitlement string
}

func NewBtpClient(ctx context.Context, cfg *Config) (*Client, error) {
	http, err := oauth2.NewOAuth2ClientWithContext(ctx, cfg.OAuth2)
	if err != nil {
		return nil, err
	}
	return &Client{
		config: cfg,

		GlobalAccounts: btpglobalaccounts.New(cfg.Endpoint.Account, http),
		SubAccounts:    btpsubaccounts.New(cfg.Endpoint.Account, http),
		Entitlements:   btpentitlements.New(cfg.Endpoint.Entitlement, http),
	}, nil
}
