NAME     := $(shell basename $(CURDIR))
VERSION  := 0.0.1
REVISION := $(shell git rev-parse --short HEAD)

SRCS := $(shell find $(CURDIR) -type f -name '*.go')

GOOS   := linux
GOARCH := amd64

LDFLAGS_NAME     := -X "main.name=$(NAME)"
LDFLAGS_VERSION  := -X "main.version=v$(VERSION)"
LDFLAGS_REVISION := -X "main.revision=$(REVISION)"
LDFLAGS          := -ldflags '-s -w $(LDFLAGS_NAME) $(LDFLAGS_VERSION) $(LDFLAGS_REVISION) -extldflags -static'

.PHONY: all
all: run

.PHONY: $(NAME)
$(NAME): $(CURDIR)/bin/$(NAME)
$(CURDIR)/bin/$(NAME): $(SRCS)
	GOOS=$(GOOS) GOARCH=$(GOARCH) go build $(LDFLAGS) -o $@

$(CURDIR)/bin/$(NAME).zip: $(CURDIR)/bin/$(NAME)
	cd $(CURDIR)/bin && zip $@ $(NAME)

.PHONY: run
run: $(CURDIR)/bin/$(NAME)
	sudo $(CURDIR)/bin/$(NAME)

.PHONY: test
test:
	go test -v

.PHONY: clean
clean:
	rm -rf $(CURDIR)/bin
