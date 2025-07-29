#!/bin/bash

# Test script to verify GoReleaser configuration
echo "Testing GoReleaser configuration..."

# Check if goreleaser is installed
if ! command -v goreleaser &> /dev/null; then
    echo "GoReleaser not found. Installing..."
    go install github.com/goreleaser/goreleaser@latest
fi

# Test the configuration without releasing
echo "Running GoReleaser check..."
goreleaser check

echo "Running GoReleaser build (snapshot mode)..."
goreleaser build --snapshot --clean

echo "Listing generated binaries..."
ls -la dist/

echo ""
echo "To create a new release:"
echo "1. Update version in your code if needed"
echo "2. Create and push a new tag: git tag v0.0.6 && git push origin v0.0.6"
echo "3. The GitHub Action will automatically:"
echo "   - Build cross-platform binaries"
echo "   - Create GitHub release with assets"
echo "   - Update Homebrew formula with binary downloads"
echo "   - Generate proper SHA checksums"
