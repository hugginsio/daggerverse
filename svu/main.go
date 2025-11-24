// Copyright (c) Kyle Huggins
// SPDX-License-Identifier: BSD-3-Clause

package main

import (
	"context"
	"dagger/svu/internal/dagger"
	"path/filepath"
	"strings"
)

type Svu struct {
	container *dagger.Container
}

// func New(
// 	// +default="3.3.0"
// 	version string,
// ) *Svu {
// 	ctr := dag.Container().From("ghcr.io/caarlos0/svu:" + version)

// 	return &Svu{
// 		container: ctr,
// 	}
// }

func (m *Svu) Next(
	ctx context.Context,
	source *dagger.Directory,
	// +default="."
	directory string,
) (string, error) {
	if version, err := dag.Container().From("ghcr.io/caarlos0/svu:"+"3.3.0").
		WithDirectory("/tmp/src", source).
		WithWorkdir(filepath.Join("/tmp/src", directory)).
		WithExec([]string{"svu", "n"}).
		CombinedOutput(ctx); err != nil {
		return "", err
	} else {
		return strings.TrimSpace(version), nil
	}
}
