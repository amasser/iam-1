package mongo

import (
	"github.com/maurofran/iam"
	"github.com/pkg/errors"
	mgo "gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

const tenants = "tenants"

type tenantRepository struct {
	client *Client
}

// Init will initialize the repository
func (r *tenantRepository) init() error {
	s := r.client.db.Copy()
	defer s.Close()
	c := s.DB(r.client.database).C(tenants)
	if err := c.EnsureIndex(mgo.Index{Key: []string{"tenantId"}, Unique: true, Name: "ixu_tenantId"}); err != nil {
		return errors.Wrap(err, "An error occurred while ensuring index ixu_tenantId")
	}
	if err := c.EnsureIndex(mgo.Index{Key: []string{"name"}, Unique: true, Name: "ixu_name"}); err != nil {
		return errors.Wrap(err, "An error occurred while ensuring index ixu_name")
	}
	return nil
}

// Add will add a tenant into the repository.
func (r *tenantRepository) Add(t *iam.Tenant) error {
	s := r.client.db.Copy()
	defer s.Close()
	c := s.DB(r.client.database).C(tenants)
	if err := c.Insert(t); err != nil {
		return errors.Wrapf(err, "An error occurred while inserting tenant %s", t)
	}
	return nil
}

// Update will update an existing tenant from repository.
func (r *tenantRepository) Update(t *iam.Tenant) error {
	s := r.client.db.Copy()
	defer s.Close()
	c := s.DB(r.client.database).C(tenants)
	if err := c.Update(bson.M{"tenantId": t.ID}, bson.M{"$set": t}); err != nil {
		return errors.Wrapf(err, "An error occurred while updating tenant %s", t)
	}
	return nil
}

// Remove will remove an existing tenant from repository.
func (r *tenantRepository) Remove(t *iam.Tenant) error {
	s := r.client.db.Copy()
	defer s.Close()
	c := s.DB(r.client.database).C(tenants)
	if err := c.Remove(bson.M{"tenantId": t.ID}); err != nil {
		return errors.Wrapf(err, "An error occurred while removing tenant %s", t)
	}
	return nil
}

// TenantName will retrieve a tenant by his name.
func (r *tenantRepository) TenantNamed(name string) (*iam.Tenant, error) {
	s := r.client.db.Copy()
	defer s.Close()
	c := s.DB(r.client.database).C(tenants)
	t := new(iam.Tenant)
	if err := c.Find(bson.M{"name": name}).One(&t); err != nil {
		if err == mgo.ErrNotFound {
			return nil, nil
		}
		return nil, errors.Wrapf(err, "An error occurred while retrieving tenant for name '%s'", name)
	}
	return t, nil
}

// TenantOfID will retrieve a tenant by his tenant identifier.
func (r *tenantRepository) TenantOfID(tID iam.TenantID) (*iam.Tenant, error) {
	s := r.client.db.Copy()
	defer s.Close()
	c := s.DB(r.client.database).C(tenants)
	t := new(iam.Tenant)
	if err := c.Find(bson.M{"tenantId": tID}).One(&t); err != nil {
		if err == mgo.ErrNotFound {
			return nil, nil
		}
		return nil, errors.Wrapf(err, "An error occurred hiwl retrieving tenant for id %s", tID)
	}
	return t, nil
}
