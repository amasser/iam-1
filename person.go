package iam

// Person is an entity represnting a person.
type Person struct {
	FullName           FullName           `bson:"fullName"`
	ContactInformation ContactInformation `bson:"contactInformation"`
}

// FullName is the value object representing the full name.
type FullName struct {
	FirstName string `bson:"firstName"`
	LastName  string `bson:"lastName"`
}

// ContactInformation is the value object holding the contact information
type ContactInformation struct {
	EmailAddress       EmailAddress  `bson:"emailAddress"`
	PostalAddress      PostalAddress `bson:"postalAddress,omitempty"`
	PrimaryTelephone   Telephone     `bson:"primaryTelephone,omitempty"`
	SecondaryTelephone Telephone     `bson:"secondaryTelephone,omitempty"`
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

// EmailAddress is the value object for email address.
type EmailAddress string

// Telephone is the value object for telephone number
type Telephone string
