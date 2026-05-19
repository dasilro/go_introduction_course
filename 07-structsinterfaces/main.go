package main

import (
	"encoding/json"
	"fmt"

	"github.com/dasilro/go_introduction_course/07-structsinterfaces/structsinterface/structs"
	"github.com/dasilro/go_introduction_course/07-structsinterfaces/structsinterface/vehicles"
)

func main() {
	var p1 structs.Product

	p1.ID = 12
	p1.Name = "test"
	fmt.Println(p1)

	p2 := structs.Product{}
	p2.ID = 12
	p2.Name = "test"
	fmt.Println(p2)

	p3 := structs.Product{2, "test 3", structs.Type{}, 0, 12.21, nil}
	fmt.Println(p3)

	p4 := structs.Product{
		ID:   2,
		Name: "test 4",
	}
	fmt.Println(p4)

	p5 := structs.Product{
		Name: "Frigorifico marca bosh",
		Type: structs.Type{
			Code:        "A",
			Description: "Electrodomestico",
		},
		Tags: []string{"frigorifico", "bosch", "heladera"},
	}

	fmt.Println(p5)

	p6 := structs.Product{
		Name: "Frigorifico marca bosh",
		Type: structs.Type{
			Code:        "A",
			Description: "Electrodomestico",
		},
		Price: 40000,
		Tags:  []string{"frigorifico", "bosch", "heladera"},
		Count: 5,
	}

	v, error := json.Marshal(p6)
	fmt.Println(error)
	fmt.Println(string(v))

	fmt.Println("Precio total: ", p6.TotalPrice())
	fmt.Println(p6)

	p6.SetName("other name")
	p6.AddTags("tag1", "tag2", "tag3")
	fmt.Println(p6)

	p7 := structs.Product{
		Name: "Horno",
		Type: structs.Type{
			Code:        "A",
			Description: "Electrodomestico",
		},
		Price: 40000,
		Tags:  []string{"horno", "bosch", "heladera"},
		Count: 5,
	}

	p8 := structs.Product{
		Name: "Secador marca bosh",
		Type: structs.Type{
			Code:        "A",
			Description: "Electrodomestico",
		},
		Price: 40000,
		Tags:  []string{"secador", "bosch", "heladera"},
		Count: 5,
	}
	c1 := structs.NewCart(1)
	c1.AddProducts(p6, p7, p8)

	fmt.Println("Products Cart")
	fmt.Println("Total Products: ", len(c1.Products))
	fmt.Printf("Total Cart: $%.2f\n", c1.Total())

	fmt.Println()
	fmt.Println("VEHICLES")

	carV := vehicles.Car{Time: 120}
	fmt.Println(carV.Distance())

	vArray := []string{"CAR", "MOTORCYCLE", "TRUCK", "MOTORCYCLE", "TRUCK", "GOKU", "DDD"}
	var d float64

	for _, v := range vArray {
		fmt.Printf("Vehicle %s\n", v)
		veh, error := vehicles.New(v, 400)
		if error != nil {
			fmt.Println("Error: ", error)
			fmt.Println()
			continue
		}
		distance := veh.Distance()
		fmt.Printf("Distance %2.f\n", distance)
		d += distance
	}
	fmt.Println("Total distance: ", d)
}
