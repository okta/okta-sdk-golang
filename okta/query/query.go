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
	"net/url"
	"strconv"
)

type Params struct {
	Q                    string `json:"q,omitempty"`
	After                string `json:"after,omitempty"`
	Limit                int64  `json:"limit,omitempty"`
	Filter               string `json:"filter,omitempty"`
	Expand               string `json:"expand,omitempty"`
	IncludeNonDeleted    *bool  `json:"includeNonDeleted,omitempty"`
	Activate             *bool  `json:"activate,omitempty"`
	TargetAid            string `json:"targetAid,omitempty"`
	QueryScope           string `json:"query_scope,omitempty"`
	RemoveUsers          *bool  `json:"removeUsers,omitempty"`
	ManagedBy            string `json:"managedBy,omitempty"`
	Until                string `json:"until,omitempty"`
	Since                string `json:"since,omitempty"`
	SortOrder            string `json:"sortOrder,omitempty"`
	Format               string `json:"format,omitempty"`
	Search               string `json:"search,omitempty"`
	Provider             string `json:"provider,omitempty"`
	NextLogin            string `json:"nextLogin,omitempty"`
	ShowAll              *bool  `json:"showAll,omitempty"`
	SendEmail            *bool  `json:"sendEmail,omitempty"`
	UpdatePhone          *bool  `json:"updatePhone,omitempty"`
	TemplateId           string `json:"templateId,omitempty"`
	TokenLifetimeSeconds int64  `json:"tokenLifetimeSeconds,omitempty"`
	TempPassword         *bool  `json:"tempPassword,omitempty"`
	OauthTokens          *bool  `json:"oauthTokens,omitempty"`
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
func WithManagedBy(managedBy string) ParamOptions {
	return func(p *Params) {
		p.ManagedBy = managedBy
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
func WithNextLogin(nextLogin string) ParamOptions {
	return func(p *Params) {
		p.NextLogin = nextLogin
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
func WithTokenLifetimeSeconds(tokenLifetimeSeconds int64) ParamOptions {
	return func(p *Params) {
		p.TokenLifetimeSeconds = tokenLifetimeSeconds
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
	qs := url.Values{}

	if p.Q != "" {
		qs.Add(`q`, p.Q)
	}
	if p.After != "" {
		qs.Add(`after`, p.After)
	}
	if p.Limit != 0 {
		qs.Add(`limit`, strconv.FormatInt(p.Limit, 10))
	}
	if p.Filter != "" {
		qs.Add(`filter`, p.Filter)
	}
	if p.Expand != "" {
		qs.Add(`expand`, p.Expand)
	}
	if p.IncludeNonDeleted != nil {
		qs.Add(`includeNonDeleted`, strconv.FormatBool(*p.IncludeNonDeleted))
	}
	if p.Activate != nil {
		qs.Add(`activate`, strconv.FormatBool(*p.Activate))
	}
	if p.TargetAid != "" {
		qs.Add(`targetAid`, p.TargetAid)
	}
	if p.QueryScope != "" {
		qs.Add(`query_scope`, p.QueryScope)
	}
	if p.RemoveUsers != nil {
		qs.Add(`removeUsers`, strconv.FormatBool(*p.RemoveUsers))
	}
	if p.ManagedBy != "" {
		qs.Add(`managedBy`, p.ManagedBy)
	}
	if p.Until != "" {
		qs.Add(`until`, p.Until)
	}
	if p.Since != "" {
		qs.Add(`since`, p.Since)
	}
	if p.SortOrder != "" {
		qs.Add(`sortOrder`, p.SortOrder)
	}
	if p.Format != "" {
		qs.Add(`format`, p.Format)
	}
	if p.Search != "" {
		qs.Add(`search`, p.Search)
	}
	if p.Provider != "" {
		qs.Add(`provider`, p.Provider)
	}
	if p.NextLogin != "" {
		qs.Add(`nextLogin`, p.NextLogin)
	}
	if p.ShowAll != nil {
		qs.Add(`showAll`, strconv.FormatBool(*p.ShowAll))
	}
	if p.SendEmail != nil {
		qs.Add(`sendEmail`, strconv.FormatBool(*p.SendEmail))
	}
	if p.UpdatePhone != nil {
		qs.Add(`updatePhone`, strconv.FormatBool(*p.UpdatePhone))
	}
	if p.TemplateId != "" {
		qs.Add(`templateId`, p.TemplateId)
	}
	if p.TokenLifetimeSeconds != 0 {
		qs.Add(`tokenLifetimeSeconds`, strconv.FormatInt(p.TokenLifetimeSeconds, 10))
	}
	if p.TempPassword != nil {
		qs.Add(`tempPassword`, strconv.FormatBool(*p.TempPassword))
	}
	if p.OauthTokens != nil {
		qs.Add(`oauthTokens`, strconv.FormatBool(*p.OauthTokens))
	}

	if len(qs) != 0 {
		return "?" + qs.Encode()
	}
	return ""
}
