/*
 * PI Web API 2018 Swagger Spec
 *
 * Swagger Spec file that describes PI Web API
 *
 * API version: 1.11.0.640
 *
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package gowebapi

type Analysis struct {
	WebId string `json:"WebId,omitempty"`

	Id string `json:"Id,omitempty"`

	Name string `json:"Name,omitempty"`

	Description string `json:"Description,omitempty"`

	Path string `json:"Path,omitempty"`

	AnalysisRulePlugInName string `json:"AnalysisRulePlugInName,omitempty"`

	AutoCreated bool `json:"AutoCreated,omitempty"`

	CategoryNames []string `json:"CategoryNames,omitempty"`

	GroupId int32 `json:"GroupId,omitempty"`

	HasNotification bool `json:"HasNotification,omitempty"`

	HasTarget bool `json:"HasTarget,omitempty"`

	HasTemplate bool `json:"HasTemplate,omitempty"`

	IsConfigured bool `json:"IsConfigured,omitempty"`

	IsTimeRuleDefinedByTemplate bool `json:"IsTimeRuleDefinedByTemplate,omitempty"`

	MaximumQueueSize int32 `json:"MaximumQueueSize,omitempty"`

	OutputTime string `json:"OutputTime,omitempty"`

	Priority string `json:"Priority,omitempty"`

	PublishResults bool `json:"PublishResults,omitempty"`

	Status string `json:"Status,omitempty"`

	TargetWebId string `json:"TargetWebId,omitempty"`

	TemplateName string `json:"TemplateName,omitempty"`

	TimeRulePlugInName string `json:"TimeRulePlugInName,omitempty"`

	Links *AnalysisLinks `json:"Links,omitempty"`

	WebException *WebException `json:"WebException,omitempty"`
}
