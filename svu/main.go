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
	Container *dagger.Container
}

// New creates a new Svu instance and loads the provided source directory.
func New(
	source *dagger.Directory,
	// +default="3.3.0@sha256:8f35fe6a6e09ee5b66a10e26753f16e31f206c919481f11b598b195e9d34b517"
	version string,
) *Svu {
	ctr := dag.Container().
		From("ghcr.io/caarlos0/svu:"+version).
		WithDirectory("/tmp/src", source)

	return &Svu{
		Container: ctr,
	}
}

func (m *Svu) Next(
	ctx context.Context,
	// +default="."
	directory string,
) (string, error) {
	m.Container.Terminal()
	if version, err := m.Container.
		WithWorkdir(filepath.Join("/tmp/src", directory)).
		WithExec([]string{"svu", "n"}).
		CombinedOutput(ctx); err != nil {
		return "", err
	} else {
		return strings.TrimSpace(version), nil
	}
}
