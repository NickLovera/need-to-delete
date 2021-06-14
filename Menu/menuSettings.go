package Menu

import (
	"bufio"
	"fmt"
	logic "github.com/NickLovera/go-apex/Mgr"
	"log"
	"os"
	"strconv"
	"time"
)

func printMenu() {
	var MENU = "---- Welcome to YMAH stat tracker ----"
	var OPTIONS = "What would you like to do\n" + "1. Get Everyone's Stats\n" + "2. Get indivudual squad meber stat"

	min, sec := logic.GetTimeTillUpdate(lastUpdate.Add(time.Minute * 5))

	fmt.Println(MENU)
	fmt.Println("Time till next update: ", min, " Min ", sec, " Sec\n") //Implement timer
	fmt.Println(OPTIONS)
}

func getChoice() int {
	scanner := bufio.NewScanner(os.Stdin)
	scanned := scanner.Scan()
	if !scanned {
		log.Fatalln("Unable to scan")
	}
	choice, err := strconv.Atoi(scanner.Text())
	if err != nil {
		log.Fatalln(err)
	}

	return choice
}
