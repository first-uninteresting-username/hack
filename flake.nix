# SPDX-FileCopyrightText: 2026 first-uninteresting-username
#
# SPDX-License-Identifier: GPL-3.0-only
{
  description = "Flake for hack";

  inputs = {
    flake-parts.url = "github:hercules-ci/flake-parts";
    nixpkgs.url = "github:NixOS/nixpkgs/nixos-unstable";
  };

  outputs = inputs @ {
    flake-parts,
    self,
    ...
  }:
    flake-parts.lib.mkFlake {inherit inputs;} {
      systems = ["x86_64-linux" "aarch64-linux" "aarch64-darwin"];

      perSystem = {
        self',
        pkgs,
        ...
      }: {
        packages = {
          hack = pkgs.callPackage ./nix/package.nix {};
          default = self'.packages.hack;
        };
      };

      flake = {
        nixosModules.default = import ./nix/nixos.nix self;
        homeManagerModules.default = import ./nix/home-manager.nix self;
      };
    };
}
