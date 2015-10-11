# sll

Tired of grepping for files and getting really long lines and/or minified junk
in the output?

<img src="https://api.monosnap.com/rpc/file/download?id=If0PPagJbqOIqosYgMnXQJOVbPLai3" alt="Yeah, the second result is not what you are looking for." />

`sll` makes it easy to strip those lines; just pipe the result of `grep` or
`ag` through `sll`.

```bash
ag check | sll
```

By default, any line longer than 1024 lines is not shown. You can change this
by passing an integer to `sll`:

```bash
ag check | sll 80
```

will filter any lines longer than 80 characters from the output.

## Installation

Run `go install github.com/kevinburke/sll` and add `"$GOPATH/bin"` to your
`$PATH` variable.

## Errata

Currently `ag` has a `--print-long-lines` flag, which [has not been
implemented][ag]. If it does, you should probably use that instead.

[ag]: https://github.com/ggreer/the_silver_searcher/issues/189
