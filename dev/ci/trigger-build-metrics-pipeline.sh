#!/usr/bin/env bash

set -eux pipefail

cat > build-metrics-pipeline.yaml <<EOF
steps:
  - label: ":eye: Build Metrics"
    trigger: devx-build-metrics
    async: true
    depends_on: "__main__::finalization"
    build:
      message: "\${BUILDKITE_MESSAGE}"
      commit: "\${BUILDKITE_COMMIT}"
      branch: "\${BUILDKITE_BRANCH}"
      env:
        BUILDKITE_PULL_REQUEST: "\${BUILDKITE_PULL_REQUEST}"
        BUILDKITE_PULL_REQUEST_BASE_BRANCH: "\${BUILDKITE_PULL_REQUEST_BASE_BRANCH}"
        BUILDKITE_PULL_REQUEST_REPO: "\${BUILDKITE_PULL_REQUEST_REPO}"
EOF

buildkite-agent pipeline upload build-metrics-pipeline.yaml
