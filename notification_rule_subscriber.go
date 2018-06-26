
package gowebapi

type NotificationRuleSubscriber struct {
	WebId string `json:"WebId,omitempty"`

	Id string `json:"Id,omitempty"`

	Name string `json:"Name,omitempty"`

	Description string `json:"Description,omitempty"`

	Path string `json:"Path,omitempty"`

	ConfigString string `json:"ConfigString,omitempty"`

	ContactTemplateName string `json:"ContactTemplateName,omitempty"`

	ContactType string `json:"ContactType,omitempty"`

	DeliveryFormatName string `json:"DeliveryFormatName,omitempty"`

	PlugInName string `json:"PlugInName,omitempty"`

	EscalationTimeout string `json:"EscalationTimeout,omitempty"`

	MaximumRetries int32 `json:"MaximumRetries,omitempty"`

	NotifyOption string `json:"NotifyOption,omitempty"`

	RetryInterval string `json:"RetryInterval,omitempty"`

	WebException *WebException `json:"WebException,omitempty"`
}
