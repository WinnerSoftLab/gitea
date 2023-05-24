// Copyright 2023 The Gitea Authors. All rights reserved.
// SPDX-License-Identifier: MIT

package git

import (
	"fmt"
)

// GetFileContent returns file content for given revision.
func (repo *Repository) GetFileContent(rev, path string) ([]byte, error) {
	cmd := NewCommand(repo.Ctx, "show")
	cmd.AddDynamicArguments(fmt.Sprintf("%s:%s", rev, path))
	stdout, _, err := cmd.RunStdBytes(&RunOpts{Dir: repo.Path})
	if err != nil {
		return nil, err
	}

	return stdout, nil
}
