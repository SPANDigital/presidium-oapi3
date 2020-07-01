#!/usr/bin/env bash

set -e

DIR="$(dirname "$0")"
source "${DIR}"/include.sh

git checkout "${TRAVIS_BRANCH}"

f_info_log "Calculating tag ${TRAVIS_BRANCH} branch..."
tag=v$(docker run --rm -v "$(pwd):/repo" gittools/gitversion:5.2.4-linux-ubuntu-18.04-netcoreapp3.1 /repo -output json -showvariable SemVer)
f_info_log "The tag for current source code is: ${tag}"

# Create tag in github
f_info_log "Pushing tag to github..."
curl -s -X POST https://api.github.com/repos/"$TRAVIS_REPO_SLUG"/git/refs -H "Authorization: token $GITHUB_TOKEN" \
  -d @- <<EOF
{
  "ref": "refs/tags/$tag",
  "sha": "$TRAVIS_COMMIT"
}
EOF
f_info_log "Source code tagged successfully as ${tag}"
