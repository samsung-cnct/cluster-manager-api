#!/bin/bash

set -o errexit
set -o nounset
set -o pipefail

THIS_DIRECTORY=$(dirname "${BASH_SOURCE}")
PROJECT_DIRECTORY=${THIS_DIRECTORY}/../..

echo
echo "Adjusting swagger to default to our API's swagger file"
sed -idummy 's/https:\/\/petstore.swagger.io\/v2\/swagger.json/..\/swagger\/api.swagger.json/' ${PROJECT_DIRECTORY}/third_party/swagger-ui/index.html
rm ${PROJECT_DIRECTORY}/third_party/swagger-ui/index.htmldummy
rm -rf ${PROJECT_DIRECTORY}/third_party/swagger-ui/*.map
