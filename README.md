Sample OpenTelemetry processor

This processor implements a FizzBuzz game on the number of traces received.

This repository also contains:

- A custom collector configuration that includes this processor
- A sample application that sends traces to the collector
- k6 scripts to test the collector

You can run the whole setup using docker compose.
