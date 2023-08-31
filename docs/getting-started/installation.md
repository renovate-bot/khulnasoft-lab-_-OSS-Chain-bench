# Installation

## Nix/NixOS

Direct issues installing `oss-chain-bench` via `nix` through the channels mentioned [here](https://nixos.wiki/wiki/Support)

You can use `nix` on Linux or macOS and on other platforms unofficially.

`nix-env --install -A nixpkgs.oss-chain-bench`

Or through your configuration as usual

NixOS:

```nix
  # your other config ...
  environment.systemPackages = with pkgs; [
    # your other packages ...
    oss-chain-bench
  ];
```

home-manager:

```nix
  # your other config ...
  home.packages = with pkgs; [
    # your other packages ...
    oss-chain-bench
  ];
```

## Binary

Download the archive file for your operating system/architecture from [here](https://github.com/khulnasoft-lab/oss-chain-bench/releases/latest).
<!-- TODO: swap to GH pages [here](https://github.com/khulnasoft-lab/oss-chain-bench/releases/tag/{{ git.tag }}). -->
Unpack the archive, and put the binary somewhere in your `$PATH` (on UNIX-y systems, `/usr/local/bin` or the like).
Make sure it has execution bits turned on.

## From source

```bash
mkdir -p $GOPATH/src/github.com/khulnasoft-lab
cd $GOPATH/src/github.com/khulnasoft-lab
git clone --depth 1 https://github.com/khulnasoft-lab/oss-chain-bench
cd oss-chain-bench/cmd/oss-chain-bench/
export GO111MODULE=on
go install
```
<!-- TODO: swap to GH pages git clone --depth 1 --branch {{ git.tag }} https://github.com/khulnasoft-lab/oss-chain-bench -->

## From source with `go install`

With a sufficient version of `go` you can install and build with `go install github.com/khulnasoft-lab/oss-chain-bench/cmd/oss-chain-bench@latest`
<!-- TODO: swap to GH pages `go install github.com/khulnasoft-lab/oss-chain-bench/cmd/oss-chain-bench@{{ git.tag }}` -->

## Docker

### Docker Hub

```bash
docker pull khulnasoft/oss-chain-bench:latest
```
<!-- TODO: swap to GH pages {{ git.tag[1:] }} -->

Example:

  ```bash
  docker run --rm khulnasoft/oss-chain-bench:latest scan --repository-url <REPOSITORY_URL> --access-token <TOKEN>
  ```
