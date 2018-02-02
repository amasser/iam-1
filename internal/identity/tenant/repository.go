package tenant

// Repository is the interface implemented by objects that expose a "collection" of tenants.
type Repository interface {
	Add(tenant *Tenant) error
	Update(tenant *Tenant) error
	Remove(tenant *Tenant) error
	FindByID(id ID) (*Tenant, error)
	FindByName(name string) (*Tenant, error)
}
