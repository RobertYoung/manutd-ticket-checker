package mutc

import (
	"github.com/go-rod/rod"
)

type UnitedPage struct {
	*rod.Page
}

func (c *UnitedPage) DeleteCookieOverlay() {
	tint_overlay, err := c.Element("#tint-overlay")

	if err == nil {
		tint_overlay.Remove()
	}

	cookie_messages, err := c.Element("#cookie-messages")

	if err == nil {
		cookie_messages.Remove()
	}
}
