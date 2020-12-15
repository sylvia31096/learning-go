package main
import ("bufio"
		"fmt"
		"log"
		"strings"
		"os")

func main(){
	var text_file string
	type Name struct{
		fname string
		lname string
	}
	s1 := make([]Name,0)
 
	fmt.Println("Enter name of text file\n")
	fmt.Scan(&text_file)

	file, err := os.Open(text_file)
 
	if err != nil {
		log.Fatalf("failed opening file: %s", err)

	}

	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanLines)
	
	var nm Name
	for scanner.Scan() {
		// split name by space
		names:= strings.Fields(scanner.Text())

		// assign each name to respective field
		nm.fname = names[0]
		nm.lname = names[1]

		// append to slice
		s1 = append(s1,nm)
	}

	for _,i:= range s1{

		fmt.Println(i)
	}
 
	file.Close()
	
}