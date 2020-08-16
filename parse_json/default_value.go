package main

import (
	"encoding/json"

	"golang.org/x/exp/errors/fmt"
)

type EditHistory struct {
	ID    int    `json:"id"`
	Name  string `json:"name"`
	Price int    `json:"price"`
}

// 例えばPriceなどがdefaultで0なのか0に設定されたのかを見分ける時に便利
type EditHistoryWithPointer struct {
	ID    int     `json:"id"`
	Name  *string `json:"name"`
	Price *int    `json:"price"`
}

func (ep EditHistoryWithPointer) String() string {
	name := "<nil>"
	if ep.Name != nil {
		name = *ep.Name
	}

	price := "<nil>"
	if ep.Price != nil {
		price = fmt.Sprintf("%d", *ep.Price)
	}

	return fmt.Sprintf("{%v %v %v}", ep.ID, name, price)
}

var jsonString = []byte(`
[
	{"id": 1, "name":"taro", "price": 1000},
	{"id": 1, "price": 1000},
	{"id": 1, "name":"taro"},
	{"id": 1, "name":"taro", "price": 0}
]`)

func main() {
	var e []EditHistory
	var ep []EditHistoryWithPointer

	if err := json.Unmarshal(jsonString, &e); err != nil {
		panic(err)
	}

	if err := json.Unmarshal(jsonString, &ep); err != nil {
		panic(err)
	}

	for i := 0; i < len(e); i++ {
		fmt.Println(i)
		fmt.Printf("EditHistory: %v\n", e[i])
		fmt.Printf("EditHistoryWithPointer: %v\n", ep[i])
	}
}
