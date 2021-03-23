package redirect

import (
	"strings"

	"github.com/Kushkaftar/leadCannon/internal/models"

	"github.com/bxcodec/faker/v3"
)

func RLead(rl *models.RedirectLead) {

	rl.CountryCode = "RU"
	rl.Name = faker.FirstName()

	phone := faker.Phonenumber()
	phone = strings.SplitAfterN(phone, "-", 2)[1]
	rl.Phone = "+7-910-" + phone

	rl.Comment = "Artur land" + faker.Email()
}
