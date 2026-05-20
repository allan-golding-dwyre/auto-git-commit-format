OS=$(uname -s | tr '[:upper:]' '[:lower:]')
ARCH=$(uname -m | sed 's/x86_64/amd64/;s/aarch64/arm64/')
VERSION=$(curl -s https://api.github.com/repos/allan-golding-dwyre/agcf/releases/latest | grep tag_name | cut -d'"' -f4)

curl -sL "https://github.com/allan-golding-dwyre/agcf/releases/download/${VERSION}/agcf_${OS}_${ARCH}.tar.gz" | tar xz
mv agcf ~/.local/bin/agcf
echo "✅ agcf ${VERSION} installé"
