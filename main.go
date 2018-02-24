package main

import (
    "os"
    "os/exec"
    "path"
    
    "github.com/mikkeloscar/aur"
    "github.com/jessevdk/go-flags"
)

var opts struct {
    Update bool `short:"u" long:"update" description:"Update all installed AUR packages"`
    Verbose bool `short:"v" long:"verbose" description:"Print more Debugging info"`
}

func main() {
    if os.Getuid() == 0 {
        ErrLog.Println("Do not run as root.")
        os.Exit(1)
    }
    
    pkgnames, err := flags.Parse(&opts)
    
    config, err := Config()
    if err != nil {
        ErrLog.Println(err)
        os.Exit(1)
    }
    if opts.Verbose { InfoLog.Printf("Configuration: %s\n", config) }
    
    // Get the packages first, that way we verify that everything is correct
    //  before we actually start builds/installs
    pkgs := []aur.Pkg{}
    for _, pkgname := range pkgnames {
        pkg, err := GetPackage(pkgname)
        if err != nil {
            ErrLog.Println(err)
            os.Exit(1)
        }
        pkgs = append(pkgs, pkg)
        if opts.Verbose { InfoLog.Printf("Got package: %s%s\n", AUR, pkg.URLPath) }
    }
    
    InfoLog.Println("Packages to build and install:")
    PrintPkgs(pkgs)
    
    for _, pkg := range pkgs {
        
        pkgBuildDir := path.Join(config.BuildDir, pkg.Name)
        pkgSrcDir := path.Join(config.SrcDir, pkg.Name)
        if err = os.MkdirAll(pkgSrcDir, 0755); err != nil {
            ErrLog.Println(err)
            os.Exit(1)
        }
        
        if err = GetPkgbuild(pkg, config.BuildDir); err != nil {
            ErrLog.Println(err)
            os.Exit(1)
        }
        
        if err = runMakePkg(pkgBuildDir, pkgSrcDir, "-si"); err != nil {
            ErrLog.Printf("makepkg: %s\n", err)
            os.Exit(1)
        }
    }
}

func runMakePkg(buildDir string, srcDir string, args ...string) error {
    cmd := exec.Command("makepkg", args...)
    cmd.Dir = buildDir
    srcdest := "SRCDEST=" + srcDir
    cmd.Env = []string{srcdest}
    cmd.Stdin, cmd.Stdout, cmd.Stderr = os.Stdin, os.Stdout, os.Stderr
    // Want to make sure this is printed
    InfoLog.Printf("makepkg %s", args)
    return cmd.Run()
}
