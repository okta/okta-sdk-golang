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

 // THIS FILE IS TEMPORARY AS A PLACEHOLDER FOR USER TO SEE HOW THE REQUEST EXECUTOR WORKS
package okta

import (
	"encoding/json"
	"fmt"
	"time"
)

type UserService service

type User struct {
	ID string `json:"id, omitempty"`
	Status string `json:"status, omitempty"`
	Created time.Time `json:"created, omitempty"`
	Activated time.Time `json:"activated, omitempty"`
	StatusChanged time.Time `json:"statusChanged, omitempty"`
	LastLogin time.Time `json:"lastLogin, omitempty"`
	LastUpdated time.Time `json:"lastUpdated, omitempty"`
	PasswordChanged	time.Time `json:"passwordChanged, omitempty"`
	Profile struct{
		FirstName string `json:"firstName, omitempty"`
		LastName string `json:"lastName, omitempty"`
		Email string `json:"email, omitempty"`
		Login string `json:"login, omitempty"`
		MobilePhone string `json:"mobilePhone, omitempty"`
	} `json:"profile, omitempty"`
}

type UserList struct {
	Users []User
}

func (u *UserService) List() []User {
	resp, err := u.client.requestExecutor.Get("/users")
	if(err != nil) {
		fmt.Printf("error: %v", err)
	}

	ul := make([]User,0)
	//fmt.Printf("%+v", resp)
	json.Unmarshal(resp, &ul)

	return ul
}
