package mock

import "github.com/maurofran/iam"

// RoleRepository is the mock struct for role repository.
type RoleRepository struct {
	AddFn            func(*iam.Role) error
	AddInvoked       bool
	UpdateFn         func(*iam.Role) error
	UpdateInvoked    bool
	RemoveFn         func(*iam.Role) error
	RemoveInvoked    bool
	RoleNamedFn      func(iam.TenantID, string) (*iam.Role, error)
	RoleNamedInvoked bool
	AllRolesFn       func(iam.TenantID) (iam.Roles, error)
	AllRolesInvoked  bool
}

// Add is the mock method.
func (r *RoleRepository) Add(role *iam.Role) error {
	r.AddInvoked = true
	return r.AddFn(role)
}

// Update is the mock method.
func (r *RoleRepository) Update(role *iam.Role) error {
	r.UpdateInvoked = true
	return r.UpdateFn(role)
}

// Remove is the mock method.
func (r *RoleRepository) Remove(role *iam.Role) error {
	r.RemoveInvoked = true
	return r.RemoveFn(role)
}

// RoleNamed is the mock method.
func (r *RoleRepository) RoleNamed(tenantID iam.TenantID, name string) (*iam.Role, error) {
	r.RoleNamedInvoked = true
	return r.RoleNamedFn(tenantID, name)
}

// AllRoles is the mock method.
func (r *RoleRepository) AllRoles(tenantID iam.TenantID) (iam.Roles, error) {
	r.AllRolesInvoked = true
	return r.AllRolesFn(tenantID)
}
