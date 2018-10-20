#!/bin/bash

set -ex

if [ ! -e ${HOME}/google-cloud-sdk/bin/gcloud ]; then
   curl https://sdk.cloud.google.com | bash >> /dev/null;
   source ${HOME}/google-cloud-sdk/path.bash.inc
fi

openssl aes-256-cbc -K $encrypted_340203539710_key -iv $encrypted_340203539710_iv -in .travis/gcloud-service-account.json.enc -out gcloud-service-account.json -d
gcloud auth activate-service-account --key-file=gcloud-service-account.json

gcloud auth configure-docker

docker build -t eu.gcr.io/myfoobarproject/fizzbuzz-api:${TRAVIS_TAG} -t eu.gcr.io/myfoobarproject/fizzbuzz-api:latest .

docker push eu.gcr.io/myfoobarproject/fizzbuzz-api:${TRAVIS_TAG}

docker push eu.gcr.io/myfoobarproject/fizzbuzz-api:latest
