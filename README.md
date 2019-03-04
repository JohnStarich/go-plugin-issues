# go-plugin-issues

## Go plugin fails to load from c-shared lib

To recreate the issue, run `make test`. This will do a few things:

1. Compile and run a Go executable named `loader`
    - This loader will open the `pow` plugin and run it with no arguments.
    - This works as expected.
2. Compile a c-shared library with an exported function `doit()` named `cloader`.
    - This loader will be dynamically opened by the `main.c` executable and then run `doit()`.
    - The `doit()` func will open the `pow` plugin and run it with no arguments.
    - This works on macOS, but **crashes** on Linux with `fatal error: runtime: no plugin module data`

You can also run `make` to recreate the issue on your native system **plus** a Linux system running in Docker.

See the most recent Travis builds here: https://travis-ci.com/JohnStarich/go-plugin-issues/builds
