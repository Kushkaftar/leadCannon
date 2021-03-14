package publiser

import "github.com/Kushkaftar/leadCannon/internal/models"

// lead ...
func lead(hash string) *models.Lead {
	var l models.Lead
	l.Name = "test"
	l.Phone = "+79100000000"
	l.CountryCode = "RU"
	l.IP = "176.215.52.206"
	l.FlowUUID = hash
	return &l
}
