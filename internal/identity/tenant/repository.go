package tenant

import db "upper.io/db.v3"

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
	panic("Not implemented")
}

func (r *repository) Update(tenant *Tenant) error {
	panic("Not implemented")
}

func (r *repository) Remove(tenant *Tenant) error {
	panic("Not implemented")
}

func (r *repository) FindByID(id ID) (*Tenant, error) {
	panic("Not implemented")
}

func (r *repository) FindByName(name string) (*Tenant, error) {
	panic("Not implemented")
}
