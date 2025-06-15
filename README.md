<p align="center">
  <h1 align="center">schemas</h1>
  <p align="center">
    <a href="https://github.com/dhth/schemas/actions/workflows/main.yml"><img alt="Build Status" src="https://img.shields.io/github/actions/workflow/status/dhth/schemas/main.yml?style=flat-square"></a>
    <a href="https://github.com/dhth/schemas/actions/workflows/vulncheck.yml"><img alt="Vulnerability Check" src="https://img.shields.io/github/actions/workflow/status/dhth/schemas/vulncheck.yml?style=flat-square&label=vulncheck"></a>
    <a href="https://github.com/dhth/schemas/releases/latest"><img alt="Latest release" src="https://img.shields.io/github/release/dhth/schemas.svg?style=flat-square"></a>
    <a href="https://github.com/dhth/schemas/releases/latest"><img alt="Commits since latest release" src="https://img.shields.io/github/commits-since/dhth/schemas/latest?style=flat-square"></a>
  </p>
</p>

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
