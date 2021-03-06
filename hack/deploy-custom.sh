#!/bin/bash
set -e

: "${DEPLOY_PLATFORM:?must be set to openshift}"


export INSTALL_METHOD="${DEPLOY_PLATFORM}-direct"
export DELETE_PVCS=${DELETE_PVCS:-true}

: "${CUSTOM_METERING_CR_FILE:?Must set \$CUSTOM_METERING_CR_FILE}"

: "${CUSTOM_HELM_OPERATOR_OVERRIDE_VALUES:?Must set \$CUSTOM_HELM_OPERATOR_OVERRIDE_VALUES}"
: "${CUSTOM_OLM_OVERRIDE_VALUES:?Must set \$CUSTOM_OLM_OVERRIDE_VALUES}"

TMP_DIR="$(mktemp -d)"
export CUSTOM_DEPLOY_MANIFESTS_DIR=${CUSTOM_DEPLOY_MANIFESTS_DIR:-"$TMP_DIR/custom-deploy-manifests"}
export CUSTOM_HELM_OPERATOR_OVERRIDE_VALUES
export CUSTOM_OLM_OVERRIDE_VALUES

ROOT_DIR=$(dirname "${BASH_SOURCE}")/..
source "${ROOT_DIR}/hack/common.sh"

export METERING_OPERATOR_IMAGE_TAG="${DEPLOY_TAG}"

echo "Creating metering manifests"
export MANIFEST_OUTPUT_DIR="$CUSTOM_DEPLOY_MANIFESTS_DIR"
"$ROOT_DIR/hack/create-metering-manifests.sh"

echo "Deploying"
export DEPLOY_MANIFESTS_DIR="$CUSTOM_DEPLOY_MANIFESTS_DIR"
"${ROOT_DIR}/hack/deploy.sh"
