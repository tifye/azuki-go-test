package azuki

type TextInputComponent struct {
	BaseComponent
	Value TextSource `json:"initialValue,omitzero"`
}

func TextInput() TextInputComponent {
	return TextInputComponent{
		BaseComponent: newBaseComponent(ComponentTypeTextInput),
	}
}

func (ti TextInputComponent) WithValue(t TextSource) TextInputComponent {
	ti.Value = t
	return ti
}
