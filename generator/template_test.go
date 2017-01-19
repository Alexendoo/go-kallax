package generator

import (
	"bytes"
	"go/ast"
	"go/importer"
	"go/parser"
	"go/token"
	"go/types"
	"testing"

	"github.com/stretchr/testify/suite"
)

type TemplateSuite struct {
	suite.Suite
	td *TemplateData
}

func (s *TemplateSuite) SetupTest() {
	s.td = &TemplateData{
		nil,
		make(map[interface{}]string),
		make(map[string]*Field),
	}
}

func (s *TemplateSuite) processSource(source string) {
	fset := &token.FileSet{}
	astFile, err := parser.ParseFile(fset, "fixture.go", source, 0)
	s.Nil(err)

	cfg := &types.Config{
		Importer: importer.For("gc", nil),
	}
	p, err := cfg.Check("foo", fset, []*ast.File{astFile}, nil)
	s.Nil(err)

	prc := NewProcessor("fixture", []string{"foo.go"})
	prc.Package = p
	s.td.Package, err = prc.processPackage()
	s.Nil(err)
}

const expectedAddresses = `case "id":
return &r.Model.ID, nil
case "foo":
return &r.Foo, nil
case "bar":
return r.Bar, nil
case "arr":
return types.Slice(&r.Arr), nil
case "json":
return types.JSON(&r.JSON), nil
`

const baseTpl = `
	package fixture

	import "github.com/src-d/go-kallax"

	type Rel struct {
		kallax.Model
		Foo string
	}

	type JSON struct {
		Foo string
	}

	type Foo struct {
		kallax.Model
		Foo string
		Bar *string
		Arr []string
		JSON JSON
		Rel Rel
	}
`

func (s *TemplateSuite) TestGenColumnAddresses() {
	s.processSource(baseTpl)

	m := findModel(s.td.Package, "Foo")
	result := s.td.GenColumnAddresses(m)
	s.Equal(expectedAddresses, result)
}

const expectedValues = `case "id":
return r.Model.ID, nil
case "foo":
return r.Foo, nil
case "bar":
return r.Bar, nil
case "aliased":
return (string)(r.Aliased), nil
case "arr":
return types.Slice(r.Arr), nil
case "json":
return types.JSON(r.JSON), nil
`

func (s *TemplateSuite) TestGenColumnValues() {
	s.processSource(`
	package fixture

	import "github.com/src-d/go-kallax"

	type Aliased string

	type Rel struct {
		kallax.Model
		Foo string
	}

	type JSON struct {
		Foo string
	}

	type Foo struct {
		kallax.Model
		Foo string
		Bar *string
		Aliased Aliased
		Arr []string
		JSON JSON
		Rel Rel
	}
	`)

	m := findModel(s.td.Package, "Foo")
	result := s.td.GenColumnValues(m)
	s.Equal(expectedValues, result)
}

const expectedColumns = `kallax.NewSchemaField("id"),
kallax.NewSchemaField("foo"),
kallax.NewSchemaField("bar"),
kallax.NewSchemaField("arr"),
kallax.NewSchemaField("json"),
`

func (s *TemplateSuite) TestGenModelColumns() {
	s.processSource(baseTpl)
	m := findModel(s.td.Package, "Foo")
	result := s.td.GenModelColumns(m)
	s.Equal(expectedColumns, result)
}

const expectedSchema = `ID kallax.SchemaField
Foo kallax.SchemaField
Bar kallax.SchemaField
Arr kallax.SchemaField
JSON *schemaFooJSON
`

const expectedSubSchemas = `type schemaFooJSON struct {
*kallax.BaseSchemaField
Foo kallax.SchemaField
}

`

func (s *TemplateSuite) TestGenModelSchema() {
	s.processSource(baseTpl)
	m := findModel(s.td.Package, "Foo")
	result := s.td.GenModelSchema(m)
	s.Equal(expectedSchema, result)
	s.Equal(expectedSubSchemas, s.td.GenSubSchemas())
}

const expectedInit = `ID:kallax.NewSchemaField("id"),
Foo:kallax.NewSchemaField("foo"),
Bar:kallax.NewSchemaField("bar"),
Arr:kallax.NewSchemaField("arr"),
JSON:&schemaFooJSON{
BaseSchemaField: kallax.NewSchemaField("json").(*kallax.BaseSchemaField),
Foo:kallax.NewSchemaField("Foo"),
},
`

func (s *TemplateSuite) TestGenSchemaInit() {
	s.processSource(baseTpl)
	m := findModel(s.td.Package, "Foo")

	s.Equal(expectedInit, s.td.GenSchemaInit(m))
}

func (s *TemplateSuite) TestExecute() {
	s.processSource(baseTpl)
	var buf bytes.Buffer
	s.Nil(Base.Execute(&buf, s.td.Package))
}

func TestTemplate(t *testing.T) {
	suite.Run(t, new(TemplateSuite))
}
