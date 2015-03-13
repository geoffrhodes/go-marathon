/*
Copyright 2014 Rohith All rights reserved.

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

package marathon

import (
	"testing"
)

func TestGroups(t *testing.T) {
	client := NewFakeMarathonEndpoint(t)
	groups, err := client.Groups()
	assertOnError(err, t)
	assertOnNull(groups, t)
	assertOnInteger(len(groups.Groups), 1, t)
	group := groups.Groups[0]
	assertOnString(group.ID, FAKE_GROUP_NAME, t)
}

func TestGroup(t *testing.T) {
	client := NewFakeMarathonEndpoint(t)
	group, err := client.Group(FAKE_GROUP_NAME)
	assertOnError(err, t)
	assertOnNull(group, t)
	assertOnInteger(len(group.Apps), 1, t)
	assertOnString(group.ID, FAKE_GROUP_NAME, t)
}
