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

## Other Ways to Implement this

There are many other ways to write this and some of them are listed here. Feel
free to add your own with a pull request. Be sure to mention how this project
is "pointless" and a "waste of time". Before merging, I will berate myself for
having foisted this waste on the unsuspecting Internet population, retreat to
my parents basement, and promise to never open source a thing I write ever
again.

- `egrep -v .{1024}`
- `cut -c 1-1024`
- `less -S`
- `sed -E 's/^(.{0,80}).*/\1/'`

## Errata

Currently `ag` has a `--print-long-lines` flag, which [has not been
implemented][ag]. If it does, you should probably use that instead.

[ag]: https://github.com/ggreer/the_silver_searcher/issues/189
