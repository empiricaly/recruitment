package metrics

import (
	"bytes"
	"fmt"
	"io"
	"net/http"
	"strings"

	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	"github.com/prometheus/common/expfmt"
	"github.com/rs/zerolog/log"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

// Logger is the expected logger for the manual print out of the metrics
type Logger func(format string, v ...interface{})

// Config configures the metrics client.
type Config struct {
	Namespace string `mapstructure:"namespace"`
	Enable    bool   `mapstructure:"enable"`
	Addr      string `mapstructure:"addr"`
	Path      string `mapstructure:"path"`
}

// ConfigFlags helps configure cobra and viper flags.
func ConfigFlags(cmd *cobra.Command, prefix, defaultNamespace, defaultHTTP, defaultPath string) error {
	if prefix == "" {
		prefix = "metrics"
	}

	viper.SetDefault(prefix, &Config{})

	if defaultHTTP == "" {
		defaultHTTP = ":9900"
	}

	if defaultPath == "" {
		defaultPath = "/debug/metrics"
	}

	flag := prefix + ".enable"
	cmd.Flags().Bool(flag, false, "Metrics enabled (Prometheus format)")
	viper.SetDefault(flag, false)

	flag = prefix + ".namespace"
	cmd.Flags().String(flag, defaultNamespace, "Prometheus metrics namespace")
	viper.SetDefault(flag, defaultNamespace)

	flag = prefix + ".addr"
	cmd.Flags().String(flag, defaultHTTP, "Addr to serve HTTP metrics on")
	viper.SetDefault(flag, defaultHTTP)

	flag = prefix + ".path"
	cmd.Flags().String(flag, defaultPath, "Path to serve HTTP metrics on")
	viper.SetDefault(flag, defaultPath)

	return nil
}

// Validate configuration is ok
func (c *Config) Validate() error {
	return nil
}

// Metrics enables app metrics
type Metrics struct {
	config      *Config
	counters    map[string]prometheus.Counter
	counterVecs map[string]*prometheus.CounterVec
	gauges      map[string]prometheus.Gauge
	gaugeVecs   map[string]*prometheus.GaugeVec
}

// New initializes the configured Metrics
func New(config *Config) (*Metrics, error) {
	m := &Metrics{
		config:      config,
		counters:    make(map[string]prometheus.Counter),
		counterVecs: make(map[string]*prometheus.CounterVec),
		gauges:      make(map[string]prometheus.Gauge),
		gaugeVecs:   make(map[string]*prometheus.GaugeVec),
	}

	if config.Enable {
		go func() {
			http.Handle("/metrics", promhttp.Handler())
			err := http.ListenAndServe(config.Addr, http.DefaultServeMux)
			log.Error().Err(err).Msg("could not start metrics HTTP server")
		}()
	}

	return m, nil
}

// Counter gets or creates an expvar Float with the given name, and
// returns an object that implements the Counter interface.
func (m *Metrics) Counter(name, help string) prometheus.Counter {
	if counter, ok := m.counters[name]; !ok {
		counter = prometheus.NewCounter(
			prometheus.CounterOpts{
				Namespace: m.config.Namespace,
				Name:      name,
				Help:      help,
			})

		prometheus.MustRegister(counter)
		return counter
	} else {
		return counter
	}
}

// CounterVec creates a new CounterVec based on the provided CounterOpts and
// partitioned by the given label names.
func (m *Metrics) CounterVec(name, help string, labels ...string) *prometheus.CounterVec {
	if counter, ok := m.counterVecs[name]; !ok {
		counter := prometheus.NewCounterVec(
			prometheus.CounterOpts{
				Namespace: m.config.Namespace,
				Name:      name,
				Help:      help,
			}, labels)

		prometheus.MustRegister(counter)
		return counter
	} else {
		return counter
	}
}

// Gauge get or creates an expvar Float with the given name, and returns
// an object that implements the Gauge interface.
func (m *Metrics) Gauge(name, help string) prometheus.Gauge {
	if gauge, ok := m.gauges[name]; !ok {
		gauge := prometheus.NewGauge(
			prometheus.GaugeOpts{
				Namespace: m.config.Namespace,
				Name:      name,
				Help:      help,
			})

		prometheus.MustRegister(gauge)
		return gauge
	} else {
		return gauge
	}
}

// GaugeVec creates a new GaugeVec based on the provided GaugeOpts and
// partitioned by the given label names.
func (m *Metrics) GaugeVec(name, help string, labels ...string) *prometheus.GaugeVec {
	if gauge, ok := m.gaugeVecs[name]; !ok {
		gauge := prometheus.NewGaugeVec(
			prometheus.GaugeOpts{
				Namespace: m.config.Namespace,
				Name:      name,
				Help:      help,
			}, labels)

		prometheus.MustRegister(gauge)
		return gauge
	} else {
		return gauge
	}
}

// Histogram counts individual observations from an event or sample stream in
// configurable buckets. Similar to a summary, it also provides a sum of
// observations and an observation count.
//
// On the Prometheus server, quantiles can be calculated from a Histogram using
// the histogram_quantile function in the query language.
//
// Note that Histograms, in contrast to Summaries, can be aggregated with the
// Prometheus query language (see the documentation for detailed procedures).
// However, Histograms require the user to pre-define suitable buckets, and
// they are in general less accurate. The Observe method of a Histogram has a
// very low performance overhead in comparison with the Observe method of a
// Summary.
func (m *Metrics) Histogram(name, help string, buckets []float64) prometheus.Histogram {
	histogram := prometheus.NewHistogram(prometheus.HistogramOpts{
		Namespace: m.config.Namespace,
		Name:      name,
		Help:      help,
		Buckets:   buckets,
	})

	prometheus.MustRegister(histogram)

	return histogram
}

// HistogramVec is a Collector that bundles a set of Histograms that all share
// the same Desc, but have different values for their variable labels. This is
// used if you want to count the same thing partitioned by various dimensions
// (e.g. HTTP request latencies, partitioned by status code and method). Create
// instances with NewHistogramVec.
func (m *Metrics) HistogramVec(name, help string, buckets []float64, labels ...string) *prometheus.HistogramVec {
	histogram := prometheus.NewHistogramVec(prometheus.HistogramOpts{
		Namespace: m.config.Namespace,
		Name:      name,
		Help:      help,
		Buckets:   buckets,
	}, labels)

	prometheus.MustRegister(histogram)

	return histogram
}

type writer struct {
	l   Logger
	all bool
}

func (w writer) Write(b []byte) (int, error) {
	r := bytes.NewBuffer(b)
	for {
		line, err := r.ReadString('\n')
		if err != nil {
			if err == io.EOF {
				return len(b), nil
			}
			return 0, err
		}

		if !w.all {
			if line[0] == '#' || strings.HasPrefix(line, "go_") || strings.HasPrefix(line, "promhttp_") || line == "" {
				continue
			}
		}

		w.l("%s", strings.TrimSpace(line))
	}
}

// Log a human readable output of the metrics. This is not optimized, only to be
// used for debugging.
func (m *Metrics) Log(l Logger, all bool) {
	if !m.config.Enable {
		return
	}

	w := writer{l, all}

	enc := expfmt.NewEncoder(w, expfmt.FmtText)

	mfs, err := prometheus.DefaultGatherer.Gather()
	if err != nil {
		fmt.Println("error gathering metrics:", err)
		return
	}

	for _, mf := range mfs {
		if err := enc.Encode(mf); err != nil {
			fmt.Println("error encoding and sending metric family:", err)
			return
		}
	}
}
