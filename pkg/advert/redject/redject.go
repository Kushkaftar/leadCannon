package redject

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/Kushkaftar/leadCannon/internal/models"
)

func Redject(lead_uuid string, c *models.Config) {
	var r models.Redject
	r.OrderID = lead_uuid
	r.RedjectReason = 1
	fmt.Println(r)
	setRedject(r, c)
}

// TODO: переделать
func setRedject(r models.Redject, c *models.Config) {

	bytesRepresentation, err := json.Marshal(r)
	if err != nil {
		log.Fatalln(err)
	}
	// req, err := http.NewRequest("POST", "/api/lead/trash", bytesRepresentation)
	resp, err := http.Post(
		c.Main.URL+"/api/lead/reject?api_key="+c.Advertiser.API_Key,
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
