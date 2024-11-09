package model

import (
	"fmt"
	"github.com/google/uuid"
)

type Alert struct {
	ID           uuid.UUID
	IsCritical   bool
	AlertMessage string
	Resolved     bool
}

func NewAlert(isCritical bool, alertMessage string) *Alert {
	return &Alert{
		ID:           uuid.New(),
		IsCritical:   isCritical,
		AlertMessage: alertMessage,
		Resolved:     false,
	}
}

func (a *Alert) UpdateMessage(newMessage string) error {
	if newMessage == "" {
		return fmt.Errorf("nova mensagem não pode ser vazia")
	}
	a.AlertMessage = newMessage
	return nil
}

func (a *Alert) Validate() error {
	if a.AlertMessage == "" {
		return fmt.Errorf("mensagem de alerta não pode ser vazia")
	}
	return nil
}

func (a *Alert) Clone() *Alert {
	return &Alert{
		ID:           uuid.New(),
		IsCritical:   a.IsCritical,
		AlertMessage: a.AlertMessage,
	}
}

func FilterCritical(alerts []*Alert) []*Alert {
	var criticalAlerts []*Alert
	for _, alert := range alerts {
		if alert.IsCritical {
			criticalAlerts = append(criticalAlerts, alert)
		}
	}
	return criticalAlerts
}

func (a *Alert) MarkResolved() {
	a.IsCritical = false
	a.Resolved = true
}

func (a *Alert) Equals(other *Alert) bool {
	return a.ID == other.ID &&
		a.IsCritical == other.IsCritical &&
		a.AlertMessage == other.AlertMessage
}

func (a *Alert) IsUrgent() bool {
	return a.IsCritical && a.AlertMessage != ""
}

func (a *Alert) GetID() uuid.UUID {
	return a.ID
}

func (a *Alert) GetIsCritical() bool {
	return a.IsCritical
}

func (a *Alert) GetAlertMessage() string {
	return a.AlertMessage
}

func (a *Alert) SetCritical(isCritical bool) {
	a.IsCritical = isCritical
}
