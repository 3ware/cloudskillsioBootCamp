# Notes from IaC Labs

## Disable the go extension in VsCode (for now)

This causes problems because go is changing from GOPATH mode to module mode. You can do both in 16, but it's being disabled in 17.

This is the first I've touch go, so not really sure what this means. See this blog: <https://blog.golang.org/go116-module-changes>

You also need to set the mode to auto or the test will moan that it cannot load the main module.

'''go env -w GO111MODULE=auto'''
