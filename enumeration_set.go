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

type EnumerationSet struct {
	WebId string `json:"WebId,omitempty"`

	Id string `json:"Id,omitempty"`

	Name string `json:"Name,omitempty"`

	Description string `json:"Description,omitempty"`

	Path string `json:"Path,omitempty"`

	Links *EnumerationSetLinks `json:"Links,omitempty"`

	SerializeDescription bool `json:"SerializeDescription,omitempty"`

	WebException *WebException `json:"WebException,omitempty"`
}
