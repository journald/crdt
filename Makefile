test:
	@go test -v ./... | \
	sed s/RUN/$$(printf "\033[34mRUN\033[0m")/ | \
	sed s/PASS/$$(printf "\033[32mPASS\033[0m")/ | \
	sed s/FAIL/$$(printf "\033[31mFAIL\033[0m")/
