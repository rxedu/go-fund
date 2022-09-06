# Go Fund

[![GitHub Actions](https://github.com/rxedu/go-fund/workflows/main/badge.svg)](https://github.com/rxedu/go-fund/actions)

## Description

[Go Programming Language: An Introductory Golang Tutorial by Brendon Hogger](https://www.toptal.com/go/go-programming-a-step-by-step-introductory-tutorial)

## Installation

Simply import this module into your project

```go
import "github.com/rxedu/go-fund"
```

Then run

```
$ go mod tidy
```

## Development and Testing

### Quickstart

```
$ git clone https://github.com/rxedu/go-fund.git
$ cd go-fund
$ make
$ make test
```

### Source code

The [source code] is hosted on GitHub.
Clone the project with

```
$ git clone git@github.com:rxedu/go-fund.git
```

[source code]: https://github.com/rxedu/go-fund

### Requirements

A [Go] version compatible with the one specified in `go.mod`,
[GoReleaser], and [golangci-lint].

[Go]: https://golang.org/
[golangci-lint]: https://golangci-lint.run/
[GoReleaser]: https://goreleaser.com/

## GitHub Actions

_GitHub Actions should already be configured: this section is for reference only._

The following repository secrets must be set on [GitHub Actions]:

- `GH_USER`: The GitHub user's username to pull and push containers.
- `GH_TOKEN`: A personal access token that can trigger workflows.
- `GPG_PRIVATE_KEY`: The GitHub user's [GPG private key].
- `GPG_PASSPHRASE`: The GitHub user's GPG passphrase.

These must be set manually.

### Secrets for Optional GitHub Actions

The version and format GitHub actions
require a user with write access to the repository,
including access to trigger workflows.
Set these additional secrets to enable the action:

- `GH_USER`: The GitHub user's username.
- `GH_TOKEN`: A personal access token for the user.
- `GIT_USER_NAME`: The GitHub user's real name.
- `GIT_USER_EMAIL`: The GitHub user's email.

[GitHub Actions]: https://github.com/features/actions
[GPG private key]: https://github.com/marketplace/actions/import-gpg#prerequisites

## Contributing

Please submit and comment on bug reports and feature requests.

To submit a patch:

1. Fork it (https://github.com/rxedu/go-fund/fork).
2. Create your feature branch (`git checkout -b my-new-feature`).
3. Make changes.
4. Commit your changes (`git commit -am 'Add some feature'`).
5. Push to the branch (`git push origin my-new-feature`).
6. Create a new Pull Request.

## License

This Go module plugin is licensed under the MIT license.

## Warranty

This software is provided by the copyright holders and contributors "as is" and
any express or implied warranties, including, but not limited to, the implied
warranties of merchantability and fitness for a particular purpose are
disclaimed. In no event shall the copyright holder or contributors be liable for
any direct, indirect, incidental, special, exemplary, or consequential damages
(including, but not limited to, procurement of substitute goods or services;
loss of use, data, or profits; or business interruption) however caused and on
any theory of liability, whether in contract, strict liability, or tort
(including negligence or otherwise) arising in any way out of the use of this
software, even if advised of the possibility of such damage.
