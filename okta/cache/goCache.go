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

package cache

import (
	"strings"
	"time"

	"github.com/okta/okta-sdk-golang/okta"
	patrickmnGoCache "github.com/patrickmn/go-cache"
)

type GoCache struct {
	Config      *okta.Config
	RootLibrary *patrickmnGoCache.Cache
}

func NewGoCache(config okta.Config) GoCache {
	c := patrickmnGoCache.New(time.Duration(config.Okta.Client.Cache.DefaultTtl),
		time.Duration(config.Okta.Client.Cache.DefaultTti))

	gc := GoCache{
		Config:      &config,
		RootLibrary: c,
	}

	return gc
}

func (c GoCache) Get(key string) interface{} {
	item, found := c.RootLibrary.Get(key)
	if found {
		return item
	}

	return nil
}

func (c GoCache) Set(key string, value interface{}) {
	c.RootLibrary.Set(key, value, time.Duration(c.Config.Okta.Client.Cache.DefaultTtl))
}

func (c GoCache) Delete(key string) {
	c.RootLibrary.Delete(key)
}

func (c GoCache) Clear() {
	c.RootLibrary.Flush()
}

func (c GoCache) Has(key string) bool {
	_, found := c.RootLibrary.Get(key)
	return found
}

func (c GoCache) CreateKey(uri string) string {
	s := uri
	s = strings.Replace(s, "/", "_", -1)
	s = strings.Replace(s, ".", "_", -1)
	s = strings.Replace(s, "-", "_", -1)
	s = strings.Replace(s, "@", "_", -1)
	s = strings.Replace(s, "+", "_", -1)
	s = strings.Replace(s, ":", "_", -1)
	return s
}
