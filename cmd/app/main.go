package main

import (
	"fmt"
	"log"

	"github.com/Kushkaftar/leadCannon/pkg/config"
	"github.com/Kushkaftar/leadCannon/pkg/publiser"
	"github.com/go-playground/validator/v10"
)

// var validate *validator.Validate

func main() {
	validate := validator.New()

	config, err := config.NewConfig()
	if err != nil {
		fmt.Println(err)
	}
	err = validate.Struct(config)
	if err != nil {
		log.Fatalln(err)
	} else {
		publiser.Run(config)
		// fmt.Println(config)
	}
	//fmt.Printf("%T\n", var)
}
