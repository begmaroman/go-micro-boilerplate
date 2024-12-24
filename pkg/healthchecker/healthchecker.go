package healthchecker

import (
	"net/http"

	"github.com/sirupsen/logrus"
)

const (
	defaultListenAddr = ":5678"
)

// Check is the type of the healthcheck function accepted by Run
type Check func() error

// Options defines the options for the healthchecker
type Options struct {
	// If set, this will be used as the listen address for the HTTP server. Defaults to :5678 otherwise
	ListenAddr string
}

// Handler turns a Check function into an http.Handler for the GET method at path "/health".
//
// The status code of the healthcheck is determined by the return value of the given check function:
// - 500 if check is nil
// - 500 if check returns a non-nil error
// - 200 otherwise
func Handler(check Check) http.Handler {
	mux := http.NewServeMux()
	mux.HandleFunc("/health", func(rw http.ResponseWriter, req *http.Request) {
		if req.Method != http.MethodGet {
			http.Error(rw, "not found", http.StatusNotFound)
			return
		}

		if check == nil {
			http.Error(rw, "health check callback is nil", http.StatusInternalServerError)
			return
		}

		if err := check(); err != nil {
			http.Error(rw, err.Error(), http.StatusInternalServerError)
			return
		}

		rw.WriteHeader(http.StatusOK)
	})

	return mux
}

// Run starts a new HTTP server that listens on opts.ListenAddr (default :5678) and responds to GET /health
// using the http.Handler returned by Handler(check).
//
// Run will terminate the program if there is any error starting up the HTTP server. Run returns a callback
// that can be used to shutdown the HTTP server.
func Run(log logrus.FieldLogger, check Check, opts *Options) (shutdown func() error) {
	if opts == nil {
		opts = &Options{}
	}

	if len(opts.ListenAddr) == 0 {
		opts.ListenAddr = defaultListenAddr
	}

	server := &http.Server{
		Addr:    opts.ListenAddr,
		Handler: Handler(check),
	}

	go func() {
		if err := server.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.WithError(err).Fatal()
		}
	}()

	return server.Close
}
