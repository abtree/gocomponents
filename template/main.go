package main

import (
	"fmt"
	"strings"
	"text/template"
)

type STest struct {
	pwd string
}

var vecs = make(map[string]interface{})

const lfu = `
func Get{{.FunName}}() (*{{.SName}}, bool) {
	f, ok := vecs["{{.Path}}"]
	if !ok {
		return nil, false
	}
	return f.(*{{.SName}}), true
}
`

func main() {
	b := &strings.Builder{}
	t := template.Must(template.New("lfn").Parse(lfu))
	params := make(map[string]string)
	path := "a/test"
	fnn := strings.ReplaceAll(path, "/", "_")
	params["FunName"] = fnn
	params["SName"] = "STest"
	params["Path"] = path
	t.Execute(b, params)
	fmt.Println(b.String())
}
