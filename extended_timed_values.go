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

type ExtendedTimedValues struct {
	Items []ExtendedTimedValue `json:"Items,omitempty"`

	UnitsAbbreviation string `json:"UnitsAbbreviation,omitempty"`

	WebException *WebException `json:"WebException,omitempty"`
}
