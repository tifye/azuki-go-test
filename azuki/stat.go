package azuki

type StatComponent struct {
	BaseComponent
	Title       TextSource `json:"title,omitzero"`
	Value       TextSource `json:"value"`
	Description TextSource `json:"description,omitempty"`
	Place       Place      `json:"place,omitzero"`
}

func Stat(value TextSource) StatComponent {
	return StatComponent{
		BaseComponent: newBaseComponent(ComponentTypeStat),
		Value:         value,
		Place:         PlaceStart,
	}
}

func (s StatComponent) WithTitle(title string) StatComponent {
	return s.WithTitleSource(Text(title))
}
func (s StatComponent) WithTitleSource(title TextSource) StatComponent {
	s.Title = title
	return s
}

func (s StatComponent) WithDescription(desc string) StatComponent {
	return s.WithDescriptionSource(Text(desc))
}
func (s StatComponent) WithDescriptionSource(desc TextSource) StatComponent {
	s.Description = desc
	return s
}

type Place string

const (
	PlaceStart  Place = "start"
	PlaceCenter Place = "center"
	PlaceEnd    Place = "end"
)

func (s StatComponent) WithPlace(place Place) StatComponent {
	s.Place = place
	return s
}
