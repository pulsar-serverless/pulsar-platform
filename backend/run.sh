#!/bin/bash

while true; do
  # Run the Golang server command
  go run cmd/main.go &

  # Capture the background process ID
  SERVER_PID=$!

  # Wait for the background process to finish
  wait $SERVER_PID

  # Check the exit status
  EXIT_STATUS=$?

  if [ $EXIT_STATUS -ne 0 ]; then
    echo "Golang server exited with code: $EXIT_STATUS. Restarting..."
  else
    echo "Golang server exited normally."
  fi

done