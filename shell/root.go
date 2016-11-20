package shell

import (
    "os"
    "errors"
    "strings"
    "os/exec"
    "fmt"
)

func Evaluate(command string) (string, error) {
    shell := os.Getenv("SHELL")

    if shell == "" {
        return "", errors.New("Can't determine shell")
    }

    if strings.Contains(shell, "bash") {
        out, err := exec.Command(shell, "-c", command).Output()

        if err != nil {
            return "", err
        }

        return string(out), nil
    } else if strings.Contains(shell, "fish") {
        out, err := exec.Command(shell, "-c", command).Output()

        if err != nil {
            return "", err
        }

        return string(out), nil
    }

    return "", errors.New(fmt.Sprintf("unknown shell: %s", shell))
}
