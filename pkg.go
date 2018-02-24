package main

import (
    "io"
    "errors"
    "net/http"
    "os"
    "os/exec"
    "path"
    
    "github.com/mikkeloscar/aur"
)

// Unpacks a tarball into it's current directory
func unpackTarball(ball string) error {
    dir, _ := path.Split(ball)
    err := exec.Command("tar", "-C", dir, "-xvf", ball).Run()
    if err != nil { return errors.New("tar: " + err.Error()) }
    os.Remove(ball)
    return nil
}

// Must provide a filename for the file!
// ^^ Bad API ^^
func downloadFile(url, dest string) error {
    outfile, err := os.Create(dest)
    
    if err != nil {
        return err
    }
    defer outfile.Close()
    
    resp, err := http.Get(url)
    if err != nil {
        return err
    }
    defer resp.Body.Close()
    
    _, err = io.Copy(outfile, resp.Body)
    return err
}

// Download and unpack a PKGBUILD tarball into dest
// (Should have a PKGBUILD in dest when complete)
func GetPkgbuild(pkg aur.Pkg, dest string) (err error) {
    tarball_name := path.Join(dest, pkg.Name + ".tar.gz")
    err = downloadFile(AUR + pkg.URLPath, tarball_name)
    if err != nil { return }
    return unpackTarball(tarball_name)
}

func GetPackage(pkgname string) (pkg aur.Pkg, err error) {
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
