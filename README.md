# Gator --- CLI Web‑Blog Aggregator

`gator` is a command‑line tool that lets you aggregate blog / RSS‑style
feeds, store them in a local database (PostgreSQL), and browse them from
the CLI.

## Prerequisites

-   Go installed.
-   PostgreSQL installed and running.

## Installation

``` bash
go install github.com/jaharbaugh/gator@latest
```

## Configuration

Create the config file:

``` bash
touch ~/.gatorconfig.json
```

Example content:

``` json
{
  "db_url": "postgres://USERNAME:PASSWORD@localhost:5432/gator?sslmode=disable",
  "current_user": "your-username"
}
```

## Usage

-   `gator register <username>`
-   `gator login <username>`
-   `gator reset`
-   `gator addfeed <url>`
-   `gator feeds`
-   `gator follow <url>`
-   `gator unfollow <url>`
-   `gator following`
-   `gator agg <interval>`
-   `gator browse [limit]`

## Quick Start

``` bash
go install github.com/jaharbaugh/gator@latest
gator reset
gator register alice
gator addfeed https://example.com/rss.xml
gator follow https://example.com/rss.xml
gator agg 60s
gator browse 10
```
