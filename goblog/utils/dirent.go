package utils

import (
    "os/exec"
    "os"
    "strings"
)

func GetCurrentPath() string {
    s, err := exec.LookPath(os.Args[0])
    if err!=nil {
        return ""
    }
    i := strings.LastIndex(s, "\\")
    path := string(s[0 : i+1])
    return path
}