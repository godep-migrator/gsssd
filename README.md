# gsssd

[![Build Status](https://travis-ci.org/ossareh/gsssd.svg?branch=master)](https://travis-ci.org/ossareh/gsssd)

This daemon reads cpu and memory information out of proc and sends the
values to statsd. It is known to work with go 1.2.


## Installing

You can install gsssd by doing the following:

 1. `curl -LO http://science.twitch.tv/debs/gsssd_1.0_amd64.deb`
 2. `sudo dpkg --install gsssd_1.0_amd64.deb`


## Running

    gsssd -address="statsd.host.com" -prefix="useful.prefix"


## Building Deb Packages

`Makefile` includes a `release` target. Calling this will ramp up a
Vagrant instance which will build the project, a .deb, and upload it
to your desired location.

    make release BUCKET="s3://your_bucket/and/path"


### Requirements for building deb packages:

 * Vagrant
 * VirtualBox
 * s3cmd


## TODO

 * Stop using the twitch internal go build
