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
lamp_channel_matrix := make ([][] int, N_FLOORS)
for i := 0; i<N_FLOORS; i++{
	lamp_channel_matrix [i] = make([]int, N_BUTTONS)
}

lamp_channel_matrix[0][0] = LIGHT_UP1
lamp_channel_matrix[1][0] = LIGHT_UP2
lamp_channel_matrix[2][0] = LIGHT_UP3
lamp_channel_matrix[3][0] = LIGHT_UP4
lamp_channel_matrix[0][1] = LIGHT_DOWN1
lamp_channel_matrix[1][1] = LIGHT_DOWN2
lamp_channel_matrix[2][1] = LIGHT_DOWN3
lamp_channel_matrix[3][1] = LIGHT_DOWN4
lamp_channel_matrix[0][2] = LIGHT_COMMAND1
lamp_channel_matrix[1][2] = LIGHT_COMMAND2
lamp_channel_matrix[2][2] = LIGHT_COMMAND3
lamp_channel_matrix[3][2] = LIGHT_COMMAND4


//Lager 4x3 buttonmatrise:
const button_channel_matrix := make ([][] int, N_FLOORS)
for i := 0; i<N_FLOORS; i++{
	button_channel_matrix [i] = make([]int, N_BUTTONS)
}

button_channel_matrix[0][0] = BUTTON_UP1
button_channel_matrix[1][0] = BUTTON_UP2
button_channel_matrix[2][0] = BUTTON_UP3
button_channel_matrix[3][0] = BUTTON_UP4
button_channel_matrix[0][1] = BUTTON_DOWN1
button_channel_matrix[1][1] = BUTTON_DOWN2
button_channel_matrix[2][1] = BUTTON_DOWN3
button_channel_matrix[3][1] = BUTTON_DOWN4
button_channel_matrix[0][2] = BUTTON_COMMAND1
button_channel_matrix[1][2] = BUTTON_COMMAND2
button_channel_matrix[2][2] = BUTTON_COMMAND3
button_channel_matrix[3][2] = BUTTON_COMMAND4

/*
button_channel_matrix = [[BUTTON_UP1, BUTTON_DOWN1, BUTTON_COMMAND1],
	[BUTTON_UP2, BUTTON_DOWN2, BUTTON_COMMAND2],
	[BUTTON_UP3, BUTTON_DOWN3, BUTTON_COMMAND3],
	[BUTTON_UP4, BUTTON_DOWN4, BUTTON_COMMAND4]]
*/


func elev_init() int {

	//Init hardware
	if (!io_init())	{return 0}
	
	// Zero all floor button lamps
	for i := 0; i<N_FLOORS; i++ {
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
	} else if dirn > 0 {
		io_clear_bit(MOTORDIR)
		io_write_analog(MOTOR, 2800)
	} else if dirn < 0 {
		io_set_bit(MOTORDIR)
		io_write_analog(MOTOR, 2800)
	}
}


func elev_set_door_open_lamp(value int) {
	if value {
		io_set_bit(LIGHT_DOOR_OPEN)
	} else {
		io_clear_bit(LIGHT_DOOR_OPEN)
	}
}


func elev_get_obstruction_signal() int {
	return io_read_bit(OBSTRUCTION)
}


func elev_get_stop_signal() int {
	return io_read_bit(STOP)
}


func elev_set_stop_lamp(value int) {
	if value {
		io_set_bit(LIGHT_STOP)
	} else {
		io_clear_bit(LIGHT_STOP)
	}
}


func elev_get_floor_sensor_signal() int {

	if (io_read_bit(SENSOR_FLOOR1)) {
		return 0 
	} else if (io_read_bit(SENSOR_FLOOR2)) {
		return 1 
	} else if (io_read_bit(SENSOR_FLOOR3)) {
		return 2 
	} else if (io_read_bit(SENSOR_FLOOR4)) {
		return 3 
	} else {
		return -1 
	}
}


func elev_set_floor_indicator(int floor) {
	
	if floor < 0 {
		Println ("Floor is lower than 0\n")
		return 0 }
	if floor > N_FLOORS {
		Println ("Floor variable is too high\n")
		return 0 }
	
	// Binary encoding. One light must always be on.
	if (floor & 0x02){
		io_set_bit(LIGHT_FLOOR_IND1)
	} else {
		io_clear_bit(LIGHT_FLOOR_IND1)
	}
		
	if (floor & 0x01){
		io_set_bit(LIGHT_FLOOR_IND2)
	} else {
		io_clear_bit(LIGHT_FLOOR_IND2)
	}
}


func elev_get_button_signal(elev_button_type_t button, int floor, chan ch int) int {
	
	if floor < 0 {
		Println ("Floor is lower than 0\n")
		return 0 }
	if floor > N_FLOORS {
		Println ("Floor variable is too high\n")
		return 0 }
	if (button == BUTTON_CALL_UP && floor == N_FLOORS-1)) {
		Println("Unvalid button call or floor\n")
		return 0 }
	if (button == BUTTON_CALL_DOWN && floor == 0)) {
		Println("Unvalid button call or floor\n")
		return 0 }
	if !(button == BUTTON_CALL_UP || button == BUTTON_CALL_DOWN || button == BUTTON_COMMAND) {
		Println("Unvalid button\n")
		return 0 }
	

	if (io_read_bit(button_channel_matrix[floor][button])){
		return 1
	} else {
		return 0 
	}
}



func elev_set_button_lamp(elev_button_type_t button, int floor, int value) {

	if floor < 0 {
		Println("Floor is negative\n")
		return }
	if floor > N_FLOORS {
		Println("Floor variable is too high\n")
		return }
	if (button == BUTTON_CALL_UP && floor == N_FLOORS-1)) {
		Println("Unvalid button call or floor")
		return }
	if (button == BUTTON_CALL_DOWN && floor == 0)) {
		Println("Unvalid button call or floor\n")
		return }
	if !(button == BUTTON_CALL_UP || button == BUTTON_CALL_DOWN || button == BUTTON_COMMAND) {
		Println("Unvalid button\n")
		return }

	
	if (value) {
		io_set_bit(lamp_channel_matrix[floor][button])
	} else {
		io_clear_bit(lamp_channel_matrix[floor][button])
	}
}

