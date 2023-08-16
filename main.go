package main

import (
	"log"
	"time"

	"github.com/googolgl/go-i2c"
	"github.com/googolgl/go-pca9685"
)

const ICM20948_Address uint16 = 0x68

func main() {
	// Create new connection to i2c-bus on 1 line with address 0x40.
	// Use i2cdetect utility to find device address over the i2c-bus
	i2c, err := i2c.New(pca9685.Address, "/dev/12c-0")
	if err != nil {
		log.Fatal(err)
	}

	pca0, err := pca9685.New(i2c, nil)
	if err != nil {
		log.Fatal(err)
	}

	// Sets a single PWM channel 0
	pca0.SetChannel(0, 0, 130)

	// Servo on channel 0
	servo0 := pca0.ServoNew(0, nil)

	// Angle in degrees. Must be in the range `0` to `Range`
	for i := 0; i < 130; i++ {
		servo0.Angle(i)
		time.Sleep(10 * time.Millisecond)
	}

	// Fraction as pulse width expressed between 0.0 `MinPulse` and 1.0 `MaxPulse`
	servo0.Fraction(0.5)
	// // Load all the drivers:
	// state, err := host.Init()
	// if err != nil {
	// 	log.Fatalf("failed initializing host: %s", err)
	// }

	// log.Printf("State Details: %+v\n\n", state)

	// // Open a handle to the first available I²C bus:
	// bus, err := i2creg.Open("/dev/i2c-1") //Might need to specify a 1
	// if err != nil {
	// 	log.Fatalf("failed opening bus: %s\n", err)
	// }
	// defer bus.Close()

	// log.Printf("Bus Details: %s\n\n", bus.String())

	// // Open a handle to device on the I²C bus
	// dev, err := icm20948.NewI2C(bus, ICM20948_Address, nil)
	// if err != nil {
	// 	log.Fatalf("failed creating icm20948 device: %s\n", err)
	// }
	// defer dev.Halt()

	// var env physic.Env
	// if err = dev.ReadGyro(&env); err != nil {
	// 	log.Fatal(err)
	// }
	// fmt.Printf("%8s %10s %9s\n", env.Temperature, env.Pressure, env.Humidity)
}
