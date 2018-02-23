package main

import (
    "io/ioutil"
    "path"
    "os"
)

const (
    AUR = "https://aur.archlinux.org"
)

//TODO: Configuration file
type Configuration struct {
    // Set $SRCDEST to this before running makepkg
    SrcDir string
    BuildDir string
}

// This guy creates a configuration based on default
//  values. It also makes sure that the directories SrcDir
//  and BuildDir exist.
func defaultConfig() (self Configuration, err error) {
    // Get the appropriate system config directory
    var cachedir string
    xdg_cache := os.Getenv("XDG_CACHE_HOME")
    if xdg_cache == "" {
        cachedir = path.Join(os.Getenv("HOME"), ".cache")
    } else {
        cachedir = xdg_cache
    }
    
    self.SrcDir = path.Join(cachedir, "yam")
    if err = os.MkdirAll(self.SrcDir, 0755); err != nil { return }
    
    self.BuildDir, err = ioutil.TempDir("", "yam")
    return
}
