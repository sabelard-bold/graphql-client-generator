package generate

import (
	"fmt"
	"strings"

	graphql "github.com/Wryte/graphql-client-generator/graphql"
)

type mutationTemplateModel struct {
	Function         functionTemplateModel
	GoErrorFieldName string
	ErrorFields      string
}

func newMutationTemplateModel(schema graphql.Schema, mutation graphql.Field) (mutationTemplateModel, error) {
	ret, ok := schema.Type(mutation.TypeName())
	if !ok {
		return mutationTemplateModel{}, fmt.Errorf("could not find type for mutation field: %s", mutation.TypeName())
	}

	mt := mutationTemplateModel{
		Function: newFunctionTemplateModel(
			fmt.Sprintf("%sMutation", makeExportedName(mutation.Name)),
			schema.Mutation.Type,
			mutation,
			ret,
		),
	}

	errField := ret.ErrorField()
	if errField != nil {
		mt.GoErrorFieldName = makeExportedName(errField.Name)

		t, ok := schema.Type(errField.TypeName())
		if !ok {
			return mutationTemplateModel{}, fmt.Errorf("could not find type for mutation user error: %s", errField.TypeName())
		}

		var errorParts []string
		for _, field := range t.Fields {
			errorParts = append(errorParts, field.Name)
		}

		mt.ErrorFields = strings.Join(errorParts, " ")
	}

	return mt, nil
}
