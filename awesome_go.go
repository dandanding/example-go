package awesome

import "github.com/codecov/example-go/awesomeness"

var result string

func Setup() {
    result = awesome.Smile()
}

func GetResult() string {
    return result
}
