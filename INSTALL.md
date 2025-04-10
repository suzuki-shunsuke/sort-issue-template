# Install

sort-issue-template is written in Go. So you only have to install a binary in your `PATH`.

There are some ways to install sort-issue-template.

1. [Homebrew](#homebrew)
1. [aqua](#aqua)
1. [GitHub Releases](#github-releases)
1. [Build an executable binary from source code yourself using Go](#build-an-executable-binary-from-source-code-yourself-using-go)

## Homebrew

You can install sort-issue-template using [Homebrew](https://brew.sh/).

```sh
brew install suzuki-shunsuke/sort-issue-template/sort-issue-template
```

## aqua

[aqua-registry >= v4.272.0](https://github.com/aquaproj/aqua-registry/releases/tag/v4.272.0)

You can install sort-issue-template using [aqua](https://aquaproj.github.io/).

```sh
aqua g -i suzuki-shunsuke/sort-issue-template
```

## Build an executable binary from source code yourself using Go

```sh
go install github.com/suzuki-shunsuke/sort-issue-template/cmd/sort-issue-template@latest
```

## GitHub Releases

You can download an asset from [GitHub Releases](https://github.com/suzuki-shunsuke/sort-issue-template/releases).
Please unarchive it and install a pre built binary into `$PATH`. 

### Verify downloaded assets from GitHub Releases

You can verify downloaded assets using some tools.

1. [GitHub CLI](https://cli.github.com/)
1. [slsa-verifier](https://github.com/slsa-framework/slsa-verifier)
1. [Cosign](https://github.com/sigstore/cosign)

### 1. GitHub CLI

You can install GitHub CLI by aqua.

```sh
aqua g -i cli/cli
```

```sh
version=v0.1.0
asset=sort-issue-template_darwin_arm64.tar.gz
gh release download -R suzuki-shunsuke/sort-issue-template "$version" -p "$asset"
gh attestation verify "$asset" \
  -R suzuki-shunsuke/sort-issue-template \
  --signer-workflow suzuki-shunsuke/go-release-workflow/.github/workflows/release.yaml
```

### 2. slsa-verifier

You can install slsa-verifier by aqua.

```sh
aqua g -i slsa-framework/slsa-verifier
```

```sh
version=v0.1.0
asset=sort-issue-template_darwin_arm64.tar.gz
gh release download -R suzuki-shunsuke/sort-issue-template "$version" -p "$asset" -p multiple.intoto.jsonl
slsa-verifier verify-artifact "$asset" \
  --provenance-path multiple.intoto.jsonl \
  --source-uri github.com/suzuki-shunsuke/sort-issue-template \
  --source-tag "$version"
```

### 3. Cosign

You can install Cosign by aqua.

```sh
aqua g -i sigstore/cosign
```

```sh
version=v0.1.0
checksum_file="sort-issue-template_${version#v}_checksums.txt"
asset=sort-issue-template_darwin_arm64.tar.gz
gh release download "$version" \
  -R suzuki-shunsuke/sort-issue-template \
  -p "$asset" \
  -p "$checksum_file" \
  -p "${checksum_file}.pem" \
  -p "${checksum_file}.sig"
cosign verify-blob \
  --signature "${checksum_file}.sig" \
  --certificate "${checksum_file}.pem" \
  --certificate-identity-regexp 'https://github\.com/suzuki-shunsuke/go-release-workflow/\.github/workflows/release\.yaml@.*' \
  --certificate-oidc-issuer "https://token.actions.githubusercontent.com" \
  "$checksum_file"
cat "$checksum_file" | sha256sum -c --ignore-missing
```
