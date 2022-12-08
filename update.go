package tgbot

import (
	"encoding/json"
	"fmt"
	"strconv"
)

func (i *New) GetUpdates(offset int) []Update {
	res := i.get(GetArgs{Action: "getUpdates", Params: []Param{{Name: "offset", Value: strconv.Itoa(offset)}}})
	updateResult := UpdateResult{}
	err := json.Unmarshal([]byte(res), &updateResult)

	if err != nil {
		fmt.Println(err)
	}
	if !updateResult.Ok {
		fmt.Println("An error occured with update result parsing")
	}

	return updateResult.Updates
}
