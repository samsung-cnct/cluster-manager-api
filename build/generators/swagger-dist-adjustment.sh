#!/bin/bash

set -o errexit
set -o nounset
set -o pipefail

THIS_DIRECTORY=$(dirname "${BASH_SOURCE}")
PROJECT_DIRECTORY=${THIS_DIRECTORY}/../..

echo
echo "Adjusting swagger to default to our API's swagger file"
echo "OK if you see \"sed: unsupported command u\" and this is running on linux/gnu-sed"
(sed -i dummy 's/http:\/\/petstore.swagger.io\/v2\/swagger.json/..\/swagger\/api.swagger.json/' ${PROJECT_DIRECTORY}/third_party/swagger-ui/index.html) || (sed -idummy 's/http:\/\/petstore.swagger.io\/v2\/swagger.json/..\/swagger\/api.swagger.json/' ${PROJECT_DIRECTORY}/third_party/swagger-ui/index.html )
rm ${PROJECT_DIRECTORY}/third_party/swagger-ui/index.htmldummy
