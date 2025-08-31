@app.route("/manual")
def manual_span():
  tracer = trace.get_tracer(__name__)

  # Start the span manually
  span = tracer.start_span("manual-span")

  try:
    # Do some work
    import time
    time.sleep(0.2)

    # You can add attributes/logs
    span.set_attribute("work.type", "manual")
    span.add_event("doing some manual work")

    result = "Work done manually!"

  finally:
    # Explicitly end the span
    span.end()

  return result


