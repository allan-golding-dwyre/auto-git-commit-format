OS=$(uname -s | tr '[:upper:]' '[:lower:]')
ARCH=$(uname -m | sed 's/x86_64/amd64/;s/aarch64/arm64/')

# Fetch the latest version tag
VERSION=$(curl -s "https://api.github.com/repos/allan-golding-dwyre/auto-git-commit-format/releases/latest" | grep -oP '"tag_name": "\K[^"]+')

# Fallback check if no release is found
if [ -z "$VERSION" ]; then
  echo "❌ Error: Could not fetch the latest version. Ensure the repo has a published release."
  exit 1
fi

# Corrected GitHub download URL (removed /repos)
DOWNLOAD_URL="https://github.com/allan-golding-dwyre/auto-git-commit-format/releases/download/${VERSION}/agcf_${OS}_${ARCH}.tar.gz"

echo "Downloading from: $DOWNLOAD_URL"

# Ensure the destination directory exists
mkdir -p ~/.local/bin

# Download and extract
curl -sL "$DOWNLOAD_URL" | tar -xz -C ~/.local/bin

echo "✅ agcf ${VERSION} installé dans ~/.local/bin"
