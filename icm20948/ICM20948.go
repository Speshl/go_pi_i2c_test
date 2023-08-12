package icm20948

import (
	"fmt"
	"log"

	"periph.io/x/conn/v3"
	"periph.io/x/conn/v3/i2c"
)

type Dev struct {
	device  conn.Conn
	options *Opts

	name string
	os   uint8
}

type Opts struct {
}

func NewI2C(b i2c.Bus, addr uint16, opts *Opts) (*Dev, error) {
	d := &Dev{device: &i2c.Dev{Bus: b, Addr: addr}, options: opts}
	if err := d.makeDev(opts); err != nil {
		return nil, err
	}
	return d, nil
}

func (d *Dev) makeDev(opts *Opts) error {
	ok, err := d.Check()
	if err != nil {
		return err
	}

	log.Printf("Check was %t\n", ok)

	return nil
}

func (d *Dev) Halt() error {
	return nil
}

func (d *Dev) ReadGyro() error {
	return nil
}

func (d *Dev) Check() (bool, error) {
	buffer := make([]byte, 0)
	err := d.readRegister(REG_ADD_WIA, buffer)
	if err != nil {
		return false, err
	}
	if REG_VAL_WIA == buffer[0] {
		return true, nil
	}
	return false, nil
}

func (d *Dev) readRegister(reg uint8, b []byte) error {
	err := d.device.Tx([]byte{reg}, b)
	if err != nil {
		return fmt.Errorf("failed reading register - %w", err)
	}
	return nil
}

func (d *Dev) writeCommands(b []byte) error {
	err := d.device.Tx(b, nil)
	if err != nil {
		return fmt.Errorf("failed writing command - %w", err)
	}
	return nil
}
