class Retry < Formula
  desc "The missing retry command"
  homepage "https://github.com/wingsuitist/retry"
  url "https://github.com/wingsuitist/retry/archive/refs/tags/v0.0.7.tar.gz"
  sha256 "fb6a81d07c32456f7da6db36cfca7339ee285e6398ee258b7003c1b971895580"
  license "MIT"
  version "0.0.7"

  # Pre-compiled binaries (preferred for speed)
  on_macos do
    if Hardware::CPU.intel?
      url "https://github.com/wingsuitist/retry/releases/download/v0.0.7/retry_0.0.7_Darwin_x86_64.tar.gz"
      sha256 "4d93a6afc50199b6d761265a524b98cc6611e0bca3a3f99e1b971746e3eadfcd"
    end
    if Hardware::CPU.arm?
      url "https://github.com/wingsuitist/retry/releases/download/v0.0.7/retry_0.0.7_Darwin_arm64.tar.gz"
      sha256 "c5490b4207cb5598d649d34d40649a4ee17bb5ec2fe9a3268a2b0fe29c52fb4b"
    end
  end

  on_linux do
    if Hardware::CPU.intel?
      if Hardware::CPU.is_64_bit?
        url "https://github.com/wingsuitist/retry/releases/download/v0.0.7/retry_0.0.7_Linux_x86_64.tar.gz"
        sha256 "54aedf0b06b8a558f033bae32ef606ccf318d3b82a869e9edd9b20fb1b4b7978"
      else
        url "https://github.com/wingsuitist/retry/releases/download/v0.0.7/retry_0.0.7_Linux_i386.tar.gz"
        sha256 "06a0527e38ab4206e78fbf8aa107a60ad725b75338c16919e1e335e048031747"
      end
    end
    if Hardware::CPU.arm?
      if Hardware::CPU.is_64_bit?
        url "https://github.com/wingsuitist/retry/releases/download/v0.0.7/retry_0.0.7_Linux_arm64.tar.gz"
        sha256 "b0a326b65a1d0f35fe93b2ad51511ac17cc406e80ac8cfea91bd7bbaf908a56b"
      end
    end
  end

  depends_on "go" => :build

  def install
    if build.bottle?
      # Install pre-compiled binary
      bin.install "retry"
    else
      # Build from source
      system "go", "build", *std_go_args(ldflags: "-s -w"), "./cmd"
    end
  end

  test do
    assert_match "command", shell_output("#{bin}/retry -h")
  end
end
