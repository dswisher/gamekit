#!/bin/bash
# create-tags.sh - Create version tags for gamekit packages
#
# This script creates git tags for all gamekit packages at the specified version.
# Tags follow the format: <package>/v<version>
#
# Usage: ./scripts/create-tags.sh

set -e

VERSION="0.4.0"
PACKAGES=("ecs" "sprites" "scenes")

echo "Creating tags for version v${VERSION}..."
echo ""

for package in "${PACKAGES[@]}"; do
    tag="${package}/v${VERSION}"
    echo "Creating tag: ${tag}"
    git tag "${tag}"
done

echo ""
echo "Tags created successfully!"
echo ""
echo "To push all tags to the remote repository, run:"
echo ""
echo "  git push origin --tags"
echo ""
