package logging

import "testing"

func TestConfig(t *testing.T) {
    logger1 := ComponentLogger("aaaa")
    logger2 := SubComponentLogger(logger1, "bbbb")

    logger1.Printf("1111")
    logger2.Printf("2222")


}
