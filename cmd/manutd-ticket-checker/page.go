package mutc

import (
	"github.com/go-rod/rod"
)

type UnitedPage struct {
	*rod.Page
}

func (c *UnitedPage) DeleteCookieOverlay() {
	tint_overlay := c.MustElement("#tint-overlay")
	tint_overlay.Remove()

	cookie_messages := c.MustElement("#cookie-messages")
	cookie_messages.Remove()
}
