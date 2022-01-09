// go test -tags=proxytest ./...

//go:build proxytest

package main

import (
	"testing"

	"github.com/stretchr/testify/require"

	"github.com/tetratelabs/proxy-wasm-go-sdk/proxywasm/proxytest"
	"github.com/tetratelabs/proxy-wasm-go-sdk/proxywasm/types"
)

func TestHttpHeaders_OnHttpRequestHeaders(t *testing.T) {
	opt := proxytest.NewEmulatorOption().WithVMContext(&vmContext{})
	host, reset := proxytest.NewHostEmulator(opt)
	defer reset()

	// Initialize http context.
	id := host.InitializeHttpContext()

	// Call OnHttpRequestHeaders.
	action := host.CallOnRequestHeaders(id, [][2]string{}, false)
	require.Equal(t, types.ActionContinue, action)

	found := false
	requestHeaders := host.GetCurrentRequestHeaders(id)
	for _, val := range requestHeaders {
		if val[0] == headerName {
			require.Equal(t, val[1], headerValue)
			found = true
		}
	}
	require.True(t, found, "header not found")

	// Call OnHttpStreamDone.
	host.CompleteHttpContext(id)

}
