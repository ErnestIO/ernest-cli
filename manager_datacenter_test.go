/* This Source Code Form is subject to the terms of the Mozilla Public
 * License, v. 2.0. If a copy of the MPL was not distributed with this
 * file, You can obtain one at http://mozilla.org/MPL/2.0/. */

package main

import (
	"testing"

	"github.com/smartystreets/goconvey/convey"
)

func TestListDatacenters(t *testing.T) {
	convey.Convey("Given I get all datacenters", t, func() {
		server := mockRequest("/datacenters/", "GET", 200, ``)
		m := Manager{URL: server.URL}
		_, err := m.ListDatacenters("token")
		convey.Convey("Then It does not fail", func() {
			convey.So(err, convey.ShouldBeNil)
		})
	})
}
