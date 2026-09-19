# SPDX-FileCopyrightText: 2026 first-uninteresting-username
#
# SPDX-License-Identifier: GPL-3.0-only
self: {
  config,
  lib,
  pkgs,
  ...
}: let
  cfg = config.programs.hack;
  tomlFormat = pkgs.formats.toml {};
in {
  imports = [(import ./module.nix self)];

  config = lib.mkIf cfg.enable {
    home.packages = [cfg.package];
    xdg.configFile."hack/config.toml".source = tomlFormat.generate "config.toml" cfg.settings;
  };
}
