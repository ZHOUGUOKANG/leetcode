package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"os"
	"strings"
)

var name = flag.String("name", "", "code name")
var id = flag.Int("id", 0, "code id")

func main() {
	flag.Parse()
	if *name == "" || *id == 0 {
		println("no name AND id")
		return
	}

	dirName := fmt.Sprintf("Num-%d-%s", *id, *name)
	err := os.Mkdir(dirName, 0777)
	if err != nil {
		println(dirName, "make dir failed", err.Error())
	}

	err = ioutil.WriteFile(fmt.Sprintf("%s/%s.go", dirName, *name), []byte(fmt.Sprintf("package Num_%d_%s \n", *id, strings.Title(*name))), 0777)
	if err != nil {
		println("create file err", err.Error())
	}

	template := `package Num_%d_%s
	
import "testing"

func TestNum_%d_%s(t *testing.T)  {

`
	template2 := `
	e := []struct{
		expected interface{}
	}{
		{},
	}
	for _,v := range e{
		o := 0
		if o != v.expected{
			t.Error(fmt.Sprintf("%v %v",v,o))
		}
	}
}`
	err = ioutil.WriteFile(fmt.Sprintf("%s/%s_test.go", dirName, *name), []byte(fmt.Sprintf(template, *id, strings.Title(*name), *id, strings.Title(*name))+template2), 0777)
	if err != nil {
		println("create file err", err.Error())
	}

	err = ioutil.WriteFile(fmt.Sprintf("%s/readme.md", dirName), nil, 0777)
	if err != nil {
		println("create file err", err.Error())
	}
}
