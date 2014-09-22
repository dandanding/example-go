package awesome

import (
    "os"
    "testing"
    "github.com/onsi/ginkgo/reporters"
)

func TestAwesome(t *testing.T) {
    RegisterFailHandler(Fail)
    junitReporter := reporters.NewJUnitReporter(os.Getenv("CI_REPORT"))
    RunSpecsWithDefaultAndCustomReporters(t, "Awesome Suite", []Reporter{junitReporter})
}
