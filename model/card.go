package model

type Deck struct {
	Cards []Door `json:"cards,omitempty"`
}

type Door struct {
	CardType     string `json:"cardType,omitempty"`
	CardCategory string `json:"cardCategory,omitempty"`
	Text         string `json:"text,omitempty"`
	Bonus        string `json:"bonus,omitempty"`
	TextAdd      string `json:"text_add,omitempty"`
}
