package gohttp

import (
	"net/http"
	"time"
)

type clientBuilder struct {
	maxIdleConnection int
	connectionTimeout time.Duration
	requestTimeout    time.Duration
	disableTimeouts   bool
	headers           http.Header
}

type ClientBuilder interface {
	SetHeaders(headers http.Header) ClientBuilder
	SetConnectionTimeout(timeout time.Duration) ClientBuilder
	SetRequestTimeout(timeout time.Duration) ClientBuilder
	SetMaxIdleConnections(connections int) ClientBuilder
	DisableTimeouts(disable bool) ClientBuilder
	Build() Client
}

func NewBuilder() ClientBuilder {
	builder := &clientBuilder{}
	return builder
}

func (c *clientBuilder) Build() Client {
	client := httpClient{
		builder: c,
	}
	return &client
}

func (c *clientBuilder) DisableTimeouts(disable bool) ClientBuilder {
	c.disableTimeouts = disable
	return c
}

func (c *clientBuilder) SetHeaders(headers http.Header) ClientBuilder {
	c.headers = headers
	return c
}

func (c *clientBuilder) SetConnectionTimeout(timeout time.Duration) ClientBuilder {
	c.connectionTimeout = timeout
	return c
}

func (c *clientBuilder) SetRequestTimeout(timeout time.Duration) ClientBuilder {
	c.requestTimeout = timeout
	return c
}

func (c *clientBuilder) SetMaxIdleConnections(connections int) ClientBuilder {
	c.maxIdleConnection = connections
	return c
}
