/*
Copyright The Kubernetes Authors.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package httpserver

import (
	"context"
	"testing"
	"time"
)

func TestNewServerDefaults(t *testing.T) {
	srv := New(context.Background(), nil)

	// The idle timeout must stay above kube-apiserver's 90s client idle
	// timeout so the client closes idle connections first; if both ends
	// raced to close, admission requests intermittently failed with EOF.
	if got, want := srv.IdleTimeout, 120*time.Second; got != want {
		t.Errorf("IdleTimeout = %v, want %v", got, want)
	}
	if got, want := srv.ReadHeaderTimeout, 32*time.Second; got != want {
		t.Errorf("ReadHeaderTimeout = %v, want %v", got, want)
	}
	if srv.MaxHeaderBytes != 1<<20 {
		t.Errorf("MaxHeaderBytes = %v, want %v", srv.MaxHeaderBytes, 1<<20)
	}
}
