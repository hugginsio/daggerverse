// Copyright (c) Kyle Huggins
// SPDX-License-Identifier: BSD-3-Clause

package main

import (
	"dagger/osv/internal/dagger"
)

type Osv struct {
	// +private
	Container *dagger.Container
}

func New(
	// The image (and repository) to use for the base container.
	//
	// +default="ghcr.io/google/osv-scanner"
	image string,
	// The image version to use for the base container.
	//
	// +default="v2.3.0@sha256:6421cdd773a54fad5fd991895b4f3840beba55a2f30a8182185919e13979a2c8"
	version string,
) *Osv {
	ctr := dag.Container().From(image + ":" + version)

	return &Osv{
		Container: ctr,
	}
}
