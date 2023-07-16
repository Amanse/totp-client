{
  description = "A very basic flake";

  inputs.nixpkgs.url = "github:nixos/nixpkgs/nixos-unstable";

  outputs = {self}: {
    defaultShell = import ./shell.nix;
  };
}
