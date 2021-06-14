package main

import (
	"fmt"
	menu "github.com/NickLovera/go-apex/Menu"
	logic "github.com/NickLovera/go-apex/Mgr"
	"time"
)

var squadStats = logic.GetStats()
var lastUpdate = time.Now()

func main() {

	for {
		menu.printMenu()
		choice := getChoice()

		if time.Since(lastUpdate).Minutes() > 5 {
			fmt.Println("Retrieving Stats......")
			squadStats = logic.GetStats()
			lastUpdate = time.Now()
		}

		switch choice {
		case 1:
			getEveryone()
		case 2:
			fmt.Println("Who would you like to view?\n" +
				"1. Hk_Dingledorf\n2. Its_SkeetR\n3. MoneyManRex937\n4. SourMonkeyy")
			playerId := getChoice()
			getIndiv(playerId)
		}
	}
}

func getEveryone() {
	logic.GetEveryone(squadStats)
}

func getIndiv(playerId int) {
	logic.GetIndivdual(squadStats, playerId)
}
