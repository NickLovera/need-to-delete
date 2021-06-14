package Mgr

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
)

//Result contains entire json
type Result struct{
	Legends []Legend `json:"data"`
}
//Legend Individual Fields in the json
type Legend struct{
	//Type string `json:"type"`
	//Attributes Attribute `json:"attributes"`
	Metas MetaData `json:"metadata"`
	//Dates string `json:"expiryDate"`
	Stats Stat `json:"stats"`
}
//Attribute contains legend id (Will use for searching legend in future)
type Attribute struct{
	LegendId string `json:"id"`
}
//MetaData name is the legends name ie. bloodhound
type MetaData struct {
	Name string `json:"name"`
}

//Stat Each field is a stat ie. kills, winningKills, .....
type Stat struct{
	KillNum Kills `json:"kills"`
	Damages Damage `json:"damage"`
	Headshots Headshot `json:"headshot"`
}

//Kills
type Kills struct{
	Rank float32 `json:"rank"`
	Value string `json:"displayValue"`
}

type Damage struct{
	Rank float32 `json:"rank"`
	Value string `json:"displayValue"`
}

type Headshot struct{
	Rank float32 `json:"rank"`
	Value string `json:"displayValue"`
}
var squad = [4]string{"HK_Dingledorf", "Its_SkeetR", "MoneyManRex937", "SourMonkeyy"}


func getStats(){
	var squadStats [4]Result

	for i:=0; i<4; i++{
		req, err := http.NewRequest("GET", "https://public-api.tracker.gg/v2/apex/standard/profile/psn/"+ squad[i] +"/segments/legend", nil)
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}

		req.Header.Set("TRN-Api-Key", "cf97cbea-dfd7-46f0-aa43-5acc8da4e47c")
		client := &http.Client{}

		resp, err := client.Do(req)
		if err != nil {
			fmt.Println("Error getting response: ",err)
			os.Exit(1)
		}

		body, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			fmt.Println("Error reading body: ",err)
		}

		var stats Result
		err = json.Unmarshal(body, &stats)
		if err != nil {
			fmt.Println("Error Unmarshalling: ",err)
		}
		//Store each bois stats
		squadStats[i] = stats
	}
}

//Get all stats of everyone
func getEveryone (squadStats [4]Result){

	for i, item := range squadStats{
		showStats(i, item)
	}
}

func showStats (i int, item Result){
	file, err := os.Create("C:\\Users\\nickl\\IdeaProjects\\go-apex\\Data\\"+squad[i])
	if err != nil {
		fmt.Println("Error opening write file: ", err)
	}
	//fmt.Println(squad[i]," stat's")
	file.WriteString(squad[i]+" stat's"+"\n\n")
	for _,each :=range item.Legends{
		//fmt.Println("Legend: ", each.Metas.Name)
		file.WriteString("Legend: "+each.Metas.Name+"\n")

		//fmt.Println("Kills: ",each.Stats.KillNum.Value, " Rank: ",each.Stats.KillNum.Rank)
		file.WriteString("Kills: "+each.Stats.KillNum.Value+" Rank: "+ fmt.Sprintf("%g",each.Stats.KillNum.Rank)+"\n")

		//fmt.Println("Damage: ", each.Stats.Damages.Value, " Rank: ", each.Stats.Damages.Rank)
		file.WriteString("Damage: " + each.Stats.Damages.Value + " Rank: "+ fmt.Sprintf("%g",each.Stats.Damages.Rank)+"\n" )

		//fmt.Println("Headshots: ", each.Stats.Headshots.Value, " Rank: "+ fmt.Sprintf("%g", each.Stats.Headshots.Rank))
		file.WriteString("Headshots: " + each.Stats.Headshots.Value+ " Rank: "+ fmt.Sprintf("%g", each.Stats.Headshots.Rank)+"\n\n")

		//fmt.Println()
	}
	//fmt.Println()
	file.WriteString("------------------------------------------------------\n")

	err = file.Close()
	if err != nil {
		fmt.Println(err)
	}
}

//Get one person's stats
func getIndivdual (squadStats [4]Result, nameId int){
	showStats(nameId, squadStats[nameId])
}
