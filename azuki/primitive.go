package azuki

import "time"

type TextSource interface {
	Type() TextSourceType
}

type TextSourceType string

const (
	TextSourceTypeString TextSourceType = "string"
	TextSourceTypeHTTP   TextSourceType = "http"
)

type TextSourceBase struct {
	SType TextSourceType `json:"type"`
}

func (ts TextSourceBase) Type() TextSourceType {
	return ts.SType
}

type String string

func (String) Type() TextSourceType {
	return TextSourceTypeString
}

type StringTextSource struct {
	TextSourceBase
	Value string `json:"value"`
}

type HTTPTextSource struct {
	TextSourceBase
	URL       string `json:"url"`
	Fieldpath string `json:"fieldpath"`
	PollRate  uint64 `json:"pollRate,omitzero"`
}

func HTTPText(url string) HTTPTextSource {
	return HTTPTextSource{
		TextSourceBase: TextSourceBase{SType: TextSourceTypeHTTP},
		URL:            url,
	}
}
func (h HTTPTextSource) WitFieldpath(fieldpath string) HTTPTextSource {
	h.Fieldpath = fieldpath
	return h
}
func (h HTTPTextSource) WithWatch(every time.Duration) HTTPTextSource {
	h.PollRate = uint64(every.Milliseconds())
	return h
}

func Text(text string) StringTextSource {
	return StringTextSource{
		TextSourceBase: TextSourceBase{SType: TextSourceTypeString},
		Value:          text,
	}
}
