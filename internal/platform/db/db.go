package db

import (
	"context"
	"time"

	"github.com/pkg/errors"
	"gopkg.in/mgo.v2"
)

// ErrInvalidDBProvided is the error returned in case an invalid database is provided.
var ErrInvalidDBProvided = errors.New("Invalid Database Provided")

// DB is used to support MongoDB connections
type DB struct {
	database *mgo.Database
	session  *mgo.Session
}

// New returns a new DB value for use with MongoDB based on a registered master session.
func New(url string, timeout time.Duration) (*DB, error) {
	if timeout == 0 {
		timeout = 60 * time.Second
	}

	ses, err := mgo.DialWithTimeout(url, timeout)
	if err != nil {
		return nil, errors.Wrapf(err, "mgo.DialWithTimeout(%s, %v)", url, timeout)
	}

	ses.SetMode(mgo.Monotonic, true)

	db := DB{
		database: ses.DB(""),
		session:  ses,
	}
	return &db, nil
}

// Close the database
func (db *DB) Close() {
	db.session.Close()
}

// Copy returns a new copy of database.
func (db *DB) Copy() (*DB, error) {
	ses := db.session.Copy()

	newDB := DB{
		database: ses.DB(""),
		session:  ses,
	}
	return &newDB, nil
}

// Execute is used to execute mongoDB commands.
func (db *DB) Execute(ctx context.Context, collName string, f func(*mgo.Collection) error) error {
	if db == nil || db.session == nil {
		return errors.Wrap(ErrInvalidDBProvided, "db == nil || db.session == nil")
	}
	return f(db.database.C(collName))
}

// ExecuteWithTimeout is used to execute a mongoDB commmand with timeout.
func (db *DB) ExecuteWithTimeout(ctx context.Context, collName string, f func(*mgo.Collection) error, timeout time.Duration) error {
	if db == nil || db.session == nil {
		return errors.Wrap(ErrInvalidDBProvided, "db == nil || db.session == nil")
	}
	db.session.SetSocketTimeout(timeout)
	return f(db.database.C(collName))
}

// StatusCheck validates the status.
func (db *DB) StatusCheck() error {
	return nil
}
