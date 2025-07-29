# The missing `retry` command

```txt
NAME:
   retry - The missing retry command

USAGE:
   retry [options] -- command

GLOBAL OPTIONS:
   --count value, -c value     Number of retries (default: 5)
   --interval value, -i value  Interval between retries (default: 3s)
   --timeout value, -t value   Timeout for each command run (default: 1s)
   --verbose, -v               Verbose output (default: false)
   --help, -h                  show help
```

## Installation

To check out all available releases visit: https://github.com/wingsuitist/retry/releases/
Follow these instructions to download and install the `retry` command on your local system.

### macOS 

You'll need to download and extract the binary before moving it to your `bin` directory:

```sh
curl -LO https://github.com/wingsuitist/retry/releases/download/v0.0.4/retry_0.0.4_darwin_arm64.tar.gz
tar xzvf retry_0.0.4_darwin_arm64.tar.gz

# to install globally use
sudo mv retry /usr/local/bin/

# to install it only for your user move it to `~/.bin/`
# make sure `~/.bin/` is in your PATH variable
mv retry ~/.bin/
```

### Linux

You'll need to download and extract the binary before moving it to your `bin` directory:

```sh
curl -LO https://github.com/wingsuitist/retry/releases/download/v0.0.4/retry_0.0.4_linux_amd64.tar.gz
tar xzvf retry_0.0.4_linux_amd64.tar.gz
sudo mv retry /usr/local/bin/
```

### Windows

For Windows, you will need a tool that can extract tar.gz archive like 7-Zip. 

1. Download your architecture's `.tar.gz` file from the [releases page](https://github.com/wingsuitist/retry/releases)
2. Extract both the outer `.tar.gz` file and the resulting `.tar` file with 7-Zip.
3. From the extracted files you can move the `retry.exe` to `C:\Windows\`.

### Homebrew (macOS and Linux)

You can install the `retry` command with [Homebrew](https://brew.sh/). Since the formula lives in this repository you need to provide its URL when tapping:

```sh
brew tap wingsuitist/retry https://github.com/wingsuitist/retry
brew install wingsuitist/retry/retry
```

The Homebrew formula is automatically updated by the release workflow whenever a
new version of `retry` is published.

## Development

### Building from Source

```sh
git clone https://github.com/wingsuitist/retry.git
cd retry
go build -o retry ./cmd/
```

### Testing the Release Process

Use the provided test script to verify the GoReleaser configuration:

```sh
./test-release.sh
```

### Creating a New Release

1. Update version information if needed
2. Create and push a new tag:
   ```sh
   git tag v0.0.6
   git push origin v0.0.6
   ```
3. The GitHub Action will automatically:
   - Build cross-platform binaries for Linux, macOS, and Windows
   - Create a GitHub release with downloadable assets
   - Update the Homebrew formula with binary downloads and correct checksums
   - Support both Intel and ARM architectures

This ensures that Linux Homebrew users get pre-compiled binaries instead of having to build from source.
