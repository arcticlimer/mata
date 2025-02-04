# mata

[![builds.sr.ht status](https://builds.sr.ht/~glorifiedgluer/mata.svg)](https://builds.sr.ht/~glorifiedgluer/mata?)

[mata](https://git.sr.ht/~glorifiedgluer/mata) is CLI tool for [mataroa.blog](https://mataroa.blog).

## Usage

Run `mata init` to get started. Read the man page to learn about all commands.

## Documentation

Also available as man pages:

- [mata(1)](https://git.sr.ht/~glorifiedgluer/mata/tree/master/item/doc/mata.1.scd)
- [mata-config(5)](https://git.sr.ht/~glorifiedgluer/mata/tree/master/item/doc/mata-config.5.scd)

## Building

Dependencies (not needed for Nix users):

- Go
- scdoc (optional, for man pages)

### From Source

For end users, a Makefile is provided:

```
make
make install
```

### From Nix

Dependencies:

- Nix 2.7 or later

You can build and run from your machine with the following:

```
nix run sourcehut:~glorifiedgluer/mata
```

## Contributing

You can find me on IRC: [#mdzk on Libera Chat](ircs://irc.libera.chat/#mdzk).

# License

MIT, see [LICENSE](https://git.sr.ht/~glorifiedgluer/mata/tree/master/LICENSE).

Copyright (C) 2022 Victor Freire
