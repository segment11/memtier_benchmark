package main

import (
	"encoding/json"
	"github.com/jaswdr/faker/v2"
	"strconv"
)

type Mocker struct {
	person     faker.Person
	company    faker.Company
	address    faker.Address
	phone      faker.Phone
	internet   faker.Internet
	fakeGroups []string
}

func NewMocker() *Mocker {
	m := &Mocker{
		fakeGroups: []string{
			"name",
			"title",
			"address",
			"company",
			"phone",
			"email",
		},
	}

	fake := faker.New()
	m.person = fake.Person()
	m.company = fake.Company()
	m.address = fake.Address()
	m.phone = fake.Phone()
	m.internet = fake.Internet()
	return m
}

func (m *Mocker) FakerValue(valueLength int) string {
	str := ""

	for len(str) < valueLength {
		for _, group := range m.fakeGroups {
			if "name" == group {
				str += m.person.Name()
			} else if "title" == group {
				str += m.person.Title()
			} else if "address" == group {
				str += m.address.StreetAddress()
			} else if "company" == group {
				str += m.company.Name()
			} else if "phone" == group {
				str += m.phone.Number()
			} else if "email" == group {
				str += m.internet.Email()
			}

			if len(str) >= valueLength {
				break
			}
		}
	}

	return str
}
func (m *Mocker) FakerValueAsJson(valueLength int) string {
	kv := make(map[string]map[string]string)
	strLen := 0

	loop := 0
	for strLen < valueLength {
		subKey := "loop" + strconv.Itoa(loop)
		subKv := make(map[string]string)
		kv[subKey] = subKv
		for _, group := range m.fakeGroups {
			if "name" == group {
				name := m.person.Name()
				subKv["person-name"] = name
				strLen += len(name)
			} else if "title" == group {
				title := m.person.Title()
				subKv["person-title"] = title
				strLen += len(title)
			} else if "address" == group {
				streetAddress := m.address.StreetAddress()
				subKv["address-streetAddress"] = streetAddress
				strLen += len(streetAddress)
			} else if "company" == group {
				name := m.company.Name()
				subKv["company-name"] = name
				strLen += len(name)
			} else if "phone" == group {
				number := m.phone.Number()
				subKv["phone-number"] = number
				strLen += len(number)
			} else if "email" == group {
				email := m.internet.Email()
				subKv["internet-email"] = email
				strLen += len(email)
			}

			if strLen >= valueLength {
				break
			}
		}
		loop++
	}

	// to json encode
	jsonValue, err := json.Marshal(kv)
	if err != nil {
		panic(err)
	}
	return string(jsonValue)
}
