.PHONY: test
test:
	# Recommended flags according to https://onsi.github.io/ginkgo/#recommended-continuous-integration-configuration
	go run github.com/onsi/ginkgo/v2/ginkgo \
	  -r \
	  --procs=1 \
	  --compilers=1 \
	  --randomize-all \
	  --randomize-suites \
	  --fail-on-pending \
	  --keep-going \
	  --cover \
	  --coverprofile=cover.profile \
	  --race \
	  --trace \
	  --json-report=report.json

.PHONY: test-no-race
test-no-race:
	# Recommended flags according to https://onsi.github.io/ginkgo/#recommended-continuous-integration-configuration,
	# but without `-race`.
	go run github.com/onsi/ginkgo/v2/ginkgo \
	  -r \
	  --procs=1 \
	  --compilers=1 \
	  --randomize-all \
	  --randomize-suites \
	  --fail-on-pending \
	  --keep-going \
	  --cover \
	  --coverprofile=cover.profile \
	  --trace \
	  --json-report=report.json
