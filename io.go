package driver  // where "driver" is the folder that contains io.go, io.c, io.h, channels.go, channels.c and driver.go
/*
#cgo LDFLAGS: -lcomedi -lm
#include "io.h"
*/
import "C"



func int io_init() {
	return int (C.io_init())
}



func io_set_bit(channel int) {
	C.io_set_bit(int channel)
}



func io_clear_bit(channel int) {
	C.io_clear_bit(int channel)
}



func io_write_analog(channel int, value int) {
	C.io_write_analog(int channel, int value)
}



func int io_read_bit(channel int) {
	return int (C.io_read_bit(int channel))
}



func int io_read_analog(channel int) {
	return int (C.io_read_analog(int cahnnel))
}
