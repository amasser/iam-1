package iam

// Person is an entity representing a person.
type Person struct {
	FullName           FullName           `bson:"fullName,omitempty"`
	ContactInformation ContactInformation `bson:"contactInformation,omitempty"`
}

// FullName is the value object representing the full name.
type FullName struct {
	FirstName string `bson:"firstName"`
	LastName  string `bson:"lastName"`
}

// ContactInformation is the value object holding the contact information
type ContactInformation struct {
	EmailAddress       EmailAddress  `bson:"emailAddress,omitempty"`
	PostalAddress      PostalAddress `bson:"postalAddress,omitempty"`
	PrimaryTelephone   Telephone     `bson:"primaryTelephone,omitempty"`
	SecondaryTelephone Telephone     `bson:"secondaryTelephone,omitempty"`
}

// WithEmailAddress will create a new contact information with supplied email address.
func (ci *ContactInformation) WithEmailAddress(emailAddress EmailAddress) ContactInformation {
	return ContactInformation{
		EmailAddress:       emailAddress,
		PostalAddress:      ci.PostalAddress,
		PrimaryTelephone:   ci.PrimaryTelephone,
		SecondaryTelephone: ci.SecondaryTelephone,
	}
}

// WithPostalAddress will create a new contact information with supplied postal address.
func (ci *ContactInformation) WithPostalAddress(postalAddress PostalAddress) ContactInformation {
	return ContactInformation{
		EmailAddress:       ci.EmailAddress,
		PostalAddress:      postalAddress,
		PrimaryTelephone:   ci.PrimaryTelephone,
		SecondaryTelephone: ci.SecondaryTelephone,
	}
}

// WithPrimaryTelephone will create a new contact information with supplied primary telephone.
func (ci *ContactInformation) WithPrimaryTelephone(primaryTelephone Telephone) ContactInformation {
	return ContactInformation{
		EmailAddress:       ci.EmailAddress,
		PostalAddress:      ci.PostalAddress,
		PrimaryTelephone:   primaryTelephone,
		SecondaryTelephone: ci.SecondaryTelephone,
	}
}

// WithSecondaryTelephone will create a new contact information with supplied secondary telephone.
func (ci *ContactInformation) WithSecondaryTelephone(secondaryTelephone Telephone) ContactInformation {
	return ContactInformation{
		EmailAddress:       ci.EmailAddress,
		PostalAddress:      ci.PostalAddress,
		PrimaryTelephone:   ci.PrimaryTelephone,
		SecondaryTelephone: secondaryTelephone,
	}
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

// EmailAddress is the type defined for an email address.
type EmailAddress string

// Telephone is the type defined for a telephone number.
type Telephone string
