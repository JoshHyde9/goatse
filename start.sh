#!/bin/bash

# Exit on any error
set -e

echo "Starting services..."

# Start docker-compose in the background
echo "Starting Docker Compose..."
docker compose up -d

# Start the client in the background
echo "Starting client with bun dev..."
cd ./client
bun dev &
CLIENT_PID=$!

# Move back to the original directory
cd ..

# Start the Go server
echo "Starting Go server..."
cd ./server
go run main.go &
SERVER_PID=$!

# Store PIDs for potential cleanup
echo "Started processes:"
echo "Client PID: $CLIENT_PID"
echo "Server PID: $SERVER_PID"

# Set up a trap to catch Ctrl+C and clean up processes
trap cleanup INT TERM
cleanup() {
    echo "Shutting down services..."

    # Kill the client and server processes
    kill $CLIENT_PID $SERVER_PID 2>/dev/null || true

    # Shutdown docker-compose
    docker compose down

    echo "All services stopped."
    exit 0
}

# Keep the script runningc
echo "All services started. Press Ctrl+C to stop all services."
wait
