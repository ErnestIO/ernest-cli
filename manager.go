/* This Source Code Form is subject to the terms of the Mozilla Public
 * License, v. 2.0. If a copy of the MPL was not distributed with this
 * file, You can obtain one at http://mozilla.org/MPL/2.0/. */

package main

import (
	"bytes"
	"crypto/tls"
	"encoding/json"
	"errors"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"runtime"

	"github.com/fatih/color"
)

// Manager manages all api communications
type Manager struct {
	URL string `json:"url"`
}

// Service ...
type Service struct {
	ID         string `json:"service_id"`
	Name       string `json:"service_name"`
	Datacenter string `json:"datacenter_id"`
	Version    string `json:"service_version"`
	Status     string `json:"service_status"`
	Definition string `json:"service_definition"`
	Result     string `json:"service_result"`
	Endpoint   string `json:"service_endpoint"`
}

// Datacenter ...
type Datacenter struct {
	ID   string `json:"datacenter_id"`
	Name string `json:"datacenter_name"`
}

// User ...
type User struct {
	ID       string `json:"user_id"`
	Name     string `json:"user_name"`
	Email    string `json:"user_email"`
	ClientID string `json:"client_id"`
	IsAdmin  bool   `json:"user_admin"`
}

// Group ...
type Group struct {
	ID   string `json:"client_id"`
	Name string `json:"client_name"`
}

// Session ...
type Session struct {
	UserID    string `json:"user_id"`
	ClientID  string `json:"client_id"`
	UserName  string `json:"user_name"`
	UserEmail string `json:"user_email"`
}

func (m *Manager) client() *http.Client {
	tr := &http.Transport{
		TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
	}
	client := &http.Client{Transport: tr}

	return client
}

func (m *Manager) doRequest(url, method string, payload []byte, token string, contentType string) (string, *http.Response, error) {
	url = m.URL + url
	req, err := http.NewRequest(method, url, bytes.NewBuffer(payload))
	if token != "" {
		req.Header.Add("X-AUTH-TOKEN", token)
	}
	if contentType != "" {
		req.Header.Set("Content-Type", contentType)
	}
	resp, err := m.client().Do(req)

	if err != nil {
		return err.Error(), resp, err
	}
	defer resp.Body.Close()
	responseBody, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		color.Red(err.Error())
	}
	body := string(responseBody)

	if resp.StatusCode != 200 {
		return string(body), resp, errors.New(resp.Status)
	}
	return string(body), resp, nil
}

func (m *Manager) createClient(token string, name string) (string, error) {
	payload := []byte(`{"client_name":"` + name + `"}`)
	body, _, err := m.doRequest("/clients/", "POST", payload, token, "")
	if err != nil {
		return body, err
	}

	color.Green("SUCCESS: Group " + name + " created")

	var client struct {
		ID string `json:"client_id"`
	}
	err = json.Unmarshal([]byte(body), &client)
	if err != nil {
		return "", errors.New("ERROR: Couldn't read response from server")
	}
	return client.ID, nil
}

func (m *Manager) createUser(token string, client string, user string, password string, email string) error {
	payload := []byte(`{"client_id": "` + client + `", "user_name":"` + user + `", "user_email": "` + email + `", "user_password": "` + password + `"}`)
	_, _, err := m.doRequest("/users/", "POST", payload, token, "")
	if err != nil {
		return err
	}
	color.Green("SUCCESS: User " + user + " created")
	return nil
}

func (m *Manager) getUser(token string, userid string) (user User, err error) {
	res, _, err := m.doRequest("/users/"+userid, "GET", nil, token, "application/yaml")
	json.Unmarshal([]byte(res), &user)
	return user, err
}

func (m *Manager) deleteUser(token string, user string) error {
	_, _, err := m.doRequest("/users/"+user, "DELETE", nil, token, "application/yaml")
	return err
}

func (m *Manager) getSession(token string) (session Session, err error) {
	res, _, err := m.doRequest("/session/", "GET", nil, token, "application/yaml")
	json.Unmarshal([]byte(res), &session)
	return session, err
}

// ********************* Login & Logout *******************

// Login does a login action against the api
func (m *Manager) Login(username string, password string) (token string, body string, err error) {
	payload := []byte(`{"user_name":"` + username + `", "user_password": "` + password + `"}`)

	body, res, err := m.doRequest("/session/", "POST", payload, "", "")
	if err != nil {
		return body, "", err
	}

	token = res.Header.Get("X-Auth-Token")

	return token, username, nil
}

// Logout clear local authentication credentials
func (m *Manager) Logout(token string) error {
	_, _, err := m.doRequest("/session/", "DELETE", nil, token, "")
	return err
}

// ********************* Update *******************

// ChangePassword ...
func (m *Manager) ChangePassword(token string, userid string, oldpassword string, newpassword string) error {
	payload := []byte(`{"old_password":"` + oldpassword + `", "new_password": "` + newpassword + `"}`)
	_, _, err := m.doRequest("/users/"+userid, "PUT", payload, token, "application/yaml")
	if err != nil {
		return err
	}
	return nil
}

// ChangePasswordByAdmin ...
func (m *Manager) ChangePasswordByAdmin(token string, userid string, newpassword string) error {
	payload := []byte(`{"new_password": "` + newpassword + `"}`)
	_, _, err := m.doRequest("/users/"+userid, "PUT", payload, token, "application/yaml")
	if err != nil {
		return err
	}
	return nil
}

// ********************* Create *******************

// CreateDatacenter ...
func (m *Manager) CreateDatacenter(token string, name string, user string, password string, url string, network string, vseURL string) (string, error) {
	payload := []byte(`{"datacenter_name": "` + name + `", "datacenter_type": "vcloud", "datacenter_region": "LON-001", "datacenter_username":"` + user + `", "datacenter_password":"` + password + `", "external_network":"` + network + `", "vcloud_url":"` + url + `", "vse_url":"` + vseURL + `"}`)
	body, _, err := m.doRequest("/datacenters/", "POST", payload, token, "")
	if err != nil {
		return body, err
	}
	color.Green("SUCCESS: Datacenter " + name + " created")
	return body, err
}

// CreateUser ...
func (m *Manager) CreateUser(name string, email string, user string, password string, adminuser string, adminpassword string) error {
	token, _, err := m.Login(adminuser, adminpassword)
	c, err := m.createClient(token, name)
	if err != nil {
		color.Red(err.Error() + ": Group " + name + " already exists")
		os.Exit(1)
	}
	res := m.createUser(token, c, user, password, email)
	return res
}

// ********************* Get *******************

// GetUser ...
func (m *Manager) GetUser(token string, userid string) (user User, err error) {
	res, _, err := m.doRequest("/users/"+userid, "GET", nil, token, "application/yaml")
	json.Unmarshal([]byte(res), &user)
	return user, err
}

// GetUUID ...
func (m *Manager) GetUUID(token string, payload []byte) string {
	id, err := buildServiceUUID(payload)
	if err != nil {
		log.Fatal(err)
	}
	body, _, err := m.doRequest("/services/uuid/", "POST", []byte(`{"id":"`+id+`"}`), token, "")
	var dat map[string]interface{}
	json.Unmarshal([]byte(body), &dat)

	if str, ok := dat["uuid"].(string); ok {
		return str
	}
	return ""
}

// ********************* Apply *******************

// Apply ...
func (m *Manager) Apply(token string, path string, monit bool) (string, error) {
	payload, err := ioutil.ReadFile(path)
	if err != nil {
		color.Red(err.Error())
		return "", nil
	}

	color.Green("Environment creation requested")
	println("Ernest will show you all output from your requested service creation")
	println("You can cancel at any moment with Ctrl+C, even the service is still being created, you won't have any output")

	streamID := m.GetUUID(token, payload)
	if streamID == "" {
		color.Red("Please log in")
		return "", nil
	}

	if monit == true {
		go Monitorize(m.URL, token, streamID)
	} else {
		println("Additionally you can trace your service on ernest monitor tool with id: " + streamID)
	}

	if body, _, err := m.doRequest("/services/", "POST", payload, token, "application/yaml"); err != nil {
		return "", errors.New(body)
	}
	if monit == true {
		runtime.Goexit()
	}
	return streamID, nil
}

// ********************* Destroy *******************

// Destroy ...
func (m *Manager) Destroy(token string, name string, monit bool) error {
	body, _, err := m.doRequest("/services/"+name, "DELETE", nil, token, "application/yaml")
	if err != nil {
		return err
	}

	var res map[string]interface{}
	json.Unmarshal([]byte(body), &res)

	if monit == true {
		if str, ok := res["stream_id"].(string); ok {
			Monitorize(m.URL, token, str)
			runtime.Goexit()
		}
	}

	return err
}

// ********************* Reset *******************

// ResetService ...
func (m *Manager) ResetService(name string, token string) error {
	_, _, err := m.doRequest("/services/"+name+"/reset/", "POST", nil, token, "application/yaml")
	return err
}

// ********************* Status *******************

// ServiceStatus ...
func (m *Manager) ServiceStatus(token string, serviceName string) (service Service, err error) {
	body, _, err := m.doRequest("/services/"+serviceName, "GET", []byte(""), token, "")
	if err != nil {
		return service, err
	}
	json.Unmarshal([]byte(body), &service)
	return service, err
}

// ServiceBuildStatus ...
func (m *Manager) ServiceBuildStatus(token string, serviceName string, serviceID string) (service Service, err error) {
	body, _, err := m.doRequest("/services/"+serviceName+"/builds/"+serviceID, "GET", []byte(""), token, "")
	if err != nil {
		return service, err
	}
	json.Unmarshal([]byte(body), &service)
	return service, err
}

// ********************* List *********************

// ListDatacenters ...
func (m *Manager) ListDatacenters(token string) (datacenters []Datacenter, err error) {
	body, _, err := m.doRequest("/datacenters/", "GET", []byte(""), token, "")
	if err != nil {
		return nil, err
	}
	json.Unmarshal([]byte(body), &datacenters)
	return datacenters, err
}

// ListServices ...
func (m *Manager) ListServices(token string) (services []Service, err error) {
	body, _, err := m.doRequest("/services/", "GET", []byte(""), token, "")
	if err != nil {
		return nil, err
	}
	json.Unmarshal([]byte(body), &services)
	return services, err
}

// ListBuilds ...
func (m *Manager) ListBuilds(name string, token string) (builds []Service, err error) {
	body, _, err := m.doRequest("/services/"+name+"/builds/", "GET", []byte(""), token, "")
	if err != nil {
		return nil, err
	}
	json.Unmarshal([]byte(body), &builds)
	return builds, err
}

// ListUsers ...
func (m *Manager) ListUsers(token string) (users []User, err error) {
	body, _, err := m.doRequest("/users/", "GET", []byte(""), token, "")
	if err != nil {
		return nil, err
	}
	json.Unmarshal([]byte(body), &users)
	return users, err
}

// ListGroups ...
func (m *Manager) ListGroups(token string) (groups []Group, err error) {
	body, _, err := m.doRequest("/clients/", "GET", []byte(""), token, "")
	if err != nil {
		return nil, err
	}
	json.Unmarshal([]byte(body), &groups)
	return groups, err
}
