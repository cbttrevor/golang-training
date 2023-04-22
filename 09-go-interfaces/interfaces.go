package main

import (
	"fmt"
	"reflect"
)

type Vehicle interface {
	Drive(uint16)
	Destroy()
	Repair(uint8) bool
}

type GolfCart struct {
	chargePercentage uint8
	manufacturer     string
	model            string
	color            string
	distanceDriven   uint16
	drivable         bool
	health           uint8
}

func (g *GolfCart) chargeGolfCart(chargeAmount uint8) {
	if g.chargePercentage+chargeAmount > 100 {
		g.chargePercentage = 100
	} else {
		g.chargePercentage += chargeAmount
	}
}

func (g *GolfCart) Drive(distanceToDrive uint16) {
	if g.drivable {
		g.distanceDriven += distanceToDrive
	} else {
		fmt.Println("This vehicle cannot be driven, because it's not drivable")
	}

}

func (g *GolfCart) Destroy() {
	g.drivable = false
	g.health = 0
}

func (g *GolfCart) Repair(repairPercentage uint8) bool {
	if g.health+repairPercentage > 100 {
		g.health = 100
	} else {
		g.health += repairPercentage
	}

	if g.health >= 60 {
		g.drivable = true
		return true
	} else {
		g.drivable = false
		return false
	}
}

func doSomeVehicleStuff(v Vehicle) {
	v.Destroy()
	v.Drive(5)
	fmt.Printf("%+v\n", v)

	v.Repair(70)
	v.Drive(5)
	fmt.Printf("%+v\n", v)

	v.Repair(50)
	fmt.Printf("%+v\n", v)

	gc := v.(*GolfCart)
	gc.chargeGolfCart(7)
	fmt.Printf("%+v\n", v)
}

func main() {
	fmt.Println("Learning about Go interfaces!")

	var myCart01 Vehicle = &GolfCart{chargePercentage: 87, manufacturer: "GolfCartMake1", color: "Green", drivable: true, health: 100, model: "GolfCartModel01"}
	fmt.Printf("%+v\n", myCart01)
	fmt.Printf("The type of myCart01 is: %v\n", reflect.TypeOf(myCart01))

	doSomeVehicleStuff(myCart01)
}
