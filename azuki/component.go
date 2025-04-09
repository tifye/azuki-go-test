package azuki

type ComponentType string

var (
	ComponentTypeButton    ComponentType = "button"
	ComponentTypeLabel     ComponentType = "label"
	ComponentTypeImage     ComponentType = "image"
	ComponentTypeStack     ComponentType = "stack"
	ComponentTypeStat      ComponentType = "stat"
	ComponentTypeTextInput ComponentType = "textInput"
)

type BaseComponent struct {
	ComponentType ComponentType `json:"type"`
}

func newBaseComponent(t ComponentType) BaseComponent {
	return BaseComponent{ComponentType: t}
}

func (c BaseComponent) Type() ComponentType {
	return c.ComponentType
}

type Component interface {
	Type() ComponentType
}
