package main

import (
	"context"
	"fmt"
	"github.com/go-resty/resty/v2"
	"net/http"
	"net/url"
	"os"
	"path"
	"strconv"
)

const APIHost = "169.254.169.254"
const APIProto = "http"
const APIVersion = "v1"

type ClientCreateOptions struct {
	HTTPClient *http.Client

	BaseURLOverride string
	VersionOverride string

	UserAgentPrefix string

	DisableTokenInit bool
}

type Client struct {
	resty *resty.Client

	apiBaseURL  string
	apiProtocol string
	apiVersion  string
	userAgent   string
}

func NewClient(ctx context.Context, opts *ClientCreateOptions) (*Client, error) {
	var result Client

	shouldUseHTTPClient := false
	shouldSkipTokenGeneration := false
	// We might need to move the version to a subpackage to prevent a cyclic dependency
	userAgent := DefaultUserAgent

	if opts != nil {
		shouldUseHTTPClient = opts.HTTPClient != nil
		shouldSkipTokenGeneration = opts.DisableTokenInit

		if opts.BaseURLOverride != "" {
			result.SetBaseURL(opts.BaseURLOverride)
		}

		if opts.VersionOverride != "" {
			result.SetVersion(opts.VersionOverride)
		}

		if opts.UserAgentPrefix != "" {
			userAgent = fmt.Sprintf("%s %s", opts.UserAgentPrefix, userAgent)
		}
	}

	if shouldUseHTTPClient {
		result.resty = resty.NewWithClient(opts.HTTPClient)
	} else {
		result.resty = resty.New()
	}

	if debugEnv, ok := os.LookupEnv("LINODE_DEBUG"); ok {
		debugBool, err := strconv.ParseBool(debugEnv)
		if err != nil {
			return nil, fmt.Errorf("failed to parse debug bool: %s", err)
		}
		result.resty.SetDebug(debugBool)
	}

	result.updateHostURL()

	result.SetUserAgent(userAgent)

	if !shouldSkipTokenGeneration {
		if _, err := result.RefreshToken(ctx); err != nil {
			return nil, fmt.Errorf("failed to refresh metadata token: %s", err)
		}
	}

	return &result, nil
}

func (c *Client) UseToken(token string) *Client {
	c.resty.SetHeader("X-Metadata-Token", token)
	return c
}

func (c *Client) RefreshToken(ctx context.Context) (*Client, error) {
	token, err := c.GenerateToken(ctx, GenerateTokenOptions{})
	if err != nil {
		return nil, fmt.Errorf("failed to generate metadata token: %w", err)
	}

	c.UseToken(token.Token)

	return c, nil
}

func (c *Client) SetBaseURL(baseURL string) *Client {
	baseURLPath, _ := url.Parse(baseURL)

	c.apiBaseURL = path.Join(baseURLPath.Host, baseURLPath.Path)
	c.apiProtocol = baseURLPath.Scheme

	c.updateHostURL()

	return c
}

func (c *Client) SetVersion(version string) *Client {
	c.apiVersion = version

	c.updateHostURL()

	return c
}

func (c *Client) updateHostURL() {
	apiProto := APIProto
	baseURL := APIHost
	apiVersion := APIVersion

	if c.apiBaseURL != "" {
		baseURL = c.apiBaseURL
	}

	if c.apiVersion != "" {
		apiVersion = c.apiVersion
	}

	if c.apiProtocol != "" {
		apiProto = c.apiProtocol
	}

	c.resty.SetBaseURL(fmt.Sprintf("%s://%s/%s", apiProto, baseURL, apiVersion))
}

// R wraps resty's R method
func (c *Client) R(ctx context.Context) *resty.Request {
	return c.resty.R().
		ExpectContentType("application/json").
		SetHeader("Content-Type", "application/json").
		SetContext(ctx)
}

func (c *Client) SetUserAgent(userAgent string) *Client {
	c.userAgent = userAgent
	c.resty.SetHeader("User-Agent", c.userAgent)

	return c
}
