package approved

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/Kushkaftar/leadCannon/internal/models"
)

func Approve(lead_uuid string, c *models.Config) {
	var a models.Approve
	a.OrderID = lead_uuid
	fmt.Println(a)
	setApprovet(a, c)
}

// TODO: переделать
func setApprovet(a models.Approve, c *models.Config) {

	bytesRepresentation, err := json.Marshal(a)
	if err != nil {
		log.Fatalln(err)
	}
	// req, err := http.NewRequest("POST", "/api/lead/trash", bytesRepresentation)
	resp, err := http.Post(
		c.Main.URL+"/api/lead/approve?api_key="+c.Advertiser.API_Key,
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
