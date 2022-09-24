package main

type HomeAssistantMatchStateAttributes struct {
	MinPrice     uint16 `json:"min_price"`
	MaxPrice     uint16 `json:"max_price"`
	FriendlyName string `json:"friendly_name"`
}
