package main

var rooms = [MaxXSize][MaxYSize]bool{}

func InitRooms() {
	rooms[0][3] = true
	rooms[1][1] = true
	rooms[1][2] = true
	rooms[1][3] = true
	rooms[1][2] = true
	rooms[2][3] = true
	rooms[3][3] = true
	rooms[4][1] = true
	rooms[4][2] = true
	rooms[4][3] = true
	rooms[4][4] = true
	rooms[5][0] = true
	rooms[5][1] = true
	rooms[5][4] = true
	rooms[5][5] = true
	rooms[6][5] = true
	rooms[7][5] = true
}
