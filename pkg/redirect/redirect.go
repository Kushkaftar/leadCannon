package redirect

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"os"

	"github.com/Kushkaftar/leadCannon/internal/models"
)

// Redirect ...
func Redirect(c *models.Config) {

	req, err := http.NewRequest("GET", "http://"+c.Redirect.Redirect_Domain+"/r/"+c.Redirect.UUID, nil)
	req.Header.Add("X-Forwarded-For", faker.r)

	resp, err := http.Get(
		"http://" + c.Redirect.Redirect_Domain + "/r/" + c.Redirect.UUID)
	if err != nil {
		log.Fatalln(err)
		// return "", err
	}
	defer resp.Body.Close()
	fmt.Println(resp.Header)
	io.Copy(os.Stdout, resp.Body)
}
