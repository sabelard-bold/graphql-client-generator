package generate

import (
	"fmt"
	"strings"

	graphql "github.com/Wryte/graphql-client-generator/graphql"
)

type mutationTemplateModel struct {
	Function    functionTemplateModel
	ErrorFields string
}

func newMutationTemplateModel(mutation graphql.Field, ret graphql.Type, parent graphql.Type, userErrors graphql.Type) mutationTemplateModel {
	mt := mutationTemplateModel{
		Function: newFunctionTemplateModel(
			fmt.Sprintf("%sMutation", makeExportedName(mutation.Name)),
			parent,
			mutation,
			ret,
		),
	}

	parts := []string{}
	for _, field := range userErrors.Fields {
		parts = append(parts, field.Name)
	}
	mt.ErrorFields = strings.Join(parts, " ")

	return mt
}
