package testjaeger

import (
	"context"
	"fmt"
	"io"

	opentracing "github.com/opentracing/opentracing-go"
	"github.com/opentracing/opentracing-go/log"
	jaeger "github.com/uber/jaeger-client-go"
	config "github.com/uber/jaeger-client-go/config"
)

/*

Tags:
- db.instance:"jdbc:mysql://127.0.0.1:3306/customers
- db.statement: "SELECT * FROM mytable WHERE foo='bar';"

Logs:
- message:"Can't connect to mysql server on '127.0.0.1'(10061)"

SpanContext:
- trace_id:"abc123"
- span_id:"xyz789"
- Baggage Items:
  - special_id:"vsid1738"

*/

func Mainflow(samplingServerURL, collectorEndpoint string) {
	helloTo := "lewis success"
	tracer, closer := InitJaeger("hello-world", samplingServerURL, collectorEndpoint)
	//fmt.Println(tracer)
	opentracing.SetGlobalTracer(tracer)
	span := tracer.StartSpan("say-hello")
	span.SetTag("hello-to", helloTo)

	// helloStr := fmt.Sprintf("Hello, %s!", helloTo)
	// span.LogFields(
	// 	log.String("event", "string-format"),
	// 	log.String("value", helloStr),
	// )

	// println(helloStr)
	// span.LogKV("event", "println")
	// span.Finish()
	defer closer.Close()
	defer span.Finish()

	ctx := opentracing.ContextWithSpan(context.Background(), span)
	helloStr := formatString(ctx, helloTo)
	printHello(ctx, helloStr)

	span2 := tracer.StartSpan("say-hello2")
	defer span2.Finish()
	span2.SetTag("hello-to2", helloTo)

	ctx2 := opentracing.ContextWithSpan(context.Background(), span2)
	helloStr = formatString(ctx2, helloTo)
	printHello(ctx2, helloStr)

}

func InitJaeger(service string, samplingServerURL, collectorEndpoint string) (opentracing.Tracer, io.Closer) {
	cfg := &config.Configuration{
		Sampler: &config.SamplerConfig{
			SamplingServerURL: samplingServerURL,
			Type:              "const",
			Param:             1,
		},
		Reporter: &config.ReporterConfig{
			CollectorEndpoint: collectorEndpoint,
			LogSpans:          true,
			User:              "admin",
			Password:          "admin",
		},
	}
	tracer, closer, err := cfg.New(service, config.Logger(jaeger.StdLogger))
	if err != nil {
		panic(fmt.Sprintf("ERROR: cannot init Jaeger: %v\n", err))
	}
	return tracer, closer
}

func formatString(ctx context.Context, helloTo string) string {
	span, _ := opentracing.StartSpanFromContext(ctx, "formatString")
	defer span.Finish()

	helloStr := fmt.Sprintf("Hello, %s!", helloTo)
	span.LogFields(
		log.String("event", "string-format"),
		log.String("value", helloStr),
	)

	return helloStr
}

func printHello(ctx context.Context, helloStr string) {
	span, _ := opentracing.StartSpanFromContext(ctx, "printHello")
	defer span.Finish()

	println(helloStr)
	span.LogKV("event", "println")
}
