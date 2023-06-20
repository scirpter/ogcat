# ogcat
Checks every single Discord username combination possible, written in Go.

## Installation
Run the install_deps.bat to install the following tools:
- [Syso](https://github.com/hallazzang/syso)
- [Garble](https://github.com/burrowers/garble)

## Build
Run the build.bat to build the project. This will produce a OGCAT.exe file in the `./bin` directory.

## Usage
Set your tokens in the `./bin/config.json` (will be initialized once the .exe was run once) and start the .exe. The rest should be self-explanatory. If you add multiple tokens, the usernames to check will be split up equally between the tokens, so the process can be sped up.