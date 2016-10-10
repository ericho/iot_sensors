package environment

import (
	"time"

	"github.com/intelsdi-x/snap/control/plugin"
	"github.com/intelsdi-x/snap/control/plugin/cpolicy"
	"github.com/intelsdi-x/snap/core"
)

const (
	Name = "environment"
	Version = 1
	Type = plugin.CollectorPluginType
)

var _ plugin.CollectorPlugin = (*Environment)(nil)

type Environment struct {
}

func (f *Environment) GetConfigPolicy() (*cpolicy.ConfigPolicy, error) {
	c := cpolicy.New()
	return c, nil
}

func (f *Environment) GetMetricTypes(_ plugin.ConfigType) ([]plugin.MetricType, error) {
	mts := []plugin.MetricType{}
	mts = append(mts, plugin.MetricType{Namespace_: core.NewNamespace("intel", "edison", "temperature")})
	mts = append(mts, plugin.MetricType{Namespace_: core.NewNamespace("intel", "edison", "light")})
	return mts, nil
}

func (f *Environment) CollectMetrics(mts []plugin.MetricType) ([]plugin.MetricType, error) {
	metrics := make([]plugin.MetricType, len(mts))

	for i, m := range mts {
		metrics[i] = plugin.MetricType{
			Namespace_: m.Namespace(),
			Data_:		Temperature(),
			Timestamp_:	time.Now(),
		}
	}
	return metrics, nil
}

func Meta() *plugin.PluginMeta {
	return plugin.NewPluginMeta(
		Name,
		Version,
		Type,
		[]string{plugin.SnapGOBContentType},
		[]string{plugin.SnapGOBContentType},
		plugin.Unsecure(true),
		plugin.RoutingStrategy(plugin.DefaultRouting),
		plugin.CacheTTL(1100*time.Millisecond),
	)
}
