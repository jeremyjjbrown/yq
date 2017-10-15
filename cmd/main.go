package main

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"os"

	"github.com/jeremyjjbrown/yq/out"
	"gopkg.in/yaml.v2"
)

func main() {

	stdinFileInfo, err := os.Stdin.Stat()

	if err != nil {
		panic("Error reading stdin")
	}

	var pipeIn []byte
	if stdinFileInfo.Mode()&os.ModeNamedPipe != 0 {
		stdinContent, _ := ioutil.ReadAll(os.Stdin)
		pipeIn = stdinContent
	}

	var t map[interface{}]interface{}
	err = yaml.Unmarshal(pipeIn, &t)
	if err != nil {
		fmt.Println(err)
	}

	buf := new(bytes.Buffer)
	out.ColorMap(buf, t, 0)
	fmt.Println(buf.String())
}
