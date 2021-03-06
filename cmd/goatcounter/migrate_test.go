// Copyright © 2019 Martin Tournoij – This file is part of GoatCounter and
// published under the terms of a slightly modified EUPL v1.2 license, which can
// be found in the LICENSE file or at https://license.goatcounter.com

package main

import (
	"testing"
)

func TestMigrate(t *testing.T) {
	_, dbc, clean := tmpdb(t)
	defer clean()

	run(t, 0, []string{"migrate",
		"-db", dbc})
}
