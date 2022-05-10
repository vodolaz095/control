export app=control
export majorVersion=1
export minorVersion=0

export arch=$(shell uname)-$(shell uname -m)
export gittip=$(shell git log --format='%h' -n 1)
export subver=$(shell hostname)_on_$(shell date -u '+%Y-%m-%d_%I:%M:%S%p')
export patchVersion=$(shell git log --format='%h' | wc -l)
export ver=$(majorVersion).$(minorVersion).$(patchVersion).$(gittip)

clean:
	rm -f build/$(app)
	go clean

deps:
	go mod download
	go mod verify
	go mod tidy

lint:
	# format code
	gofmt -w=true -s=true -l=true ./models/
	gofmt -w=true -s=true -l=true ./components/
	gofmt -w=true -s=true -l=true main.go
	# run basic code quality and sanity check
	golint ./...
	go vet ./...

test: check

check: lint
# ran unit tests with coverage report                                                                                                                                |
	go test -v -coverprofile=cover.out ./...

start:
	forego start

build:
	go build -ldflags "-X main.Version=$(ver)" -o build/$(app) main.go
	upx build/$(app)
	./build/$(app)

rpm:
	which rpmbuild
	rpmbuild --version
	mkdir -p ~/rpmbuild/BUILD
	mkdir -p ~/rpmbuild/BUILDROOT
	mkdir -p ~/rpmbuild/RPMS
	mkdir -p ~/rpmbuild/SOURCES
	mkdir -p ~/rpmbuild/SPECS
	mkdir -p ~/rpmbuild/SRPMS
	git archive --format=tar.gz --prefix=$(app)-$(majorVersion).$(minorVersion).$(patchVersion).$(gittip)/ -o ~/rpmbuild/SOURCES/$(app)-$(majorVersion).$(minorVersion).$(patchVersion).$(gittip).tar.gz HEAD
	ls -l ~/rpmbuild/SOURCES/
	rpmbuild -ba --define 'version $(majorVersion).$(minorVersion).$(patchVersion).$(gittip)' $(app).spec
	rpm --addsign ~/rpmbuild/RPMS/x86_64/$(app)-$(majorVersion).$(minorVersion).$(patchVersion).$(gittip)*

server:
	forego start server

agent:
	forego start agent


.PHONY: build
