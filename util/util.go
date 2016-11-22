package util

import (
    "os"
    "strconv"
    "errors"
)

func String2FileMode(val string) (os.FileMode, error) {
    val2, err := strconv.ParseUint(val, 0, 32)

    if err != nil {
        return os.FileMode(0), err
    }

    val3 := os.FileMode(val2)
    return os.FileMode(val3), nil
}

func Error2String(err error) string {
    if err == nil {
        return ""
    } else {
        return err.Error()
    }
}

func String2Error(err string) error {
    if err == "" {
        return nil
    } else {
        return errors.New(err)
    }
}
