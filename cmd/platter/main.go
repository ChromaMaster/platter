package main

import (
	"database/sql"
	"fmt"
	_ "github.com/mattn/go-sqlite3"
	"github.com/urfave/cli/v2"
	"log"
	"os"
	"platter/internal/model"
	"platter/internal/repository/dish"
	"platter/internal/repository/ingredient"
)

func main() {

	db, err := openDB()
	if err != nil {
		log.Fatal(err)
	}

	defer func(db *sql.DB) { _ = closeDB(db) }(db)

	ingredientsRepository := ingredient.NewInDBRepository(db)
	dishesRepository := dish.NewInDBRepository(db)

	if err := ingredientsRepository.Init(); err != nil {
		panic(err)
	}
	if err := dishesRepository.Init(); err != nil {
		panic(err)
	}

	app := &cli.App{
		Commands: []*cli.Command{
			{
				Name:  "ingredient",
				Usage: "manage ingredients",
				Subcommands: []*cli.Command{
					{
						Name:  "list",
						Usage: "list all the ingredients",
						Action: func(ctx *cli.Context) error {
							fmt.Println("Listing all ingredients...")
							ingredients, err := ingredientsRepository.GetAll()
							if err != nil {
								panic(err)
							}

							for _, i := range ingredients {
								fmt.Printf("%d - %s\n", i.ID, i.Name)
							}

							return nil
						},
					},
					{
						Name:      "create",
						Usage:     "add an ingredient",
						ArgsUsage: "ingredientName",
						Action: func(ctx *cli.Context) error {
							if ctx.NArg() == 0 {
								return fmt.Errorf("missing ingredient name")
							}

							fmt.Println("Adding the ingredient...")
							name := ctx.Args().First()

							if err := ingredientsRepository.Create(model.NewIngredient(0, name)); err != nil {
								return fmt.Errorf("cannot add the ingredient: %w", err)
							}

							return nil
						},
					},
				},
			},
			{
				Name:  "dish",
				Usage: "manage dishes",
				Subcommands: []*cli.Command{
					{
						Name:  "list",
						Usage: "list all the dishes",
						Action: func(ctx *cli.Context) error {
							fmt.Println("Listing all dishes...")
							dishes, err := dishesRepository.GetAll()
							if err != nil {
								panic(err)
							}

							for _, d := range dishes {
								fmt.Printf("%d - %s\n", d.ID, d.Name)
							}

							return nil
						},
					},
					{
						Name:      "create",
						Usage:     "add a dish",
						ArgsUsage: "dishName",
						Action: func(ctx *cli.Context) error {
							if ctx.NArg() == 0 {
								return fmt.Errorf("missing dish name")
							}

							fmt.Println("Adding the dish...")
							name := ctx.Args().First()

							if err := dishesRepository.Create(model.NewDish(0, name)); err != nil {
								return fmt.Errorf("cannot add the dish: %w", err)
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

func openDB() (*sql.DB, error) {
	databaseName := "platter.db"
	db, err := sql.Open("sqlite3", databaseName)
	if err != nil {
		return nil, fmt.Errorf("cannot open the database: %w", err)
	}

	return db, nil
}

func closeDB(db *sql.DB) error {
	if err := db.Close(); err != nil {
		return fmt.Errorf("cannot close the database: %w", err)
	}

	return nil
}
