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

type ItemsAnnotation struct {
	Items []Annotation `json:"Items,omitempty"`

	Links *PaginationLinks `json:"Links,omitempty"`
}
