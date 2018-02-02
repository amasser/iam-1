package tenant

import "github.com/maurofran/kit/assert"
import "github.com/google/uuid"
import "fmt"

// ID is the unique identifier of a tenant.
type ID struct {
	value string
}

// NewID is the factory function that can be used to create a new ID instance with supplied value.
func NewID(id string) (ID, error) {
	if err := assert.NotEmpty(id, "id"); err != nil {
		return ID{}, err
	}
	return ID{value: id}, nil
}

// RandomID will create a new random ID instance.
func RandomID() (ID, error) {
	return NewID(uuid.New().String())
}

// Value will retrieve the unique identifier of this value as a string.
func (id ID) Value() string {
	return id.value
}

// IsZero will check if receiver ID is zero.
func (id ID) IsZero() bool {
	return id.value == ""
}

func (id ID) String() string {
	return fmt.Sprintf("ID [id=%s]", id.value)
}
