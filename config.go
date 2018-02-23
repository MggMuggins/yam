package main

import (
    "io/ioutil"
    "path"
    "os"
    
    "github.com/BurntSushi/toml"
)

const (
    AUR = "https://aur.archlinux.org"
)

//TODO: Configuration file
type Configuration struct {
    // Set $SRCDEST to this before running makepkg
    SrcDir   string
    BuildDir string
}

// The default configuration location for yam is $XDG_CONFIG_HOME/yam.toml
//  (preferred) or $HOME/.config/yam.toml.
func configFromFile() (self Configuration, err error) {
    configfile := path.Join(xdgDir("XDG_CONFIG_HOME", ".config"), "yam.toml")
    Info("Config file location:", configfile)
    
    //TODO: Complain if there is a config file but we can't access...
    if _, statErr := os.Stat(configfile); statErr == nil {
        _, err = toml.DecodeFile(configfile, &self)
    }
    return
}

// Get a configuration for yam. This returns a default config
//  plus whatever values where specified in the config file
func Config() (self Configuration, err error) {
    self, err = configFromFile()
    if err != nil { return }
    err = self.fillDefault()
    return
}

// Fills all values in the Configuration sturct with their defaults
//  if they have not been changed from their zero values
func (self *Configuration) fillDefault() (err error) {
    if self.SrcDir == "" {
        srcdir := path.Join(xdgDir("XDG_CACHE_HOME", ".cache"), "yam")
        if err = os.MkdirAll(self.SrcDir, 0755); err != nil { return }
        self.SrcDir = srcdir
    }
    
    if self.BuildDir == "" {
        // nasty workaround
        var tmpdir string
        tmpdir, err = ioutil.TempDir("", "yam")
        if err != nil { return }
        self.BuildDir = tmpdir
    }
    return
}

// Get an XDG dir. If the variable envV is not set, suffix
//  path.Join'd with $HOME is used instead
func xdgDir(envV, suffix string) (dir string) {
    xdg := os.Getenv(envV)
    if xdg == "" {
        dir = path.Join(os.Getenv("HOME"), suffix)
    } else {
        dir = xdg
    }
    return
}
