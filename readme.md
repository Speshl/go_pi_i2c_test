https://www.waveshare.com/wiki/Sense_HAT_(B)

Pi sense Hat B

ADS1015: AD conversion demo (STM32, BCM2835, WringPi and Python four demos) Device address: 0x48 

ICM-20948: Device address: 0x68  - 9-axis sensor demo (STM32, BCM2835, WringPi and Python four demos)

LPS22HB: Device address: 0x5C

SHTC3: Device address: 0x70

TCS34725: Device address: 0x29


https://periph.io/platform/raspberrypi/


The BCM238x has two I²C buses but it is recommended to only use the second. Enabling /dev/i2c-1 permanently:

sudo raspi-config nonint do_i2c 0