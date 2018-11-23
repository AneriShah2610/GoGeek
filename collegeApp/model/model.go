package model

type (
	// College model
	College struct {
		ID      int    `json:"id"`
		Name    string `json:"name"`
		Address string `json:"address"`
	}
)
