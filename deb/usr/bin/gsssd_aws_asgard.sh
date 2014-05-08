#!/bin/bash --
set -e -u -o pipefail

eval "$(curl 169.254.169.254/latest/user-data/)"
export PUBLIC_IP="$(curl 169.254.169.254/latest/meta-data/public-ipv4 | sed 's/\./-/g')"

/usr/bin/gsssd \
    -address="ec2-54-186-66-175.us-west-2.compute.amazonaws.com" \
    -prefix="${CLOUD_APP}.${CLOUD_ENVIRONMENT}.${EC2_REGION}.${CLOUD_AUTO_SCALE_GROUP##*-}.${PUBLIC_IP}" &
