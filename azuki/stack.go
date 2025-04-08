package azuki

type Orientation string

const (
	Vertical   Orientation = "vertical"
	Horizontal Orientation = "horizontal"
)

type StackComponent struct {
	BaseComponent
	Orientation Orientation `json:"orientation"`
	Children    []Component `json:"children"`
	Gap         uint        `json:"gap,omitzero"`
}

func Stack(orientation Orientation) StackComponent {
	return StackComponent{
		BaseComponent: newBaseComponent(ComponentTypeStack),
		Orientation:   orientation,
		Children:      make([]Component, 0),
	}
}

func (s StackComponent) WithChildren(children ...Component) StackComponent {
	s.Children = append(s.Children, children...)
	return s
}

func (s StackComponent) WithGap(n uint) StackComponent {
	s.Gap = n
	return s
}
