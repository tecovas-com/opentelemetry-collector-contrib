// Copyright The OpenTelemetry Authors
// SPDX-License-Identifier: Apache-2.0

package googlecloudstorageexporter

import (
	"testing"

	"github.com/stretchr/testify/require"
	"go.opentelemetry.io/collector/component"
	"go.opentelemetry.io/collector/exporter/exportertest"

	"github.com/open-telemetry/opentelemetry-collector-contrib/exporter/googlecloudstorageexporter/internal/metadata"
)

func TestCreateExporters(t *testing.T) {
	factory := NewFactory()
	cfg := factory.CreateDefaultConfig().(*Config)
	cfg.Bucket.Name = "bucket"
	cfg.Bucket.Region = "region"
	cfg.Bucket.ProjectID = "project"
	set := exportertest.NewNopSettings(metadata.Type)

	logs, err := factory.CreateLogs(t.Context(), set, cfg)
	require.NoError(t, err)
	require.NotNil(t, logs)

	traces, err := factory.CreateTraces(t.Context(), set, cfg)
	require.NoError(t, err)
	require.NotNil(t, traces)

	metrics, err := factory.CreateMetrics(t.Context(), set, cfg)
	require.NoError(t, err)
	require.NotNil(t, metrics)
	require.False(t, metrics.Capabilities().MutatesData)
	require.Equal(t, component.StabilityLevelDevelopment, factory.MetricsStability())
}
