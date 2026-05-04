// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package cercago

import (
	"context"
	"net/http"
	"slices"

	"github.com/matrices/cerca-go/internal/apijson"
	"github.com/matrices/cerca-go/internal/requestconfig"
	"github.com/matrices/cerca-go/option"
)

// ModelService contains methods and other services that help with interacting with
// the cerca API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewModelService] method instead.
type ModelService struct {
	Options []option.RequestOption
}

// NewModelService generates a new service that applies the given options to each
// request. These options are applied after the parent client's options (if there
// is one), and before any request-specific options.
func NewModelService(opts ...option.RequestOption) (r *ModelService) {
	r = &ModelService{}
	r.Options = opts
	return
}

// List models
func (r *ModelService) List(ctx context.Context, opts ...option.RequestOption) (res *ModelListResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	path := "models"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return res, err
}

type ModelDescriptor struct {
	ID                     string                  `json:"id" api:"required"`
	Capabilities           []string                `json:"capabilities" api:"required"`
	ContextWindowTokens    float64                 `json:"contextWindowTokens" api:"required"`
	DefaultMaxOutputTokens float64                 `json:"defaultMaxOutputTokens" api:"required"`
	Name                   string                  `json:"name" api:"required"`
	Provider               ModelDescriptorProvider `json:"provider" api:"required"`
	RequestHeadroomTokens  float64                 `json:"requestHeadroomTokens" api:"required"`
	JSON                   modelDescriptorJSON     `json:"-"`
}

// modelDescriptorJSON contains the JSON metadata for the struct [ModelDescriptor]
type modelDescriptorJSON struct {
	ID                     apijson.Field
	Capabilities           apijson.Field
	ContextWindowTokens    apijson.Field
	DefaultMaxOutputTokens apijson.Field
	Name                   apijson.Field
	Provider               apijson.Field
	RequestHeadroomTokens  apijson.Field
	raw                    string
	ExtraFields            map[string]apijson.Field
}

func (r *ModelDescriptor) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r modelDescriptorJSON) RawJSON() string {
	return r.raw
}

type ModelDescriptorProvider string

const (
	ModelDescriptorProviderAnthropic ModelDescriptorProvider = "anthropic"
	ModelDescriptorProviderOpenAI    ModelDescriptorProvider = "openai"
	ModelDescriptorProviderGoogle    ModelDescriptorProvider = "google"
	ModelDescriptorProviderFireworks ModelDescriptorProvider = "fireworks"
)

func (r ModelDescriptorProvider) IsKnown() bool {
	switch r {
	case ModelDescriptorProviderAnthropic, ModelDescriptorProviderOpenAI, ModelDescriptorProviderGoogle, ModelDescriptorProviderFireworks:
		return true
	}
	return false
}

type ModelListResponse struct {
	Models []ModelDescriptor     `json:"models" api:"required"`
	JSON   modelListResponseJSON `json:"-"`
}

// modelListResponseJSON contains the JSON metadata for the struct
// [ModelListResponse]
type modelListResponseJSON struct {
	Models      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ModelListResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r modelListResponseJSON) RawJSON() string {
	return r.raw
}
