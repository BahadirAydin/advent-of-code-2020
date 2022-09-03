package handleInput

import(
	"fmt"
	"os"
	"strconv"
	"bufio"
)

func readIntegerLines() {
	f, err := os.Open("input.txt")
	if err != nil {
		fmt.Println("Error while reading.")
	}
	defer f.Close()
	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		converted, err := strconv.Atoi(scanner.Text())
		if err != nil {
			fmt.Println("Error while converting.")
		}
		fmt.Println(converted)
	}
}