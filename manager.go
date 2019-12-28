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

// Manager is responsible for managing and persisting policies.
type Manager interface {
	// Initialize the manager.
	Initialize() error

	// Create persists the policy.
	Create(policy Policy) error

	// Update updates an existing policy.
	Update(policy Policy) error

	// Get retrieves a policy.
	Get(id string) (Policy, error)

	// Delete removes a policy.
	Delete(id string) error

	// GetAll retrieves all policies.
	GetAll(limit, offset int) (Policies, error)

	// FindRequestCandidates returns candidates that could match the request object. It either returns
	// a set that exactly matches the request, or a superset of it. If an error occurs, it returns nil and
	// the error.
	FindRequestCandidates(r *Request) (Policies, error)

	// FindPoliciesForSubject returns policies that could match the subject in a paginated fashion.
	FindPoliciesForSubject(subject string, limit int, offset int) (Policies, error)

	// FindPoliciesForResource returns policies that could match the resource. It either returns
	// a set of policies that apply to the resource, or a superset of it.
	// If an error occurs, it returns nil and the error.
	FindPoliciesForResource(resource string) (Policies, error)
}
