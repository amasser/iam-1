package mongo

import (
	"github.com/maurofran/iam/internal/app/domain/model"
	"github.com/pkg/errors"
	db "upper.io/db.v3"
)

const groupCollection = "groups"

// GroupRepository is the repository of group.
type GroupRepository struct {
	Database db.Database
}

func (r *GroupRepository) collection() db.Collection {
	return r.Database.Collection(groupCollection)
}

// Add will add the supplied group to repository.
func (r *GroupRepository) Add(group *model.Group) error {
	_, err := r.collection().Insert(group)
	if err != nil {
		return errors.Wrapf(err, "unable to insert group into collection %s", groupCollection)
	}
	return nil
}

// Update will update a group already in repository.
func (r *GroupRepository) Update(group *model.Group) error {
	res := r.collection().Find(db.Cond{"tenantId": group.TenantID, "name": group.Name})
	err := res.Update(group)
	if err != nil {
		return errors.Wrapf(err, "unable to update group in collection %s", groupCollection)
	}
	return nil
}

// Remove will remove the group from repository.
func (r *GroupRepository) Remove(group *model.Group) error {
	res := r.collection().Find(db.Cond{"tenantId": group.TenantID, "name": group.Name})
	err := res.Delete()
	if err != nil {
		return errors.Wrapf(err, "unable to delete group from collection %s", groupCollection)
	}
	return nil
}

// AllGroups will retrieve all the groups fro supplied tenant.
func (r *GroupRepository) AllGroups(tenantID model.TenantID) ([]*model.Group, error) {
	res := r.collection().Find(db.Cond{"tenantId": tenantID}).OrderBy("name")
	var groups []*model.Group
	err := res.All(&groups)
	if err != nil {
		return nil, errors.Wrapf(err, "unable to retrieve groups for tenant %s from collection %s", tenantID, groupCollection)
	}
	return groups, nil
}

// GroupNamed will retrieve a single group by his name.
func (r *GroupRepository) GroupNamed(tenantID model.TenantID, name string) (*model.Group, error) {
	res := r.collection().Find(db.Cond{"tenantId": tenantID, "name": name})
	var group model.Group
	err := res.One(&group)
	if err == db.ErrNoMoreRows {
		return nil, nil
	}
	if err != nil {
		return errors.Wrapf(err, "unable to retrieve group for tenant %s and name %s from collection %s", tenantID, name, groupCollection)
	}
	return &group, nil
}
