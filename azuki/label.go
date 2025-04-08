package azuki

type LabelComponent struct {
	BaseComponent
	Text TextSource `json:"text"`
	Size int        `json:"size,omitzero"`
}

func Label(text TextSource) LabelComponent {
	return LabelComponent{
		BaseComponent: newBaseComponent(ComponentTypeLabel),
		Text:          text,
	}
}

func (l LabelComponent) WithSize(n int) LabelComponent {
	l.Size = n
	return l
}
