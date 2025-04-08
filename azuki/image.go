package azuki

type ImageComponent struct {
	BaseComponent
	Source TextSource `json:"source"`
}

func Image(source TextSource) ImageComponent {
	return ImageComponent{
		BaseComponent: newBaseComponent(ComponentTypeImage),
		Source:        source,
	}
}
