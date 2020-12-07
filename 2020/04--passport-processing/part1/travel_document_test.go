package part1

import "testing"

func TestTravelDocument_passport1(t *testing.T) {
	input := `ecl:gry pid:860033327 eyr:2020 hcl:#fffffd
	byr:1937 iyr:2017 cid:147 hgt:183cm`

	travelDocument := ParseTravelDocument(input)
	wantTravelDocument := TravelDocument{
		BirthYear:      "1937",
		IssueYear:      "2017",
		ExpirationYear: "2020",
		Height:         "183cm",
		HairColor:      "#fffffd",
		EyeColor:       "gry",
		PassportID:     "860033327",
		CountryID:      "147",
	}
	if travelDocument != wantTravelDocument {
		t.Errorf("got %v want %v", travelDocument, wantTravelDocument)
	}
	if !travelDocument.Passport() {
		t.Errorf("got invalid passport want valid passport")
	}
}

func TestTravelDocument_passport2(t *testing.T) {
	input := `iyr:2013 ecl:amb cid:350 eyr:2023 pid:028048884
	hcl:#cfa07d byr:1929`

	travelDocument := ParseTravelDocument(input)
	wantTravelDocument := TravelDocument{
		BirthYear:      "1929",
		IssueYear:      "2013",
		ExpirationYear: "2023",
		HairColor:      "#cfa07d",
		EyeColor:       "amb",
		PassportID:     "028048884",
		CountryID:      "350",
	}
	if travelDocument != wantTravelDocument {
		t.Errorf("got %v want %v", travelDocument, wantTravelDocument)
	}
	if travelDocument.Passport() {
		t.Errorf("got valid passport want invalid passport")
	}
}

func TestTravelDocument_passport3(t *testing.T) {
	input := `hcl:#ae17e1 iyr:2013
	eyr:2024
	ecl:brn pid:760753108 byr:1931
	hgt:179cm`

	travelDocument := ParseTravelDocument(input)
	wantTravelDocument := TravelDocument{
		BirthYear:      "1931",
		IssueYear:      "2013",
		ExpirationYear: "2024",
		Height:         "179cm",
		HairColor:      "#ae17e1",
		EyeColor:       "brn",
		PassportID:     "760753108",
	}
	if travelDocument != wantTravelDocument {
		t.Errorf("got %v want %v", travelDocument, wantTravelDocument)
	}
	if !travelDocument.Passport() {
		t.Errorf("got invalid passport want valid passport")
	}
}

func TestTravelDocument_passport4(t *testing.T) {
	input := `hcl:#cfa07d eyr:2025 pid:166559648
	iyr:2011 ecl:brn hgt:59in`

	travelDocument := ParseTravelDocument(input)
	wantTravelDocument := TravelDocument{
		IssueYear:      "2011",
		ExpirationYear: "2025",
		Height:         "59in",
		HairColor:      "#cfa07d",
		EyeColor:       "brn",
		PassportID:     "166559648",
	}
	if travelDocument != wantTravelDocument {
		t.Errorf("got %v want %v", travelDocument, wantTravelDocument)
	}
	if travelDocument.Passport() {
		t.Errorf("got valid passport want invalid passport")
	}
}
