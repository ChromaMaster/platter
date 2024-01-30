package main

import (
	"fmt"
	"github.com/urfave/cli/v2"
	"log"
	"os"
	"platter/internal/model"
	"platter/internal/repository/ingredient"
)

func main() {
	ingredientsRepository := ingredient.NewInMemIngredientRepository()
	if err := ingredientsRepository.Create(model.NewIngredient(0, "Foo")); err != nil {
		panic(err)
	}

	app := &cli.App{
		Commands: []*cli.Command{
			{
				Name:  "ingredient",
				Usage: "manage ingredients",
				Subcommands: []*cli.Command{
					{
						Name:  "show",
						Usage: "show all the ingredients",
						Action: func(ctx *cli.Context) error {
							fmt.Println("Listing all ingredients...")
							ingredients, err := ingredientsRepository.GetAll()
							if err != nil {
								panic(err)
							}

							for _, i := range ingredients {
								fmt.Printf("%d - %s", i.ID, i.Name)
							}

							return nil
						},
					},
				},
			},
		},
	}

	if err := app.Run(os.Args); err != nil {
		log.Fatal(err)
	}
}
