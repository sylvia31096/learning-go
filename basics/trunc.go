package main
import "fmt"

func main(){
	fmt.Printf("Enter floating point number\n")

	var floatnum float32

	fmt.Scan(&floatnum)

	fmt.Printf("%.2f\n",floatnum)
}