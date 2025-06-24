package fizzbuzzprocessor

import (
	"context"
	"fmt"

	"go.opentelemetry.io/collector/consumer"
	"go.opentelemetry.io/collector/pdata/ptrace"
	"go.opentelemetry.io/collector/processor"
	"go.opentelemetry.io/collector/processor/processorhelper"
)

type fizzbuzzProcessor struct {
	counter int
}

func newFizzBuzzProcessor(
	ctx context.Context,
	set processor.Settings,
	cfg *Config,
	nextConsumer consumer.Traces,
) (processor.Traces, error) {
	tp := &fizzbuzzProcessor{counter: 0}
	return processorhelper.NewTraces(
		ctx,
		set,
		cfg,
		nextConsumer,
		tp.processTraces,
		processorhelper.WithCapabilities(consumer.Capabilities{MutatesData: true}))
}

func (tp *fizzbuzzProcessor) processTraces(ctx context.Context, td ptrace.Traces) (ptrace.Traces, error) {
	tp.counter += 1
	if tp.counter%3 == 0 && tp.counter%5 == 0 {
		fmt.Println("FizzBuzz")
	} else if tp.counter%3 == 0 {
		fmt.Println("Fizz")
	} else if tp.counter%5 == 0 {
		fmt.Println("Buzz")
	} else {
		fmt.Println(tp.counter)
	}
	return td, nil
}
