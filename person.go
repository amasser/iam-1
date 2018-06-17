package iam

import (
	"fmt"
	"regexp"
)

// Person is an entity represnting a person.
type Person struct {
	FullName           FullName           `bson:"fullName"`
	ContactInformation ContactInformation `bson:"contactInformation"`
}

func (p *Person) String() string {
	return fmt.Sprintf(
		"Person[fullName=%s, contactInformation=%s]",
		p.FullName,
		p.ContactInformation,
	)
}

// NewPerson will create a new person.
// It returns an EINVALID error if either full name or contact information are zero valued.
func NewPerson(fullName FullName, contactInformation ContactInformation) (*Person, error) {
	if fullName.IsZero() {
		return nil, &Error{Code: EINVALID, Message: "full name is required", Op: "NewPerson"}
	}
	if contactInformation.IsZero() {
		return nil, &Error{Code: EINVALID, Message: "contact information is required", Op: "NewPerson"}
	}
	return &Person{
		FullName:           fullName,
		ContactInformation: contactInformation,
	}, nil
}

// FullName is the value object representing the full name.
type FullName struct {
	FirstName  string `bson:"firstName"`
	MiddleName string `bson:"middelName"`
	LastName   string `bson:"lastName"`
}

// IsZero will check if full name is zero value.
func (fn *FullName) IsZero() bool {
	return *fn == FullName{}
}

func (fn *FullName) String() string {
	return fmt.Sprintf(
		"FullName [firstName=%s, middleName=%s, lastName=%s]",
		fn.FirstName,
		fn.MiddleName,
		fn.LastName,
	)
}

// MakeFullName will make a new full name.
func MakeFullName(firstName, middleName, lastName string) (FullName, error) {
	if firstName == "" {
		return FullName{}, &Error{Code: EINVALID, Message: "First name is required", Op: "MakeFullName"}
	}
	if lastName == "" {
		return FullName{}, &Error{Code: EINVALID, Message: "Last name is required", Op: "MakeFullName"}
	}
	return FullName{
		FirstName:  firstName,
		MiddleName: middleName,
		LastName:   lastName,
	}, nil
}

// ContactInformation is the value object holding the contact information
type ContactInformation struct {
	EmailAddress       EmailAddress  `bson:"emailAddress"`
	PostalAddress      PostalAddress `bson:"postalAddress,omitempty"`
	PrimaryTelephone   Telephone     `bson:"primaryTelephone,omitempty"`
	SecondaryTelephone Telephone     `bson:"secondaryTelephone,omitempty"`
}

// IsZero will check if contact information is zero value.
func (ci *ContactInformation) IsZero() bool {
	return *ci == ContactInformation{}
}

func (ci *ContactInformation) String() string {
	return fmt.Sprintf(
		"ContactInformation [emailAddress=%s, postalAddress=%s, primaryTelephone=%s, secondaryTelephone=%s]",
		ci.EmailAddress,
		ci.PostalAddress,
		ci.PrimaryTelephone,
		ci.SecondaryTelephone,
	)
}

// MakeContactInformation will make a contact information instance.
func MakeContactInformation(
	emailAddress EmailAddress,
	postalAddress PostalAddress,
	primaryTelephone Telephone,
	secondaryTelephone Telephone,
) (ContactInformation, error) {
	if emailAddress.IsZero() {
		return ContactInformation{}, &Error{Code: EINVALID, Message: "emailAddress is required", Op: "MakeContactInformation"}
	}
	if postalAddress.IsZero() {
		return ContactInformation{}, &Error{Code: EINVALID, Message: "postalAddress is required", Op: "MakeContactInformation"}
	}
	if primaryTelephone.IsZero() {
		return ContactInformation{}, &Error{Code: EINVALID, Message: "primaryTelephone is required", Op: "MakeContactInformation"}
	}
	return ContactInformation{
		EmailAddress:       emailAddress,
		PostalAddress:      postalAddress,
		PrimaryTelephone:   primaryTelephone,
		SecondaryTelephone: secondaryTelephone,
	}, nil
}

// PostalAddress is the value object used to manage a postal address.
type PostalAddress struct {
	StreetName     string `bson:"streetName"`
	BuildingNumber string `bson:"buildingNumber"`
	PostalCode     string `bson:"postalCode"`
	Town           string `bson:"town"`
	StateProvince  string `bson:"stateProvince"`
	CountryCode    string `bson:"countryCode"`
}

// IsZero check if postal address is zero valued.
func (pa *PostalAddress) IsZero() bool {
	return *pa == PostalAddress{}
}

func (pa *PostalAddress) String() string {
	return fmt.Sprintf(
		"PostalAddress [streetName=%s, buildingNumber=%s, postalCode=%s, town=%s, stateProvince=%s, countryCode=%s]",
		pa.StreetName,
		pa.BuildingNumber,
		pa.PostalCode,
		pa.Town,
		pa.StateProvince,
		pa.CountryCode,
	)
}

// MakePostalAddress will make a new postal address with supplied data.
// It returns an EINVALID error if one of street name, postal code, town, state province or country code
// is an empty string.
func MakePostalAddress(streetName, buildingNumber, postalCode, town, stateProvince, countryCode string) (PostalAddress, error) {
	if streetName == "" {
		return PostalAddress{}, &Error{Code: EINVALID, Message: "Street name is required", Op: "MakePostalAddress"}
	}
	if postalCode == "" {
		return PostalAddress{}, &Error{Code: EINVALID, Message: "Postal code is required", Op: "MakePostalAddress"}
	}
	if town == "" {
		return PostalAddress{}, &Error{Code: EINVALID, Message: "Town is required", Op: "MakePostalAddress"}
	}
	if stateProvince == "" {
		return PostalAddress{}, &Error{Code: EINVALID, Message: "State province is required", Op: "MakePostalAddress"}
	}
	if countryCode == "" {
		return PostalAddress{}, &Error{Code: EINVALID, Message: "Country code is required", Op: "MakePostalAddress"}
	}
	return PostalAddress{
		StreetName:     streetName,
		BuildingNumber: buildingNumber,
		PostalCode:     postalCode,
		Town:           town,
		StateProvince:  stateProvince,
		CountryCode:    countryCode,
	}, nil
}

// EmailAddress is the value object for email address.
type EmailAddress string

// Value will return the email address value as string.
func (ea EmailAddress) Value() string {
	return string(ea)
}

// IsZero check if the email address is zero value
func (ea EmailAddress) IsZero() bool {
	return ea == ""
}

func (ea EmailAddress) String() string {
	return fmt.Sprintf("EmailAddress [value=%s]", string(ea))
}

var emailAddress = regexp.MustCompile("^[a-zA-Z0-9.!#$%&’*+/=?^_`{|}~-]+@[a-zA-Z0-9-]+(?:\\.[a-zA-Z0-9-]+)*$")

// MakeEmailAddress will create a new email address value object.
// It return an EINVALID error if the supplied value is an empty string or an invalid email address.
func MakeEmailAddress(value string) (EmailAddress, error) {
	if value == "" {
		return "", &Error{Code: EINVALID, Message: "Email address is required", Op: "MakeEmailAddress"}
	}
	if !emailAddress.MatchString(value) {
		return "", &Error{Code: EINVALID, Message: "Email address format is invalid", Op: "MakeEmailAddress"}
	}
	return EmailAddress(value), nil
}

// Telephone is the value object for telephone number
type Telephone string

// Value will return the telephone value as string
func (t Telephone) Value() string {
	return string(t)
}

// IsZero will check if telephone is zero value.
func (t Telephone) IsZero() bool {
	return t == ""
}

func (t Telephone) String() string {
	return fmt.Sprintf("Telephone [value=%s]", string(t))
}

var telephone = regexp.MustCompile("^(?:\\+\\d{2,3})?\\d{3,}$")

// MakeTelephone will build a new telephone value object instance.
// It returns an EINVALID error if supplied value is empty or has an invalid format.
func MakeTelephone(value string) (Telephone, error) {
	if value == "" {
		return "", &Error{Code: EINVALID, Message: "Telephone number is required", Op: "MakeTelephone"}
	}
	if !telephone.MatchString(value) {
		return "", &Error{Code: EINVALID, Message: "Telephone number or its format is invalid", Op: "MakeTelephone"}
	}
	return Telephone(value), nil
}
