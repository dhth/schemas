# schemas

✨ Overview
---

Inspect schemas of postgres tables in the terminal.

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
