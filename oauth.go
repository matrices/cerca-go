// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package cercago

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"slices"

	"github.com/matrices/cerca-go/internal/apijson"
	"github.com/matrices/cerca-go/internal/param"
	"github.com/matrices/cerca-go/internal/requestconfig"
	"github.com/matrices/cerca-go/option"
)

// OAuthService contains methods and other services that help with interacting with
// the cerca API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewOAuthService] method instead.
type OAuthService struct {
	Options []option.RequestOption
}

// NewOAuthService generates a new service that applies the given options to each
// request. These options are applied after the parent client's options (if there
// is one), and before any request-specific options.
func NewOAuthService(opts ...option.RequestOption) (r *OAuthService) {
	r = &OAuthService{}
	r.Options = opts
	return
}

func (r *OAuthService) Connect(ctx context.Context, provider string, body OAuthConnectParams, opts ...option.RequestOption) (res *OAuthConnectResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	if provider == "" {
		err = errors.New("missing required provider parameter")
		return nil, err
	}
	path := fmt.Sprintf("oauth/connect/%s", provider)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return res, err
}

type OAuthConnectResponse struct {
	// Provider authorization URL that the client should open.
	AuthorizationURL string `json:"authorizationUrl" api:"required"`
	// Origin hosting the OAuth callback endpoint.
	CallbackOrigin string `json:"callbackOrigin" api:"required"`
	// OAuth state expiration timestamp, when returned.
	ExpiresAt string `json:"expiresAt"`
	// OAuth state nonce, when returned.
	StateNonce string                   `json:"stateNonce"`
	JSON       oauthConnectResponseJSON `json:"-"`
}

// oauthConnectResponseJSON contains the JSON metadata for the struct
// [OAuthConnectResponse]
type oauthConnectResponseJSON struct {
	AuthorizationURL apijson.Field
	CallbackOrigin   apijson.Field
	ExpiresAt        apijson.Field
	StateNonce       apijson.Field
	raw              string
	ExtraFields      map[string]apijson.Field
}

func (r *OAuthConnectResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r oauthConnectResponseJSON) RawJSON() string {
	return r.raw
}

type OAuthConnectParams struct {
	// HTTP(S) origin that receives the OAuth completion message.
	ReturnOrigin param.Field[string] `json:"returnOrigin" api:"required"`
	// Credential connection scope to attach the OAuth account to.
	Scope param.Field[string] `json:"scope" api:"required"`
	// Provider-specific OAuth scopes. Empty entries are ignored after trimming.
	Scopes param.Field[[]string] `json:"scopes"`
}

func (r OAuthConnectParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}
