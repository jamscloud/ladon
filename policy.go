/*
 * Copyright © 2016-2018 Aeneas Rekkas <aeneas+oss@aeneas.io>
 *
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 *
 *     http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 *
 * @author		Aeneas Rekkas <aeneas+oss@aeneas.io>
 * @copyright 	2015-2018 Aeneas Rekkas <aeneas+oss@aeneas.io>
 * @license 	Apache-2.0
 */

package ladon

import (
	"time"
)

// Policies is an array of policies.
type Policies []Policy

// Policy represent a policy model.
type Policy interface {
	// GetID returns the policies id.
	GetID() string

	// GetDescription returns the policies description.
	GetDescription() string

	// GetSubjects returns the policies subjects.
	GetSubjects() []string

	// AllowAccess returns true if the policy effect is allow, otherwise false.
	AllowAccess() bool

	// GetEffect returns the policies effect which might be 'allow' or 'deny'.
	GetEffect() string

	// GetResources returns the policies resources.
	GetResources() []string

	// GetActions returns the policies actions.
	GetActions() []string

	// GetConditions returns the policies conditions.
	GetConditions() Conditions

	// GetMeta returns the policies arbitrary metadata set by the user.
	GetMeta() []byte

	// GetStartDelimiter returns the delimiter which identifies the beginning of a regular expression.
	GetStartDelimiter() byte

	// GetEndDelimiter returns the delimiter which identifies the end of a regular expression.
	GetEndDelimiter() byte

	// Retrieves CreatedAt timestamp.
	GetCreatedAt() time.Time

	// Retrieves UpdatedAt timestamp.
	GetUpdatedAt() time.Time

	// Sets CreatedAt timestamp.
	SetCreatedAt(time.Time)

	// Sets UpdatedAt timestamp.
	SetUpdatedAt(time.Time)
}

// The AccessHawk default policy implementation.
type DefaultPolicy struct {
	ID          string     `json:"id" dynamo:"id,hash"`
	Description string     `json:"description,omitempty" dynamo:"description,omitempty"`
	Subjects    []string   `json:"subjects" dynamo:"subjects"`
	Effect      string     `json:"effect" dynamo:"effect"`
	Resources   []string   `json:"resources" dynamo:"resources"`
	Actions     []string   `json:"actions" dynamo:"actions"`
	Conditions  Conditions `json:"conditions,omitempty" dynamo:"conditions,omitempty"`
	Meta        []byte     `json:"meta,omitempty" dynamo:"conditions,omitempty"`
	CreatedAt   time.Time  `json:"created_at" dynamo:"created_at"`
	UpdatedAt   time.Time  `json:"updated_at" dynamo:"updated_at"`
}

// GetID returns the policies id.
func (policy *DefaultPolicy) GetID() string {
	return policy.ID
}

// GetDescription returns the policies description.
func (policy *DefaultPolicy) GetDescription() string {
	return policy.Description
}

// GetSubjects returns the policies subjects.
func (policy *DefaultPolicy) GetSubjects() []string {
	return policy.Subjects
}

// AllowAccess returns true if the policy effect is allow, otherwise false.
func (policy *DefaultPolicy) AllowAccess() bool {
	return policy.Effect == AllowAccess
}

// GetEffect returns the policies effect which might be 'allow' or 'deny'.
func (policy *DefaultPolicy) GetEffect() string {
	return policy.Effect
}

// GetResources returns the policies resources.
func (policy *DefaultPolicy) GetResources() []string {
	return policy.Resources
}

// GetActions returns the policies actions.
func (policy *DefaultPolicy) GetActions() []string {
	return policy.Actions
}

// GetConditions returns the policies conditions.
func (policy *DefaultPolicy) GetConditions() Conditions {
	return policy.Conditions
}

// GetMeta returns the policies arbitrary metadata set by the user.
func (policy *DefaultPolicy) GetMeta() []byte {
	return policy.Meta
}

// Get the CreatedAt timestamp.
func (policy *DefaultPolicy) GetCreatedAt() time.Time {
	return policy.CreatedAt
}

// Get the UpdatedAt timestamp.
func (policy *DefaultPolicy) GetUpdatedAt() time.Time {
	return policy.UpdatedAt
}

// GetEndDelimiter returns the delimiter which identifies the end of a regular expression.
func (policy *DefaultPolicy) GetEndDelimiter() byte {
	return '>'
}

// GetStartDelimiter returns the delimiter which identifies the beginning of a regular expression.
func (policy *DefaultPolicy) GetStartDelimiter() byte {
	return '<'
}

// Set the CreatedAt timestamp.
func (policy *DefaultPolicy) SetCreatedAt(createdAt time.Time) {
	policy.CreatedAt = createdAt
}

// Set the UpdatedAt timestamp.
func (policy *DefaultPolicy) SetUpdatedAt(updatedAt time.Time) {
	policy.UpdatedAt = updatedAt
}
