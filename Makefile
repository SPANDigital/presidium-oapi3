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


# test runs all tests
test:
	@mkdir -p reports
	go test -p 1 -v $(TESTDIRS) -coverprofile=reports/tests-cov.out

test_reports:
	@mkdir -p reports
	@go test -p 1 -v $(TESTDIRS) -coverprofile=reports/tests-cov.out -json > reports/tests.json

coverage_report:
	@go tool cover -html=reports/tests-cov.out


.PHONY: default clean test test_reports coverage_report
