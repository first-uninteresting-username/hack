# SPDX-FileCopyrightText: 2026 first-uninteresting-username
#
# SPDX-License-Identifier: GPL-3.0-only
{
  lib,
  buildGoModule,
  fetchFromGitHub,
  nix-update-script,
}:
buildGoModule (finalAttrs: {
  pname = "hack";
  version = "5";
  structuredAttrs = true;

  src = fetchFromGitHub {
    owner = "first-uninteresting-username";
    repo = "hack";
    tag = finalAttrs.version;
    hash = "sha256-h9qF3IYGXDW7yeKlc40NV4Be9VdYBzHJ8eZ/EwWDm0Q=";
  };

  vendorHash = "sha256-gTj1xJwj/Qf+v6wY5FEheWxFbHQ4NEV68FflgGbqnDc=";

  ldflags = ["-s"];

  env = {
     CGO_ENABLED = 0;
  };

  passthru.updateScript = nix-update-script {
    extraArgs = [ "--flake" ];
  };

  meta = {
    description = "CLI tool for interacting with LLMs";
    homepage = "https://github.com/first-uninteresting-username/hack";
    license = lib.licenses.gpl3Only;
    # maintainers = with lib.maintainers; [ ];
    mainProgram = "hack";
  };
})
