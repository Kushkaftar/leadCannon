package redirect

import (
	"fmt"
	"log"
	"net/http"
	"strings"

	"github.com/Kushkaftar/leadCannon/internal/models"
	"github.com/bxcodec/faker/v3"
)

// TODO: переделать
// Redirect ...
func Redirect(c *models.Config) {
	client := &http.Client{
		CheckRedirect: func(req *http.Request, via []*http.Request) error {
			return http.ErrUseLastResponse
		}}
	req, err := http.NewRequest("GET", "http://"+c.Redirect.Redirect_Domain+"/r/"+c.Redirect.UUID, nil)
	req.Header.Add("X-Forwarded-For", faker.IPv4())

	resp, err := client.Do(req)
	if err != nil {
		log.Fatalln(err)
		// return "", err
	}
	defer resp.Body.Close()
	redirectURL, err := resp.Location()
	if err != nil {
		//log.Fatalln(err)
		// return "", err
		// TODO: переделать
		fmt.Println(err)
	} else {
		fmt.Println(redirectURL)
		q := redirectURL.Query()
		// fmt.Println(q["r_id"])
		var rl models.RedirectLead
		rl.RID = strings.Join(q["r_id"], "")
		CreateLead(rl, c)
	}

}
