package healthchecker_test

import (
	"errors"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/require"

	"github.com/begmaroman/go-micro-boilerplate/utils/healthchecker"
)

var (
	successChecker = func() error {
		return nil
	}

	errorChecker = func() error {
		return errors.New("something went wrong")
	}
)

func TestSuccess(t *testing.T) {
	t.Run("check callback succeeds", func(t *testing.T) {
		server := httptest.NewServer(healthchecker.Handler(successChecker))
		defer server.Close()

		requireStatus(t, server.URL, http.StatusOK)
	})
}

func TestFailure(t *testing.T) {
	t.Run("check callback is nil", func(t *testing.T) {
		server := httptest.NewServer(healthchecker.Handler(nil))
		defer server.Close()

		requireStatus(t, server.URL, http.StatusInternalServerError)
	})

	t.Run("check callback errors", func(t *testing.T) {
		server := httptest.NewServer(healthchecker.Handler(errorChecker))
		defer server.Close()

		requireStatus(t, server.URL, http.StatusInternalServerError)
	})
}

func requireStatus(t require.TestingT, url string, expected int) {
	res, err := http.Get(url + "/health")
	require.NoError(t, err)
	require.Equal(t, expected, res.StatusCode)
}
