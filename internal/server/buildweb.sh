#!/bin/bash

set -e

cd ../../web

yarn install
export NODE_ENV=production
yarn build