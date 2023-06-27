#!/bin/sh
# reference: https://qiita.com/mazgi/items/585348b6cdff3e320726

# example: $1="http://localhost:8080/api/v1/ping"
url=$1
status=`curl -LI $url -o /dev/null -w '%{http_code}\n' -s`

echo "response status(${url}): ${status}"
