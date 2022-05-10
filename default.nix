{ pkgs ? import <nixpkgs-unstable> { } }:

pkgs.mkShell {
  NIX_SHELL = "terraform-provider-meshapi";
  shellHook = ''
    echo starting terraform-provider-meshapi dev shell
  '';

  buildInputs = [
    pkgs.go_1_16

    pkgs.terraform

    # script dependencies
    pkgs.just
  ];
}
