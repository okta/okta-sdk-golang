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

// AUTO-GENERATED!  DO NOT EDIT FILE DIRECTLY

package query

import (
	"encoding/json"
	"net/url"
	"strconv"
)

type Params struct {
	Q string `json:"q,omitempty"`
	After string `json:"after,omitempty"`
	Limit int64 `json:"limit,omitempty"`
	Filter string `json:"filter,omitempty"`
	Expand string `json:"expand,omitempty"`
	IncludeNonDeleted *bool `json:"includeNonDeleted,omitempty"`
	Activate *bool `json:"activate,omitempty"`
	TargetAid string `json:"targetAid,omitempty"`
	QueryScope string `json:"query_scope,omitempty"`
	RemoveUsers *bool `json:"removeUsers,omitempty"`
	Until string `json:"until,omitempty"`
	Since string `json:"since,omitempty"`
	SortOrder string `json:"sortOrder,omitempty"`
	Format string `json:"format,omitempty"`
	Search string `json:"search,omitempty"`
	Provider string `json:"provider,omitempty"`
	ShowAll *bool `json:"showAll,omitempty"`
	SendEmail *bool `json:"sendEmail,omitempty"`
	UpdatePhone *bool `json:"updatePhone,omitempty"`
	TemplateId string `json:"templateId,omitempty"`
	TempPassword *bool `json:"tempPassword,omitempty"`
	OauthTokens *bool `json:"oauthTokens,omitempty"`
}

func NewQueryParams(paramOpt ...ParamOptions) *Params {
	p := &Params{}

	for _, par := range paramOpt {
		par(p)
	}

	return p
}

type ParamOptions func(*Params)

func WithQ(q string) ParamOptions {
	return func(p *Params) {
		p.Q = q
	}
}
func WithAfter(after string) ParamOptions {
	return func(p *Params) {
		p.After = after
	}
}
func WithLimit(limit int64) ParamOptions {
	return func(p *Params) {
		p.Limit = limit
	}
}
func WithFilter(filter string) ParamOptions {
	return func(p *Params) {
		p.Filter = filter
	}
}
func WithExpand(expand string) ParamOptions {
	return func(p *Params) {
		p.Expand = expand
	}
}
func WithIncludeNonDeleted(includeNonDeleted bool) ParamOptions {
	return func(p *Params) {
		b := new(bool)
		*b = includeNonDeleted
		p.IncludeNonDeleted = b
	}
}
func WithActivate(activate bool) ParamOptions {
	return func(p *Params) {
		b := new(bool)
		*b = activate
		p.Activate = b
	}
}
func WithTargetAid(targetAid string) ParamOptions {
	return func(p *Params) {
		p.TargetAid = targetAid
	}
}
func WithQueryScope(query_scope string) ParamOptions {
	return func(p *Params) {
		p.QueryScope = query_scope
	}
}
func WithRemoveUsers(removeUsers bool) ParamOptions {
	return func(p *Params) {
		b := new(bool)
		*b = removeUsers
		p.RemoveUsers = b
	}
}
func WithUntil(until string) ParamOptions {
	return func(p *Params) {
		p.Until = until
	}
}
func WithSince(since string) ParamOptions {
	return func(p *Params) {
		p.Since = since
	}
}
func WithSortOrder(sortOrder string) ParamOptions {
	return func(p *Params) {
		p.SortOrder = sortOrder
	}
}
func WithFormat(format string) ParamOptions {
	return func(p *Params) {
		p.Format = format
	}
}
func WithSearch(search string) ParamOptions {
	return func(p *Params) {
		p.Search = search
	}
}
func WithProvider(provider string) ParamOptions {
	return func(p *Params) {
		p.Provider = provider
	}
}
func WithShowAll(showAll bool) ParamOptions {
	return func(p *Params) {
		b := new(bool)
		*b = showAll
		p.ShowAll = b
	}
}
func WithSendEmail(sendEmail bool) ParamOptions {
	return func(p *Params) {
		b := new(bool)
		*b = sendEmail
		p.SendEmail = b
	}
}
func WithUpdatePhone(updatePhone bool) ParamOptions {
	return func(p *Params) {
		b := new(bool)
		*b = updatePhone
		p.UpdatePhone = b
	}
}
func WithTemplateId(templateId string) ParamOptions {
	return func(p *Params) {
		p.TemplateId = templateId
	}
}
func WithTempPassword(tempPassword bool) ParamOptions {
	return func(p *Params) {
		b := new(bool)
		*b = tempPassword
		p.TempPassword = b
	}
}
func WithOauthTokens(oauthTokens bool) ParamOptions {
	return func(p *Params) {
		b := new(bool)
		*b = oauthTokens
		p.OauthTokens = b
	}
}

func (p *Params) String() string {
	s, _ := json.Marshal(p)

	m := map[string]interface{}{}
	qs := ""
	i := 1

	json.Unmarshal(s, &m)

	for key, value := range m {
		qs = qs + string(key) + "="
		switch v := value.(type) {
		case bool:
			qs = qs + strconv.FormatBool(v)
		default:
			qs = qs + url.QueryEscape(value.(string))
		}
		if i < len(m) {
			qs = qs + "&"
		}
		i++
	}

	if qs != "" {
		qs = "?" + qs
	}
	return qs
}

