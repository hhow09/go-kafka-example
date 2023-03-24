#!/bin/bash
set -e

export ID=$HOSTNAME
echo "start app"
exec "/app/consumer"