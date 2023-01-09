// Code generated by go-swagger; DO NOT EDIT.

package rest_client_zrok

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"github.com/go-openapi/runtime"
	httptransport "github.com/go-openapi/runtime/client"
	"github.com/go-openapi/strfmt"

	"github.com/openziti-test-kitchen/zrok/rest_client_zrok/account"
	"github.com/openziti-test-kitchen/zrok/rest_client_zrok/admin"
	"github.com/openziti-test-kitchen/zrok/rest_client_zrok/environment"
	"github.com/openziti-test-kitchen/zrok/rest_client_zrok/invite"
	"github.com/openziti-test-kitchen/zrok/rest_client_zrok/metadata"
	"github.com/openziti-test-kitchen/zrok/rest_client_zrok/share"
)

// Default zrok HTTP client.
var Default = NewHTTPClient(nil)

const (
	// DefaultHost is the default Host
	// found in Meta (info) section of spec file
	DefaultHost string = "localhost"
	// DefaultBasePath is the default BasePath
	// found in Meta (info) section of spec file
	DefaultBasePath string = "/api/v1"
)

// DefaultSchemes are the default schemes found in Meta (info) section of spec file
var DefaultSchemes = []string{"http"}

// NewHTTPClient creates a new zrok HTTP client.
func NewHTTPClient(formats strfmt.Registry) *Zrok {
	return NewHTTPClientWithConfig(formats, nil)
}

// NewHTTPClientWithConfig creates a new zrok HTTP client,
// using a customizable transport config.
func NewHTTPClientWithConfig(formats strfmt.Registry, cfg *TransportConfig) *Zrok {
	// ensure nullable parameters have default
	if cfg == nil {
		cfg = DefaultTransportConfig()
	}

	// create transport and client
	transport := httptransport.New(cfg.Host, cfg.BasePath, cfg.Schemes)
	return New(transport, formats)
}

// New creates a new zrok client
func New(transport runtime.ClientTransport, formats strfmt.Registry) *Zrok {
	// ensure nullable parameters have default
	if formats == nil {
		formats = strfmt.Default
	}

	cli := new(Zrok)
	cli.Transport = transport
	cli.Account = account.New(transport, formats)
	cli.Admin = admin.New(transport, formats)
	cli.Environment = environment.New(transport, formats)
	cli.Invite = invite.New(transport, formats)
	cli.Metadata = metadata.New(transport, formats)
	cli.Share = share.New(transport, formats)
	return cli
}

// DefaultTransportConfig creates a TransportConfig with the
// default settings taken from the meta section of the spec file.
func DefaultTransportConfig() *TransportConfig {
	return &TransportConfig{
		Host:     DefaultHost,
		BasePath: DefaultBasePath,
		Schemes:  DefaultSchemes,
	}
}

// TransportConfig contains the transport related info,
// found in the meta section of the spec file.
type TransportConfig struct {
	Host     string
	BasePath string
	Schemes  []string
}

// WithHost overrides the default host,
// provided by the meta section of the spec file.
func (cfg *TransportConfig) WithHost(host string) *TransportConfig {
	cfg.Host = host
	return cfg
}

// WithBasePath overrides the default basePath,
// provided by the meta section of the spec file.
func (cfg *TransportConfig) WithBasePath(basePath string) *TransportConfig {
	cfg.BasePath = basePath
	return cfg
}

// WithSchemes overrides the default schemes,
// provided by the meta section of the spec file.
func (cfg *TransportConfig) WithSchemes(schemes []string) *TransportConfig {
	cfg.Schemes = schemes
	return cfg
}

// Zrok is a client for zrok
type Zrok struct {
	Account account.ClientService

	Admin admin.ClientService

	Environment environment.ClientService

	Invite invite.ClientService

	Metadata metadata.ClientService

	Share share.ClientService

	Transport runtime.ClientTransport
}

// SetTransport changes the transport on the client and all its subresources
func (c *Zrok) SetTransport(transport runtime.ClientTransport) {
	c.Transport = transport
	c.Account.SetTransport(transport)
	c.Admin.SetTransport(transport)
	c.Environment.SetTransport(transport)
	c.Invite.SetTransport(transport)
	c.Metadata.SetTransport(transport)
	c.Share.SetTransport(transport)
}
