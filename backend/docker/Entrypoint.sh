#!/bin/bash
set -e

echo "[ENTRYPOINT] Starting ava-sesisenai Go app..."

go build -o ava-api ./cmd/api

exec ./ava-api

