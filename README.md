# PG-Trunc
![Lines-of-Code](https://img.shields.io/badge/lines--of--code-345-brightgreen)

CLI for PostgreSQL Truncation specified by Schemas

## Why?
- DB Tests can fail and leave data
- The purpose of this tool is to remove leftover data

## Commands
```sh
# creates the pg function for truncation across schemas
pgtrunc --init

# explicitly state the truncation as an argument for safety reasons
pgtrunc --trunc
```

## Schemas
- Schemas, in which the truncation happens, have to be specified in the config file!
- This guarantees that no unwanted schema is truncated
