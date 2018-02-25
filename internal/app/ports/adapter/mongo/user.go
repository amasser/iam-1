package mongo

import (
	"github.com/maurofran/iam/internal/app/domain/model"
	"github.com/pkg/errors"
	db "upper.io/db.v3"
)

const userCollection = "users"

// UserRepository is the implmentation of repository for users.
type UserRepository struct {
	Database db.Database
}

func (r *UserRepository) collection() db.Collection {
	return r.Database.Collection(userCollection)
}

// Add a user to virtual collection.
func (r *UserRepository) Add(user *model.User) error {
	_, err := r.collection().Insert(user)
	if err != nil {
		return errors.Wrapf(err, "error occurred while inserting user to collection %s", userCollection)
	}
	return nil
}

// Update will update a user already in virtual collection.
func (r *UserRepository) Update(user *model.User) error {
	res := r.collection().Find(db.Cond{"tenantId": user.TenantID, "username": user.Username})
	err := res.Update(user)
	if err != nil {
		return errors.Wrapf(err, "error occurred while updateing user in collection %s", userCollection)
	}
	return nil
}

// Remove will remove a user from virtual collection.
func (r *UserRepository) Remove(user *model.User) error {
	res := r.collection().Find(db.Cond{"tenantId": user.TenantID, "username": user.Username})
	err := res.Delete()
	if err != nil {
		return errors.Wrapf(err, "error occurred while removing user from collection %s", userCollection)
	}
	return nil
}

// AllUsers will retrieve all the users of virtual collection for tenant.
func (r *UserRepository) AllUsers(tenantID model.TenantID) ([]*model.User, error) {
	res := r.collection().Find(db.Cond{"tenantId": tenantID}).OrderBy("username")
	var users []*model.User
	err := res.All(&users)
	if err != nil {
		return nil, errors.Wrapf(
			err,
			"error occurred while retrieving users of tenant %s in collection %s",
			tenantID,
			userCollection,
		)
	}
	return users, nil
}

// UserWithUsername will retrieve a user for a username.
func (r *UserRepository) UserWithUsername(tenantID model.TenantID, username string) (*model.User, error) {
	res := r.collection().Find(db.Cond{"tenantId": tenantID, "username": username})
	var user model.User
	err := res.One(&user)
	if err == db.ErrNoMoreRows {
		return nil, nil
	}
	if err != nil {
		return nil, errors.Wrapf(
			err,
			"error occurred while retrieving user with username %s in tenant %s from collection %s",
			username,
			tenantID,
			userCollection,
		)
	}
	return &user, nil
}

// UserWithCredentials will retrieve a user with specific credentials.
func (r *UserRepository) UserWithCredentials(tenantID model.TenantID, username string, password string) (*model.User, error) {
	res := r.collection().Find(db.Cond{"tenantId": tenantID, "username": username, "password": password})
	var user model.User
	err := res.One(&user)
	if err == db.ErrNoMoreRows {
		return nil, nil
	}
	if err != nil {
		return nil, errors.Wrapf(
			err,
			"error occurred while retrieveing user with username %s in tenant %s from collection %s",
			tenantID,
			username,
			userCollection,
		)
	}
	return &user, nil
}
