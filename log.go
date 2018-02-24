package main

import (
    "fmt"
    "log"
    "os"
    
    "github.com/mikkeloscar/aur"
    . "github.com/logrusorgru/aurora"
)

// Prefixes
var (
    bold_p = Bold(":: ")
    err_p  = bold_p.Red()
    warn_p = bold_p.Brown()
    info_p = bold_p.Blue()
    ok_p   = bold_p.Green()
)

var (
    ErrLog  = log.New(os.Stderr, err_p .String(), 0)
    WarnLog = log.New(os.Stdout, warn_p.String(), 0)
    InfoLog = log.New(os.Stdout, info_p.String(), 0)
    OkLog   = log.New(os.Stdout, ok_p  .String(), 0)
)

func PrintPkgs(pkgs []aur.Pkg) (bytes int, err error) {
    for _, pkg := range pkgs {
        writ, err_ := fmt.Printf("\t%s %s\n", pkg.Name, pkg.Version)
        bytes += writ
        err = err_
        if err != nil {
            return
        }
    }
    fmt.Printf("\n")
    return
}
