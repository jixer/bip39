.DEFAULT_GOAL := compile

compile: 
	@echo "Compiling..."
	@mkdir -p ./targets
	@go build -o ./targets -v .
	@echo "Done."

distribute_one:
	@mkdir -p ./targets/distributions/$(GOOS)/$(GOARCH)
	@echo "Compiling for $(GOOS) $(GOARCH)..."
	@GOOS=$(GOOS) GOARCH=$(GOARCH) go build -o ./targets/distributions/$(GOOS)/$(GOARCH) .
	@zip ./targets/distributions/$(GOOS)-$(GOARCH).zip ./targets/distributions/$(GOOS)/$(GOARCH)/bip39*
	@rm -rf ./targets/distributions/$(GOOS)

distribute:
	@echo "Distributing for all target architectures..."
	
	@echo "Compiling for linux..."
	@GOOS=linux GOARCH=amd64 make distribute_one
	@GOOS=linux GOARCH=386 make distribute_one
	@GOOS=linux GOARCH=arm make distribute_one
	@GOOS=linux GOARCH=arm64 make distribute_one

	@echo "Compiling for OSX..."
	@GOOS=darwin GOARCH=amd64 make distribute_one
	@GOOS=darwin GOARCH=arm64 make distribute_one

	@echo "Compiling for Windows..."
	@GOOS=windows GOARCH=386 make distribute_one
	@GOOS=windows GOARCH=amd64 make distribute_one

	@echo "Done."