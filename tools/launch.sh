#!/usr/bin/env bash
set -eu

outfile="../docs/methods.todo.md"
echo "# Scrumwise API method list" > "$outfile"
echo "" >> "$outfile"
clojure user.clj >> "$outfile"
