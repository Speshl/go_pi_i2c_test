package icm20948

const I2C_ADD_ICM20948 byte = 0x68
const I2C_ADD_ICM20948_AK09916 byte = 0x0C
const I2C_ADD_ICM20948_AK09916_READ byte = 0x80
const I2C_ADD_ICM20948_AK09916_WRITE byte = 0x00

// define ICM-20948 Register
// user bank 0 register
const REG_ADD_WIA byte = 0x00
const REG_VAL_WIA byte = 0xEA
const REG_ADD_USER_CTRL byte = 0x03
const REG_VAL_BIT_DMP_EN byte = 0x80
const REG_VAL_BIT_FIFO_EN byte = 0x40
const REG_VAL_BIT_I2C_MST_EN byte = 0x20
const REG_VAL_BIT_I2C_IF_DIS byte = 0x10
const REG_VAL_BIT_DMP_RST byte = 0x08
const REG_VAL_BIT_DIAMOND_DMP_RST byte = 0x04
const REG_ADD_PWR_MIGMT_1 byte = 0x06
const REG_VAL_ALL_RGE_RESET byte = 0x80
const REG_VAL_RUN_MODE byte = 0x01 // Non low-power mode
const REG_ADD_LP_CONFIG byte = 0x05
const REG_ADD_PWR_MGMT_1 byte = 0x06
const REG_ADD_PWR_MGMT_2 byte = 0x07
const REG_ADD_ACCEL_XOUT_H byte = 0x2D
const REG_ADD_ACCEL_XOUT_L byte = 0x2E
const REG_ADD_ACCEL_YOUT_H byte = 0x2F
const REG_ADD_ACCEL_YOUT_L byte = 0x30
const REG_ADD_ACCEL_ZOUT_H byte = 0x31
const REG_ADD_ACCEL_ZOUT_L byte = 0x32
const REG_ADD_GYRO_XOUT_H byte = 0x33
const REG_ADD_GYRO_XOUT_L byte = 0x34
const REG_ADD_GYRO_YOUT_H byte = 0x35
const REG_ADD_GYRO_YOUT_L byte = 0x36
const REG_ADD_GYRO_ZOUT_H byte = 0x37
const REG_ADD_GYRO_ZOUT_L byte = 0x38
const REG_ADD_EXT_SENS_DATA_00 byte = 0x3B
const REG_ADD_REG_BANK_SEL byte = 0x7F
const REG_VAL_REG_BANK_0 byte = 0x00
const REG_VAL_REG_BANK_1 byte = 0x10
const REG_VAL_REG_BANK_2 byte = 0x20
const REG_VAL_REG_BANK_3 byte = 0x30

// user bank 1 register
// user bank 2 register
const REG_ADD_GYRO_SMPLRT_DIV byte = 0x00
const REG_ADD_GYRO_CONFIG_1 byte = 0x01
const REG_VAL_BIT_GYRO_DLPCFG_2 byte = 0x10   // bit[5:3]
const REG_VAL_BIT_GYRO_DLPCFG_4 byte = 0x20   // bit[5:3]
const REG_VAL_BIT_GYRO_DLPCFG_6 byte = 0x30   // bit[5:3]
const REG_VAL_BIT_GYRO_FS_250DPS byte = 0x00  // bit[2:1]
const REG_VAL_BIT_GYRO_FS_500DPS byte = 0x02  // bit[2:1]
const REG_VAL_BIT_GYRO_FS_1000DPS byte = 0x04 // bit[2:1]
const REG_VAL_BIT_GYRO_FS_2000DPS byte = 0x06 // bit[2:1]
const REG_VAL_BIT_GYRO_DLPF byte = 0x01       // bit[0]
const REG_ADD_ACCEL_SMPLRT_DIV_2 byte = 0x11
const REG_ADD_ACCEL_CONFIG byte = 0x14
const REG_VAL_BIT_ACCEL_DLPCFG_2 byte = 0x10 // bit[5:3]
const REG_VAL_BIT_ACCEL_DLPCFG_4 byte = 0x20 // bit[5:3]
const REG_VAL_BIT_ACCEL_DLPCFG_6 byte = 0x30 // bit[5:3]
const REG_VAL_BIT_ACCEL_FS_2g byte = 0x00    // bit[2:1]
const REG_VAL_BIT_ACCEL_FS_4g byte = 0x02    // bit[2:1]
const REG_VAL_BIT_ACCEL_FS_8g byte = 0x04    // bit[2:1]
const REG_VAL_BIT_ACCEL_FS_16g byte = 0x06   // bit[2:1]
const REG_VAL_BIT_ACCEL_DLPF byte = 0x01     // bit[0]

// user bank 3 register
const REG_ADD_I2C_SLV0_ADDR byte = 0x03
const REG_ADD_I2C_SLV0_REG byte = 0x04
const REG_ADD_I2C_SLV0_CTRL byte = 0x05
const REG_VAL_BIT_SLV0_EN byte = 0x80
const REG_VAL_BIT_MASK_LEN byte = 0x07
const REG_ADD_I2C_SLV0_DO byte = 0x06
const REG_ADD_I2C_SLV1_ADDR byte = 0x07
const REG_ADD_I2C_SLV1_REG byte = 0x08
const REG_ADD_I2C_SLV1_CTRL byte = 0x09
const REG_ADD_I2C_SLV1_DO byte = 0x0A

// define ICM-20948 Register  end

// define ICM-20948 MAG Register
const REG_ADD_MAG_WIA1 byte = 0x00
const REG_VAL_MAG_WIA1 byte = 0x48
const REG_ADD_MAG_WIA2 byte = 0x01
const REG_VAL_MAG_WIA2 byte = 0x09
const REG_ADD_MAG_ST2 byte = 0x10
const REG_ADD_MAG_DATA byte = 0x11
const REG_ADD_MAG_CNTL2 byte = 0x31
const REG_VAL_MAG_MODE_PD byte = 0x00
const REG_VAL_MAG_MODE_SM byte = 0x01
const REG_VAL_MAG_MODE_10HZ byte = 0x02
const REG_VAL_MAG_MODE_20HZ byte = 0x04
const REG_VAL_MAG_MODE_50HZ byte = 0x05
const REG_VAL_MAG_MODE_100HZ byte = 0x08
const REG_VAL_MAG_MODE_ST byte = 0x10

//define ICM-20948 MAG Register  end

const MAG_DATA_LEN byte = 6
