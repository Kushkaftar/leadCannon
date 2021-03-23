package publiser

import (
	"fmt"
	"log"
	"math/rand"

	"github.com/Kushkaftar/leadCannon/internal/models"
	"github.com/Kushkaftar/leadCannon/pkg/advert/approved"
	"github.com/Kushkaftar/leadCannon/pkg/advert/redject"
	"github.com/Kushkaftar/leadCannon/pkg/advert/trash"
	"github.com/Kushkaftar/leadCannon/pkg/redirect"
)

func Run(c *models.Config) {

	fmt.Println(c)
	for i := 0; i < c.Main.Sum; i++ {

		r, err := request(c)
		if err != nil {
			log.Fatal(err)
		}
		// fmt.Println(r)
		// false, rand, trash, redject, approved

		// TODO: переделать
		status := c.Main.Advertiser
		if status == "rand" {
			option := []string{"false", "trash", "redject", "approved"}
			status = option[rand.Intn(len(option))]
		}
		switch status {
		case "false":
			fmt.Println("false")
		case "trash":
			trash.Trash(r, c)
		case "approved":
			approved.Approve(r, c)
		case "redject":
			redject.Redject(r, c)
		}
	}

	if c.Redirect.Run {
		redirect.Redirect(c)
	}
}
