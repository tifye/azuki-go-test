package azuki

type ButtonComponent struct {
	BaseComponent
	Target             string     `json:"target"`
	Text               TextSource `json:"text"`
	InvalidatesTargets []string   `json:"invalidatesTargets,omitzero"`
}

func (c ButtonComponent) WithText(text string) ButtonComponent {
	c.Text = Text(text)
	return c
}

func (c ButtonComponent) WithTextSource(source TextSource) ButtonComponent {
	c.Text = source
	return c
}

func (c ButtonComponent) WithInvalidatesTargets(targets ...string) ButtonComponent {
	if c.InvalidatesTargets == nil {
		c.InvalidatesTargets = make([]string, 0, len(targets))
	}
	c.InvalidatesTargets = append(c.InvalidatesTargets, targets...)
	return c
}

func Button(target string) ButtonComponent {
	return ButtonComponent{
		BaseComponent: newBaseComponent(ComponentTypeButton),
		Target:        target,
	}
}
