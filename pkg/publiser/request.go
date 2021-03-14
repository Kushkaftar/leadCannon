package publiser

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/Kushkaftar/leadCannon/internal/models"
)

func request(c *models.Config) (string, error) {

	l := lead(c.Publisher.Hash)

	bytesRepresentation, err := json.Marshal(l)
	if err != nil {
		log.Fatalln(err)
		return "", err
	}

	// TODO: переделать
	resp, err := http.Post(
		c.Main.URL+"/api/lead/create/publisher?api_key="+c.Publisher.API_Key,
		"application/json",
		bytes.NewBuffer(bytesRepresentation))
	if err != nil {
		log.Fatalln(err)
		return "", err
	}

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatalln(err)
		return "", err
	}

	var r models.Response
	// TODO: переделать
	err = json.Unmarshal(body, &r)
	if err != nil {
		fmt.Println(err.Error())
	}
	// fmt.Println(r)
	var data models.DataResponse
	if r.Success {

		err = json.Unmarshal(body, &data)
		if err != nil {
			fmt.Println(err.Error())
		}
		// fmt.Println(data.Data.LeadUUID)
	}
	// str := string(body)
	str := data.Data.LeadUUID
	return str, nil
}