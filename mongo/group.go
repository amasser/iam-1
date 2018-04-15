package mongo

import (
	"github.com/maurofran/iam"
	"github.com/pkg/errors"
	mgo "gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

const groups = "groups"

type groupRepository struct {
	client *Client
}

func (r *groupRepository) init() error {
	s := r.client.db.Copy()
	defer s.Close()
	c := s.DB(r.client.database).C(groups)
	if err := c.EnsureIndex(mgo.Index{Key: []string{"tenantId", "name"}, Unique: true, Name: "ixu_tenantId_name"}); err != nil {
		return errors.Wrap(err, "An error occurred while ensuring index ixu_tenantId_name")
	}
	return nil
}

// Add will add a group to repository.
func (r *groupRepository) Add(g *iam.Group) error {
	s := r.client.db.Copy()
	defer s.Close()
	c := s.DB(r.client.database).C(groups)
	if err := c.Insert(g); err != nil {
		return errors.Wrapf(err, "An error occurred while adding group %s", g)
	}
	return nil
}

// Update will update a group from repository.
func (r *groupRepository) Update(g *iam.Group) error {
	s := r.client.db.Copy()
	defer s.Close()
	c := s.DB(r.client.database).C(groups)
	if err := c.Update(bson.M{"tenantId": g.TenantID, "name": g.Name}, bson.M{"$set": g}); err != nil {
		return errors.Wrapf(err, "An error occurred while updating group %s", g)
	}
	return nil
}

// Remove will remove a group from repository.
func (r *groupRepository) Remove(g *iam.Group) error {
	s := r.client.db.Copy()
	defer s.Close()
	c := s.DB(r.client.database).C(groups)
	if err := c.Remove(bson.M{"tenantId": g.TenantID, "name": g.Name}); err != nil {
		return errors.Wrapf(err, "An error occurred while removing group %s", g)
	}
	return nil
}

// GroupNamed will retrieve a group by tenant id and name.
func (r *groupRepository) GroupNamed(tID iam.TenantID, name string) (*iam.Group, error) {
	s := r.client.db.Copy()
	defer s.Close()
	c := s.DB(r.client.database).C(groups)
	g := new(iam.Group)
	if err := c.Find(bson.M{"tenantId": tID, "name": name}).One(&g); err != nil {
		if err == mgo.ErrNotFound {
			return nil, nil
		}
		return nil, errors.Wrapf(err, "An error occurred while retrieving group for id %s and name %s", tID, name)
	}
	return g, nil
}

// AllGroups will retrieve all groups for tenant id
func (r *groupRepository) AllGroups(tID iam.TenantID) (iam.Groups, error) {
	s := r.client.db.Copy()
	defer s.Close()
	c := s.DB(r.client.database).C(groups)
	var gg iam.Groups
	if err := c.Find(bson.M{"tenantId": tID}).Sort("name").All(gg); err != nil {
		return nil, errors.Wrapf(err, "An error occurred while retrieving groups for id %s", tID)
	}
	return gg, nil
}
