#!/usr/bin/env bash

set -eux pipefail

cat > build-metrics-pipeline.yaml <<EOF
steps:
  - label: ":eye: Build Metrics"
    trigger: devx-build-metrics
    async: true
    depends_on: "__main__::finalization"
EOF

buildkite-agent pipeline upload build-metrics-pipeline.yaml
