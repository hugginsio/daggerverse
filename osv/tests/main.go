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
		return m.New(ectx)
	})

	return eg.Wait()
}

func (m *Tests) New(ctx context.Context) error {
	osv := dag.Osv()

	if osv == nil {
		return fmt.Errorf("failed to create")
	}

	return nil
}
