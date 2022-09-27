package graphql

// Schema the representation of the GraphQL schema
type Schema struct {
	Query struct {
		Type Type
		Name string `json:"name"`
	} `json:"queryType"`
	Mutation struct {
		Type Type
		Name string `json:"name"`
	} `json:"mutationType"`
	Types        []Type `json:"types"`
	typeLookup   map[string]Type
	errorsLookup map[string]bool
}

// Type retrieve a type by name
func (s *Schema) Type(name string) (Type, bool) {
	t, ok := s.typeLookup[name]

	return t, ok
}

// IsError check if a type name is an error
func (s *Schema) IsError(name string) bool {
	return s.errorsLookup[name]
}

func (s *Schema) Init() {
	s.typeLookup = map[string]Type{}
	s.errorsLookup = map[string]bool{}

	for _, t := range s.Types {
		s.typeLookup[t.Name] = t

		if t.Name == s.Mutation.Name {
			s.Mutation.Type = t
		}

		if t.Name == s.Query.Name {
			s.Query.Type = t
		}

		for _, field := range t.Fields {
			if field.IsError() {
				s.errorsLookup[field.TypeName()] = true
			}
		}
	}
}
