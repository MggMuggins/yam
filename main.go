package main

import (
    "errors"
    "os"
    "os/exec"
    "path"
    
    "github.com/mikkeloscar/aur"
)

func main() {
    if os.Getuid() == 0 {
        Err("Do not run as root.")
        os.Exit(1)
    }
    
    //TODO: Real Arg parsing
    var pkgname string
    if len(os.Args) != 2 {
        Err("Not enough arguments.")
        os.Exit(1)
    } else {
        pkgname = os.Args[1]
    }
    
    config, err := Config()
    if err != nil {
        Err(err)
        os.Exit(1)
    }
    Info("Configuration:", config)
    
    pkg, err := getPackage(pkgname)
    if err != nil {
        Err(err)
        os.Exit(1)
    }
    Info("Got package:", AUR + pkg.URLPath)
    
    pkgBuildDir := path.Join(config.BuildDir, pkg.Name)
    pkgSrcDir := path.Join(config.SrcDir, pkg.Name)
    if err = os.MkdirAll(pkgSrcDir, 0755); err != nil {
        Err(err)
        os.Exit(1)
    }
    
    if err = getPkgbuild(pkg, config.BuildDir); err != nil {
        Err(err)
        os.Exit(1)
    }
    Info("Got PKGBUILD")
    
    if err = runMakePkg(pkgBuildDir, pkgSrcDir, "-si"); err != nil {
        Err("makepkg:", err)
        os.Exit(1)
    }
}

func getPackage(pkgname string) (pkg aur.Pkg, err error) {
    // Get package, check things
    pkgs, err := aur.Search(pkgname)
    if err != nil { return }
    
    for _, pkgopt := range pkgs {
        if pkgopt.Name == pkgname {
            pkg = pkgopt
            return
        }
    }
    err = errors.New("Package not found: " + pkgname)
    return
}

func runMakePkg(buildDir string, srcDir string, args ...string) error {
    cmd := exec.Command("makepkg", args...)
    cmd.Dir = buildDir
    srcdest := "SRCDEST=" + srcDir
    cmd.Env = []string{srcdest}
    cmd.Stdin, cmd.Stdout, cmd.Stderr = os.Stdin, os.Stdout, os.Stderr
    Info("makepkg", args)
    return cmd.Run()
}
