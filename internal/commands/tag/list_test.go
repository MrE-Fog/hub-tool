/*
   Copyright 2020 Docker Hub Tool authors

   Licensed under the Apache License, Version 2.0 (the "License");
   you may not use this file except in compliance with the License.
   You may obtain a copy of the License at

       http://www.apache.org/licenses/LICENSE-2.0

   Unless required by applicable law or agreed to in writing, software
   distributed under the License is distributed on an "AS IS" BASIS,
   WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
   See the License for the specific language governing permissions and
   limitations under the License.
*/

package tag

import (
	"testing"

	"gotest.tools/v3/assert"
)

func TestMappingSortFieldToOrderingAPI(t *testing.T) {
	testCases := []struct {
		name          string
		sort          string
		ordering      string
		expectedError string
	}{
		{
			name:     "default",
			sort:     "",
			ordering: "",
		},
		{
			name:          "invalid sort by",
			sort:          "invalid",
			ordering:      "",
			expectedError: `unknown sorting column "invalid": should be either "name" or "updated"`,
		},
		{
			name:     "ascending order by default",
			sort:     "updated",
			ordering: "last_updated",
		},
		{
			name:     "updated ascending",
			sort:     "updated=asc",
			ordering: "last_updated",
		},
		{
			name:     "updated descending",
			sort:     "updated=desc",
			ordering: "-last_updated",
		},
		{
			name:     "name ascending by default",
			sort:     "name",
			ordering: "-name",
		},
		{
			name:     "name ascending",
			sort:     "name=asc",
			ordering: "-name",
		},
		{
			name:     "name descending",
			sort:     "name=desc",
			ordering: "name",
		},
	}
	for _, testCase := range testCases {
		t.Run(testCase.name, func(t *testing.T) {
			actualOrdering, actualError := mapOrdering(testCase.sort)
			if testCase.expectedError != "" {
				assert.Error(t, actualError, testCase.expectedError)
			} else {
				assert.Equal(t, actualOrdering, testCase.ordering)
			}
		})
	}
}
