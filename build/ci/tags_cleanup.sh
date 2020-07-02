#!/usr/bin/env bash

set -e

git checkout "${TRAVIS_BRANCH}"

f_info_log "Cleaning old RFV tags"
# Delete old alpha tags on release
numTags=$(git tag -l "*-rfv.*" | wc -l | xargs)
if [ $numTags -gt 0 ]; then
  git push -q -d "https://${GITHUB_KEY}@github.com/${TRAVIS_REPO_SLUG}" $(git tag -l "*-rfv.*")
fi
f_info_log "RFV tags cleaned successfully"
