package out

import (
	"bytes"
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
	"gopkg.in/yaml.v2"
)

type ColorSuite struct {
	suite.Suite
}

func (suite *ColorSuite) SetupTest() {
}

func (suite *ColorSuite) TestColorMap() {
	buf := new(bytes.Buffer)
	var t map[interface{}]interface{}
	yamlStr := `a: 1
b: 2
c: 3
`
	err := yaml.Unmarshal([]byte(yamlStr), &t)
	if err != nil {
		fmt.Println(err)
	}

	ColorMap(buf, t, 0)
	fmt.Println(buf.String())
	//TODO find out how to assert color bytes
	assert.Equal(suite.T(), yamlStr, buf.String())
}

func (suite *ColorSuite) TestColorSlice() {
	buf := new(bytes.Buffer)
	var t []interface{}
	yamlStr := `- 1
- 2
- 3
`
	err := yaml.Unmarshal([]byte(yamlStr), &t)
	if err != nil {
		fmt.Println(err)
	}

	ColorSlice(buf, t, 0)
	fmt.Println(buf.String())
	//TODO find out how to assert color bytes
	assert.Equal(suite.T(), yamlStr, buf.String())
}

func TestColorSuite(t *testing.T) {
	suite.Run(t, new(ColorSuite))
}
