package metriccounter
package mtscount
package customprocessor
import (
	"context"

	"github.com/open-telemetry/opentelemetry-collector-contrib/pkg/translator/opencensus"
        "go.opentelemetry.io/collector/model/pdata"
)

type metricCounterProcessor struct {
}

func newMetricCounterProcessor() *metricCounterProcessor {
	return &metricCounterProcessor{}
}

func (p *metricCounterProcessor) ConsumeMetrics(ctx context.Context, md pdata.Metrics) error {
	metricCount := 0
	resourceMetrics := md.ResourceMetrics()
	for i := 0; i < resourceMetrics.Len(); i++ {
		resourceMetric := resourceMetrics.At(i)
		ilms := resourceMetric.InstrumentationLibraryMetrics()
		for j := 0; j < ilms.Len(); j++ {
			ilm := ilms.At(j)
			metrics := ilm.Metrics()
			metricCount += metrics.Len()
		}
	}

	// Update a counter with the metricCount value or perform any necessary operation with the count.

	// Pass the metrics along unchanged.
	return nil
}

func (p *metricCounterProcessor) Capabilities() component.ProcessorCapabilities {
	return component.ProcessorCapabilities{MutatesData: false}
}

func (p *metricCounterProcessor) Start(_ context.Context, _ component.Host) error {
	return nil
}

func (p *metricCounterProcessor) Shutdown(context.Context) error {
	return nil
}

