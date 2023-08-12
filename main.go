package main

import (
	"log"

	"github.com/Speshl/go_pi_i2c_test/icm20948"
	"periph.io/x/conn/v3/i2c/i2creg"
	"periph.io/x/host/v3"
)

const ICM20948_Address = 0x68

func main() {
	// Load all the drivers:
	state, err := host.Init()
	if err != nil {
		log.Fatalf("failed initializing host: %s", err)
	}

	log.Printf("State Details: %+v\n\n", state)

	// Open a handle to the first available I²C bus:
	bus, err := i2creg.Open("") //Might need to specify a 1
	if err != nil {
		log.Fatalf("failed opening bus: %s", err)
	}
	defer bus.Close()

	log.Printf("Bus Details: %s\n\n", bus.String())

	// Open a handle to device on the I²C bus
	dev, err := icm20948.NewI2C(bus, ICM20948_Address, nil)
	if err != nil {
		log.Fatalf("failed creating icm20948 device: %s", err)
	}
	defer dev.Halt()

	// var env physic.Env
	// if err = dev.ReadGyro(&env); err != nil {
	// 	log.Fatal(err)
	// }
	// fmt.Printf("%8s %10s %9s\n", env.Temperature, env.Pressure, env.Humidity)
}
