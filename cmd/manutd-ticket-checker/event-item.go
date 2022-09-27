package mutc

import (
	"errors"
	"fmt"
	"strings"
	"time"

	models "github.com/robertyoung/manutd-ticket-checker/v2/cmd/manutd-ticket-checker/models"
	haas "github.com/robertyoung/manutd-ticket-checker/v2/pkg/home-assistant"

	"github.com/go-rod/rod"
)

type UnitedEventItem struct {
	*rod.Element

	config   *Config
	haas_api *haas.HomeAssistantAPI

	MinPrice, MaxPrice uint16
	NotificationSentAt time.Time
}

func (e *UnitedEventItem) Uuid() string {
	id, err := e.Attribute("data-id")

	if err != nil {
		return "0"
	}

	return *id
}

func (e *UnitedEventItem) Name() string {
	name, err := e.Attribute("data-name")

	if err != nil {
		return "unknown"
	}

	return *name
}

func (e *UnitedEventItem) Opponent() string {
	name := e.Name()
	split := strings.Split(name, "v")

	if len(split) == 2 {
		return split[1]
	}

	return "unknown"
}

func (e *UnitedEventItem) EntityId() string {
	value := e.Opponent()
	value = strings.Trim(value, " ")
	value = strings.ToLower(value)
	value = strings.ReplaceAll(value, " ", "_")
	value = strings.ReplaceAll(value, ".", "")

	return value
}

func (e *UnitedEventItem) State() string {
	if e.MinPrice >= uint16(e.config.MaxPrice) {
		return "unavailable"
	}

	_, err := e.FindBuyButton()

	if err != nil {
		return "unavailable"
	}

	return "available"
}

func (e *UnitedEventItem) FindBuyButton() (*rod.Element, error) {
	element, err := e.Element.Element("div.addToBasket:not([style*='display: none']) > a")

	if err != nil {
		return nil, err
	}

	if strings.ToUpper(element.MustText()) == UNITED_BUY_BUTTON_TEXT {
		return element, err
	}

	return nil, errors.New("buy button not found")
}

func (e *UnitedEventItem) IsPremierLeagueEvent() bool {
	button, err := e.Element.Element("img.item_image.otherImageMediumImageUrl")

	if err != nil {
		return false
	}

	if strings.Contains(*button.MustAttribute("src"), UNITED_PREMIER_IMAGE_ID) {
		return true
	}

	return false
}

func (e *UnitedEventItem) BuyButton() *rod.Element {
	return e.MustElement("div.addToBasket > a")
}

func (e *UnitedEventItem) LoadEventDetailPage(event *UnitedEventItem) {
	buy_button := event.BuyButton()
	buy_button.MustEval(`() => this.target="_blank"`)
	buy_button.MustClick()
}

func (e *UnitedEventItem) UpdateState() {
	request := haas.HomeAssistantStateUpdateRequest{
		State: e.State(),
		Attribute: HomeAssistantMatchStateAttributes{
			MinPrice:     e.MinPrice,
			MaxPrice:     e.MaxPrice,
			FriendlyName: e.Name(),
		},
	}
	e.haas_api.StateUpdate(fmt.Sprintf("entity.united_ticket_home_%s", e.EntityId()), request)
}

func (e *UnitedEventItem) NotificationSent() {
	e.NotificationSentAt = time.Now()
}

func (e *UnitedEventItem) ToEventModel() models.EventModel {
	return models.EventModel{
		Uuid:               e.Uuid(),
		Name:               e.Name(),
		MinPrice:           e.MinPrice,
		MaxPrice:           e.MaxPrice,
		NotificationSentAt: e.NotificationSentAt,
	}
}
