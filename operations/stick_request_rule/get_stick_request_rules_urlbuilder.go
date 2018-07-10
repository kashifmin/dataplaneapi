// Code generated by go-swagger; DO NOT EDIT.

package stick_request_rule

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the generate command

import (
	"errors"
	"net/url"
	golangswaggerpaths "path"
)

// GetStickRequestRulesURL generates an URL for the get stick request rules operation
type GetStickRequestRulesURL struct {
	Backend       string
	TransactionID *string

	_basePath string
	// avoid unkeyed usage
	_ struct{}
}

// WithBasePath sets the base path for this url builder, only required when it's different from the
// base path specified in the swagger spec.
// When the value of the base path is an empty string
func (o *GetStickRequestRulesURL) WithBasePath(bp string) *GetStickRequestRulesURL {
	o.SetBasePath(bp)
	return o
}

// SetBasePath sets the base path for this url builder, only required when it's different from the
// base path specified in the swagger spec.
// When the value of the base path is an empty string
func (o *GetStickRequestRulesURL) SetBasePath(bp string) {
	o._basePath = bp
}

// Build a url path and query string
func (o *GetStickRequestRulesURL) Build() (*url.URL, error) {
	var result url.URL

	var _path = "/services/haproxy/configuration/stick_request_rules"

	_basePath := o._basePath
	if _basePath == "" {
		_basePath = "/v1"
	}
	result.Path = golangswaggerpaths.Join(_basePath, _path)

	qs := make(url.Values)

	backend := o.Backend
	if backend != "" {
		qs.Set("backend", backend)
	}

	var transactionID string
	if o.TransactionID != nil {
		transactionID = *o.TransactionID
	}
	if transactionID != "" {
		qs.Set("transaction_id", transactionID)
	}

	result.RawQuery = qs.Encode()

	return &result, nil
}

// Must is a helper function to panic when the url builder returns an error
func (o *GetStickRequestRulesURL) Must(u *url.URL, err error) *url.URL {
	if err != nil {
		panic(err)
	}
	if u == nil {
		panic("url can't be nil")
	}
	return u
}

// String returns the string representation of the path with query string
func (o *GetStickRequestRulesURL) String() string {
	return o.Must(o.Build()).String()
}

// BuildFull builds a full url with scheme, host, path and query string
func (o *GetStickRequestRulesURL) BuildFull(scheme, host string) (*url.URL, error) {
	if scheme == "" {
		return nil, errors.New("scheme is required for a full url on GetStickRequestRulesURL")
	}
	if host == "" {
		return nil, errors.New("host is required for a full url on GetStickRequestRulesURL")
	}

	base, err := o.Build()
	if err != nil {
		return nil, err
	}

	base.Scheme = scheme
	base.Host = host
	return base, nil
}

// StringFull returns the string representation of a complete url
func (o *GetStickRequestRulesURL) StringFull(scheme, host string) string {
	return o.Must(o.BuildFull(scheme, host)).String()
}
