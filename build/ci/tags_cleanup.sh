#!/usr/bin/env bash

set -e

# Delete old alpha tags on release
numTags=$(git tag -l "*-alpha.*" | wc -l | xargs)
if [ $numTags -gt 0 ]; then
  git push -q -d "https://${GITHUB_KEY}@github.com/${TRAVIS_REPO_SLUG}" $(git tag -l "*-rfv.*")
fi
