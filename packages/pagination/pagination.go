// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package pagination

import (
	"net/http"

	"github.com/matrices/cerca-go/internal/apijson"
	"github.com/matrices/cerca-go/internal/requestconfig"
	"github.com/matrices/cerca-go/option"
)

type AgentsCursorPage[T any] struct {
	Agents  []T                  `json:"agents"`
	Cursor  string               `json:"cursor" api:"nullable"`
	HasMore bool                 `json:"hasMore"`
	JSON    agentsCursorPageJSON `json:"-"`
	cfg     *requestconfig.RequestConfig
	res     *http.Response
}

// agentsCursorPageJSON contains the JSON metadata for the struct
// [AgentsCursorPage[T]]
type agentsCursorPageJSON struct {
	Agents      apijson.Field
	Cursor      apijson.Field
	HasMore     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AgentsCursorPage[T]) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r agentsCursorPageJSON) RawJSON() string {
	return r.raw
}

// GetNextPage returns the next page as defined by this pagination style. When
// there is no next page, this function will return a 'nil' for the page value, but
// will not return an error
func (r *AgentsCursorPage[T]) GetNextPage() (res *AgentsCursorPage[T], err error) {
	if len(r.Agents) == 0 {
		return nil, nil
	}

	if !r.JSON.HasMore.IsMissing() && r.HasMore == false {
		return nil, nil
	}
	next := r.Cursor
	if len(next) == 0 {
		return nil, nil
	}
	cfg := r.cfg.Clone(r.cfg.Context)
	err = cfg.Apply(option.WithQuery("cursor", next))
	if err != nil {
		return nil, err
	}
	var raw *http.Response
	cfg.ResponseInto = &raw
	cfg.ResponseBodyInto = &res
	err = cfg.Execute()
	if err != nil {
		return nil, err
	}
	res.SetPageConfig(cfg, raw)
	return res, nil
}

func (r *AgentsCursorPage[T]) SetPageConfig(cfg *requestconfig.RequestConfig, res *http.Response) {
	if r == nil {
		r = &AgentsCursorPage[T]{}
	}
	r.cfg = cfg
	r.res = res
}

type AgentsCursorPageAutoPager[T any] struct {
	page *AgentsCursorPage[T]
	cur  T
	idx  int
	run  int
	err  error
}

func NewAgentsCursorPageAutoPager[T any](page *AgentsCursorPage[T], err error) *AgentsCursorPageAutoPager[T] {
	return &AgentsCursorPageAutoPager[T]{
		page: page,
		err:  err,
	}
}

func (r *AgentsCursorPageAutoPager[T]) Next() bool {
	if r.page == nil || len(r.page.Agents) == 0 {
		return false
	}
	if r.idx >= len(r.page.Agents) {
		r.idx = 0
		r.page, r.err = r.page.GetNextPage()
		if r.err != nil || r.page == nil || len(r.page.Agents) == 0 {
			return false
		}
	}
	r.cur = r.page.Agents[r.idx]
	r.run += 1
	r.idx += 1
	return true
}

func (r *AgentsCursorPageAutoPager[T]) Current() T {
	return r.cur
}

func (r *AgentsCursorPageAutoPager[T]) Err() error {
	return r.err
}

func (r *AgentsCursorPageAutoPager[T]) Index() int {
	return r.run
}

type ApprovalRequestsCursorPage[T any] struct {
	Approvals []T                            `json:"approvals"`
	Cursor    string                         `json:"cursor" api:"nullable"`
	HasMore   bool                           `json:"hasMore"`
	JSON      approvalRequestsCursorPageJSON `json:"-"`
	cfg       *requestconfig.RequestConfig
	res       *http.Response
}

// approvalRequestsCursorPageJSON contains the JSON metadata for the struct
// [ApprovalRequestsCursorPage[T]]
type approvalRequestsCursorPageJSON struct {
	Approvals   apijson.Field
	Cursor      apijson.Field
	HasMore     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ApprovalRequestsCursorPage[T]) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r approvalRequestsCursorPageJSON) RawJSON() string {
	return r.raw
}

// GetNextPage returns the next page as defined by this pagination style. When
// there is no next page, this function will return a 'nil' for the page value, but
// will not return an error
func (r *ApprovalRequestsCursorPage[T]) GetNextPage() (res *ApprovalRequestsCursorPage[T], err error) {
	if len(r.Approvals) == 0 {
		return nil, nil
	}

	if !r.JSON.HasMore.IsMissing() && r.HasMore == false {
		return nil, nil
	}
	next := r.Cursor
	if len(next) == 0 {
		return nil, nil
	}
	cfg := r.cfg.Clone(r.cfg.Context)
	err = cfg.Apply(option.WithQuery("cursor", next))
	if err != nil {
		return nil, err
	}
	var raw *http.Response
	cfg.ResponseInto = &raw
	cfg.ResponseBodyInto = &res
	err = cfg.Execute()
	if err != nil {
		return nil, err
	}
	res.SetPageConfig(cfg, raw)
	return res, nil
}

func (r *ApprovalRequestsCursorPage[T]) SetPageConfig(cfg *requestconfig.RequestConfig, res *http.Response) {
	if r == nil {
		r = &ApprovalRequestsCursorPage[T]{}
	}
	r.cfg = cfg
	r.res = res
}

type ApprovalRequestsCursorPageAutoPager[T any] struct {
	page *ApprovalRequestsCursorPage[T]
	cur  T
	idx  int
	run  int
	err  error
}

func NewApprovalRequestsCursorPageAutoPager[T any](page *ApprovalRequestsCursorPage[T], err error) *ApprovalRequestsCursorPageAutoPager[T] {
	return &ApprovalRequestsCursorPageAutoPager[T]{
		page: page,
		err:  err,
	}
}

func (r *ApprovalRequestsCursorPageAutoPager[T]) Next() bool {
	if r.page == nil || len(r.page.Approvals) == 0 {
		return false
	}
	if r.idx >= len(r.page.Approvals) {
		r.idx = 0
		r.page, r.err = r.page.GetNextPage()
		if r.err != nil || r.page == nil || len(r.page.Approvals) == 0 {
			return false
		}
	}
	r.cur = r.page.Approvals[r.idx]
	r.run += 1
	r.idx += 1
	return true
}

func (r *ApprovalRequestsCursorPageAutoPager[T]) Current() T {
	return r.cur
}

func (r *ApprovalRequestsCursorPageAutoPager[T]) Err() error {
	return r.err
}

func (r *ApprovalRequestsCursorPageAutoPager[T]) Index() int {
	return r.run
}

type ConnectionsCursorPage[T any] struct {
	Connections []T                       `json:"connections"`
	Cursor      string                    `json:"cursor" api:"nullable"`
	HasMore     bool                      `json:"hasMore"`
	JSON        connectionsCursorPageJSON `json:"-"`
	cfg         *requestconfig.RequestConfig
	res         *http.Response
}

// connectionsCursorPageJSON contains the JSON metadata for the struct
// [ConnectionsCursorPage[T]]
type connectionsCursorPageJSON struct {
	Connections apijson.Field
	Cursor      apijson.Field
	HasMore     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ConnectionsCursorPage[T]) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r connectionsCursorPageJSON) RawJSON() string {
	return r.raw
}

// GetNextPage returns the next page as defined by this pagination style. When
// there is no next page, this function will return a 'nil' for the page value, but
// will not return an error
func (r *ConnectionsCursorPage[T]) GetNextPage() (res *ConnectionsCursorPage[T], err error) {
	if len(r.Connections) == 0 {
		return nil, nil
	}

	if !r.JSON.HasMore.IsMissing() && r.HasMore == false {
		return nil, nil
	}
	next := r.Cursor
	if len(next) == 0 {
		return nil, nil
	}
	cfg := r.cfg.Clone(r.cfg.Context)
	err = cfg.Apply(option.WithQuery("cursor", next))
	if err != nil {
		return nil, err
	}
	var raw *http.Response
	cfg.ResponseInto = &raw
	cfg.ResponseBodyInto = &res
	err = cfg.Execute()
	if err != nil {
		return nil, err
	}
	res.SetPageConfig(cfg, raw)
	return res, nil
}

func (r *ConnectionsCursorPage[T]) SetPageConfig(cfg *requestconfig.RequestConfig, res *http.Response) {
	if r == nil {
		r = &ConnectionsCursorPage[T]{}
	}
	r.cfg = cfg
	r.res = res
}

type ConnectionsCursorPageAutoPager[T any] struct {
	page *ConnectionsCursorPage[T]
	cur  T
	idx  int
	run  int
	err  error
}

func NewConnectionsCursorPageAutoPager[T any](page *ConnectionsCursorPage[T], err error) *ConnectionsCursorPageAutoPager[T] {
	return &ConnectionsCursorPageAutoPager[T]{
		page: page,
		err:  err,
	}
}

func (r *ConnectionsCursorPageAutoPager[T]) Next() bool {
	if r.page == nil || len(r.page.Connections) == 0 {
		return false
	}
	if r.idx >= len(r.page.Connections) {
		r.idx = 0
		r.page, r.err = r.page.GetNextPage()
		if r.err != nil || r.page == nil || len(r.page.Connections) == 0 {
			return false
		}
	}
	r.cur = r.page.Connections[r.idx]
	r.run += 1
	r.idx += 1
	return true
}

func (r *ConnectionsCursorPageAutoPager[T]) Current() T {
	return r.cur
}

func (r *ConnectionsCursorPageAutoPager[T]) Err() error {
	return r.err
}

func (r *ConnectionsCursorPageAutoPager[T]) Index() int {
	return r.run
}

type EntriesCursorPage[T any] struct {
	Entries []T                   `json:"entries"`
	Cursor  string                `json:"cursor" api:"nullable"`
	HasMore bool                  `json:"hasMore"`
	JSON    entriesCursorPageJSON `json:"-"`
	cfg     *requestconfig.RequestConfig
	res     *http.Response
}

// entriesCursorPageJSON contains the JSON metadata for the struct
// [EntriesCursorPage[T]]
type entriesCursorPageJSON struct {
	Entries     apijson.Field
	Cursor      apijson.Field
	HasMore     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *EntriesCursorPage[T]) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r entriesCursorPageJSON) RawJSON() string {
	return r.raw
}

// GetNextPage returns the next page as defined by this pagination style. When
// there is no next page, this function will return a 'nil' for the page value, but
// will not return an error
func (r *EntriesCursorPage[T]) GetNextPage() (res *EntriesCursorPage[T], err error) {
	if len(r.Entries) == 0 {
		return nil, nil
	}

	if !r.JSON.HasMore.IsMissing() && r.HasMore == false {
		return nil, nil
	}
	next := r.Cursor
	if len(next) == 0 {
		return nil, nil
	}
	cfg := r.cfg.Clone(r.cfg.Context)
	err = cfg.Apply(option.WithQuery("cursor", next))
	if err != nil {
		return nil, err
	}
	var raw *http.Response
	cfg.ResponseInto = &raw
	cfg.ResponseBodyInto = &res
	err = cfg.Execute()
	if err != nil {
		return nil, err
	}
	res.SetPageConfig(cfg, raw)
	return res, nil
}

func (r *EntriesCursorPage[T]) SetPageConfig(cfg *requestconfig.RequestConfig, res *http.Response) {
	if r == nil {
		r = &EntriesCursorPage[T]{}
	}
	r.cfg = cfg
	r.res = res
}

type EntriesCursorPageAutoPager[T any] struct {
	page *EntriesCursorPage[T]
	cur  T
	idx  int
	run  int
	err  error
}

func NewEntriesCursorPageAutoPager[T any](page *EntriesCursorPage[T], err error) *EntriesCursorPageAutoPager[T] {
	return &EntriesCursorPageAutoPager[T]{
		page: page,
		err:  err,
	}
}

func (r *EntriesCursorPageAutoPager[T]) Next() bool {
	if r.page == nil || len(r.page.Entries) == 0 {
		return false
	}
	if r.idx >= len(r.page.Entries) {
		r.idx = 0
		r.page, r.err = r.page.GetNextPage()
		if r.err != nil || r.page == nil || len(r.page.Entries) == 0 {
			return false
		}
	}
	r.cur = r.page.Entries[r.idx]
	r.run += 1
	r.idx += 1
	return true
}

func (r *EntriesCursorPageAutoPager[T]) Current() T {
	return r.cur
}

func (r *EntriesCursorPageAutoPager[T]) Err() error {
	return r.err
}

func (r *EntriesCursorPageAutoPager[T]) Index() int {
	return r.run
}

type EventsCursorPage[T any] struct {
	Events  []T                  `json:"events"`
	Cursor  string               `json:"cursor" api:"nullable"`
	HasMore bool                 `json:"hasMore"`
	JSON    eventsCursorPageJSON `json:"-"`
	cfg     *requestconfig.RequestConfig
	res     *http.Response
}

// eventsCursorPageJSON contains the JSON metadata for the struct
// [EventsCursorPage[T]]
type eventsCursorPageJSON struct {
	Events      apijson.Field
	Cursor      apijson.Field
	HasMore     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *EventsCursorPage[T]) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r eventsCursorPageJSON) RawJSON() string {
	return r.raw
}

// GetNextPage returns the next page as defined by this pagination style. When
// there is no next page, this function will return a 'nil' for the page value, but
// will not return an error
func (r *EventsCursorPage[T]) GetNextPage() (res *EventsCursorPage[T], err error) {
	if len(r.Events) == 0 {
		return nil, nil
	}

	if !r.JSON.HasMore.IsMissing() && r.HasMore == false {
		return nil, nil
	}
	next := r.Cursor
	if len(next) == 0 {
		return nil, nil
	}
	cfg := r.cfg.Clone(r.cfg.Context)
	err = cfg.Apply(option.WithQuery("cursor", next))
	if err != nil {
		return nil, err
	}
	var raw *http.Response
	cfg.ResponseInto = &raw
	cfg.ResponseBodyInto = &res
	err = cfg.Execute()
	if err != nil {
		return nil, err
	}
	res.SetPageConfig(cfg, raw)
	return res, nil
}

func (r *EventsCursorPage[T]) SetPageConfig(cfg *requestconfig.RequestConfig, res *http.Response) {
	if r == nil {
		r = &EventsCursorPage[T]{}
	}
	r.cfg = cfg
	r.res = res
}

type EventsCursorPageAutoPager[T any] struct {
	page *EventsCursorPage[T]
	cur  T
	idx  int
	run  int
	err  error
}

func NewEventsCursorPageAutoPager[T any](page *EventsCursorPage[T], err error) *EventsCursorPageAutoPager[T] {
	return &EventsCursorPageAutoPager[T]{
		page: page,
		err:  err,
	}
}

func (r *EventsCursorPageAutoPager[T]) Next() bool {
	if r.page == nil || len(r.page.Events) == 0 {
		return false
	}
	if r.idx >= len(r.page.Events) {
		r.idx = 0
		r.page, r.err = r.page.GetNextPage()
		if r.err != nil || r.page == nil || len(r.page.Events) == 0 {
			return false
		}
	}
	r.cur = r.page.Events[r.idx]
	r.run += 1
	r.idx += 1
	return true
}

func (r *EventsCursorPageAutoPager[T]) Current() T {
	return r.cur
}

func (r *EventsCursorPageAutoPager[T]) Err() error {
	return r.err
}

func (r *EventsCursorPageAutoPager[T]) Index() int {
	return r.run
}

type FleetsCursorPage[T any] struct {
	Fleets  []T                  `json:"fleets"`
	Cursor  string               `json:"cursor" api:"nullable"`
	HasMore bool                 `json:"hasMore"`
	JSON    fleetsCursorPageJSON `json:"-"`
	cfg     *requestconfig.RequestConfig
	res     *http.Response
}

// fleetsCursorPageJSON contains the JSON metadata for the struct
// [FleetsCursorPage[T]]
type fleetsCursorPageJSON struct {
	Fleets      apijson.Field
	Cursor      apijson.Field
	HasMore     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *FleetsCursorPage[T]) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r fleetsCursorPageJSON) RawJSON() string {
	return r.raw
}

// GetNextPage returns the next page as defined by this pagination style. When
// there is no next page, this function will return a 'nil' for the page value, but
// will not return an error
func (r *FleetsCursorPage[T]) GetNextPage() (res *FleetsCursorPage[T], err error) {
	if len(r.Fleets) == 0 {
		return nil, nil
	}

	if !r.JSON.HasMore.IsMissing() && r.HasMore == false {
		return nil, nil
	}
	next := r.Cursor
	if len(next) == 0 {
		return nil, nil
	}
	cfg := r.cfg.Clone(r.cfg.Context)
	err = cfg.Apply(option.WithQuery("cursor", next))
	if err != nil {
		return nil, err
	}
	var raw *http.Response
	cfg.ResponseInto = &raw
	cfg.ResponseBodyInto = &res
	err = cfg.Execute()
	if err != nil {
		return nil, err
	}
	res.SetPageConfig(cfg, raw)
	return res, nil
}

func (r *FleetsCursorPage[T]) SetPageConfig(cfg *requestconfig.RequestConfig, res *http.Response) {
	if r == nil {
		r = &FleetsCursorPage[T]{}
	}
	r.cfg = cfg
	r.res = res
}

type FleetsCursorPageAutoPager[T any] struct {
	page *FleetsCursorPage[T]
	cur  T
	idx  int
	run  int
	err  error
}

func NewFleetsCursorPageAutoPager[T any](page *FleetsCursorPage[T], err error) *FleetsCursorPageAutoPager[T] {
	return &FleetsCursorPageAutoPager[T]{
		page: page,
		err:  err,
	}
}

func (r *FleetsCursorPageAutoPager[T]) Next() bool {
	if r.page == nil || len(r.page.Fleets) == 0 {
		return false
	}
	if r.idx >= len(r.page.Fleets) {
		r.idx = 0
		r.page, r.err = r.page.GetNextPage()
		if r.err != nil || r.page == nil || len(r.page.Fleets) == 0 {
			return false
		}
	}
	r.cur = r.page.Fleets[r.idx]
	r.run += 1
	r.idx += 1
	return true
}

func (r *FleetsCursorPageAutoPager[T]) Current() T {
	return r.cur
}

func (r *FleetsCursorPageAutoPager[T]) Err() error {
	return r.err
}

func (r *FleetsCursorPageAutoPager[T]) Index() int {
	return r.run
}

type GrantsCursorPage[T any] struct {
	Grants  []T                  `json:"grants"`
	Cursor  string               `json:"cursor" api:"nullable"`
	HasMore bool                 `json:"hasMore"`
	JSON    grantsCursorPageJSON `json:"-"`
	cfg     *requestconfig.RequestConfig
	res     *http.Response
}

// grantsCursorPageJSON contains the JSON metadata for the struct
// [GrantsCursorPage[T]]
type grantsCursorPageJSON struct {
	Grants      apijson.Field
	Cursor      apijson.Field
	HasMore     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *GrantsCursorPage[T]) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r grantsCursorPageJSON) RawJSON() string {
	return r.raw
}

// GetNextPage returns the next page as defined by this pagination style. When
// there is no next page, this function will return a 'nil' for the page value, but
// will not return an error
func (r *GrantsCursorPage[T]) GetNextPage() (res *GrantsCursorPage[T], err error) {
	if len(r.Grants) == 0 {
		return nil, nil
	}

	if !r.JSON.HasMore.IsMissing() && r.HasMore == false {
		return nil, nil
	}
	next := r.Cursor
	if len(next) == 0 {
		return nil, nil
	}
	cfg := r.cfg.Clone(r.cfg.Context)
	err = cfg.Apply(option.WithQuery("cursor", next))
	if err != nil {
		return nil, err
	}
	var raw *http.Response
	cfg.ResponseInto = &raw
	cfg.ResponseBodyInto = &res
	err = cfg.Execute()
	if err != nil {
		return nil, err
	}
	res.SetPageConfig(cfg, raw)
	return res, nil
}

func (r *GrantsCursorPage[T]) SetPageConfig(cfg *requestconfig.RequestConfig, res *http.Response) {
	if r == nil {
		r = &GrantsCursorPage[T]{}
	}
	r.cfg = cfg
	r.res = res
}

type GrantsCursorPageAutoPager[T any] struct {
	page *GrantsCursorPage[T]
	cur  T
	idx  int
	run  int
	err  error
}

func NewGrantsCursorPageAutoPager[T any](page *GrantsCursorPage[T], err error) *GrantsCursorPageAutoPager[T] {
	return &GrantsCursorPageAutoPager[T]{
		page: page,
		err:  err,
	}
}

func (r *GrantsCursorPageAutoPager[T]) Next() bool {
	if r.page == nil || len(r.page.Grants) == 0 {
		return false
	}
	if r.idx >= len(r.page.Grants) {
		r.idx = 0
		r.page, r.err = r.page.GetNextPage()
		if r.err != nil || r.page == nil || len(r.page.Grants) == 0 {
			return false
		}
	}
	r.cur = r.page.Grants[r.idx]
	r.run += 1
	r.idx += 1
	return true
}

func (r *GrantsCursorPageAutoPager[T]) Current() T {
	return r.cur
}

func (r *GrantsCursorPageAutoPager[T]) Err() error {
	return r.err
}

func (r *GrantsCursorPageAutoPager[T]) Index() int {
	return r.run
}

type ResultsCursorPage[T any] struct {
	Results []T                   `json:"results"`
	Cursor  string                `json:"cursor" api:"nullable"`
	HasMore bool                  `json:"hasMore"`
	JSON    resultsCursorPageJSON `json:"-"`
	cfg     *requestconfig.RequestConfig
	res     *http.Response
}

// resultsCursorPageJSON contains the JSON metadata for the struct
// [ResultsCursorPage[T]]
type resultsCursorPageJSON struct {
	Results     apijson.Field
	Cursor      apijson.Field
	HasMore     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ResultsCursorPage[T]) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r resultsCursorPageJSON) RawJSON() string {
	return r.raw
}

// GetNextPage returns the next page as defined by this pagination style. When
// there is no next page, this function will return a 'nil' for the page value, but
// will not return an error
func (r *ResultsCursorPage[T]) GetNextPage() (res *ResultsCursorPage[T], err error) {
	if len(r.Results) == 0 {
		return nil, nil
	}

	if !r.JSON.HasMore.IsMissing() && r.HasMore == false {
		return nil, nil
	}
	next := r.Cursor
	if len(next) == 0 {
		return nil, nil
	}
	cfg := r.cfg.Clone(r.cfg.Context)
	err = cfg.Apply(option.WithQuery("cursor", next))
	if err != nil {
		return nil, err
	}
	var raw *http.Response
	cfg.ResponseInto = &raw
	cfg.ResponseBodyInto = &res
	err = cfg.Execute()
	if err != nil {
		return nil, err
	}
	res.SetPageConfig(cfg, raw)
	return res, nil
}

func (r *ResultsCursorPage[T]) SetPageConfig(cfg *requestconfig.RequestConfig, res *http.Response) {
	if r == nil {
		r = &ResultsCursorPage[T]{}
	}
	r.cfg = cfg
	r.res = res
}

type ResultsCursorPageAutoPager[T any] struct {
	page *ResultsCursorPage[T]
	cur  T
	idx  int
	run  int
	err  error
}

func NewResultsCursorPageAutoPager[T any](page *ResultsCursorPage[T], err error) *ResultsCursorPageAutoPager[T] {
	return &ResultsCursorPageAutoPager[T]{
		page: page,
		err:  err,
	}
}

func (r *ResultsCursorPageAutoPager[T]) Next() bool {
	if r.page == nil || len(r.page.Results) == 0 {
		return false
	}
	if r.idx >= len(r.page.Results) {
		r.idx = 0
		r.page, r.err = r.page.GetNextPage()
		if r.err != nil || r.page == nil || len(r.page.Results) == 0 {
			return false
		}
	}
	r.cur = r.page.Results[r.idx]
	r.run += 1
	r.idx += 1
	return true
}

func (r *ResultsCursorPageAutoPager[T]) Current() T {
	return r.cur
}

func (r *ResultsCursorPageAutoPager[T]) Err() error {
	return r.err
}

func (r *ResultsCursorPageAutoPager[T]) Index() int {
	return r.run
}

type SourcesCursorPage[T any] struct {
	Sources []T                   `json:"sources"`
	Cursor  string                `json:"cursor" api:"nullable"`
	HasMore bool                  `json:"hasMore"`
	JSON    sourcesCursorPageJSON `json:"-"`
	cfg     *requestconfig.RequestConfig
	res     *http.Response
}

// sourcesCursorPageJSON contains the JSON metadata for the struct
// [SourcesCursorPage[T]]
type sourcesCursorPageJSON struct {
	Sources     apijson.Field
	Cursor      apijson.Field
	HasMore     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SourcesCursorPage[T]) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r sourcesCursorPageJSON) RawJSON() string {
	return r.raw
}

// GetNextPage returns the next page as defined by this pagination style. When
// there is no next page, this function will return a 'nil' for the page value, but
// will not return an error
func (r *SourcesCursorPage[T]) GetNextPage() (res *SourcesCursorPage[T], err error) {
	if len(r.Sources) == 0 {
		return nil, nil
	}

	if !r.JSON.HasMore.IsMissing() && r.HasMore == false {
		return nil, nil
	}
	next := r.Cursor
	if len(next) == 0 {
		return nil, nil
	}
	cfg := r.cfg.Clone(r.cfg.Context)
	err = cfg.Apply(option.WithQuery("cursor", next))
	if err != nil {
		return nil, err
	}
	var raw *http.Response
	cfg.ResponseInto = &raw
	cfg.ResponseBodyInto = &res
	err = cfg.Execute()
	if err != nil {
		return nil, err
	}
	res.SetPageConfig(cfg, raw)
	return res, nil
}

func (r *SourcesCursorPage[T]) SetPageConfig(cfg *requestconfig.RequestConfig, res *http.Response) {
	if r == nil {
		r = &SourcesCursorPage[T]{}
	}
	r.cfg = cfg
	r.res = res
}

type SourcesCursorPageAutoPager[T any] struct {
	page *SourcesCursorPage[T]
	cur  T
	idx  int
	run  int
	err  error
}

func NewSourcesCursorPageAutoPager[T any](page *SourcesCursorPage[T], err error) *SourcesCursorPageAutoPager[T] {
	return &SourcesCursorPageAutoPager[T]{
		page: page,
		err:  err,
	}
}

func (r *SourcesCursorPageAutoPager[T]) Next() bool {
	if r.page == nil || len(r.page.Sources) == 0 {
		return false
	}
	if r.idx >= len(r.page.Sources) {
		r.idx = 0
		r.page, r.err = r.page.GetNextPage()
		if r.err != nil || r.page == nil || len(r.page.Sources) == 0 {
			return false
		}
	}
	r.cur = r.page.Sources[r.idx]
	r.run += 1
	r.idx += 1
	return true
}

func (r *SourcesCursorPageAutoPager[T]) Current() T {
	return r.cur
}

func (r *SourcesCursorPageAutoPager[T]) Err() error {
	return r.err
}

func (r *SourcesCursorPageAutoPager[T]) Index() int {
	return r.run
}

type SubscriptionsCursorPage[T any] struct {
	Subscriptions []T                         `json:"subscriptions"`
	Cursor        string                      `json:"cursor" api:"nullable"`
	HasMore       bool                        `json:"hasMore"`
	JSON          subscriptionsCursorPageJSON `json:"-"`
	cfg           *requestconfig.RequestConfig
	res           *http.Response
}

// subscriptionsCursorPageJSON contains the JSON metadata for the struct
// [SubscriptionsCursorPage[T]]
type subscriptionsCursorPageJSON struct {
	Subscriptions apijson.Field
	Cursor        apijson.Field
	HasMore       apijson.Field
	raw           string
	ExtraFields   map[string]apijson.Field
}

func (r *SubscriptionsCursorPage[T]) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r subscriptionsCursorPageJSON) RawJSON() string {
	return r.raw
}

// GetNextPage returns the next page as defined by this pagination style. When
// there is no next page, this function will return a 'nil' for the page value, but
// will not return an error
func (r *SubscriptionsCursorPage[T]) GetNextPage() (res *SubscriptionsCursorPage[T], err error) {
	if len(r.Subscriptions) == 0 {
		return nil, nil
	}

	if !r.JSON.HasMore.IsMissing() && r.HasMore == false {
		return nil, nil
	}
	next := r.Cursor
	if len(next) == 0 {
		return nil, nil
	}
	cfg := r.cfg.Clone(r.cfg.Context)
	err = cfg.Apply(option.WithQuery("cursor", next))
	if err != nil {
		return nil, err
	}
	var raw *http.Response
	cfg.ResponseInto = &raw
	cfg.ResponseBodyInto = &res
	err = cfg.Execute()
	if err != nil {
		return nil, err
	}
	res.SetPageConfig(cfg, raw)
	return res, nil
}

func (r *SubscriptionsCursorPage[T]) SetPageConfig(cfg *requestconfig.RequestConfig, res *http.Response) {
	if r == nil {
		r = &SubscriptionsCursorPage[T]{}
	}
	r.cfg = cfg
	r.res = res
}

type SubscriptionsCursorPageAutoPager[T any] struct {
	page *SubscriptionsCursorPage[T]
	cur  T
	idx  int
	run  int
	err  error
}

func NewSubscriptionsCursorPageAutoPager[T any](page *SubscriptionsCursorPage[T], err error) *SubscriptionsCursorPageAutoPager[T] {
	return &SubscriptionsCursorPageAutoPager[T]{
		page: page,
		err:  err,
	}
}

func (r *SubscriptionsCursorPageAutoPager[T]) Next() bool {
	if r.page == nil || len(r.page.Subscriptions) == 0 {
		return false
	}
	if r.idx >= len(r.page.Subscriptions) {
		r.idx = 0
		r.page, r.err = r.page.GetNextPage()
		if r.err != nil || r.page == nil || len(r.page.Subscriptions) == 0 {
			return false
		}
	}
	r.cur = r.page.Subscriptions[r.idx]
	r.run += 1
	r.idx += 1
	return true
}

func (r *SubscriptionsCursorPageAutoPager[T]) Current() T {
	return r.cur
}

func (r *SubscriptionsCursorPageAutoPager[T]) Err() error {
	return r.err
}

func (r *SubscriptionsCursorPageAutoPager[T]) Index() int {
	return r.run
}

type ThreadsCursorPage[T any] struct {
	Threads []T                   `json:"threads"`
	Cursor  string                `json:"cursor" api:"nullable"`
	HasMore bool                  `json:"hasMore"`
	JSON    threadsCursorPageJSON `json:"-"`
	cfg     *requestconfig.RequestConfig
	res     *http.Response
}

// threadsCursorPageJSON contains the JSON metadata for the struct
// [ThreadsCursorPage[T]]
type threadsCursorPageJSON struct {
	Threads     apijson.Field
	Cursor      apijson.Field
	HasMore     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ThreadsCursorPage[T]) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r threadsCursorPageJSON) RawJSON() string {
	return r.raw
}

// GetNextPage returns the next page as defined by this pagination style. When
// there is no next page, this function will return a 'nil' for the page value, but
// will not return an error
func (r *ThreadsCursorPage[T]) GetNextPage() (res *ThreadsCursorPage[T], err error) {
	if len(r.Threads) == 0 {
		return nil, nil
	}

	if !r.JSON.HasMore.IsMissing() && r.HasMore == false {
		return nil, nil
	}
	next := r.Cursor
	if len(next) == 0 {
		return nil, nil
	}
	cfg := r.cfg.Clone(r.cfg.Context)
	err = cfg.Apply(option.WithQuery("cursor", next))
	if err != nil {
		return nil, err
	}
	var raw *http.Response
	cfg.ResponseInto = &raw
	cfg.ResponseBodyInto = &res
	err = cfg.Execute()
	if err != nil {
		return nil, err
	}
	res.SetPageConfig(cfg, raw)
	return res, nil
}

func (r *ThreadsCursorPage[T]) SetPageConfig(cfg *requestconfig.RequestConfig, res *http.Response) {
	if r == nil {
		r = &ThreadsCursorPage[T]{}
	}
	r.cfg = cfg
	r.res = res
}

type ThreadsCursorPageAutoPager[T any] struct {
	page *ThreadsCursorPage[T]
	cur  T
	idx  int
	run  int
	err  error
}

func NewThreadsCursorPageAutoPager[T any](page *ThreadsCursorPage[T], err error) *ThreadsCursorPageAutoPager[T] {
	return &ThreadsCursorPageAutoPager[T]{
		page: page,
		err:  err,
	}
}

func (r *ThreadsCursorPageAutoPager[T]) Next() bool {
	if r.page == nil || len(r.page.Threads) == 0 {
		return false
	}
	if r.idx >= len(r.page.Threads) {
		r.idx = 0
		r.page, r.err = r.page.GetNextPage()
		if r.err != nil || r.page == nil || len(r.page.Threads) == 0 {
			return false
		}
	}
	r.cur = r.page.Threads[r.idx]
	r.run += 1
	r.idx += 1
	return true
}

func (r *ThreadsCursorPageAutoPager[T]) Current() T {
	return r.cur
}

func (r *ThreadsCursorPageAutoPager[T]) Err() error {
	return r.err
}

func (r *ThreadsCursorPageAutoPager[T]) Index() int {
	return r.run
}

type ThreadMessagesCursorPage[T any] struct {
	Messages []T                          `json:"messages"`
	Cursor   string                       `json:"cursor" api:"nullable"`
	HasMore  bool                         `json:"hasMore"`
	JSON     threadMessagesCursorPageJSON `json:"-"`
	cfg      *requestconfig.RequestConfig
	res      *http.Response
}

// threadMessagesCursorPageJSON contains the JSON metadata for the struct
// [ThreadMessagesCursorPage[T]]
type threadMessagesCursorPageJSON struct {
	Messages    apijson.Field
	Cursor      apijson.Field
	HasMore     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ThreadMessagesCursorPage[T]) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r threadMessagesCursorPageJSON) RawJSON() string {
	return r.raw
}

// GetNextPage returns the next page as defined by this pagination style. When
// there is no next page, this function will return a 'nil' for the page value, but
// will not return an error
func (r *ThreadMessagesCursorPage[T]) GetNextPage() (res *ThreadMessagesCursorPage[T], err error) {
	if len(r.Messages) == 0 {
		return nil, nil
	}

	if !r.JSON.HasMore.IsMissing() && r.HasMore == false {
		return nil, nil
	}
	next := r.Cursor
	if len(next) == 0 {
		return nil, nil
	}
	cfg := r.cfg.Clone(r.cfg.Context)
	err = cfg.Apply(option.WithQuery("cursor", next))
	if err != nil {
		return nil, err
	}
	var raw *http.Response
	cfg.ResponseInto = &raw
	cfg.ResponseBodyInto = &res
	err = cfg.Execute()
	if err != nil {
		return nil, err
	}
	res.SetPageConfig(cfg, raw)
	return res, nil
}

func (r *ThreadMessagesCursorPage[T]) SetPageConfig(cfg *requestconfig.RequestConfig, res *http.Response) {
	if r == nil {
		r = &ThreadMessagesCursorPage[T]{}
	}
	r.cfg = cfg
	r.res = res
}

type ThreadMessagesCursorPageAutoPager[T any] struct {
	page *ThreadMessagesCursorPage[T]
	cur  T
	idx  int
	run  int
	err  error
}

func NewThreadMessagesCursorPageAutoPager[T any](page *ThreadMessagesCursorPage[T], err error) *ThreadMessagesCursorPageAutoPager[T] {
	return &ThreadMessagesCursorPageAutoPager[T]{
		page: page,
		err:  err,
	}
}

func (r *ThreadMessagesCursorPageAutoPager[T]) Next() bool {
	if r.page == nil || len(r.page.Messages) == 0 {
		return false
	}
	if r.idx >= len(r.page.Messages) {
		r.idx = 0
		r.page, r.err = r.page.GetNextPage()
		if r.err != nil || r.page == nil || len(r.page.Messages) == 0 {
			return false
		}
	}
	r.cur = r.page.Messages[r.idx]
	r.run += 1
	r.idx += 1
	return true
}

func (r *ThreadMessagesCursorPageAutoPager[T]) Current() T {
	return r.cur
}

func (r *ThreadMessagesCursorPageAutoPager[T]) Err() error {
	return r.err
}

func (r *ThreadMessagesCursorPageAutoPager[T]) Index() int {
	return r.run
}
