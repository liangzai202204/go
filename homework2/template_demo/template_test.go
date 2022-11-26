package template_demo

import (
	"bytes"
	"fmt"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"html/template"
	"testing"
)

func TestHelloWorld(t *testing.T) {
	type User struct {
		Name string
	}
	tpl := template.New("hello-world")
	tpl, err := tpl.Parse(`Hello,{{ .Name}}`)
	require.NoError(t, err)
	buffer := &bytes.Buffer{}
	err = tpl.Execute(buffer, User{Name: "Ning"})
	require.NoError(t, err)
	assert.Equal(t, `Hello,Ning`, buffer.String())

}
func TestMapData(t *testing.T) {

	tpl := template.New("hello-world")
	tpl, err := tpl.Parse(`Hello,{{ .Name}}`)
	require.NoError(t, err)
	buffer := &bytes.Buffer{}
	err = tpl.Execute(buffer, map[string]string{"Name": "宁"})
	require.NoError(t, err)
	assert.Equal(t, `Hello,宁`, buffer.String())
}
func TestSliceData(t *testing.T) {

	tpl := template.New("hello-world")
	tpl, err := tpl.Parse(`Hello,{{index . 0}}`)
	require.NoError(t, err)
	buffer := &bytes.Buffer{}
	err = tpl.Execute(buffer, []string{"宁"})
	require.NoError(t, err)
	assert.Equal(t, `Hello,宁`, buffer.String())
}
func TestFucn(t *testing.T) {
	tpl := template.New("hello-world")
	tpl, err := tpl.Parse(`lenght:{{len .Slice}},Hello,{{.Hello "学" "宁"}}`)
	require.NoError(t, err)
	buffer := &bytes.Buffer{}
	err = tpl.Execute(buffer, FuncCall{
		Slice: []string{"a"},
	})
	require.NoError(t, err)
	assert.Equal(t, `lenght:1,Hello,学,宁`, buffer.String())

}

type FuncCall struct {
	Slice []string
}

func (f FuncCall) Hello(first string, last string) string {

	return fmt.Sprintf("%s,%s", first, last)
}
