{
  description = "Fyne Practice";

  inputs.nixpkgs.url = "github:nixos/nixpkgs/nixos-unstable";

  outputs =
    {
      self,
      nixpkgs,
    }:
    let
      system = "x86_64-linux";
      pkgs = nixpkgs.legacyPackages.${system};
    in
    {
      devShells.${system}.default = pkgs.mkShell {
        packages = with pkgs; [
          go # runtime
          gopls # lsp
          golangci-lint # lint
          gofumpt # fmt

          # Fyne components
          libGL
          pkg-config
          libx11.dev
          libxcursor
          libxi
          libxinerama
          libxrandr
          libxxf86vm
          libxkbcommon
          wayland

          #make
          gnumake
          checkmake
          mbake
        ];
      };
    };
}
