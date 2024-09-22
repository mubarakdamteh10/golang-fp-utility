run-test:
	go test -race ./... -failfast -count=1
	# golangci-lint run