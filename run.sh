#!/bin/sh

go build .

#export $(grep -v '^#' alpha_car.env | xargs)

sudo ./go_pi_i2c_test