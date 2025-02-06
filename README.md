# schemas

[![Build Workflow Status](https://img.shields.io/github/actions/workflow/status/dhth/schemas/build.yml?style=flat-square)](https://github.com/dhth/schemas/actions/workflows/build.yml)
[![Vulncheck Workflow Status](https://img.shields.io/github/actions/workflow/status/dhth/schemas/vulncheck.yml?style=flat-square&label=vulncheck)](https://github.com/dhth/schemas/actions/workflows/vulncheck.yml)
[![Latest Release](https://img.shields.io/github/release/dhth/schemas.svg?style=flat-square)](https://github.com/dhth/schemas/releases/latest)
[![Commits Since Latest Release](https://img.shields.io/github/commits-since/dhth/schemas/latest?style=flat-square)](https://github.com/dhth/schemas/releases)

`schemas` lets you inspect schemas of postgres tables via a TUI.

<p align="center">
  <img src="./schemas.gif?raw=true" alt="Usage" />
</p>


💾 Installation
---

**homebrew**:

```sh
brew install dhth/tap/schemas
```

**go**:

```sh
go install github.com/dhth/schemas@latest
```

⚡️ Usage
---

```bash
DATABASE_ADDRESS='ABC' \
DATABASE_PORT='ABC' \
DATABASE_USERNAME='ABC' \
DATABASE_PASSWORD='ABC' \
DATABASE_DBNAME='ABC' \
schemas
```

Acknowledgements
---

`schemas` is built using [bubbletea][1].

[1]: https://github.com/charmbracelet/bubbletea
