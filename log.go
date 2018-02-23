package main

import (
    "fmt"
    "github.com/mikkeloscar/aur"
    . "github.com/logrusorgru/aurora"
)

// Logging helper
var (
    err_p = Bold(Red("::"))
    warn_p = Bold(Brown("::"))
    info_p = Bold(Blue("::"))
    ok_p = Bold(Green("::"))
)

func Err(msgs ...interface{}) (int, error) {
    return fmt.Printf("%s%s\n", combine(err_p, msgs)...)
}

func Warn(msgs ...interface{}) (int, error) {
    return fmt.Printf("%s%s\n", combine(warn_p, msgs)...)
}

func Info(msgs ...interface{}) (int, error) {
    return fmt.Printf("%s%s\n", combine(info_p, msgs)...)
}

func Ok(msgs ...interface{}) (int, error) {
    return fmt.Printf("%s%s\n", combine(ok_p, msgs)...)
}

func PrintPkgs(pkgs []aur.Pkg) (bytes int, err error) {
    for _, pkg := range pkgs {
        writ, err_ := fmt.Printf("\t%s %s\n", pkg.Name, pkg.Version)
        bytes += writ
        err = err_
        if err != nil {
            return
        }
    }
    return
}

func combine(pref interface{}, msgs []interface{}) []interface{} {
    return append([]interface{}{pref}, msgs)
}
