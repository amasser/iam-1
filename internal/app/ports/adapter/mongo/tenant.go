package mongo

import (
	"github.com/maurofran/iam/internal/app/domain/model"
	"github.com/pkg/errors"
	db "upper.io/db.v3"
)

const tenantCollection = "tenants"

// TenantRepository implements the tenant repository interfance.
type TenantRepository struct {
	Database db.Database `inject:""`
}

func (r *TenantRepository) collection() db.Collection {
	return r.Database.Collection(tenantCollection)
}

// Add will add a tenant to tenant repository virtual collection
func (r *TenantRepository) Add(tenant *model.Tenant) error {
	_, err := r.collection().Insert(tenant)
	if err != nil {
		return errors.Wrapf(err, "error while adding tenant to collection %s", tenantCollection)
	}
	return nil
}

// Update will update an existing tenant.
func (r *TenantRepository) Update(tenant *model.Tenant) error {
	res := r.collection().Find(db.Cond{"tenantId": tenant.ID})
	err := res.Update(tenant)
	if err != nil {
		return errors.Wrapf(err, "error while updating tenant in collection %s", tenantCollection)
	}
	return err
}

// Remove will remove a tenant from repository.
func (r *TenantRepository) Remove(tenant *model.Tenant) error {
	res := r.collection().Find(db.Cond{"tenantId": tenant.ID})
	err := res.Delete()
	if err != nil {
		return errors.Wrapf(err, "error while removing tenant from collection %s", tenantCollection)
	}
	return err
}

// TenantWithID will find a tenant by his unique identifier.
func (r *TenantRepository) TenantWithID(tenantID model.TenantID) (*model.Tenant, error) {
	res := r.collection().Find(db.Cond{"tenantId": tenantID})
	var tenant model.Tenant
	err := res.One(&tenant)
	if err == db.ErrNoMoreRows {
		return nil, nil
	}
	if err != nil {
		return nil, errors.Wrapf(err, "error while retrieving tenant from collection %s", tenantCollection)
	}
	return &tenant, nil
}

// TenantWithName will find a tenant by his unique name.
func (r *TenantRepository) TenantWithName(name string) (*model.Tenant, error) {
	res := r.collection().Find(db.Cond{"name": name})
	var tenant model.Tenant
	err := res.One(&tenant)
	if err == db.ErrNoMoreRows {
		return nil, nil
	}
	if err != nil {
		return nil, errors.Wrapf(err, "error while retrieving tenant from collection %s", tenantCollection)
	}
	return &tenant, nil
}
