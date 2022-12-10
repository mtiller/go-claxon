## Claxon Format

A the moment, these are some ideas for encoding hypermedia related data. The
goal is to find convenient ways to encode the metadata associated with a normal
response so that it can be easily handled both client and server side in
multiple languages.

Initially, I'll probably focus on Go and TypeScript but the idea is to establish
conventions and formats that are easy to implement in other languages.

Key goals here making it convenient to marshal, unmarshal and query hypermedia
metadata in HTTP requests and responses such that working with hypermedia data
puts as little burden on the developer as possible and ideally requires very
little in the way of additional tooling beyond some basic libraries. A related
goal is to make these formats and conventions as compatible with existing
standards like OpenAPI.

The name `Claxon` is meant as an homage ot the `Siren` format developed by Kevin
Swiber since it has been not just an inspiration for this work but also a
valuable collection of insights that have inspired much of my work over the last
decade.
