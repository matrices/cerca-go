// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package cercago

import (
	"github.com/matrices/cerca-go/internal/apijson"
	"github.com/matrices/cerca-go/option"
)

// EnvironmentService contains methods and other services that help with
// interacting with the cerca API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewEnvironmentService] method instead.
type EnvironmentService struct {
	Options     []option.RequestOption
	ToolSources *EnvironmentToolSourceService
	Webhooks    *EnvironmentWebhookService
	Events      *EnvironmentEventService
}

// NewEnvironmentService generates a new service that applies the given options to
// each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewEnvironmentService(opts ...option.RequestOption) (r *EnvironmentService) {
	r = &EnvironmentService{}
	r.Options = opts
	r.ToolSources = NewEnvironmentToolSourceService(opts...)
	r.Webhooks = NewEnvironmentWebhookService(opts...)
	r.Events = NewEnvironmentEventService(opts...)
	return
}

type Environment struct {
	ID        string          `json:"id" api:"required"`
	CreatedAt string          `json:"createdAt" api:"required"`
	Name      string          `json:"name" api:"required"`
	OrgID     string          `json:"orgId" api:"required"`
	UpdatedAt string          `json:"updatedAt" api:"required"`
	JSON      environmentJSON `json:"-"`
}

// environmentJSON contains the JSON metadata for the struct [Environment]
type environmentJSON struct {
	ID          apijson.Field
	CreatedAt   apijson.Field
	Name        apijson.Field
	OrgID       apijson.Field
	UpdatedAt   apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *Environment) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r environmentJSON) RawJSON() string {
	return r.raw
}
