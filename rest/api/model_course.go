/*
 * Mandatory Assignemnt 1
 *
 * Mandotory Assignment 1
 *
 * API version: 0.1
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package api

type Course struct {
	Id int64 `json:"id,omitempty"`

	Title string `json:"title,omitempty"`

	Teachers []int64 `json:"teachers,omitempty"`

	Ects int32 `json:"ects,omitempty"`

	SatisfactionScore int32 `json:"satisfaction_score,omitempty"`
}
