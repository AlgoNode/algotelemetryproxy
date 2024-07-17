FROM gcr.io/distroless/cc
COPY algotelemetry /app/
CMD ["/app/algotelemetry"]
