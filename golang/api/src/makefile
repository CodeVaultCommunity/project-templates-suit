APP_NAME=mod_name
SONAR_PROJECT_KEY=project-key
SONAR_HOST_URL=http://localhost:9000
SONAR_TOKEN=sqp_e76d9a54106716199c6bf886cc79f0fd9c18ef0b

ifeq ($(OS),Windows_NT)
    MKDIR = if not exist .report mkdir
	RM_DIR_CMD = if exist $(1) rmdir /s /q $(1)
else
    MKDIR = mkdir -p
	RM_DIR_CMD = rm -rf $(1)
endif

define RMDIR
	$(call RM_DIR_CMD,$(1))
endef

runserver:
	swag init
	go run .

run:
	go run ./cmd/$(APP_NAME)

build:
	go build -o bin/$(APP_NAME) ./cmd/$(APP_NAME)

lint:
	golangci-lint run

test:
	go test ./...

prepare-report:
	$(MKDIR) .report

lint-report: prepare-report
	golangci-lint run

test-report: prepare-report
	go test ./... -coverprofile=.report/coverage.out

test-coverview: test-report
	go tool cover -html=./.report/coverage.out

generate-report:
	$(MKDIR) .report
	- golangci-lint run
	- go test ./... -coverprofile=.report/coverage.out

clean:
	$(call RMDIR,bin)
	$(call RMDIR,.report)

sonar: generate-report
	sonar-scanner \
	-Dsonar.projectKey="$(SONAR_PROJECT_KEY)" \
	-Dsonar.host.url="$(SONAR_HOST_URL)" \
	-Dsonar.login="$(SONAR_TOKEN)"