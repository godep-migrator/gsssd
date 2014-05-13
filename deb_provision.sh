#!/bin/bash
set -e -u -o pipefail

sudo apt-get update -y
sudo apt-get install -y build-essential git mercurial unzip

echo "Host git.internal.justin.tv" >> ~/.ssh/config
echo "  StrictHostKeyChecking no" >> ~/.ssh/config

git clone git@git.internal.justin.tv:release/debpkg /tmp/debpkg
sudo dpkg -i /tmp/debpkg/twitch-golang_1.2-1_amd64.deb

export GOPATH=/home/vagrant/go
export PATH=${GOPATH}/bin:${PATH}
SRCPATH=${GOPATH}/src/github.com/ossareh/gsssd

mkdir -p ${GOPATH}/{src,pkg,bin}
mkdir -p ${SRCPATH}
cp -R /home/vagrant/src/* ${SRCPATH}
cd ${SRCPATH}
make clean prep
sudo chown -R root.root deb/
sudo make deb
cp ./gsssd_*.deb ~/src
