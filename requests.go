package tgbot

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"strings"
)

type Param struct {
	Name  string
	Value string
}

type GetArgs struct {
	Action string
	Params []Param
}

type PostArgs struct {
	Action string
	Body   interface{}
}

func (i *New) buildQuery(action string, params []Param) string {
	var qParams []string

	if params == nil {
		params = []Param{}
	}

	for _, elem := range params {
		qParams = append(qParams, elem.Name+"="+elem.Value)
	}

	return fmt.Sprintf("https://api.telegram.org/bot%s/%s?%s", i.Token, action, strings.Join(qParams, "&"))
}

func (i *New) get(args GetArgs) string {
	q := i.buildQuery(args.Action, args.Params)

	res, err := http.Get(q)

	if err != nil {
		fmt.Println(err)
	}

	defer res.Body.Close()

	body, err := io.ReadAll(res.Body)

	if err != nil {
		fmt.Println(err)
	}

	return string(body)
}

func (i *New) post(args *PostArgs) string {
	q := i.buildQuery(args.Action, []Param{})
	body, err := json.Marshal(args.Body)
	if err != nil {
		fmt.Println(err)
	}

	res, err := http.Post(q, "application/json", bytes.NewReader(body))
	if err != nil {
		fmt.Println(err)
	}

	body, err = io.ReadAll(res.Body)
	if err != nil {
		fmt.Println(err)
	}

	defer res.Body.Close()

	return string(body)
}

func (i *New) getFilePath(fileId string) string {
	res := i.get(GetArgs{Action: "getFile", Params: []Param{{Name: "file_id", Value: fileId}}})
	var file struct {
		Ok   bool
		File File `json:"result"`
	}
	err := json.Unmarshal([]byte(res), &file)

	if err != nil {
		fmt.Println(err)
	}

	return file.File.Path
}

func (i *New) GetFile(fileId string) io.ReadCloser {
	q := fmt.Sprintf("https://api.telegram.org/file/bot%s/%s", i.Token, i.getFilePath(fileId))

	res, err := http.Get(q)
	if err != nil {
		fmt.Println(err)
	}

	return res.Body
}
