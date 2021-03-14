package publiser

import (
	"strings"

	"github.com/Kushkaftar/leadCannon/internal/models"
	"github.com/bxcodec/faker/v3"
)

// lead ...
func lead(c *models.Config) *models.Lead {
	var l models.Lead
	l.Name = faker.FirstName()

	phone := faker.Phonenumber()
	phone = strings.SplitAfterN(phone, "-", 2)[1]
	l.Phone = "+7-910-" + phone
	// l.Phone = "+7-910-" + strings.SplitAfterN(faker.Phonenumber(), "-", 2)[1]

	l.CountryCode = "RU"
	l.IP = faker.IPv4()
	l.FlowUUID = c.Publisher.Hash

	if c.Main.Label {
		l.SubID1 = faker.FirstName()
		l.SubID2 = faker.LastName()
		l.SubID3 = faker.Email()
		l.UtmCampaign = "camp " + faker.Currency()
		l.UtmContent = "content " + faker.Currency()
		l.UtmMedium = "medium " + faker.Currency()
		l.UtmSource = "source " + faker.Currency()
		l.UtmTerm = "term " + faker.Currency()
	}

	if c.Main.Subaccaunt {
		l.SubAccount = faker.DomainName()
	}
	// fmt.Println(l)
	return &l
}
