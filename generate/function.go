package generate

import (
	"fmt"
	"html/template"

	graphql "github.com/Wryte/graphql-client-generator/graphql"
)

type functionTemplateModel struct {
	Name           string
	OperationName  string
	GoName         string
	Description    template.HTML
	Args           []functionTemplateArg
	OperationType  string
	PayloadRoot    string
	ReturnDeref    string
	ReturnPrefix   string
	ReturnType     string
	ArgsDefinition string
}

type functionTemplateArg struct {
	Name           string
	GoName         string
	Description    template.HTML
	Nullable       bool
	GoParamDef     template.HTML
	GQLParamDef    string
	GQLArgumentDef string
}

func newFunctionTemplateModel(name string, parent graphql.Type, function graphql.Field, ret graphql.Type) functionTemplateModel {
	ft := functionTemplateModel{
		OperationName: function.Name,
		Name:          name,
		GoName:        makeExportedName(function.Name),
		Description:   template.HTML(addComments(function.Description, "")),
		PayloadRoot:   parent.Name,
		ReturnType:    mapToGoScalar(makeExportedName(function.TypeName())),
		ReturnDeref:   "*",
	}

	if function.IsList() {
		ft.ReturnPrefix = "[]"
		ft.ReturnDeref = ""
	}

	if function.IsScalar() {
		ft.ReturnPrefix = ""
		ft.ReturnDeref = ""
	}

	for i, a := range function.Args {
		var (
			prefix    = "*"
			suffix    string
			nullable  = true
			comma     = ", "
			omitEmpty = ",omitempty"
		)

		if a.IsNonNull() {
			nullable = false
			prefix = ""
			suffix = "!"
			omitEmpty = ""
		}

		if a.IsList() {
			prefix = "[]"
		}

		argType := mapToGoScalar(a.TypeName())
		if argType == a.TypeName() {
			argType = makeExportedName(argType)
		}

		if i == len(function.Args)-1 {
			comma = ""
		}

		ft.Args = append(ft.Args, functionTemplateArg{
			Name:           a.Name,
			GoName:         makeExportedName(a.Name),
			Description:    template.HTML(addComments(a.Description, "\t")),
			Nullable:       nullable,
			GoParamDef:     template.HTML(fmt.Sprintf("%s %s%s `json:\"%s%s\"`", makeExportedName(a.Name), prefix, argType, a.Name, omitEmpty)),
			GQLParamDef:    fmt.Sprintf("$%s: %s%s%s", a.Name, a.TypeName(), suffix, comma),
			GQLArgumentDef: fmt.Sprintf("%s: $%s%s", a.Name, a.Name, comma),
		})
	}

	return ft
}
