# The zaidlang Programming Language

[![Test](https://github.com/zaid-language/zaid-lang/actions/workflows/test.yml/badge.svg)](https://github.com/zaid-language/zaid-lang/actions/workflows/test.yml)

zaidlang is a small, object-oriented, embeddable toy scripting language. While object-oriented, zaidlang also supports procedural and functional programming styles as well.

zaidlang is dynamically typed, runs by a tree-walking interpreter, and has automatic memory management thanks to its implementation through the Go programming language.

## Status

> Currently in beta, vetting out the language and seeing how it feels writing/running. Major changes are still possible at this stage.

## Documentation

You will find robust, user friendly, and updated documentation on our [website](https://zaidlang.tech/).

## Versioning

We use [SemVer](http://semver.org/) for versioning. For the versions available, see the [tags on this repository](https://github.com/zaid-language/zaid-lang/tags).

## Installation

### Brew

If you're on mac, you may use `homebrew`:

```
$ brew tap zaid-language/zaid
$ brew install zaid-language/zaid/zaid
```

### Go Install

If you have Go installed, you may use `go install`:

```
go install zaidlang
```

### Direct Download

You may download the compiled binaries for your platform from our GitHub [releases](https://github.com/zaid-language/zaid-lang/releases) page.

## Development

- To build and execute zaidlang, run `make`.
- To build zaidlang, run `make build`.
- To execute tests, run `make test`.

```
$  git clone git@github.com:zaid-language/zaid.git
$  cd zaid
$  make
   zaidlang (x.x)
   Press Ctrl + C to exit

>>
```

## CLI

You can execute code written in zaidlang in various ways using the CLI.

### REPL

zaidlang includes a simple REPL to write and execute zaidlang code directly in your terminal. To enter the REPL environment, run `zaid` on its own:

```
$  zaid
   zaidlang (x.x)
   Press Ctrl + C to exit

>>
```

### Executing Files

To execute a zaidlang source file (`.zaid`), pass either the relative or absolute path of the file to `zaid`. The source file will be executed and then exit back to the terminal.

```
$  zaid examples/fibtc.zaid
   9227465
$
```

## Releasing

zaidlang is hosted and distributed through GitHub. We utilize [GoReleaser](https://goreleaser.com) to automate the release process. GoReleaser will build all the necessary binaries, publish the release and publish the brew tap formula. The following steps outline the process for maintainers of zaidlang:

1. Ensure you have a GitHub token with `repo` access saved to your environment:

```
export GITHUB_TOKEN="YOUR_GH_TOKEN"
```

2. Ensure the internal version reference is updated:

   ```go
   // version/version.go

   var (
      Version = "x.y.z"
   )
   ```

3. Create a new tag:

```
$ git tag -a vx.y.z -m "Release description"
$ git push origin vx.y.z
```

4. Run GoReleaser:

```
$ goreleaser
```

## Credits

- [Crafting Interpreters](https://craftinginterpreters.com/)
- [Writing An Interpreter In Go](https://interpreterbook.com/)

## License

zaidlang is open-sourced software licensed under the MIT license. See the [LICENSE](LICENSE) file for complete details.
