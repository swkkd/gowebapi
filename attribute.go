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

type Attribute struct {
	WebId string `json:"WebId,omitempty"`

	Id string `json:"Id,omitempty"`

	Name string `json:"Name,omitempty"`

	Description string `json:"Description,omitempty"`

	Path string `json:"Path,omitempty"`

	Type_ string `json:"Type,omitempty"`

	TypeQualifier string `json:"TypeQualifier,omitempty"`

	DefaultUnitsName string `json:"DefaultUnitsName,omitempty"`

	DisplayDigits int32 `json:"DisplayDigits,omitempty"`

	DataReferencePlugIn string `json:"DataReferencePlugIn,omitempty"`

	ConfigString string `json:"ConfigString,omitempty"`

	IsConfigurationItem bool `json:"IsConfigurationItem,omitempty"`

	IsExcluded bool `json:"IsExcluded,omitempty"`

	IsHidden bool `json:"IsHidden,omitempty"`

	IsManualDataEntry bool `json:"IsManualDataEntry,omitempty"`

	HasChildren bool `json:"HasChildren,omitempty"`

	CategoryNames []string `json:"CategoryNames,omitempty"`

	Step bool `json:"Step,omitempty"`

	TraitName string `json:"TraitName,omitempty"`

	DefaultUnitsNameAbbreviation string `json:"DefaultUnitsNameAbbreviation,omitempty"`

	Span float32 `json:"Span,omitempty"`

	Zero float32 `json:"Zero,omitempty"`

	Links *AttributeLinks `json:"Links,omitempty"`

	WebException *WebException `json:"WebException,omitempty"`
}
