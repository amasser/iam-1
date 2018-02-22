package model

import (
	"fmt"
	"regexp"

	"github.com/maurofran/kit/assert"
)

var (
	emailPattern     = regexp.MustCompile("\\w+([-+.']\\w+)*@\\w+([-.]\\w+)*\\.\\w+([-.]\\w+)*")
	telephonePattern = regexp.MustCompile("((\\(\\d{3}\\))|(\\d{3}-))\\d{3}-\\d{4}")
)

// EmailAddress is the value object representing an e-mail address.
//
// The default value of EmailAddress value object is can be used to represent an empty email address.
type EmailAddress string

// MakeEmailAddress will build a new e-mail address, returning an error if the supplied value
// is not a valid e-mail address.
func MakeEmailAddress(address string) (EmailAddress, error) {
	if err := assert.NotEmpty(address, "address"); err != nil {
		return "", err
	}
	if err := assert.Matches(address, emailPattern, "address"); err != nil {
		return "", err
	}
	return EmailAddress(address), nil
}

// IsZero will check if the receiver email address is zero value.
func (e EmailAddress) IsZero() bool {
	return e == ""
}

// Telephone is the value object representing a telephone number.
//
// The default value of Telephone can be used to represent an empty telephone number.
type Telephone string

// MakeTelephone will build a telephone number, returning an error if supplied value is not
// a valid telephone number.
func MakeTelephone(number string) (Telephone, error) {
	if err := assert.NotEmpty(number, "number"); err != nil {
		return "", err
	}
	if err := assert.Matches(number, telephonePattern, "number"); err != nil {
		return "", err
	}
	return Telephone(number), nil
}

// IsZero will check if the receiver telephone number is zero value.
func (t Telephone) IsZero() bool {
	return t == ""
}

// PostalAddress is the value object representing a postal address.
//
// The default value of PostalAddress can be used to represent an empty postal address.
type PostalAddress struct {
	StreetName     string
	BuildingNumber string
	PostalCode     string
	Town           string
	StateProvince  string
	CountryCode    string
}

// MakePostalAddress will make a new postal address with supplied values.
func MakePostalAddress(streetName, buildingNumber, postalCode, town, stateProvince, countryCode string) (PostalAddress, error) {
	if err := assert.NotEmpty(streetName, "streetName"); err != nil {
		return PostalAddress{}, nil
	}
	if err := assert.NotEmpty(postalCode, "postalCode"); err != nil {
		return PostalAddress{}, nil
	}
	if err := assert.NotEmpty(town, "town"); err != nil {
		return PostalAddress{}, nil
	}
	if err := assert.NotEmpty(stateProvince, "stateProvince"); err != nil {
		return PostalAddress{}, nil
	}
	if err := assert.NotEmpty(countryCode, "countryCode"); err != nil {
		return PostalAddress{}, nil
	}
	return PostalAddress{}, nil
}

// IsZero will check if receiver postal address is zero value.
func (pa PostalAddress) IsZero() bool {
	return pa.StreetName == "" && pa.BuildingNumber == "" && pa.PostalCode == "" && pa.Town == "" && pa.StateProvince == "" &&
		pa.CountryCode == ""
}

// FullName is the value object used to provide a full name of person.
//
// The default value of FullName value object is not a valid value and cannot be used.
type FullName struct {
	FirstName string
	LastName  string
}

// MakeFullName will build a new full name with supplied parameters.
// It returns an error if either first name or last name are empty.
func MakeFullName(firstName, lastName string) (FullName, error) {
	if err := assert.NotEmpty(firstName, "firstName"); err != nil {
		return FullName{}, err
	}
	if err := assert.NotEmpty(lastName, "lastName"); err != nil {
		return FullName{}, err
	}
	return FullName{firstName, lastName}, nil
}

// FormattedName will return the formatted version of full name.
func (fn FullName) FormattedName() string {
	return fmt.Sprintf("%s %s", fn.FirstName, fn.LastName)
}

// IsZero will check if receiver full name is zero value.
func (fn FullName) IsZero() bool {
	return fn.FirstName == "" && fn.LastName == ""
}
