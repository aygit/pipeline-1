/*
 * Pipeline API
 *
 * Pipeline v0.3.0 swagger
 *
 * API version: 0.3.0
 * Contact: info@banzaicloud.com
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package client

type AddClusterProfileRequest struct {
	Name string `json:"name,omitempty"`
	Location string `json:"location,omitempty"`
	Cloud string `json:"cloud,omitempty"`
	Properties map[string]interface{} `json:"properties,omitempty"`
}