package model

import (
	"time"

	"github.com/maurofran/kit/assert"
)

// ContactInformation is the value object used in order to provide content information for a person.
type ContactInformation struct {
	PostalAddress      PostalAddress `bson:"postalAddress"`
	EmailAddress       EmailAddress  `bson:"emailAddress"`
	PrimaryTelephone   Telephone     `bson:"primaryTelephone"`
	SecondaryTelephone Telephone     `bson:"secondaryTelephone"`
}

// MakeContactInformation will create a new contact information structure with supplied parameters.
func MakeContactInformation(postalAddress PostalAddress, emailAddress EmailAddress,
	primaryTelephone, secondaryTelephone Telephone) (ContactInformation, error) {
	if err := assert.NotZero(emailAddress, "emailAddress"); err != nil {
		return ContactInformation{}, err
	}
	return ContactInformation{postalAddress, emailAddress, primaryTelephone, secondaryTelephone}, nil
}

// IsZero will check if supplied contact information is zero value.
func (ci ContactInformation) IsZero() bool {
	return ci.PostalAddress.IsZero() && ci.EmailAddress.IsZero() && ci.PrimaryTelephone.IsZero() &&
		ci.SecondaryTelephone.IsZero()
}

// WithEmailAddress return a new contact information with supplied email address.
func (ci ContactInformation) WithEmailAddress(emailAddress EmailAddress) (ContactInformation, error) {
	return MakeContactInformation(ci.PostalAddress, emailAddress, ci.PrimaryTelephone, ci.SecondaryTelephone)
}

// WithPostalAddress return a new contact information with supplied postal address.
func (ci ContactInformation) WithPostalAddress(postalAddress PostalAddress) (ContactInformation, error) {
	return MakeContactInformation(postalAddress, ci.EmailAddress, ci.PrimaryTelephone, ci.SecondaryTelephone)
}

// WithPrimaryTelephone return a new contact information with changed primary telephone.
func (ci ContactInformation) WithPrimaryTelephone(primaryTelephone Telephone) (ContactInformation, error) {
	return MakeContactInformation(ci.PostalAddress, ci.EmailAddress, primaryTelephone, ci.SecondaryTelephone)
}

// WithSecondaryTelephone return a new contact information with changed secondary telephone.
func (ci ContactInformation) WithSecondaryTelephone(secondaryTelephone Telephone) (ContactInformation, error) {
	return MakeContactInformation(ci.PostalAddress, ci.EmailAddress, ci.PrimaryTelephone, secondaryTelephone)
}

// Person is an entity used in order to provide data about a person.
type Person struct {
	Name               FullName           `bson:"name"`
	ContactInformation ContactInformation `bson:"contactInformation"`
}

// NewPerson will create a new person for supplied data.
func NewPerson(fullName FullName, contactInformation ContactInformation) (*Person, error) {
	if err := assert.NotZero(fullName, "fullName"); err != nil {
		return nil, err
	}
	if err := assert.NotZero(contactInformation, "contactInformation"); err != nil {
		return nil, err
	}
	return &Person{fullName, contactInformation}, nil
}

// EmailAddress will retrieve the e-mail address of person.
func (p *Person) EmailAddress() EmailAddress {
	return p.ContactInformation.EmailAddress
}

func (p *Person) changeContactInformation(contactInformation ContactInformation) (bool, error) {
	if err := assert.NotZero(contactInformation, "contactInformation"); err != nil {
		return false, err
	}
	if p.ContactInformation != contactInformation {
		p.ContactInformation = contactInformation
		return true, nil
	}
	return false, nil
}

func (p *Person) changeName(fullName FullName) (bool, error) {
	if err := assert.NotZero(fullName, "fullName"); err != nil {
		return false, err
	}
	if p.Name != fullName {
		p.Name = fullName
		return true, nil
	}
	return false, nil
}

// PersonContactInformationChanged is the event raised when person contact information changed.
type PersonContactInformationChanged struct {
	EventVersion       int
	OccurredOn         time.Time
	TenantID           TenantID
	Username           string
	ContactInformation ContactInformation
}

// PersonNameChanged is the event raised when a person name is changed.
type PersonNameChanged struct {
	EventVersion int
	OccurredOn   time.Time
	TenantID     TenantID
	Username     string
	Name         FullName
}
