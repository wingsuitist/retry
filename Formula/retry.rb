class Retry < Formula
  desc "The missing retry command"
  homepage "https://github.com/wingsuitist/retry"
  version "0.0.6"
  license "MIT"

  on_macos do
    if Hardware::CPU.intel?
      url "https://github.com/wingsuitist/retry/releases/download/v0.0.6/retry_0.0.6_Darwin_x86_64.tar.gz"
      sha256 "c8c775135ccbbeda36b9652b6fcb594b508396fd8b8320e4265f118b7d07ce8c"
    end
    if Hardware::CPU.arm?
      url "https://github.com/wingsuitist/retry/releases/download/v0.0.6/retry_0.0.6_Darwin_arm64.tar.gz"
      sha256 "d4db7e7199eb99b76d482efdfd5d18e0aee4e3a75f151ecdfe9ef3f0a871d963"
    end
  end

  on_linux do
    if Hardware::CPU.intel?
      if Hardware::CPU.is_64_bit?
        url "https://github.com/wingsuitist/retry/releases/download/v0.0.6/retry_0.0.6_Linux_x86_64.tar.gz"
        sha256 "68f075fdade4d35c2caeed328a4e0aea7e51a80a19c3208cce5a939cad6017fd"
      else
        url "https://github.com/wingsuitist/retry/releases/download/v0.0.6/retry_0.0.6_Linux_i386.tar.gz"
        sha256 "a87a095ca8214b6aae270f9fbfcaa7551c7d251993b04742b27fdad0823d2d9f"
      end
    end
    if Hardware::CPU.arm?
      if Hardware::CPU.is_64_bit?
        url "https://github.com/wingsuitist/retry/releases/download/v0.0.6/retry_0.0.6_Linux_arm64.tar.gz"
        sha256 "cd5c8992b0b9aff7557d62c327668cf7d21f0daeb35af154a471e0e53a61b05a"
      end
    end
  end

  def install
    bin.install "retry"
  end

  test do
    assert_match "command", shell_output("#{bin}/retry -h")
  end
end
