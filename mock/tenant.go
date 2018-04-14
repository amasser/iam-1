package mock

import "github.com/maurofran/iam"

// TenantRepository is the mock struct for tenant repository.
type TenantRepository struct {
	AddFn              func(*iam.Tenant) error
	AddInvoked         bool
	UpdateFn           func(*iam.Tenant) error
	UpdateInvoked      bool
	RemoveFn           func(*iam.Tenant) error
	RemoveInvoked      bool
	TenantNamedFn      func(string) (*iam.Tenant, error)
	TenantNamedInvoked bool
	TenantOfIDFn       func(iam.TenantID) (*iam.Tenant, error)
	TenantOfIDInvoked  bool
}

// Add will mock the tenant repository add method.
func (t *TenantRepository) Add(tenant *iam.Tenant) error {
	t.AddInvoked = true
	return t.AddFn(tenant)
}

// Update will mock the tenant repository update method.
func (t *TenantRepository) Update(tenant *iam.Tenant) error {
	t.UpdateInvoked = true
	return t.UpdateFn(tenant)
}

// Remove will mock the tenant repository remove method.
func (t *TenantRepository) Remove(tenant *iam.Tenant) error {
	t.RemoveInvoked = true
	return t.RemoveFn(tenant)
}

// TenantNamed will mock the tenant repository tenant named method.
func (t *TenantRepository) TenantNamed(name string) (*iam.Tenant, error) {
	t.TenantNamedInvoked = true
	return t.TenantNamedFn(name)
}

// TenantOfID will mock the tenant repository tenant of id method.
func (t *TenantRepository) TenantOfID(tenantID iam.TenantID) (*iam.Tenant, error) {
	t.TenantOfIDInvoked = true
	return t.TenantOfIDFn(tenantID)
}
