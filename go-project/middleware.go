package main

/*
// Adapter wraps an http.Handler with additional
// functionality.
type Adapter func(http.Handler) http.Handler

func Logging(l *log.Logger) Adapter {
	return func(h http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			l.Println(r.Method, r.URL.Path)
			h.ServeHTTP(w, r)
		})
	}
}

// Adapt h with all specified adapters.
func Adapt(h http.Handler, adapters ...Adapter) http.Handler {
	for _, adapter := range adapters {
		h = adapter(h)
	}
	return h
}

// WithHeader is an Adapter that sets an HTTP handler.
func WithHeader(key, value string) Adapter {
	return func(h http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			w.Header.Add(key, value)
			h.ServeHTTP(w, r)
		})
	}
}

// SupportXHTTPMethodOverride adds support for the X-HTTP-Method-Override
// header.
func SupportXHTTPMethodOverride() Adapter {
	return func(h http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			m := r.Header.Get("X-HTTP-Method-Override")
			if len(m) > 0 {
				r.Method = m
			}
			h.ServeHTTP(w, r)
		})
	}
}

func main() {

	router := mux.NewRouter()

	// adapt a single route
	router.Handle("/", Adapt(myHandler, WithHeader("X-Something", "Specific")))

	// adapt all handlers
	http.Handle("/", Adapt(router,
		SupportXHTTPMethodOverride(),
		WithHeader("Server", "MyApp v1"),
		Logging(logger),
	))
}
*/
