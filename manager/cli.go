package manager

import (
	"github.com/ernestio/ernest-cli/model"

	eclient "github.com/ernestio/ernest-go-sdk/client"
	econfig "github.com/ernestio/ernest-go-sdk/config"
)

// Client : ...
type Client struct {
	cli          *eclient.Client
	cfg          *model.Config
	user         *User
	session      *Session
	notification *Notification
	policy       *Policy
	role         *Role
	project      *Project
	env          *Environment
	build        *Build
	logger       *Logger
	report       *Report
}

// NewFromCreds ...
func NewFromCreds(config *model.Config) *Client {
	client := eclient.New(
		econfig.New(config.URL).WithCredentials(config.User, config.Password),
	)
	return &Client{cli: client, cfg: config}
}

// NewFromCredsAndVerification ...
func NewFromCredsAndVerification(config *model.Config) *Client {
	client := eclient.New(
		econfig.New(config.URL).
			WithCredentialsAndVerification(config.User, config.Password, config.Verification),
	)
	return &Client{cli: client, cfg: config}
}

// New : ...
func New(config *model.Config) *Client {
	client := eclient.New(
		econfig.New(config.URL).WithToken(config.Token),
	)
	return &Client{cli: client, cfg: config}
}

// User : User wrapper lazy load
func (c *Client) User() *User {
	if c.user == nil {
		c.user = &User{cli: c.cli}
	}
	return c.user
}

// Session : Session wrapper lazy load
func (c *Client) Session() *Session {
	if c.session == nil {
		c.session = &Session{cli: c.cli}
	}
	return c.session
}

// Notification : Notification wrapper lazy load
func (c *Client) Notification() *Notification {
	if c.notification == nil {
		c.notification = &Notification{cli: c.cli}
	}
	return c.notification
}

// Policy : Policy wrapper lazy load
func (c *Client) Policy() *Policy {
	if c.notification == nil {
		c.policy = &Policy{cli: c.cli}
	}
	return c.policy
}

// Role : Role wrapper lazy load
func (c *Client) Role() *Role {
	if c.role == nil {
		c.role = &Role{cli: c.cli}
	}
	return c.role
}

// Project : Project wrapper lazy load
func (c *Client) Project() *Project {
	if c.project == nil {
		c.project = &Project{cli: c.cli}
	}
	return c.project
}

// Environment : Environment wrapper lazy load
func (c *Client) Environment() *Environment {
	if c.env == nil {
		c.env = &Environment{cli: c.cli}
	}
	return c.env
}

// Build : Build wrapper lazy load
func (c *Client) Build() *Build {
	if c.build == nil {
		c.build = &Build{cli: c.cli}
	}
	return c.build
}

// Logger : Logger wrapper lazy load
func (c *Client) Logger() *Logger {
	if c.logger == nil {
		c.logger = &Logger{cli: c.cli}
	}
	return c.logger
}

// Report : Report wrapper lazy load
func (c *Client) Report() *Report {
	if c.report == nil {
		c.report = &Report{cli: c.cli}
	}
	return c.report
}

// Cli : gets the internal eclient.Client
func (c *Client) Cli() *eclient.Client {
	return c.cli
}

// Config : ...
func (c *Client) Config() *model.Config {
	return c.cfg
}
