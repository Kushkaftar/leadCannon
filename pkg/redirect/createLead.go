package redirect

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/Kushkaftar/leadCannon/internal/models"
)

// CreateLead ...
func CreateLead(rl models.RedirectLead, c *models.Config) {
	RLead(&rl)
	bytesRepresentation, err := json.Marshal(rl)
	if err != nil {
		log.Fatalln(err)
	}
	// req, err := http.NewRequest("POST", "/api/lead/trash", bytesRepresentation)
	resp, err := http.Post(
		"http://"+c.Redirect.Redirect_Domain+"/l/create",
		"application/json",
		bytes.NewBuffer(bytesRepresentation))
	if err != nil {
		log.Fatalln(err)
	}

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatalln(err)
	}

	fmt.Println(string(body))
}
