# SPDX-FileCopyrightText: 2026 first-uninteresting-username
#
# SPDX-License-Identifier: GPL-3.0-only
self: {
  config,
  lib,
  pkgs,
  ...
}: let
  inherit (pkgs.stdenv.hostPlatform) system;
  tomlFormat = pkgs.formats.toml {};
in {
  options.programs.hack = {
    enable = lib.mkEnableOption "hack, CLI tool for interacting with LLMs";

    package = lib.mkOption {
      type = lib.types.package;
      default = self.packages.${system}.hack
        or (throw "hack: no package available for system '${system}'");
      defaultText = lib.literalExpression "self.packages.\${pkgs.stdenv.hostPlatform.system}.hack";
      description = "The hack package to install.";
    };
    settings = lib.mkOption {
      type = lib.types.submodule {
        freeformType = tomlFormat.type;

        options = {
          model = lib.mkOption {
            type = lib.types.nullOr lib.types.str;
            default = null;
            example = "anthropic/claude-fable-5";
            description = "Model to be used if no overwrite was provided";
          };

          base_url = lib.mkOption {
            type = lib.types.nullOr lib.types.str;
            default = null;
            example = "https://ai.hackclub.com/proxy/v1";
            description = "API base to be used if no overwrite was provided";
          };

          api_key_path = lib.mkOption {
            type = lib.types.nullOr lib.types.str;
            default = null;
            example = "\${config.sops.secrets.HACK_CLUB_AI_API_KEY.path}";
            description = ''
              Runtime path to a file containing the API key, e.g. a sops
              secret. Str, not path, so it is not resolved to the Nix store.
            '';
          };
        };
      };

      default = {};
      description = "Configuration written to {file}`config.toml`.";
    };
  };
}
