#!/bin/bash
ENV_FILE="$(dirname "$0")/../.env"

if [ -f "$ENV_FILE" ]; then
    echo "Loading secrets from $ENV_FILE..."
    # This magic line exports all variables from .env into this script's session
    export $(grep -v '^#' "$ENV_FILE" | xargs)
else
    echo "Error: .env file not found at $ENV_FILE"
    exit 1
fi

if [ -z "$REDIS_PORT" ]; then
    echo "Error: REDIS_PORT is not set in your .env file!"
    exit 1
fi
if [ -z "$REDIS_PASSWORD" ]; then
    echo "Error: REDIS_PASSWORD is not set in your .env file!"
    exit 1
fi

echo "Starting Redis with the password from .env..."

docker run -d \
  --name redis-thestral \
  --restart always \
  -p $REDIS_PORT:6379 \
  -v redis-thestral:/data \
  redis:alpine \
  redis-server --appendonly yes --requirepass "$REDIS_PASSWORD"