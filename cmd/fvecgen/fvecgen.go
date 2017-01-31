// Generates a "compound" object from an input string of the form
//
//  {name}=ssi,{name}=mssi,...
//
//  name is the field name that will also be the cassandra column name
//  scalars are of the form
//
//  {name}=d -> float64 -> double
//  {name}=i -> int64  -> bigint
//  {name}=s -> string -> varchar
//

package main

import (
	"bytes"
	"flag"
	"github.com/wyndhblb/fvec"
	"io"
	"os"
	"regexp"
	"strings"
	"text/template"
	"reflect"
)

var FVEC_PACKAGE string = "fvec"

var matchFirstCap = regexp.MustCompile("(.)([A-Z][a-z]+)")
var matchAllCap = regexp.MustCompile("([a-z0-9])([A-Z])")
var camelingRegex = regexp.MustCompile("[0-9A-Za-z]+")

func ToCamelCase(src string) string {
	byteSrc := []byte(src)
	chunks := camelingRegex.FindAll(byteSrc, -1)
	for idx, val := range chunks {
		if idx > 0 {
			chunks[idx] = bytes.Title(val)
		}
	}
	return string(bytes.Title(bytes.Join(chunks, nil)))
}
func ToSnakeCase(str string) string {
	snake := matchFirstCap.ReplaceAllString(str, "${1}_${2}")
	snake = matchAllCap.ReplaceAllString(snake, "${1}_${2}")
	return strings.ToLower(snake)
}

type GoFieldVars struct {
	FieldName  string
	FieldType  string
	ColumnName string
	IsScalar   bool
	IsVector   bool
	IsCounter  bool
}

type GoTpl struct {
	ClassName   string
	PackageName string
	FieldStr    string
	HaveCounters bool

	Fields []GoFieldVars
	tpl    *template.Template
}

func (g *GoTpl) ComposeVars() {
	g.Fields = make([]GoFieldVars, 0)
	g.ClassName = ToCamelCase(g.ClassName)

	var fns = template.FuncMap{
		"last": func(x int, a interface{}) bool {
			return x == reflect.ValueOf(a).Len() - 1
		},
	}

	g.tpl = template.Must(template.New("go-lang-fvec-gen").Funcs(fns).Parse(GOLANG_TEMPLATE))

	for _, s := range strings.Split(g.FieldStr, ",") {
		nmv := strings.Split(strings.Replace(s, " ", "", -1), "=")
		ftype := ""
		var ok bool = false
		isscal := true
		if ftype, ok = fvec.SCALAR_SHORT_NAME_MAP[nmv[1]]; !ok {
			isscal = false
			if ftype, ok = fvec.SHORT_NAME_MAP[nmv[1]]; !ok {
				panic(nmv[1] + " is not valid type")
			}
		}
		if fvec.IsCounter(ftype){
			g.HaveCounters = true
		}
		g.Fields = append(g.Fields, GoFieldVars{
			FieldName:  ToCamelCase(nmv[0]),
			FieldType:  ftype,
			ColumnName: ToSnakeCase(nmv[0]),
			IsScalar:   isscal,
			IsVector:   !isscal,
			IsCounter:  fvec.IsCounter(ftype),
		})
	}
}

func (g *GoTpl) Render(out io.Writer) {
	g.tpl.Execute(out, g)
}

func main() {
	mode := flag.String("mode", "class", "Mode: for now only class")
	clsname := flag.String("classname", "", "the output struct name")
	pkgname := flag.String("pkgname", "", "the output package name")
	pstr := flag.String("format", "", "the shorthand for the items in the class (like {varname}=s,{varname}=mssi,{varname}=i")

	flag.Parse()

	if len(*pstr) == 0 {
		panic("format cannot be blank")
	}

	if *mode != "class" {
		panic("mode must be class for now")
	}

	goGen := &GoTpl{
		PackageName: *pkgname,
		ClassName:   *clsname,
		FieldStr:    *pstr,
		HaveCounters: false,
	}

	goGen.ComposeVars()
	goGen.Render(os.Stdout)

}
