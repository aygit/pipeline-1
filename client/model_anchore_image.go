/*
 * Pipeline API
 *
 * Pipeline v0.3.0 swagger
 *
 * API version: 0.3.0
 * Contact: info@banzaicloud.com
 */

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package client

import (
	"time"
)

// A unique image in the engine. May have multiple tags or references. Unique to an image content across registries or repositories.
type AnchoreImage struct {
	ImageContent ImageContent `json:"imageContent,omitempty"`
	// Details specific to an image reference and type such as tag and image source
	ImageDetail []ImageDetail `json:"imageDetail,omitempty"`
	LastUpdated time.Time     `json:"lastUpdated,omitempty"`
	CreatedAt   time.Time     `json:"createdAt,omitempty"`
	ImageDigest string        `json:"imageDigest,omitempty"`
	UserId      string        `json:"userId,omitempty"`
	// State of the image
	ImageStatus string `json:"imageStatus,omitempty"`
	// A state value for the current status of the analysis progress of the image
	AnalysisStatus string `json:"analysisStatus,omitempty"`
}
