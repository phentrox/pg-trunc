# PG-Trunc
![Lines-of-Code](https://img.shields.io/badge/lines--of--code-345-brightgreen)  
  
PostgreSQL Truncation Tool based on Schemas

## Why?
- Database tests can fail and leave data
- The purpose of this tool is to remove leftover data

## Configuration
```yaml
# pgtrunc.yaml
title: "Test-DB"
host: 
port: "5432"
user: "postgres"
password: "password"
database: "test"
schemas:
  - contacts
  - it
  - sales
```

## Commands
```sh
# creates the pg function for truncation across schemas
pgtrunc --init

# explicitly state the truncation as an argument for safety reasons!
pgtrunc --trunc
```

## Schemas
- Schemas, in which the truncation happens, have to be specified in the config file!
- This guarantees that no unwanted schema is truncated
