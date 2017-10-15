package out

import (
	"bytes"
	"fmt"
	"reflect"
	"sort"
	"strings"

	"github.com/fatih/color"
)

func ColorMap(buf *bytes.Buffer, in map[interface{}]interface{}, indent int) {

	keys := make([]string, 0, len(in))
	for k := range in {
		keys = append(keys, k.(string))
	}
	sort.Strings(keys)

	for _, k := range keys {
		color.New(color.FgBlue).Fprint(buf, fmt.Sprintf("%v%v", strings.Repeat(" ", indent), k))
		color.New(color.FgWhite).Fprint(buf, ": ")
		ColorValue(buf, in[k], indent)
	}

}

func ColorSlice(buf *bytes.Buffer, in []interface{}, indent int) {

	for _, i := range in {
		color.New(color.FgWhite).Fprint(buf, fmt.Sprintf("%v- ", strings.Repeat(" ", indent)))
		ColorValue(buf, i, indent)
	}

}

func ColorValue(buf *bytes.Buffer, in interface{}, indent int) {
	if in == nil {
		color.New(color.FgWhite).Fprint(buf, "null\n")
		return
	}
	v := reflect.ValueOf(in)
	switch v.Kind() {
	case reflect.Slice:
		color.New(color.FgWhite).Fprint(buf, "\n")
		ColorSlice(buf, in.([]interface{}), indent+2)
	case reflect.Map:
		color.New(color.FgWhite).Fprint(buf, "\n")
		ColorMap(buf, in.(map[interface{}]interface{}), indent+2)
	default:
		color.New(color.FgMagenta).Fprint(buf, fmt.Sprintf("%v\n", v))
	}
}
