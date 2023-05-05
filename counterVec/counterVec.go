package counterVec

import (
	"github.com/prometheus/client_golang/prometheus"
)

type CounterVec struct {
	counterVec *prometheus.CounterVec
}

const (
	Hit  = "hit"
	Fail = "failure"
)

var types = []string{Hit, Fail}

func New() (ICounterVec, error) {
	counterVec := prometheus.NewCounterVec(prometheus.CounterOpts{
		Name: "org_requests_total",
		Help: "Total number of requests hits/failures.",
	}, []string{"type", "org_uid"})

	for _, value := range types {
		counterVec.WithLabelValues(value, "")
	}

	prometheus.Unregister(counterVec)

	if err := prometheus.Register(counterVec); err != nil {
		return nil, err
	}

	return &CounterVec{
		counterVec: counterVec,
	}, nil
}

func (c *CounterVec) CreateHit(orgUid string) {
	c.counterVec.WithLabelValues(Hit, orgUid).Inc()
}

func (c *CounterVec) CreateFailure(orgUid string) {
	c.counterVec.WithLabelValues(Fail, orgUid).Inc()
}
