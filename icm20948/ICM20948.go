package icm20948

import (
	"fmt"
	"log"
	"time"

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
	log.Printf("Device details: %s\n", d.device.String())

	ok, err := d.Check()
	if err != nil {
		return err
	}
	if !ok {
		return fmt.Errorf("device return false for check")
	}

	time.Sleep(500 * time.Millisecond)
	//user bank 0 register
	err = d.writeCommands([]byte{REG_ADD_REG_BANK_SEL, REG_VAL_REG_BANK_0})
	if err != nil {
		return err
	}

	err = d.writeCommands([]byte{REG_ADD_PWR_MIGMT_1, REG_VAL_ALL_RGE_RESET})
	if err != nil {
		return err
	}

	time.Sleep(100 * time.Millisecond)

	err = d.writeCommands([]byte{REG_ADD_PWR_MIGMT_1, REG_VAL_RUN_MODE})
	if err != nil {
		return err
	}
	//user bank 2 register
	err = d.writeCommands([]byte{REG_ADD_REG_BANK_SEL, REG_VAL_REG_BANK_2})
	if err != nil {
		return err
	}

	err = d.writeCommands([]byte{REG_ADD_GYRO_SMPLRT_DIV, 0x07})
	if err != nil {
		return err
	}

	err = d.writeCommands([]byte{REG_ADD_GYRO_CONFIG_1, REG_VAL_BIT_GYRO_DLPCFG_6 | REG_VAL_BIT_GYRO_FS_1000DPS | REG_VAL_BIT_GYRO_DLPF})
	if err != nil {
		return err
	}

	err = d.writeCommands([]byte{REG_ADD_ACCEL_SMPLRT_DIV_2, 0x07})
	if err != nil {
		return err
	}

	err = d.writeCommands([]byte{REG_ADD_ACCEL_CONFIG, REG_VAL_BIT_ACCEL_DLPCFG_6 | REG_VAL_BIT_ACCEL_FS_2g | REG_VAL_BIT_ACCEL_DLPF})
	if err != nil {
		return err
	}

	//user bank 0 register
	err = d.writeCommands([]byte{REG_ADD_REG_BANK_SEL, REG_VAL_REG_BANK_0})
	if err != nil {
		return err
	}

	time.Sleep(100 * time.Millisecond)

	return d.gyroOffset()
}

func (d *Dev) gyroOffset() error {
	return d.gyroRead()
}

func (d *Dev) gyroRead() error {
	err := d.writeCommands([]byte{REG_ADD_REG_BANK_SEL, REG_VAL_REG_BANK_0})
	if err != nil {
		return err
	}

	//read 12 bytes
	buffer := make([]byte, 12)
	err = d.readRegister(REG_ADD_ACCEL_XOUT_H, buffer)
	if err != nil {
		return err
	}

	err = d.writeCommands([]byte{REG_ADD_REG_BANK_SEL, REG_VAL_REG_BANK_2})
	if err != nil {
		return err
	}

	log.Printf("Read Gyro Bytes - %+v\n", buffer)

	return nil
}

func (d *Dev) Halt() error {
	return nil
}

func (d *Dev) ReadGyro() error {
	return nil
}

func (d *Dev) Check() (bool, error) {
	buffer := make([]byte, 1)
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
