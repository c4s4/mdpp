#!/usr/bin/env python
#
# Filter that processes given Markdown file and replaces :
#
# - "#(file)" with the content of the file (if path is relative, this
#   is relative to the Markdown file).
# - "?(command)" with the output of the command (running in the Markdown
#   file directory.
#
# Output is written on command line.

import re
import sys
import subprocess


def execute(match):
    process = subprocess.Popen(match.group(1),
                               stdout=subprocess.PIPE,
                               stderr=subprocess.PIPE)
    output, errput = process.communicate()
    if process.returncode != 0:
        print("Error running command '%s':\n%s" % (match, errput))
        sys.exit(2)
    return output.strip()


def readfile(match):
    with open(match.group(1)) as stream:
        return stream.read().strip()


def main(filename):
    with open(filename) as stream:
        content = stream.read()
    content = re.sub(r'^\?\((.+)\)$', execute, content, flags=re.MULTILINE)
    content = re.sub(r'^#\((.+)\)$', readfile, content, flags=re.MULTILINE)
    print(content)


if __name__ == '__main__':
    if len(sys.argv) != 2:
        print('You must pass Markdow file to process on command line')
        sys.exit(1)
    main(sys.argv[1])
