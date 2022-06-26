{ pkgs ? import <nixpkgs> {} }:
  pkgs.mkShell {
    # nativeBuildInputs is usually what you want -- tools you need to run
    nativeBuildInputs = [ pkgs.nodejs pkgs.redis  ];
    SERVERL_URL="http://localhost:8080";
}
