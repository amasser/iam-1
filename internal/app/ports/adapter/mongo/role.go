package mongo

import (
	"github.com/maurofran/iam/internal/app/domain/model"
	"github.com/pkg/errors"
	db "upper.io/db.v3"
)

const roleCollection = "roles"

// RoleRepository is the MongoDB implementation of role repository.
type RoleRepository struct {
	Database db.Database `inject:""`
}

func (r *RoleRepository) collection() db.Collection {
	return r.Database.Collection(roleCollection)
}

// Add will add a new role to repository.
func (r *RoleRepository) Add(role *model.Role) error {
	_, err := r.collection().Insert(role)
	if err != nil {
		return errors.Wrapf(err, "error while adding role to collection %s", roleCollection)
	}
	return nil
}

// Update will update a role already in repository.
func (r *RoleRepository) Update(role *model.Role) error {
	res := r.collection().Find(db.Cond{"tenantId": role.TenantID, "name": role.Name})
	err := res.Update(role)
	if err != nil {
		return errors.Wrapf(err, "error while updating role in collection %s", roleCollection)
	}
	return nil
}

// Remove will remove the role from repository.
func (r *RoleRepository) Remove(role *model.Role) error {
	res := r.collection().Find(db.Cond{"tenantId": role.TenantID, "name": role.Name})
	err := res.Delete()
	if err != nil {
		return errors.Wrapf(err, "error while removing role from collection %s", roleCollection)
	}
	return nil
}

// AllRoles will retrieve all the roles for specific tenant.
func (r *RoleRepository) AllRoles(tenantID model.TenantID) ([]*model.Role, error) {
	res := r.collection().Find(db.Cond{"tenantId": tenantID})
	var roles []*model.Role
	err := res.All(&roles)
	if err != nil {
		return nil, errors.Wrapf(err, "error while retrieving roles for tenant %s in collection %s", tenantID, roleCollection)
	}
	return roles, nil
}

// RoleNamed will retrieve the role with supplied name.
func (r *RoleRepository) RoleNamed(tenantID model.TenantID, name string) (*model.Role, error) {
	res := r.collection().Find(db.Cond{"tenantId": tenantID, "name": name})
	var role model.Role
	err := res.One(&role)
	if err == db.ErrNoMoreRows {
		return nil, nil
	}
	if err != nil {
		return nil, errors.Wrapf(err, "error while retrieving role with name %s in tenant %s from collection %s", name, tenantID, roleCollection)
	}
	return &role, nil
}
