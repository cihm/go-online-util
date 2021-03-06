package testjaeger

import (
	"context"
	"net/http"
	"net/url"
	"time"

	"github.com/opentracing/opentracing-go"
	"github.com/opentracing/opentracing-go/ext"
	"github.com/opentracing/opentracing-go/log"
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

func MainCompleteflow(samplingServerURL, collectorEndpoint string) {

	interval := 5 * time.Second
	go func() {
		c := time.Tick(interval)
		loop := true
		for loop {
			select {
			case <-c:

				tracer, closer := InitJaeger("hello-world", samplingServerURL, collectorEndpoint)
				defer closer.Close()
				opentracing.SetGlobalTracer(tracer)

				helloTo := "lewis success"
				greeting := "lewis greeting"

				span := tracer.StartSpan("say-hello")
				span.SetTag("hello-to", helloTo)
				span.SetBaggageItem("greeting", greeting)
				//defer closer.Close()

				ctx := opentracing.ContextWithSpan(context.Background(), span)

				helloStr := formatStringc(ctx, helloTo)
				printHelloc(ctx, helloStr)
				span.Finish()
			}
		}
	}()

	interval2 := 7 * time.Second
	go func() {
		c := time.Tick(interval2)
		loop := true
		for loop {
			select {
			case <-c:
				go Mainflow(samplingServerURL, collectorEndpoint)

			}
		}
	}()
}

func formatStringc(ctx context.Context, helloTo string) string {
	span, _ := opentracing.StartSpanFromContext(ctx, "formatString")
	defer span.Finish()

	v := url.Values{}
	v.Set("helloTo", helloTo)
	url := "http://localhost:8081/format?" + v.Encode()
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		panic(err.Error())
	}

	ext.SpanKindRPCClient.Set(span)
	ext.HTTPUrl.Set(span, url)
	ext.HTTPMethod.Set(span, "GET")
	span.Tracer().Inject(
		span.Context(),
		opentracing.HTTPHeaders,
		opentracing.HTTPHeadersCarrier(req.Header),
	)

	resp, err := Do(req)
	if err != nil {
		panic(err.Error())
	}

	helloStr := string(resp)

	span.LogFields(
		log.String("event", "string-format"),
		log.String("value", helloStr),
	)

	return helloStr
}

func printHelloc(ctx context.Context, helloStr string) {
	span, _ := opentracing.StartSpanFromContext(ctx, "printHello")
	defer span.Finish()

	v := url.Values{}
	v.Set("helloStr", helloStr)
	url := "http://localhost:8082/publish?" + v.Encode()
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		panic(err.Error())
	}

	ext.SpanKindRPCClient.Set(span)
	ext.HTTPUrl.Set(span, url)
	ext.HTTPMethod.Set(span, "GET")
	span.Tracer().Inject(span.Context(), opentracing.HTTPHeaders, opentracing.HTTPHeadersCarrier(req.Header))

	if _, err := Do(req); err != nil {
		panic(err.Error())
	}
}
