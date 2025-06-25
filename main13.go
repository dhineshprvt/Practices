package main
import "fmt"
func main() {
	var myInt8 int8 = 86
	var myInt = 2300
	var myUint uint = 600
	var myHexNumber = 0xFC  
	var myOctalNumber = 344 
	var myFloat32 float32 = 6.3
	var myFloat = 10.33 
	fmt.Printf("%d, %d, %d, %#x, %#o %f %f\n", myInt8, myInt, myUint, myHexNumber, myOctalNumber, myFloat32, myFloat)
}
