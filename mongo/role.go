package mongo

import (
	"github.com/maurofran/iam"
	"github.com/pkg/errors"
	mgo "gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

const roles = "roles"

type roleRepository struct {
	client *Client
}

func (r *roleRepository) init() error {
	s := r.client.db.Copy()
	defer s.Close()
	c := s.DB(r.client.database).C(roles)
	if err := c.EnsureIndex(mgo.Index{Key: []string{"tenantId", "name"}, Unique: true, Name: "ixu_tenantId_name"}); err != nil {
		return errors.Wrap(err, "An error occurred while ensuring index ixu_tenantId_name")
	}
	return nil
}

// Add will add a role to repository.
func (r *roleRepository) Add(rl *iam.Role) error {
	s := r.client.db.Copy()
	defer s.Close()
	c := s.DB(r.client.database).C(roles)
	if err := c.Insert(rl); err != nil {
		return errors.Wrapf(err, "An error occurred while adding role %s", rl)
	}
	return nil
}

// Update will update a role from repository.
func (r *roleRepository) Update(rl *iam.Role) error {
	s := r.client.db.Copy()
	defer s.Close()
	c := s.DB(r.client.database).C(roles)
	if err := c.Update(bson.M{"tenantId": rl.TenantID, "name": rl.Name}, bson.M{"$set": rl}); err != nil {
		return errors.Wrapf(err, "An error occurred while updating role %s", rl)
	}
	return nil
}

// Remove will remove a role from repository.
func (r *roleRepository) Remove(rl *iam.Role) error {
	s := r.client.db.Copy()
	defer s.Close()
	c := s.DB(r.client.database).C(roles)
	if err := c.Remove(bson.M{"tenantId": rl.TenantID, "name": rl.Name}); err != nil {
		return errors.Wrapf(err, "An error occurred while removing role %s", rl)
	}
	return nil
}

// RoleNamed will retrieve a role by tenant id and name.
func (r *roleRepository) RoleNamed(tID iam.TenantID, name string) (*iam.Role, error) {
	s := r.client.db.Copy()
	defer s.Close()
	c := s.DB(r.client.database).C(roles)
	rl := new(iam.Role)
	if err := c.Find(bson.M{"tenantId": tID, "name": name}).One(&rl); err != nil {
		if err == mgo.ErrNotFound {
			return nil, nil
		}
		return nil, errors.Wrapf(err, "An error occurred while retrieving role for id %s and name %s", tID, name)
	}
	return rl, nil
}

// Allroles will retrieve all roles for tenant id
func (r *roleRepository) AllRoles(tID iam.TenantID) (iam.Roles, error) {
	s := r.client.db.Copy()
	defer s.Close()
	c := s.DB(r.client.database).C(roles)
	var rr iam.Roles
	if err := c.Find(bson.M{"tenantId": tID}).Sort("name").All(rr); err != nil {
		return nil, errors.Wrapf(err, "An error occurred while retrieving roles for id %s", tID)
	}
	return rr, nil
}
