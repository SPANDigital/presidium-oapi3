TESTDIRS=`go list ./... |grep -v "vendor/"`
devops_path=$DEFAULT_DEVOPS

default:
	rm -rf dist
	mkdir -p dist
	go build -o dist/presidium-oapi3 main.go

generate:
	go-bindata --nometadata -pkg tpl -o pkg/tpl/tpl.go templates/...

clean:
	rm -rf dist
	rm -rf reports
	rm -rf tmp

test:
	@mkdir -p reports
	go test -failfast -coverprofile reports/coverage.out -v ./...

# Renders the generated markdown with Hugo and the pinned Presidium theme,
# then checks every internal link and anchor. Needs hugo and network access.
render-test:
	PRESIDIUM_RENDER_TEST=1 go test -count=1 -run TestRenderedSiteHasNoBrokenLinks -v ./pkg/markdown/

test_reports:
	@mkdir -p reports
	@go test -v $(TESTDIRS) -coverprofile=reports/tests-cov.out -json > reports/tests.json

coverage_report:
	@go tool cover -html=reports/tests-cov.out


.PHONY: default clean test test_reports coverage_report

# Markdown linting
markdown-lint:
	docker run -v $(shell pwd):/workdir ghcr.io/igorshubovych/markdownlint-cli:latest README.md

markdown-lint-fix:
	docker run -v $(shell pwd):/workdir ghcr.io/igorshubovych/markdownlint-cli:latest README.md --fix
