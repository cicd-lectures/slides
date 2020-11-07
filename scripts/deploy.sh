#!/bin/bash

set -eux

ZIP_FILE=gh-pages.zip
ARCHIVE_URL="${REPOSITORY_BASE_URL}/archive/${ZIP_FILE}"
CURRENT_DIR="$(cd "$(dirname "${0}")" && pwd -P)"
DOCS_DIR="${CURRENT_DIR}/../docs"

# Rebuild the docs directory which will be uploaded to gh-pages
rm -rf "${DOCS_DIR}"
if curl -sSLI --fail "${ARCHIVE_URL}"
then
    curl -sSLO "${ARCHIVE_URL}"
    unzip -o "./${ZIP_FILE}"
    mv ./"$(basename "${REPOSITORY_BASE_URL}")"-${ZIP_FILE%%.*} "${DOCS_DIR}" # No ".zip" at the end
    rm -f "./${ZIP_FILE}"
else
    echo "== No gh-pages found, I assume this is the first time."
    mkdir -p "${DOCS_DIR}"
fi

# If a tag triggered the deploy, we deploy to a folder having the tag name
# Same if it is a branch different of "gh-pages" or "main"
# otherwise we are on main and we deploy into latest
set +u
if [ -n "${GIT_TAG}" ]; then
    echo "== Using tag ${GIT_TAG}"
    DEPLOY_DIR="${DOCS_DIR}/${GIT_TAG}"
    # Generate QRCode and overwrite the default one
    make qrcode
elif [ -n "${GIT_BRANCH}" ] && [ "${GIT_BRANCH}" != "main" ]; then
    echo "== Using branch ${GIT_BRANCH}"
    DEPLOY_DIR="${DOCS_DIR}/${GIT_BRANCH}"
    # Generate QRCode and overwrite the default one
    make qrcode
else
    DEPLOY_DIR="${DOCS_DIR}"
fi
set -u

rsync -av ./dist/ "${DEPLOY_DIR}"
