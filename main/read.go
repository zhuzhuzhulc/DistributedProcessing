package main
import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {

	reader := bufio.NewReader(os.Stdin)
	fmt.Println("Simple Shell")
	fmt.Println("---------------------")

	for {
		fmt.Print("To which node and what message: ")
		fmt.Print("-> ")
		text, _ := reader.ReadString('\n')
		// convert CRLF to LF
		var a[2]string
		a[0] = strings.Split(text," ")[0]
		a[1] = strings.Split(text," ")[1]
		fmt.Print("To " + a[0] +" msg: "+a[1])
	}

}


func check(e error) {
	if e != nil {
		panic(e)
	}
}