package driver  // where "driver" is the folder that contains io.go, io.c, io.h, channels.go, channels.c and driver.go
/*
#cgo LDFLAGS: -lcomedi -lm
#include "io.h"
*/
import "C"

import (
	."fmt"
)


var N_BUTTONS := 3;
var N_FLOORS := 4;

//Lager 4x3 lysmatrise:
const lamp_channel_matrix := make ([][] int, N_FLOORS)
for i := range lamp_channel_matrix {
	a[i] = make([]int, N_BUTTONS)
}

//Implementerer lysmatrise for anntall etasjer:???? string med LIGHT_UP, LIGHT_DOWN og LIGHT_COMMAND med tall for antall etasjer?
for i := range N_FLOORS {
	for j := range N_BUTTONS {
		if j=0 {
			lamp_channel_matrix[i][j] = LIGHT_UP +j
		}
		else if j=1 {
			lamp_channel_matrix[i][j] = LIGHT_DOWN +j
		}
		else if j=2 {
			lamp_channel_matrix[i][j] = LIGHT_COMMAND +j
		}
	}
}

//Samme opplegg for button_channel_matrix
const button_channel_matrix := make ([][] int, N_FLOORS)
for i := range button_channel_matrix {
	a[i] = make([]int, N_BUTTONS)
}

//....implementer....




func elev_init() int {			//bruker chan istedet for return

	//Init hardware
	if (!io_init())	{return 0}
	
	// Zero all floor button lamps
	for i := range N_FLOORS {
		if i != 0 {
			elev_set_button_lamp(BUTTON_CALL_DOWN, i, 0)
		}
		if i != N_FLOORS - 1 {
			elev_set_button_lamp(BUTTON_CALL_UP, i, 0)
		}
		
		elev_set_button_lamp(BUTTON_COMMAND, i, 0)
	}
	
	// Clear stop lamp, door open lamp, and set floor indicator to ground floor.
	elev_set_stop_lamp(0);
	elev_set_door_open_lamp(0);
	elev_set_floor_indicator(0);


	// Return success
	return 1;
}



func elev_set_motor_direction(elev_motor_direvtion_t dirn) {
	if dirn == 0 {
		io_write_analog(MOTOR, 0)
	}
	else if dirn > 0 {
		io_clear_bit(MOTORDIR)
		io_write_analog(MOTOR, 2800)
	}
	else if dirn < 0 {
		io_set_bit(MOTORDIR)
		io_write_analog(MOTOR, 2800)
	}
}


func elev_set_door_open_lamp(value int) {
	if value {
		io_set_bit(LIGHT_DOOR_OPEN)}
	else {
		io_clear_bit(LIGHT_DOOR_OPEN)}
}


func elev_get_obstruction_signal() int {
	return io_read_bit(OBSTRUCTION)
}


func elev_get_stop_signal() int {
	return io_read_bit(STOP)
}


func elev_set_stop_lamp(value int) {
	if value {
		io_set_bit(LIGHT_STOP)}
	else {
		io_clear_bit(LIGHT_STOP)}
}


func elev_get_floor_sensor_signal() int {
	if (io_read_bit(SENSOR_FLOOR1)) {
		return 0 }
	else if (io_read_bit(SENSOR_FLOOR2)) {
		return 1 }
	else if (io_read_bit(SENSOR_FLOOR3)) {
		return 2 }
	else if (io_read_bit(SENSOR_FLOOR4)) {
		return 3 }
	else {
		return -1 }
}


func elev_set_floor_indicator(int floor) {
	assert(floor >= 0)
	assert(floor < N_FLOORS)

	// Binary encoding. One light must always be on.
	if (floor & 0x02){
		io_set_bit(LIGHT_FLOOR_IND1)}
	else {
		io_clear_bit(LIGHT_FLOOR_IND1)}
	if (floor & 0x01){
		io_set_bit(LIGHT_FLOOR_IND2)}
	else {
		io_clear_bit(LIGHT_FLOOR_IND2)}
}


func elev_get_button_signal(elev_button_type_t button, int floor, chan ch int) int {
	
	if floor >= 0 {
		Println ("Floor is lower than 0\n")}
	if floor < N_FLOORS {
		Println ("Floor variable is too high\n")}
	if !(button == BUTTON_CALL_UP && floor == N_FLOORS-1)) {
		Println("Unvalid button call or floor\n")}
	if !(button == BUTTON_CALL_DOWN && floor == 0)) {
		Println("Unvalid button call or floor\n")}
	if (button == BUTTON_CALL_UP || button == BUTTON_CALL_DOWN || button == BUTTON_COMMAND) {
		Println("Unvalid button\n")}


	if (io_read_bit(button_channel_matrix[floor][button])){
		return 1 }
	else {
		return 0 }
}



func elev_set_button_lamp(elev_button_type_t button, int floor, int value) {

	if floor >= 0 {
		Println("Floor is negative\n")}
	if floor < N_FLOORS {
		Println("Floor variable is too high\n")}
	if !(button == BUTTON_CALL_UP && floor == N_FLOORS-1)) {
		Println("Unvalid button call or floor")}
	if !(button == BUTTON_CALL_DOWN && floor == 0)) {
		Println("Unvalid button call or floor\n")}
	if (button == BUTTON_CALL_UP || button == BUTTON_CALL_DOWN || button == BUTTON_COMMAND) {
		Println("Unvalid button\n")}

	
	if (value) {
		io_set_bit(lamp_channel_matrix[floor][button])}
	else {
		io_clear_bit(lamp_channel_matrix[floor][button])}
}




