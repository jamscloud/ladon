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

package ladon_test

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"

	"github.com/smw-104/ladon"
	. "github.com/smw-104/ladon"
)

var policyConditions = Conditions{
	"owner": &EqualsSubjectCondition{},
}

var policyCases = []*DefaultPolicy{
	{
		ID:          "1",
		Description: "description",
		Subjects:    []string{"user"},
		Effect:      AllowAccess,
		Resources:   []string{"articles:<[0-9]+>"},
		Actions:     []string{"create", "update"},
		Conditions:  policyConditions,
	},
	{
		Effect:     DenyAccess,
		Conditions: make(Conditions),
	},
}

type TestMeta struct {
	Key string `json:"key"`
}

func TestHasAccess(t *testing.T) {
	assert.True(t, policyCases[0].AllowAccess())
	assert.False(t, policyCases[1].AllowAccess())
}

func TestDefaultPolicy(t *testing.T) {
	policy := &DefaultPolicy{
		ID:          "id",
		Description: "desc",
		Subjects:    []string{"subjects"},
		Effect:      "effect",
		Resources:   []string{"resources"},
		Actions:     []string{"actions"},
		Conditions:  ladon.Conditions{"key": &ladon.StringEqualCondition{Equals: "value"}},
		Meta:        []byte("meta"),
		CreatedAt:   time.Time{},
		UpdatedAt:   time.Time{},
	}

	assert.Equal(t, "id", policy.GetID(), "ID not equal")
	assert.Equal(t, "desc", policy.GetDescription(), "Description not equal")
	assert.Equal(t, []string{"subjects"}, policy.GetSubjects(), "Subject not equal")
	assert.Equal(t, "effect", policy.GetEffect(), "Effect not equal")
	assert.Equal(t, []string{"resources"}, policy.GetResources(), "Resources not equal")
	assert.Equal(t, []string{"actions"}, policy.GetActions(), "Actions not equal")
	assert.Equal(t, []byte("meta"), policy.GetMeta(), "Meta not equal")
	assert.Equal(t, time.Time{}, policy.CreatedAt, "CreatedAt not equal")
	assert.Equal(t, time.Time{}, policy.UpdatedAt, "UpdatedAt not equal")
	assert.Equal(t, ladon.Conditions{"key": &ladon.StringEqualCondition{Equals: "value"}},
		policy.Conditions, "Conditions not equal")
}
