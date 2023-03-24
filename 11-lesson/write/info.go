package write

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
)

type Currency struct {
	ID       int    `json:"id"`
	Code     string `json:"Code"`
	Ccy      string `json:"Ccy"`
	CcyNmRU  string `json:"CcyNm_RU"`
	CcyNmUZ  string `json:"CcyNm_UZ"`
	CcyNmUZC string `json:"CcyNm_UZC"`
	CcyNmEN  string `json:"CcyNm_EN"`
	Nominal  string `json:"Nominal"`
	Rate     string `json:"Rate"`
	Diff     string `json:"Diff"`
	Date     string `json:"Date"`
}

//var id, code int

func ReadJson() {
	var str string

	data, err := ioutil.ReadFile("write/info.json")
	if err != nil {
		fmt.Println("Error reading file", err)
		return
	}

	var menu []Currency
	err = json.Unmarshal(data, &menu)
	if err != nil {
		fmt.Println("Error unmarshaling JSON:", err)
		return
	}

	fmt.Println("Enter")
	fmt.Scanln(&str)

	for _, v := range menu {
		if v.Ccy == str {
			fmt.Println(v)
			break
		}
	}
}

func SendToMain() {
	ReadJson()
}
