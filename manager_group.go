/* This Source Code Form is subject to the terms of the Mozilla Public
 * License, v. 2.0. If a copy of the MPL was not distributed with this
 * file, You can obtain one at http://mozilla.org/MPL/2.0/. */

package main

import (
	"encoding/json"
	"fmt"

	"github.com/fatih/color"
)

// Group ...
type Group struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
}

// CreateGroup ...
func (m *Manager) CreateGroup(token string, group string) error {
	payload := []byte(`{"name": "` + group + `"}`)
	_, _, err := m.doRequest("/api/groups/", "POST", payload, token, "")
	if err != nil {
		fmt.Println(err)
		return err
	}
	color.Green("SUCCESS: Group " + group + " created")
	return nil
}

// ListGroups ...
func (m *Manager) ListGroups(token string) (groups []Group, err error) {
	body, _, err := m.doRequest("/api/groups/", "GET", []byte(""), token, "")
	if err != nil {
		return nil, err
	}
	err = json.Unmarshal([]byte(body), &groups)
	if err != nil {
		return nil, err
	}
	return groups, err
}

// GroupAddUser ...
func (m *Manager) GroupAddUser(token string, userid string, groupid string) error {
	payload := []byte(`{"userid": "` + userid + `", "groupid": "` + groupid + `"}`)
	_, _, err := m.doRequest("/api/groups/"+groupid+"/users/", "POST", payload, token, "")
	if err != nil {
		fmt.Println(err)
		return err
	}
	color.Green("SUCCESS: Added user " + userid + " to group " + groupid)
	return nil
}