class Retry < Formula
  desc "The missing retry command"
  homepage "https://github.com/wingsuitist/retry"
  url "https://github.com/wingsuitist/retry/archive/refs/tags/v0.0.5.tar.gz"
  sha256 "09042c64d36f33985d31d6b9fab93688cf13272fa15c9ba66b91f1723b499e23"
  license "MIT"

  depends_on "go" => :build

  def install
    system "go", "build", *std_go_args(ldflags: "-s -w", output: bin/"retry"), "./cmd"
  end

  test do
    assert_match "command", shell_output("#{bin}/retry -h")
  end
end
