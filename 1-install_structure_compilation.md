# Installation, Basic Compilation and Structure

In this first instalment, in order to get up and running quickly for the following sessions, 
we'll cover the installation of Go, some information around structuring projects and some basic compilation.

It might seem like an odd place to start, but the purpose is to make you feel comfortable that this is super easy.

## Installation - Standard

Simply visit the Go downloads page at https://go.dev/dl/ and download the version for your operating system.
Follow the installation instructions which are super simple, so I won't replicate them here.

Once installed, to check it's ok, run `go version` and you should see something similar to the following, 
depending on your operating system:

```sh
$ go version
go version go1.18.1 darwin/arm64
```

## Installation - Docker

You can instead use Docker if you have it installed and running and don't wish to install Go just yet. This method
would also allow you to test your code in different versions.

```
$ docker run --rm -it --name go golang:latest
root@a8b3162813d8:/go# go version
go version go1.18.1 linux/arm64
```

You can also mount your local drive as a volume using the `-v` option so you can work on your code without losing it when the container exits:

```
$ docker run --rm -it -v $PWD:/go/code --name go golang:latest 
root@a8b3162813d8:/go# go version
go version go1.17.6 linux/arm64
```

You can use whatever folder structure you like for your volume now that [modules](https://go.dev/blog/using-go-modules) are fully supported in go, they don't have to be in the "`$GOPATH`".
I won't go into this any further, but if it's worth a separate session, please let me know.

## Basic Compilation

To get started we'll use the canonical hello world application as an example to show how simple and fast compilation is, to get it out of the way now and prove it's neither time consuming nor scary.  We'll most likely cover it in more detail in future sessions when we get into cross-compilation and such like.

You can use any editor to create the following code in a directory of your choice.  

> As a side note, if you're using Visual Studio Code with the Go extensions, you can use the `pkgm` shortcut to create this basic structure -- it will also auto add any imports you need too.

Typically, as the main entry point into your application, the `main` package will be in a file called `main.go`:

```go
package main

import "fmt"

func main() {
	fmt.Println("Hello, world!")
}
```

So a couple of things to point out here are: 

* The `main` package is the main entry point into a binary.  If you're writing a library, this might be `package mysuperlibary` instead;
* The `fmt` import is part of the "standard library" which you'll hear a lot about with Go.  It's very comprehensive and contains a lot of what you need to get things done;
*  You can have multiple imports. Third-party imports typically use the url from github or similar.

We'll cover a lot more about this later.  For now, simply run:

```sh
$ go run main.go
Hello, world!
$
```

This has compiled and executed your binary in one command.  It avoids the need to compile and output a binary before executing it.  You'll type this a lot.  If your code is split across multiple files you must specify them all, but there are easy ways to do that, we'll cover it more in the future.

If you want to create a binary which you can run, you can simply use `go build`, but before you do that, we'll need to let Go know that this is a "module".  These days most Go code is written in [modules](https://go.dev/blog/using-go-modules), so standard workflow is to create a directory and run the `go mod init` command to initialise the directory as a go module before writing any code. 

The `go mod init` command creates a `go.mod` file which keeps track of the name and any dependencies for your application.  You run the command specifying the name of the module.  This can be anything you like, but typically is the url of the repository you'll be keeping the code.  Here are a couple of examples.

This is fine for simple scripts that are not going anywhere and won't need to be imported into any other applications:

```sh
$ go mod init hello
$ cat go.mod
module hello

go 1.17
$
```

This is typically what you'd do if you're storing your code in a repository:

```sh
$ go mod init github.com/darrenparkinson/golang-for-network-engineers
go: creating new go.mod: module github.com/darrenparkinson/golang-for-network-engineers
$ cat go.mod
module github.com/darrenparkinson/golang-for-network-engineers

go 1.17
$
```

Now that we have initialised our module, we can compile our application:

```sh 
$ go build
$ ls -a
.  ..  go.mod  hello  main.go
$ file hello
hello: Mach-O 64-bit executable arm64
$ ./hello
Hello, world!
```

You can see that by default, the executable is named after the module. You can change this by specifying the `-o` flag:

```sh
$ go build -o blah
$ ls -a
.  ..  blah  go.mod  hello  main.go
$ file blah
hello: Mach-O 64-bit executable arm64
$ ./blah
Hello, world!
```

Finally if your script is baked and you just want to be able to use it from anywhere, you can use the `go install` command.  This will essentially compile your code and put it in the go bin directory which should be in your PATH:

```sh
$ go install
$ hello
Hello, world!
```

You can use the `go env GOPATH` command to see where the bin directory is held.

That's probably enough for now and we'll see more examples as we move on.  Now we'll just cover a bit about the basic structure of go applications.

## Basic Structure

Essentially the structure of a go application is pretty simple but there are some conventions that are followed by the majority of users.

For example, any code typically goes into the top level folder of your repository, whether you're writing a library or a binary. If you combine the two, you might put your library at the top level and any binaries in `/cmd`, or you might have your binary at the top level and any packages/libraries you're writing in the `internal` or `pkg` directories, but as I said, this is purely convention and not required.  

The `internal` directory however does have some special meaning though in that go will *prevent one package from being imported by another unless both share a common ancestor* -- for more on this see [this post](https://dave.cheney.net/2019/10/06/use-internal-packages-to-reduce-your-public-api-surface).  It's useful for reducing the public interface of any libraries you write, and we'll probably touch on this more in later sessions.

Finally, if you want further detail, there are many resources regarding application structure.  Again, we'll cover it in more detail as we progress.  However if you need more, there is [this opinionated repository](https://github.com/golang-standards/project-layout) regarding appliciation structure which can be quite heavily debated, so don't get too hung up on it -- as it says, it's not official (despite the official looking url).

## A Note On Virtual Environments

In Go there isn't really a virtual environment as such like you get with `venv` in python, but if we look at the main reasons for a virtual environment, it's really to ensure you don't have global dependencies which cause clutter and version conflicts.  

This isn't typically an issue in Go since you specify the version of the library you want in your `go.mod` file and the package management is MUCH MUCH better than python (in my non-biased opinion).

If however you want to keep a copy of the actual dependencies you rely on in your repository in order to prevent them disappearing in the future (which is unlikely given the way the package management works), you can use the `go mod vendor` command.  It's outside of scope for this post, but something we can cover in future if necessary.
