Markdown Preprocessor
=====================

MDPP (for Makdown Preprocessor) is a tool to process Markdown files to:

- Replace `?(command)` with the result of the `command` running in the Markdown file directory.
- Replace `@(filename)` with the content of the *filename*. If *filename* is a relative path, it is relative to the directory of the Markdown file.

Installation
------------

Drop the binary for your platform in the *bin* directory of the archive somewhere in your `PATH`.

Usage
-----

To process Markdown file *file.md*, you must type:

```bash
$ mdpp file.md
```

This will output processed file on the console, thus to save result in file *processed.md*, you must type:

```bash
$ mdpp file.md > processed.md
```

*Enjoy!*
