package main

type Room struct {
	name    string
	cleaned bool
}

func checkCleanliness(rooms [4]Room) {
	for i := 0; i < len(rooms); i++ {
		room := rooms[i]
		if room.cleaned {
			println(room.name, " is clean")
		} else {
			println(room.name, " is not clean")
		}
	}
}

func main() {
	room := [...]Room{
		{
			name:    "Room 1",
			cleaned: true,
		},
		{
			name:    "Room 2",
			cleaned: false,
		},
		{
			name:    "Room 3",
			cleaned: true,
		},
		{
			name:    "Room 4",
			cleaned: false,
		},
	}
	checkCleanliness(room)
}
