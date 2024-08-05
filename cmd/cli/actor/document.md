HREADME.md > -- ## Install using a package manager ## Install using a package manager
123
2
3 You can install the protocol compiler, protoc, with a package manager under Linux or macOS using th following commands.
45
6
Linux, using apt or apt-get, for example:

````bash
$ apt install -y protobuf-compiler
$protoc ---version # Ensure compiler-version is 3+

MacOS, using Homebrew:

```bash

$ brew install protobuf
I

$protoc --version # Ensure compiler version is 3+
## Install pre-compiled binaries (any OS)
1. Manually download from `github.com/google/protobuf/releases the zip file. bash
$ PB_REL="https://github.com/protocolbuffers/protobuf/releases"

$ curl -LO $PB_REL/download/v25.1/protoc-25.1-osx-aarch_64.zip

unzip protoc-25.1-osx-aarch_64.zip -d $HOME/protoc
www

Unzip the file under $HOME/.local or a directory of your choice. For example: ```bash
$ unzip protoc-25.1-linux-x86_64.zip -d $HOME/.local
Update your environment's path variable to include the path to the protoc executable. For example: ```bash
Go plugins for the protocol compiler:
35
$ export PATH="$PATH:$HOME/.local/bin"
````
