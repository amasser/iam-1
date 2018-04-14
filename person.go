package iam

// Person is an entity represnting a person.
type Person struct {
	ID                 int
	FullName           FullName
	ContactInformation ContactInformation
}

// FullName is the value object representing the full name.
type FullName struct {
	FirstName string
	LastName  string
}

// ContactInformation is the value object holding the contact information
type ContactInformation struct {
	EmailAddress       EmailAddress
	PostalAddress      PostalAddress
	PrimaryTelephone   Telephone
	SecondaryTelephone Telephone
}

// PostalAddress is the value object used to manage a postal address.
type PostalAddress struct {
	StreetName     string
	BuildingNumber string
	PostalCode     string
	Town           string
	StateProvince  string
	CountryCode    string
}

// EmailAddress is the value object for email address.
type EmailAddress string

// Telephone is the value object for telephone number
type Telephone string
