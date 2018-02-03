package tenant

import (
	"github.com/pkg/errors"
	db "upper.io/db.v3"
)

const collectionName = "tenants"

// Repository is the interface implemented by objects that expose a "collection" of tenants.
type Repository interface {
	Add(tenant *Tenant) error
	Update(tenant *Tenant) error
	Remove(tenant *Tenant) error
	FindByID(id ID) (*Tenant, error)
	FindByName(name string) (*Tenant, error)
}

// NewRepository will return a new instance of tenant repository.
func NewRepository(db db.Database) Repository {
	return &repository{db.Collection(collectionName)}
}

type repository struct {
	c db.Collection
}

func (r *repository) Add(tenant *Tenant) error {
	_, err := r.c.Insert(tenant)
	if err != nil {
		return errors.Wrapf(err, "Error occurred while inserting tenant %v", tenant)
	}
	return nil
}

func (r *repository) Update(tenant *Tenant) error {
	res := r.c.Find("tenantId", tenant.ID.Value())
	err := res.Update(tenant)
	if err != nil {
		return errors.Wrapf(err, "Error occurred while updating tenant %v", tenant)
	}
	return nil
}

func (r *repository) Remove(tenant *Tenant) error {
	res := r.c.Find("tenantId", tenant.ID.Value())
	err := res.Delete()
	if err != nil {
		return errors.Wrapf(err, "Error occurred while removing tenant %v", tenant)
	}
	return nil
}

func (r *repository) FindByID(id ID) (*Tenant, error) {
	res := r.c.Find("tenantId", id)
	var tenant Tenant
	if err := res.One(&tenant); err != nil {
		return nil, errors.Wrapf(err, "Error occurred while retrieving tenant for id %v", id)
	}
	return &tenant, nil
}

func (r *repository) FindByName(name string) (*Tenant, error) {
	res := r.c.Find("name", name)
	tenant := new(Tenant)
	if err := res.One(tenant); err != nil {
		return nil, errors.Wrapf(err, "Error occurred while retrieving tenant for name '%s'", name)
	}
	return tenant, nil
}
