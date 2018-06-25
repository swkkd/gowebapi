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

type Response struct {
	Status int32 `json:"Status,omitempty"`

	Headers map[string]string `json:"Headers,omitempty"`

	Content *interface{} `json:"Content,omitempty"`
}
