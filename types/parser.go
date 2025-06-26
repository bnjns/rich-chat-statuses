package types

type StatusPreset struct {
	Events       []string `json:"events"`
	Emoji        *string  `json:"emoji"`
	DoNotDisturb *bool    `json:"doNotDisturb"`
	Away         *bool    `json:"away"`
}
