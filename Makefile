build::
	go get -d -v ./...
	go build -v ./...
	go test -v ./...

.PHONY:
gsssd: build
	go build -o gsssd

.PHONY: prep
prep: gsssd
	rm -rf .deb/usr/bin/gsssd
	mkdir -p ./deb/usr/bin
	cp ./gsssd ./deb/usr/bin/
	test -f ./deb/usr/bin/gsssd

.PHONY: deb
deb:
	dpkg-deb -Z gzip -b ./deb .

.PHONY: clean
clean:
	go clean
	-rm gsssd_*.deb

BUCKET = "_"

.PHONY: upload
upload:
	s3cmd put gsssd_*.deb $(BUCKET)

.PHONY: vagrant
vagrant:
	vagrant up
	vagrant destroy -f

.PHONY: release
release: vagrant upload
