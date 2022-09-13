// package generate

// import (
// 	"fmt"
// 	"html/template"

// 	graphql "github.com/Wryte/graphql-client-generator/graphql"
// )

// type queryTemplateModel struct {
// 	Name           string
// 	FunctionName   string
// 	GoName         string
// 	Description    template.HTML
// 	Args           []queryTemplateArg
// 	OperationType  string
// 	PayloadRoot    string
// 	ReturnDeref    string
// 	ReturnPrefix   string
// 	ReturnType     string
// 	ArgsDefinition string
// }

// type queryTemplateArg struct {
// 	Name           string
// 	GoName         string
// 	Description    template.HTML
// 	Nullable       bool
// 	GoParamDef     template.HTML
// 	GQLParamDef    string
// 	GQLArgumentDef string
// }

// func newQueryTemplateModel(query graphql.Field, ret graphql.Type, parent graphql.Type) queryTemplateModel {
// 	qt := queryTemplateModel{
// 		Name:         query.Name,
// 		FunctionName: fmt.Sprintf("Query%s", makeExportedName(query.Name)),
// 		GoName:       makeExportedName(query.Name),
// 		Description:  template.HTML(addComments(query.Description, "")),
// 		PayloadRoot:  parent.Name,
// 		ReturnType:   mapToGoScalar(makeExportedName(query.TypeName())),
// 		ReturnDeref:  "*",
// 	}

// 	if query.IsList() {
// 		qt.ReturnPrefix = "[]"
// 		qt.ReturnDeref = ""
// 	}

// 	if query.IsScalar() {
// 		qt.ReturnPrefix = ""
// 	}

// 	for i, a := range query.Args {
// 		var (
// 			prefix    = "*"
// 			suffix    string
// 			nullable  = true
// 			comma     = ", "
// 			omitEmpty = ",omitempty"
// 		)

// 		if a.IsNonNull() {
// 			nullable = false
// 			prefix = ""
// 			suffix = "!"
// 			omitEmpty = ""
// 		}

// 		if a.IsList() {
// 			prefix = "[]"
// 		}

// 		argType := mapToGoScalar(a.TypeName())
// 		if argType == a.TypeName() {
// 			argType = makeExportedName(argType)
// 		}

// 		if i == len(query.Args)-1 {
// 			comma = ""
// 		}

// 		qt.Args = append(qt.Args, queryTemplateArg{
// 			Name:           a.Name,
// 			GoName:         makeExportedName(a.Name),
// 			Description:    template.HTML(addComments(a.Description, "\t")),
// 			Nullable:       nullable,
// 			GoParamDef:     template.HTML(fmt.Sprintf("%s %s%s `json:\"%s%s\"`", makeExportedName(a.Name), prefix, argType, a.Name, omitEmpty)),
// 			GQLParamDef:    fmt.Sprintf("$%s: %s%s%s", a.Name, a.TypeName(), suffix, comma),
// 			GQLArgumentDef: fmt.Sprintf("%s: $%s%s", a.Name, a.Name, comma),
// 		})
// 	}

// 	return qt
// }

package generate

import (
	"fmt"

	graphql "github.com/Wryte/graphql-client-generator/graphql"
)

type queryTemplateModel struct {
	Function functionTemplateModel
}

func newQueryTemplateModel(query graphql.Field, ret graphql.Type, parent graphql.Type) queryTemplateModel {
	qt := queryTemplateModel{
		Function: newFunctionTemplateModel(
			fmt.Sprintf("Query%s", makeExportedName(query.Name)),
			parent,
			query,
			ret,
		),
	}

	return qt
}
