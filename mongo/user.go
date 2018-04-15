package mongo

import (
	"github.com/maurofran/iam"
	"github.com/pkg/errors"
	mgo "gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

const users = "users"

type userRepository struct {
	client *Client
}

func (r *userRepository) init() error {
	s := r.client.db.Copy()
	defer s.Close()
	c := s.DB(r.client.database).C(users)
	if err := c.EnsureIndex(mgo.Index{Key: []string{"tenantId", "username"}, Unique: true, Name: "ixu_tenantId_username"}); err != nil {
		return errors.Wrap(err, "An error occurred while ensuring index ixu_tenantId_username")
	}
	return nil
}

// Add will add a user to repository.
func (r *userRepository) Add(u *iam.User) error {
	s := r.client.db.Copy()
	defer s.Close()
	c := s.DB(r.client.database).C(users)
	if err := c.Insert(u); err != nil {
		return errors.Wrapf(err, "An error occurred while inserting user %s", u)
	}
	return nil
}

// Update will update a user in repository.
func (r *userRepository) Update(u *iam.User) error {
	s := r.client.db.Copy()
	defer s.Close()
	c := s.DB(r.client.database).C(users)
	if err := c.Update(bson.M{"tenantId": u.TenantID, "username": u.Username}, bson.M{"$set": u}); err != nil {
		return errors.Wrapf(err, "An error occurred while updating user %s", u)
	}
	return nil
}

// Remove will remove a user from repository.
func (r *userRepository) Remove(u *iam.User) error {
	s := r.client.db.Copy()
	defer s.Close()
	c := s.DB(r.client.database).C(users)
	if err := c.Remove(bson.M{"tenantId": u.TenantID, "username": u.Username}); err != nil {
		return errors.Wrapf(err, "An error occurred while removing user %s", u)
	}
	return nil
}

// UserWithUsername will retrieve a user by his unique username.
func (r *userRepository) UserWithUsername(tID iam.TenantID, username string) (*iam.User, error) {
	s := r.client.db.Copy()
	defer s.Close()
	c := s.DB(r.client.database).C(users)
	u := new(iam.User)
	if err := c.Find(bson.M{"tenantId": tID, "username": username}).One(&u); err != nil {
		if err == mgo.ErrNotFound {
			return nil, nil
		}
		return nil, errors.Wrapf(err, "An error occurred while retrieving user for tenant %s and username %s", tID, username)
	}
	return u, nil
}

// AllSimilarlyNamedUsers will retrieve all users by his first name and last name prefix
func (r *userRepository) AllSimilarlyNamedUsers(tID iam.TenantID, firstNamePrefix, lastNamePrefix string) (iam.Users, error) {
	s := r.client.db.Copy()
	defer s.Close()
	c := s.DB(r.client.database).C(users)
	var uu iam.Users
	query := bson.M{
		"tenantId":                  tID,
		"person.fullName.firstName": bson.M{"$regex": bson.RegEx{Pattern: "^" + firstNamePrefix}},
		"person.fullName.lastName":  bson.M{"$regex": bson.RegEx{Pattern: "^" + lastNamePrefix}},
	}
	if err := c.Find(query).Sort("username").All(uu); err != nil {
		return nil, errors.Wrapf(err, "An error occurred while retrieving users of tenant %s", tID)
	}
	return uu, nil
}
