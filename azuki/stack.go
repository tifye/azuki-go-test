package azuki

type Orientation string

const (
	Vertical   Orientation = "vertical"
	Horizontal Orientation = "horizontal"
)

type ChildrenSourceType string

var (
	ChildrenSourceTypeSlice ChildrenSourceType = "slice"
	ChildrenSourceTypeKey   ChildrenSourceType = "key"
)

type ChildrenSource interface {
	Type() ChildrenSourceType
}

type ChildrenSlice []Component

func (c ChildrenSlice) Type() ChildrenSourceType {
	return ChildrenSourceTypeSlice
}

type ChildrenKey string

func (c ChildrenKey) Type() ChildrenSourceType {
	return ChildrenSourceTypeKey
}

type StackComponent struct {
	BaseComponent
	Orientation Orientation    `json:"orientation"`
	Children    ChildrenSource `json:"children"`
	Gap         uint           `json:"gap,omitzero"`
}

func Stack(orientation Orientation) StackComponent {
	return StackComponent{
		BaseComponent: newBaseComponent(ComponentTypeStack),
		Orientation:   orientation,
	}
}

func (s StackComponent) WithChildren(children ...Component) StackComponent {
	s.Children = ChildrenSlice(children)
	return s
}

func (s StackComponent) WithChildrenKey(key string) StackComponent {
	s.Children = ChildrenKey(key)
	return s
}

func (s StackComponent) WithGap(n uint) StackComponent {
	s.Gap = n
	return s
}
