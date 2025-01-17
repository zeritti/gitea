// Copyright 2020 The Gitea Authors. All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package project

import (
	"path/filepath"
	"testing"

	"code.gitea.io/gitea/models/unittest"

	_ "code.gitea.io/gitea/models/repo"
)

func TestMain(m *testing.M) {
	unittest.MainTest(m, filepath.Join("..", ".."),
		"project.yml",
		"project_board.yml",
		"project_issue.yml",
		"repository.yml",
	)
}
