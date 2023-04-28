package metriccounter
package customprocessor
package mtscount
import (
	"context"

	"github.com/open-telemetry/opentelemetry-collector-contrib/pkg/translator/opencensus"
	"go.opentelemetry.io/collector/model/pdata"

)

const (
	typeStr = "metric_counter"
)

type factory struct {
}

func NewFactory() component.ProcessorFactory {
	return &factory{}
}

func (f *factory) Type() configmodels.Type {
	return configmodels.Type(typeStr)
}

func (f *factory) CreateDefaultConfig() configmodels.Processor {
	return &configmodels.ProcessorSettings{
		TypeVal: configmodels.Type(typeStr),
		NameVal: typeStr,
	}
}

func (f *factory) CreateMetricsProcessor(
	ctx context.Context,
	params component.ProcessorCreateParams,
	cfg configmodels.Processor,
	nextConsumer component.MetricsConsumer) (component.MetricsProcessor, error) {

	return newMetricCounterProcessor(), nil
}

