// Copyright The OpenTelemetry Authors
// SPDX-License-Identifier: Apache-2.0

package confighttp

import (
	"context"
	"net/http"

	"go.opentelemetry.io/collector/component/componenttest"
)

func ExampleServerConfigTcp() {
	settings := NewDefaultServerConfig()
	settings.Endpoint = "localhost:443"
	settings.Network = ""

	s, err := settings.ToServer(
		context.Background(),
		componenttest.NewNopHost(),
		componenttest.NewNopTelemetrySettings(),
		http.HandlerFunc(func(http.ResponseWriter, *http.Request) {}))
	if err != nil {
		panic(err)
	}

	l, err := settings.ToListener(context.Background())
	if err != nil {
		panic(err)
	}
	if err = s.Serve(l); err != nil {
		panic(err)
	}
}

func ExampleServerConfigUnix() {
	settings := NewDefaultServerConfig()
	settings.Endpoint = "unix://something.sock"
	settings.Network = "unix"

	s, err := settings.ToServer(
		context.Background(),
		componenttest.NewNopHost(),
		componenttest.NewNopTelemetrySettings(),
		http.HandlerFunc(func(http.ResponseWriter, *http.Request) {}))
	if err != nil {
		panic(err)
	}

	l, err := settings.ToListener(context.Background())
	if err != nil {
		panic(err)
	}
	if err = s.Serve(l); err != nil {
		panic(err)
	}
}
