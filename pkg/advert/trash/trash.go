package trash

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/Kushkaftar/leadCannon/internal/models"
)

func Trash(lead_uuid string, c *models.Config) {
	var t models.Trash
	t.OrderID = lead_uuid
	t.TrashReason = 1
	fmt.Println(t)
	setTresh(t, c)
}

// TODO: переделать
func setTresh(t models.Trash, c *models.Config) {

	bytesRepresentation, err := json.Marshal(t)
	if err != nil {
		log.Fatalln(err)
	}
	// req, err := http.NewRequest("POST", "/api/lead/trash", bytesRepresentation)
	resp, err := http.Post(
		c.Main.URL+"/api/lead/trash?api_key="+c.Advertiser.API_Key,
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
