package part1

import "strings"

type TravelDocument struct {
	BirthYear      string
	IssueYear      string
	ExpirationYear string
	Height         string
	HairColor      string
	EyeColor       string
	PassportID     string
	CountryID      string
}

func ParseTravelDocument(input string) TravelDocument {
	fields := strings.Fields(input)
	fieldsMap := map[string]string{}
	for _, f := range fields {
		parts := strings.Split(f, ":")
		key := parts[0]
		value := parts[1]
		fieldsMap[key] = value
	}
	return TravelDocument{
		BirthYear:      fieldsMap["byr"],
		IssueYear:      fieldsMap["iyr"],
		ExpirationYear: fieldsMap["eyr"],
		Height:         fieldsMap["hgt"],
		HairColor:      fieldsMap["hcl"],
		EyeColor:       fieldsMap["ecl"],
		PassportID:     fieldsMap["pid"],
		CountryID:      fieldsMap["cid"],
	}
}

func (td TravelDocument) Passport() bool {
	return td.BirthYear != "" &&
		td.IssueYear != "" &&
		td.ExpirationYear != "" &&
		td.Height != "" &&
		td.HairColor != "" &&
		td.EyeColor != "" &&
		td.PassportID != ""
}
