// Copyright Tharsis Labs Ltd.(Black)
// SPDX-License-Identifier:ENCL-1.0(https://github.com/fury-labs/black/blob/main/LICENSE)

package config

func EnableObservability() error {
	return nil
	// if true {
	// Temporarily disabling this until we can configure out port reuse
	// fast enough or enabling observability through the config.
	// Please see https://github.com/fury-labs/black/v9/issues/84
	// return nil
	// }

	// pe, err := prometheus.NewExporter(prometheus.Options{
	// 	Namespace: "black",
	// })
	// if err != nil {
	// 	return fmt.Errorf("cmd/config: failed to create the OpenCensus Prometheus exporter: %w", err)
	// }

	// views := app.ObservabilityViews()
	// if err := view.Register(views...); err != nil {
	// 	return fmt.Errorf("cmd/config: failed to register OpenCensus views: %w", err)
	// }
	// view.RegisterExporter(pe)

	// mux := http.NewServeMux()
	// mux.Handle("/metrics", pe)

	// // TODO: Derive the Prometheus observability exporter from the Black config.
	// addr := ":8877"
	// go func() {
	// 	println("Serving the Prometheus observability exporter at", addr)
	// 	if err := http.ListenAndServe(addr, mux); err != nil {
	// 		panic(err)
	// 	}
	// }()

	// return nil
}
