package hexcoresdk

import (
	"context"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/require"
	openapiclient "github.com/CE-RISE-software/hex-core-sdk-go"
)

func TestSmoke_AdminHealth(t *testing.T) {
	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		require.Equal(t, "/admin/health", r.URL.Path)
		w.Header().Set("Content-Type", "application/json")
		_, _ = w.Write([]byte(`{"status":"ok"}`))
	}))
	defer srv.Close()

	cfg := openapiclient.NewConfiguration()
	cfg.Servers = openapiclient.ServerConfigurations{{URL: srv.URL}}
	client := openapiclient.NewAPIClient(cfg)

	resp, httpRes, err := client.AdminAPI.Health(context.Background()).Execute()
	require.NoError(t, err)
	require.NotNil(t, httpRes)
	require.Equal(t, 200, httpRes.StatusCode)
	require.NotNil(t, resp)
}

func TestSmoke_DiscoveryListModels(t *testing.T) {
	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		require.Equal(t, "/models", r.URL.Path)
		w.Header().Set("Content-Type", "application/json")
		_, _ = w.Write([]byte(`{"models":[{"id":"demo-model","version":"0.0.1"}]}`))
	}))
	defer srv.Close()

	cfg := openapiclient.NewConfiguration()
	cfg.Servers = openapiclient.ServerConfigurations{{URL: srv.URL}}
	client := openapiclient.NewAPIClient(cfg)

	resp, httpRes, err := client.DiscoveryAPI.ListModels(context.Background()).Execute()
	require.NoError(t, err)
	require.NotNil(t, httpRes)
	require.Equal(t, 200, httpRes.StatusCode)
	require.NotNil(t, resp)
}
