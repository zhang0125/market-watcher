# Market watcher

The purpose of this project is to observe the transaction data of the exchange and store it in a database for post-event analysis

## Build

```shell
make build
```

## Configurator

Modify the database configuration in conf/default.toml.

Initialize the database tables

```shell
market-watcher -home .  migrate 
```

## Launch

```shell
market-watcher -home .  start 
```