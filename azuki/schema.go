package azuki

type Schema struct {
	Components []Component `json:"components"`
}

func NewSchema(components ...Component) Schema {
	return Schema{
		Components: components,
	}
}
