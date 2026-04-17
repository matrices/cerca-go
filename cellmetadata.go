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

// CellMetadataService contains methods and other services that help with
// interacting with the cerca API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewCellMetadataService] method instead.
type CellMetadataService struct {
	Options []option.RequestOption
}

// NewCellMetadataService generates a new service that applies the given options to
// each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewCellMetadataService(opts ...option.RequestOption) (r *CellMetadataService) {
	r = &CellMetadataService{}
	r.Options = opts
	return
}

func (r *CellMetadataService) Update(ctx context.Context, cellID string, body CellMetadataUpdateParams, opts ...option.RequestOption) (res *Cell, err error) {
	opts = slices.Concat(r.Options, opts)
	if cellID == "" {
		err = errors.New("missing required cellId parameter")
		return nil, err
	}
	path := fmt.Sprintf("cells/%s/metadata", cellID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPatch, path, body, &res, opts...)
	return res, err
}

type CellMetadataUpdateParams struct {
	// Arbitrary string metadata stored on a cell. Runtime enforces maximum key and
	// value sizes.
	Metadata param.Field[CellMetadataParam] `json:"metadata" api:"required"`
}

func (r CellMetadataUpdateParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}
