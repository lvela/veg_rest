package main

import (
	"fmt"
	"log"
	"os"
	"strconv"

	"veg_rest2/client"

	"github.com/codegangsta/cli"
)

func main() {

	app := cli.NewApp()
	app.Name = "vegetable cli"
	app.Usage = "cli to work with the `vegetable` microservice"
	app.Version = "0.0.1"

	app.Flags = []cli.Flag{
		cli.StringFlag{
			Name:  "host",
			Value: "http://localhost:8080",
			Usage: "Vegetable service host"},
	}

	app.Commands = []cli.Command{
		{
			Name:  "add",
			Usage: "(title amount_per_100) create a vegetable",
			Action: func(c *cli.Context) {
				title := c.Args().Get(0)
				desc := c.Args().Get(1)

				host := c.GlobalString("host")

				client := client.VegetableClient{Host: host}

				vegetable, err := client.CreateVegetable(title, desc)
				if err != nil {
					log.Fatal(err)
					return
				}
				fmt.Printf("%+v\n", vegetable)
			},
		},
		{
			Name:  "ls",
			Usage: "list all vegetable",
			Action: func(c *cli.Context) {

				host := c.GlobalString("host")

				client := client.VegetableClient{Host: host}

				vegetable, err := client.GetAllVegetables()
				if err != nil {
					log.Fatal(err)
					return
				}
				for _, vegetable := range vegetable {
					fmt.Printf("%+v\n", vegetable)
				}
			},
		},
		{
			Name:  "doing",
			Usage: "(id) update a vegetable status to 'doing'",
			Action: func(c *cli.Context) {
				idStr := c.Args().Get(0)
				id, err := strconv.Atoi(idStr)
				if err != nil {
					log.Print(err)
					return
				}

				host := c.GlobalString("host")

				client := client.VegetableClient{Host: host}

				vegetable, err := client.UpdateVegetableStatus(int32(id), "doing")
				if err != nil {
					log.Fatal(err)
					return
				}
				fmt.Printf("%+v\n", vegetable)
			},
		},
		{
			Name:  "done",
			Usage: "(id) update a vegetable status to 'done'",
			Action: func(c *cli.Context) {
				idStr := c.Args().Get(0)
				id, err := strconv.Atoi(idStr)
				if err != nil {
					log.Print(err)
					return
				}

				host := c.GlobalString("host")

				client := client.VegetableClient{Host: host}

				vegetable, err := client.UpdateVegetableStatus(int32(id), "done")
				if err != nil {
					log.Fatal(err)
					return
				}
				fmt.Printf("%+v\n", vegetable)
			},
		},
		{
			Name:  "save",
			Usage: "(id name amount_per_100) update a vegetable name and amount_per_100",
			Action: func(c *cli.Context) {
				idStr := c.Args().Get(0)
				id, err := strconv.Atoi(idStr)
				if err != nil {
					log.Print(err)
					return
				}
				name := c.Args().Get(1)
				amount_per_100 := c.Args().Get(2)

				host := c.GlobalString("host")

				client := client.VegetableClient{Host: host}

				vegetable, err := client.GetVegetable(int32(id))
				if err != nil {
					log.Fatal(err)
					return
				}

				vegetable.Name = name
				vegetable.Amount_per_100 = amount_per_100

				vegetable2, err := client.UpdateVegetable(vegetable)
				if err != nil {
					log.Fatal(err)
					return
				}

				fmt.Printf("%+v\n", vegetable2)
			},
		},
	}
	app.Run(os.Args)

}
