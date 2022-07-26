/*
 * Copyright 2018 - Present Okta, Inc.
 *
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 *
 *      http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 */

package integration

import (
	"sync"
	"testing"
	"time"

	"github.com/okta/okta-sdk-golang/v2/okta"
	"github.com/stretchr/testify/assert"
)

// TestGroupProfileGoroutineSafe relates to Issue 296
// https://github.com/okta/okta-sdk-golang/issues/296
func TestGroupProfileGoroutineSafe(t *testing.T) {
	gp := &okta.GroupProfile{
		Name: testName("SDK_TEST Get Test Group"),
		// GroupProfileMap: map[string]interface{}{},
	}

	var wg sync.WaitGroup
	for i := 0; i < 500; i++ {
		wg.Add(1)
		go func(wg *sync.WaitGroup, key, value string, profile *okta.GroupProfile) {
			// NOTE: before GroupProfileMap was made a sync.Map this loop
			// calling many go routines would cause a "fatal error: concurrent
			// map writes" panic as GroupProfileMap was a golang map.

			// profile.GroupProfileMap[kjey] = name

			profile.GroupProfileMap.Store(key, value)
			time.Sleep(time.Duration(time.Millisecond))
			wg.Done()
		}(&wg, "example", randomTestString(), gp)
	}
	_, found := gp.GroupProfileMap.Load("example")
	assert.True(t, found, "example value not found in GroupProfileMap")
	wg.Wait()
}
