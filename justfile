install-asdf:
  asdf plugin add golang python
  asdf install

install-docs:
  pip install mkdocs-material

install: install-asdf install-docs

[working-directory: 'cmd']
run-cmd:
  go run .

run-docs:
  mkdocs serve

build-cmd:
  go build -C cmd -o dist/rich-chat-statuses

build-image:
  docker build -t ghcr.io/bnjns/rich-chat-statuses:latest .

test-app:
  go test -v ./...

[working-directory: 'cmd']
test-cmd:
  go test -v ./...

test-calendars:
  @find calendars -name go.mod -execdir pwd \; \
    | xargs -I {} bash -c 'cd {} && go test -v ./...'

test-clients:
  @find clients -name go.mod -execdir pwd \; \
      | xargs -I {} bash -c 'cd {} && go test -v ./...'

test-all: test-app test-cmd test-calendars test-clients

lint-app:
  golangci-lint run .

lint-cmd:
  golangci-lint run cmd

lint-calendars:
  @find calendars -name go.mod -execdir pwd \; \
    | xargs -I {} bash -c 'echo {} && golangci-lint run {}'

lint-clients:
  @find clients -name go.mod -execdir pwd \; \
      | xargs -I {} bash -c 'echo {} && golangci-lint run {}'

lint-all: lint-app lint-cmd lint-calendars lint-clients
