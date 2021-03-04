#!/usr/bin/env bash

set -e

DIR="$(dirname "$0")"
source "${DIR}"/include.sh

git checkout "${TRAVIS_BRANCH}"

f_info_log "Calculating tag ${TRAVIS_BRANCH} branch..."
if [ "${TRAVIS_BRANCH}" = "master" ]; then
  tag=v$(docker run --rm -v "$(pwd):/repo" gittools/gitversion:5.6.4-debian.9-x64-5.0 /repo -output json -showvariable MajorMinorPatch)
elif [ "${TRAVIS_BRANCH}" = "develop" ]; then
  tag=v$(docker run --rm -v "$(pwd):/repo" gittools/gitversion:5.6.4-debian.9-x64-5.0 /repo -output json -showvariable SemVer)
fi
f_info_log "The tag for current source code is: ${tag}"

# Bump up npm version if is a master release
if [ "${TRAVIS_BRANCH}" = "master" ]; then
  npm version "$tag" --no-git-tag-version --allow-same-version
  git add package.json
  git commit -m "[NPM version] Bump up NPM version to ${tag}"
  git push -q "https://${GITHUB_TOKEN}@github.com/${TRAVIS_REPO_SLUG}" "$TRAVIS_BRANCH"
  # Override commit hash given the npm commit should be the one tagged
  TRAVIS_COMMIT=$(git rev-parse HEAD)
fi

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
