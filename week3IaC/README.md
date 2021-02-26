# Notes from IaC Labs

During the first attempt of this lab there were lots of errors.

![GOPATHerrors](resources/gopathErrors2.png)

This is due to the fact that, as of version 1.16 (the version I recently installed to complete this lab), go builds packages in **module-aware** mode by default. The errors I was seeing all related **GOPATH** mode. See this blog for more info: <https://blog.golang.org/go116-module-changes>

I wanted to complete the lab first and foremost, so after some poking around I followed the steps in the section below to resolve the errors and complete the lab.

## GOPATH mode

To use **GOPATH** mode, disable the VsCode extension.

You also need to set the mode to auto, as described in the blog above, or the test will moan that it cannot load the main module. Run this terminal command:

```console
go env -w GO111MODULE=auto
```

## Converting from GOPATH to module-aware mode

After completing the lab in, I wanted to get to the bottom of **module-aware** mode, because it's the only way to use go from the next release.

I'm not a go expert by any means, but, after reading [here](https://terratest.gruntwork.io/docs/getting-started/quick-start/) and referring go help:

```console
go help mod
```

I felt I could give this a *go*. I followed the terratest quick-start guide and ran the following command, in my project's root directory (the folder I have open in VsCode), to create a go.mod file for my testing module.

```console
go mod init "github.com/3ware/cloudskillsioBootCamp"
```

The errors about not being able to import terratest resurfaced. To resolve that I ran:

```console
go get "github.com/gruntwork-io/terratest/modules/terraform"
```

which added a **require** statement to my go.mod file. It also created a go.sum file.

And that seems to work, I ran the test to make sure and it worked as expected.
