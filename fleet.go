// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package cercago

import (
	"github.com/matrices/cerca-go/internal/apijson"
	"github.com/matrices/cerca-go/option"
)

// FleetService contains methods and other services that help with interacting with
// the cerca API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewFleetService] method instead.
type FleetService struct {
	Options []option.RequestOption
}

// NewFleetService generates a new service that applies the given options to each
// request. These options are applied after the parent client's options (if there
// is one), and before any request-specific options.
func NewFleetService(opts ...option.RequestOption) (r *FleetService) {
	r = &FleetService{}
	r.Options = opts
	return
}

type Fleet struct {
	ID        string    `json:"id" api:"required"`
	CreatedAt string    `json:"createdAt" api:"required"`
	Name      string    `json:"name" api:"required"`
	OrgID     string    `json:"orgId" api:"required"`
	UpdatedAt string    `json:"updatedAt" api:"required"`
	JSON      fleetJSON `json:"-"`
}

// fleetJSON contains the JSON metadata for the struct [Fleet]
type fleetJSON struct {
	ID          apijson.Field
	CreatedAt   apijson.Field
	Name        apijson.Field
	OrgID       apijson.Field
	UpdatedAt   apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *Fleet) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r fleetJSON) RawJSON() string {
	return r.raw
}
