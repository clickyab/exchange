#!/bin/bash
set -x
set -eo pipefail

# This job is from jenkins. so kill it if it is a pull request
exit_message() {
    echo ${1:-'exiting...'}
    code=${2:-1}
    exit ${code}
}

env
APP=${APP:-}
BRANCH=${BRANCH_NAME:-master}
BRANCH=${CHANGE_TARGET:-${BRANCH}}

[ -z ${CHANGE_AUTHOR} ] || exit_message "It's a PR, bail out" 0
if [[ ( "${BRANCH}" != "master" ) && ( "${BRANCH}" != "deploy" ) ]]; then
    exit_message "Its not on correct branch, bail out" 0
fi
[ -z ${APP} ] && exit_message "The APP is not defined." # WTF, the APP NAME is important

SCRIPT_DIR=$(readlink -f $(dirname ${BASH_SOURCE[0]}))

SOURCE_DIR=${1:-}
[ -z ${SOURCE_DIR} ] && exit_message "Must pass the source directory as the first parameter" 1
SOURCE_DIR=$(cd "${SOURCE_DIR}/" && pwd)

BUILD_DIR=${2:-${SOURCE_DIR}-build}
CACHE_DIR=${3:-${SOURCE_DIR}-cache}
ENV_DIR=$(mktemp -d)

mkdir -p "${BUILD_DIR}" "${CACHE_DIR}" "${ENV_DIR}"
BUILD=$(cd "${BUILD_DIR}/" && pwd)
CACHE=$(cd "${CACHE_DIR}/" && pwd)
VARS=$(cd "${ENV_DIR}/" && pwd)

#chown $(id -u):$(id -g) $ENV_DIR
#chown $(id -u):$(id -g) $CACHE_DIR

BUILD_PACKS_DIR=$(mktemp -d)

# Extract build data
pushd ${SOURCE_DIR}
GIT_WORK_TREE=${BUILD} git checkout -f ${BRANCH}

export LONGHASH=$(git log -n1 --pretty="format:%H" | cat)
export SHORTHASH=$(git log -n1 --pretty="format:%h"| cat)
export COMMITDATE=$(git log -n1 --date="format:%D-%H-%I-%S" --pretty="format:%cd"| sed -e "s/\//-/g")
export IMPDATE=$(date +%Y%m%d)
export COMMITCOUNT=$(git rev-list HEAD --count| cat)
export BUILDDATE=$(date "+%D/%H/%I/%S"| sed -e "s/\//-/g")
popd

# Populate env for herokuish
env -0 | while IFS='=' read -r -d '' n v; do
    echo "${v}">"${VARS}/${n}";
done< <(env -0)

TEMPORARY=$(mktemp -d)

# Create Rockerfile to build with rocker (the Dockerfile enhancer tool)
cat > ${TEMPORARY}/Rockerfile <<EOF
FROM gliderlabs/herokuish

MOUNT {{ .Build }}:/tmp/app
MOUNT {{ .EnvDir }}:/tmp/env
MOUNT {{ .Target }}:/tmp/build
MOUNT {{ .Cache }}:/tmp/cache

RUN /bin/herokuish buildpack build && rm -rf /app/pkg && rm -rf /app/tmp
TAG registry.clickyab.ae/clickyab/{{ .App }}:{{ .Version }}
PUSH registry.clickyab.ae/clickyab/{{ .App }}:{{ .Version }}
EOF

TARGET=$(mktemp -d)
pushd ${TEMPORARY}
# Actual build
rocker build --push -var Build=${BUILD} -var EnvDir=${VARS} -var Cache=${CACHE} -var Target=${TARGET} -var Version=${BRANCH}.${COMMITCOUNT} -var App=${APP}
popd

rm -rf ${VARS} || true
rm -rf ${TARGET} || true

for WRK_TYP in web winner impression demand show aggregator
do
    kubectl -n ${APP} set image deployment  ${APP}-${WRK_TYP} ${APP}-${BRANCH}=registry.clickyab.ae/clickyab/${APP}:${BRANCH}.${COMMITCOUNT} --record
done

fi


