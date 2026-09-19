# hack

Dead-simple, UNIX style tool for interacting with LLMs on command line.

It is intended for Hackclub AI API, but there's nothing stopping you from using it with any other Open AI completions compatible API.

![Hackatime Badge](https://hackatime.hackclub.com/api/v1/badge/U0A9Y38B28H/first-uninteresting-username/hack-ai-cli)

The tech stack consists of 5 pieces:

- [Go](https://go.dev/), the underlying language behind the project
- [Cobra](https://github.com/spf13/cobra), Go library for creating CLI tools
- [Go-OpenAI](https://github.com/sashabaranov/go-openai), Go library for making requests to OpenAI compatible APIs
- [Nix](https://nixos.org/), language providing reproductible builds and software distribution
- [Devenv](https://devenv.sh/), nix based dev environment solution

Notes and roadmap are available in [NOTES.md](/NOTES.md)

## Installation and usage

### Installation

```bash
# Clone this repo
gh repo clone first-uninteresting-username/hack && cd hack
# Build the binary
go build
```

Alternatively, get the binary for your system from [github releases](https://github.com/first-uninteresting-username/hack/releases)

In both cases, put the binary somewhere on your path and create the config in your preffered location.

There's also a Nix package available. Add:

```nix
hack.url = "github:first-uninteresting-username/hack";
```

to your flake inputs.

Then you could either that package to your package list (either home manager or system):

```nix
home.packages = [
  inputs.hack.packages.${pkgs.stdenv.hostPlatform.system}.hack
];
```

or use the module:

```nix
# Add that to your nixos imports
inputs.hack.nixosModules.default
# Or this to your home-manager imports
inputs.hack.homeManagerModules.default
# And this to your nixos/home manager configuration
# The syntax is the same for both
programs.hack = {
  enable = true;
  settings = {
    # Example values
    base_url = "https://ai.hackclub.com/proxy/v1";
    model = "deepseek/deepseek-v4-pro";
    api_key_path = config.sops.secrets.HACK_CLUB_AI_API_KEY.path;
  };
};
```

If you want to, there's a cachix cache available:

```nix
nix.settings = {
  extra-substituters = ["https://dirtree-db.cachix.org"];
  extra-trusted-public-keys = [
    "dirtree-db.cachix.org-1:geR/eeJBzFUNhj3mwjHm1EK/mzXIG/PF3Bg48YlF1ys="
  ];
};
```

### Configuration

Create a file named `config.toml` either in `$HOME/.config/hack` or `/etc/hack`.
The user dir takes precedence over system dir.

File contents (example values):

```toml
# Base URL (without /chat/completions)
base_url = "https://ai.hackclub.com/proxy/v1"
# Full name of the model in your provider
model = "deepseek/deepseek-v4-pro"
# Your API key
api_key = "sk-xxx"
# Path to a file containing your API key. It must be readable by the user you will be running hack with
api_key_path = "/run/secrets/HACKCLUB_AI_API_KEY"
```

Example config is available as [config.toml.example](config.toml.example) in the root of this repo

### Usage

```bash
hack --help

Interact with LLMs from the command line

hack is a simple tool for interacting with LLMs.
It is made to be scriptable, extensible and easy to use.
There're no agentic capabilities built in,
but because of how it works, it's possible to create an agent based on it.

Example usage:
        hack -p "your prompt here"
        echo "some content" | hack -p "summarize this"
        ls -la | hack -sp "delete the largest file"

Modes:
        shell   (-s/--shell)    Generate shell commands from a prompt
        code    (-w/--write)    Output executable code (jq, python3, bash, or POSIX sh)
        normal                  Standard prompt-and-response

Usage:
  hack [flags]

Flags:
  -b, --base string       base URL for the API (without /chat/completions)
  -c, --config string     config file path, $HOME/.config/hack-ai/config.toml if not provided
  -h, --help              help for hack
  -k, --key string        API key for selected provider
  -f, --key-file string   Path to file containing API key
  -m, --model string      LLM used for response
  -p, --prompt string     prompt for the LLM
  -s, --shell             enable command mode
  -v, --version           version for hack
  -w, --write             enable code mode
```
