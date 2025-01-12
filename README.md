# PG Trunc
CLI for PostgreSQL Truncation specified by Schemas

## Why?
- DB Tests can fail and leave data
- The purpose of this tool is to remove leftover data

## Commands
```sh
# creates the pg function for truncation across schemas
pgTrunc --init

# explicitly state the truncation as an argument for safety reasons
pgTrunc --truncate
```

## Schemas
- Schemas, in which the truncation happens, have to be specified in the config file!
- This guarantees that no unwanted schema is truncated
