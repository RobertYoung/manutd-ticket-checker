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
	HasSeatsAvailable  bool
	NotificationSentAt time.Time
	Model              *models.EventModel
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
	_, err := e.FindBuyButton()

	if err != nil {
		return "unavailable"
	}

	if e.HasSeatsAvailable {
		return "available"
	}

	return "unavailable"
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

func (e *UnitedEventItem) BuyButton() (*rod.Element, error) {
	return e.Timeout(5 * time.Second).Element("div.addToBasketXX > a")
}

func (e *UnitedEventItem) LoadEventDetailPage(event *UnitedEventItem) error {
	buy_button, err := event.BuyButton()

	if err != nil {
		return err
	}

	buy_button.MustEval(`() => this.target="_blank"`)
	buy_button.MustClick()

	return nil
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
	notification_sent_at := e.NotificationSentAt

	if e.NotificationSentAt.IsZero() && e.Model != nil {
		notification_sent_at = e.Model.NotificationSentAt
	}

	return models.EventModel{
		Uuid:               e.Uuid(),
		Name:               e.Name(),
		MinPrice:           e.MinPrice,
		MaxPrice:           e.MaxPrice,
		NotificationSentAt: notification_sent_at,
	}
}
