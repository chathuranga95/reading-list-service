// Copyright 2025 The OpenChoreo Authors
// SPDX-License-Identifier: Apache-2.0

package config

import (
	"os"
	"path/filepath"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	"reading-list-service/internal/models"
)

func TestLoadConfigUsesInitialDataPathFromEnv(t *testing.T) {
	t.Setenv(initialDataPath, "/tmp/books.json")

	cfg, err := LoadConfig()

	require.NoError(t, err)
	assert.Equal(t, "/tmp/books.json", cfg.InitialDataPath)
}

func TestResolveInitialDataPathUsesDefaultMountedFile(t *testing.T) {
	defaultPath := filepath.Join(t.TempDir(), "initial_data.json")
	require.NoError(t, os.WriteFile(defaultPath, []byte(`{"books":[]}`), 0644))

	path := resolveInitialDataPath("", defaultPath, os.Stat)

	assert.Equal(t, defaultPath, path)
}

func TestLoadInitialDataReadsMountedBookList(t *testing.T) {
	dataPath := filepath.Join(t.TempDir(), "books.json")
	require.NoError(t, os.WriteFile(dataPath, []byte(`{
		"books": [
			{
				"id": "mount-demo-1",
				"title": "Designing Data-Intensive Applications",
				"author": "Martin Kleppmann",
				"status": "reading"
			}
		]
	}`), 0644))
	t.Setenv(initialDataPath, dataPath)
	_, err := LoadConfig()
	require.NoError(t, err)

	initialData := LoadInitialData()

	require.Len(t, initialData.Books, 1)
	assert.Equal(t, "mount-demo-1", initialData.Books[0].Id)
	assert.Equal(t, "Designing Data-Intensive Applications", initialData.Books[0].Title)
	assert.Equal(t, models.ReadStatusReading, initialData.Books[0].Status)
}
