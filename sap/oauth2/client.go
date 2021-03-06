package oauth2

import (
	"errors"
	"fmt"
	"github.com/nnicora/sap-sdk-go/internal/utils"
	"golang.org/x/net/context"
	"golang.org/x/oauth2"
	"golang.org/x/oauth2/clientcredentials"
	"net/http"
	"net/url"
	"time"
)

// OAuth2 Config allowing to use
type Config struct {
	GrantType string

	// Username used only when grant_type == password.
	Username string

	// Password  used only when grant_type == password
	Password string

	// ClientID is the application's ID.
	ClientID string

	// ClientSecret is the application's secret.
	ClientSecret string

	// AuthURL is the resource server's token endpoint
	// Host. This is a constant specific to each server.
	// Used only when grant_type ==  authorization_code
	AuthURL string

	// RedirectURL is the Host to redirect users going through
	// the OAuth flow, after the resource owner's URLs.
	// Used only when grant_type ==  authorization_code
	RedirectURL string

	// TokenURL is the resource server's token endpoint
	// Host. This is a constant specific to each server.
	TokenURL string

	// Scope specifies optional requested permissions.
	Scopes []string

	// EndpointParams specifies additional parameters for requests to the token endpoint.
	EndpointParams url.Values

	// AuthStyle optionally specifies how the endpoint wants the
	// client ID & client secret sent. The zero value means to
	// auto-detect.
	AuthStyle oauth2.AuthStyle

	Timeout time.Duration

	//DefaultHttpClient *http.Client
}

func NewOAuth2Client(conf *Config) (*http.Client, error) {
	return NewOAuth2ClientWithContext(context.Background(), conf)
}
func NewOAuth2ClientWithContext(ctx context.Context, conf *Config) (*http.Client, error) {
	httpClient := &http.Client{Timeout: conf.Timeout}
	ctx = context.WithValue(ctx, oauth2.HTTPClient, httpClient)

	if ok, err := utils.HostAlive(conf.TokenURL); err != nil {
		return nil, fmt.Errorf("oauth2 token url '%s' unreachable; %v", conf.TokenURL, err)
	} else if !ok {
		return nil, fmt.Errorf("oauth2 token url '%s' unreachable", conf.TokenURL)
	}

	if conf.GrantType == "client_credentials" {
		config := &clientcredentials.Config{
			ClientID:       conf.ClientID,
			ClientSecret:   conf.ClientSecret,
			TokenURL:       conf.TokenURL,
			Scopes:         conf.Scopes,
			EndpointParams: conf.EndpointParams,
			AuthStyle:      conf.AuthStyle,
		}

		return config.Client(ctx), nil
	}

	// Redirect user to consent page to ask for permission
	// for the scopes specified above.
	//url := config.AuthCodeURL("state", oauth2.AccessTypeOffline)
	//fmt.Printf("Visit the Host for the auth dialog: %v", url)

	if conf.GrantType == "authorization_code" {
		// Using the authorization code that is pushed to the redirect
		// Host. Exchange will do the handshake to retrieve the
		// initial access token. The Http Client returned by
		// conf.Client will refresh the token as necessary.

		if ok, err := utils.HostAlive(conf.AuthURL); err != nil {
			return nil, fmt.Errorf("oauth2 authorization url unreachable; %v", err)
		} else if !ok {
			return nil, fmt.Errorf("oauth2 authorization url unreachable")
		}

		config := &oauth2.Config{
			ClientID:     conf.ClientID,
			ClientSecret: conf.ClientSecret,
			Scopes:       conf.Scopes,
			RedirectURL:  conf.RedirectURL,
			Endpoint: oauth2.Endpoint{
				AuthURL:   conf.AuthURL,
				TokenURL:  conf.TokenURL,
				AuthStyle: conf.AuthStyle,
			},
		}

		var code string
		if _, err := fmt.Scan(&code); err != nil {
			return nil, err
		}

		tok, err := config.Exchange(ctx, code)
		if err != nil {
			return nil, err
		}

		return config.Client(ctx, tok), nil
	}

	if conf.GrantType == "password" {
		config := &oauth2.Config{
			ClientID:     conf.ClientID,
			ClientSecret: conf.ClientSecret,
			Scopes:       conf.Scopes,
			Endpoint: oauth2.Endpoint{
				TokenURL:  conf.TokenURL,
				AuthStyle: conf.AuthStyle,
			},
		}

		tok, err := config.PasswordCredentialsToken(ctx, conf.Username, conf.Password)
		if err != nil {
			return nil, err
		}

		return config.Client(ctx, tok), nil
	}

	return nil, errors.New("unsupported grant type")
}

func (c *Config) Clone() *Config {
	return &Config{
		GrantType:      c.GrantType,
		Username:       c.Username,
		Password:       c.Password,
		ClientID:       c.ClientID,
		ClientSecret:   c.ClientSecret,
		AuthURL:        c.AuthURL,
		RedirectURL:    c.RedirectURL,
		TokenURL:       c.TokenURL,
		Scopes:         c.Scopes,
		EndpointParams: c.EndpointParams,
		AuthStyle:      c.AuthStyle,
		Timeout:        c.Timeout,
		//DefaultHttpClient: c.DefaultHttpClient,
	}
}
