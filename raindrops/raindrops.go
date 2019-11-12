package raindrops

import(
	"strconv"
)

func Convert(num int) string {
	/*
		Convert the numbers into the sound of raindrops (3 = Pling, 5 = Plang, 7 = Plong). 
		Othersise, return the int.
	*/

	// list := [3]int{}
	// list[0] = 3
	// list[1] = 5
	// list[2] = 7

	outputString := "" 

	divis := [3]int{3,5,7}
	rainSounds := [3]string{"Pling", "Plang", "Plong"}	

	// CHECK IF INPUT NUM MATCHES DIVIS
	for i := 0; i < len(divis); i++ {
		if num%divis[i] == 0 {
			outputString += rainSounds[i]
		}
	}

	// CHECK OUTPUT STRING LENGTH
	if len(outputString) == 0 {
		return strconv.Itoa(num) 
	}

	return outputString
}