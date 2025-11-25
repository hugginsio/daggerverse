// Copyright (c) Kyle Huggins
// SPDX-License-Identifier: BSD-3-Clause

package main

import (
	"context"
	"fmt"

	"golang.org/x/sync/errgroup"
)

type Tests struct{}

func (m *Tests) All(ctx context.Context) error {
	eg, ectx := errgroup.WithContext(ctx)

	eg.Go(func() error {
		return m.NextDefaultDir(ectx)
	})

	return eg.Wait()
}

func (m *Tests) NextDefaultDir(ctx context.Context) error {
	currentDirectory := dag.CurrentModule().Source().Directory("./testdata")

	svu := dag.Svu(currentDirectory)

	if svu == nil {
		return fmt.Errorf("failed to create svu")
	}

	res, err := svu.Next(ctx)

	if err != nil {
		return fmt.Errorf("failed to get next: %w", err)
	}

	if res != "v0.1.0" {
		return fmt.Errorf("invalid semantic version: %s", res)
	}

	return nil
}
