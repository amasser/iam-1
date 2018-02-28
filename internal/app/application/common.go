package application

import "github.com/maurofran/iam/internal/app/domain/model"

func makeContactInformation(email, streetName, buildingNumber, postalCode, town, stateProvince, countryCode,
	primaryTelephoneNumber, secondaryTelephoneNumber string) (model.ContactInformation, error) {
	emailAddress, err := model.MakeEmailAddress(email)
	if err != nil {
		return model.ContactInformation{}, err
	}
	postalAddress, err := model.MakePostalAddress(streetName, buildingNumber, postalCode, town, stateProvince, countryCode)
	if err != nil {
		return model.ContactInformation{}, err
	}
	primaryTelephone, err := model.MakeTelephone(primaryTelephoneNumber)
	if err != nil {
		return model.ContactInformation{}, err
	}
	secondaryTelephone := model.Telephone("")
	if secondaryTelephoneNumber != "" {
		secondaryTelephone, err = model.MakeTelephone(secondaryTelephoneNumber)
		if err != nil {
			return model.ContactInformation{}, err
		}
	}
	return model.MakeContactInformation(postalAddress, emailAddress, primaryTelephone, secondaryTelephone)
}
