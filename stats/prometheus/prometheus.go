package prometheus

import (
	"fmt"

	"github.com/prometheus/client_golang/prometheus"
	"github.com/raymondwongso/gogox/stats"
)

// Stats define prometheus stats.Stats implementor
type Stats struct {
	namespace string
	baseTags  stats.Tags

	counterVecMap   counterVecMap
	gaugeVecMap     gaugeVecMap
	histogramVecMap histogramVecMap
}

type promCounter struct {
	vec       *prometheus.CounterVec
	labelKeys []string
}

type promHistogram struct {
	vec       *prometheus.HistogramVec
	labelKeys []string
}

type promGauge struct {
	vec       *prometheus.GaugeVec
	labelKeys []string
}

type counterVecMap map[string]*promCounter
type gaugeVecMap map[string]*promGauge
type histogramVecMap map[string]*promHistogram

// New creates new prometheus stats
func New(namespace string, baseTags stats.Tags) *Stats {
	return &Stats{
		namespace:       namespace,
		baseTags:        baseTags,
		counterVecMap:   counterVecMap{},
		gaugeVecMap:     gaugeVecMap{},
		histogramVecMap: histogramVecMap{},
	}
}

// metricKey generate unique metric key for cache mapping.
func metricKey(metric string, tags stats.Tags) string {
	return fmt.Sprintf("%s:%s", metric, tags)
}

// tagKeys generate array of keys from tags
func tagKeys(tags stats.Tags) []string {
	res := []string{}
	for k := range tags {
		res = append(res, k)
	}
	return res
}

// tagValues extract tag values from sorted keys
func tagValues(tags stats.Tags, tagKeys []string) []string {
	res := []string{}
	for _, tk := range tagKeys {
		v := tags[tk]
		res = append(res, v)
	}
	return res
}

// Increment implements Increment stats interface
func (s *Stats) Increment(metric string, opt stats.Option) error {
	mk := metricKey(metric, opt.Tags)

	if s.counterVecMap[mk] == nil {
		tk := tagKeys(opt.Tags)

		s.counterVecMap[mk] = &promCounter{
			vec: prometheus.NewCounterVec(prometheus.CounterOpts{
				Namespace:   s.namespace,
				Name:        metric,
				ConstLabels: prometheus.Labels(s.baseTags),
			}, tk),
			labelKeys: tk,
		}

		prometheus.MustRegister(s.counterVecMap[mk].vec)
	}

	c := s.counterVecMap[mk]
	c.vec.WithLabelValues(tagValues(opt.Tags, c.labelKeys)...).Inc()
	return nil
}

// Histogram implements Histogram stats interface
func (s *Stats) Histogram(metric string, value float64, opt stats.Option) error {
	mk := metricKey(metric, opt.Tags)

	if s.histogramVecMap[mk] == nil {
		tk := tagKeys(opt.Tags)

		s.histogramVecMap[mk] = &promHistogram{
			vec: prometheus.NewHistogramVec(prometheus.HistogramOpts{
				Namespace:   s.namespace,
				Name:        metric,
				ConstLabels: prometheus.Labels(s.baseTags),
			}, tk),
			labelKeys: tk,
		}

		prometheus.MustRegister(s.histogramVecMap[mk].vec)
	}

	h := s.histogramVecMap[mk]
	h.vec.WithLabelValues(tagValues(opt.Tags, h.labelKeys)...).Observe(value)
	return nil
}

// Gauge implements Gauge stats interface
func (s *Stats) Gauge(metric string, value float64, opt stats.Option) error {
	mk := metricKey(metric, opt.Tags)

	if s.gaugeVecMap[mk] == nil {
		tk := tagKeys(opt.Tags)

		s.gaugeVecMap[mk] = &promGauge{
			vec: prometheus.NewGaugeVec(prometheus.GaugeOpts{
				Namespace:   s.namespace,
				Name:        metric,
				ConstLabels: prometheus.Labels(s.baseTags),
			}, tk),
			labelKeys: tk,
		}

		prometheus.MustRegister(s.gaugeVecMap[mk].vec)
	}

	g := s.gaugeVecMap[mk]
	g.vec.WithLabelValues(tagValues(opt.Tags, g.labelKeys)...).Set(value)
	return nil
}

// Add implements Add stats interface
func (s *Stats) Add(metric string, value float64, opt stats.Option) error {
	mk := metricKey(metric, opt.Tags)

	if s.counterVecMap[mk] == nil {
		tk := tagKeys(opt.Tags)

		s.counterVecMap[mk] = &promCounter{
			vec: prometheus.NewCounterVec(prometheus.CounterOpts{
				Namespace:   s.namespace,
				Name:        metric,
				ConstLabels: prometheus.Labels(s.baseTags),
			}, tk),
			labelKeys: tk,
		}

		prometheus.MustRegister(s.counterVecMap[mk].vec)
	}

	c := s.counterVecMap[mk]
	c.vec.WithLabelValues(tagValues(opt.Tags, c.labelKeys)...).Add(value)
	return nil
}
