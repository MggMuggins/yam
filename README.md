## Yam
Yet another AUR helper but without the pacman wrapping.

This program is build on the idea that a source package manager and a binary package manager are two different things and operations that one can do with one inherently do not apply to the other.

## Goals
 - Provide easy access and installation of packages from the AUR
 - Minimal user interaction. Execptions:
    - Use of sudo
    - Edit PKGBUILD?
 - Use of sudo for root actions. Do as little as root as is humanly possible.
 - Be stateless
 - Make efficient use of network
 - Look nice! (and be consistent with pacman's look 'n feel)

## Non-Goals
 - Wrap Pacman. All actions will only apply to the AUR

## Notes
In order to more efficiently use bandwidth and disk space, yam downloads package sources into `$XDG_CACHE_HOME` (or "$HOME/.cache"), but does builds in `/tmp`, to save space.

## Configuration
A configuration can be found in `$XDG_CONFIG_HOME/yam.toml` or `$HOME/.config/yam.toml`. If your XDG variable is set, the location in HOME will not be used. Paths cannot include environment variables. Standard Configuration:
```Toml
# This configuration is redundant because it specifies default values

# This is the directory where makepkg will download and save sources
SrcDir = "/home/user/.cache/yam"
# This is where actual builds will take place
BuildDir = "/home/user/.cache/yam"
```
