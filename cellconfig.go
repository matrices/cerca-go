// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package cercago

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"slices"

	"github.com/matrices/cerca-go/internal/requestconfig"
	"github.com/matrices/cerca-go/option"
)

// CellConfigService contains methods and other services that help with interacting
// with the cerca API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewCellConfigService] method instead.
type CellConfigService struct {
	Options []option.RequestOption
}

// NewCellConfigService generates a new service that applies the given options to
// each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewCellConfigService(opts ...option.RequestOption) (r *CellConfigService) {
	r = &CellConfigService{}
	r.Options = opts
	return
}

func (r *CellConfigService) Get(ctx context.Context, cellID string, opts ...option.RequestOption) (res *CellConfiguration, err error) {
	opts = slices.Concat(r.Options, opts)
	if cellID == "" {
		err = errors.New("missing required cellId parameter")
		return nil, err
	}
	path := fmt.Sprintf("cells/%s/config", cellID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return res, err
}
