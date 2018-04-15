package mongo

import (
	"github.com/maurofran/iam"
	"github.com/pkg/errors"
	"gopkg.in/mgo.v2"
)

// Client is the access to client repository implementation
type Client struct {
	db       *mgo.Session
	url      string
	database string
	tr       tenantRepository
	ur       userRepository
	gr       groupRepository
	rr       roleRepository
}

// NewClient will create a new client instance.
func NewClient(url string) *Client {
	c := &Client{url: url}
	c.tr.client = c
	c.ur.client = c
	c.gr.client = c
	c.rr.client = c
	return c
}

// WithDatabase will enrich the client connection with supplied database name.
func (c *Client) WithDatabase(database string) *Client {
	c.database = database
	return c
}

// TenantRepository is the accessor for tenant repository implementation with MongoDB.
func (c *Client) TenantRepository() iam.TenantRepository {
	return &c.tr
}

// UserRepository is the accessor for user repository implementation with MongoDB.
func (c *Client) UserRepository() iam.UserRepository {
	return &c.ur
}

// GroupRepository is the accessor for the group repository implementation with MongoDB.
func (c *Client) GroupRepository() iam.GroupRepository {
	return &c.gr
}

// RoleRepository is the accessor for the role repository implementation with MongoDB.
func (c *Client) RoleRepository() iam.RoleRepository {
	return &c.rr
}

// Open will open the database connection.
func (c *Client) Open() error {
	db, err := mgo.Dial(c.url)
	if err != nil {
		return errors.Wrapf(err, "Error occurred while connecting to %s", c.url)
	}
	c.db = db
	if err := c.tr.init(); err != nil {
		return err
	}
	if err := c.ur.init(); err != nil {
		return err
	}
	if err := c.gr.init(); err != nil {
		return err
	}
	return c.rr.init()
}

// Close will close the underlying mongo session.
func (c *Client) Close() {
	c.db.Close()
}
