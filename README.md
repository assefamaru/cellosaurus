# The Cellosaurus API

[![Build Status](https://travis-ci.org/assefamaru/cellosaurus-api.svg?branch=master)](https://travis-ci.org/assefamaru/cellosaurus-api)
[![Build status](https://ci.appveyor.com/api/projects/status/ssw9ljftsj3pbom5?svg=true)](https://ci.appveyor.com/project/assefamaru/cellosaurus-api)
[![GoDoc](https://godoc.org/github.com/assefamaru/cellosaurus-api/src?status.svg)](https://godoc.org/github.com/assefamaru/cellosaurus-api/src)

[Cellosaurus](https://web.expasy.org/cellosaurus/) is a knowledge resource on cell lines. It attempts to describe all cell lines used in biomedical research.

This API aims to make the data provided by Cellosaurus as integrable as possible, by providing programmatic access to the full database.

## Accessing the API

All calls are made to the following base URL, adding required endpoints for specific services.

```
https://api.cellosaur.us/v1/
```

## Endpoints

The following endpoints are currently supported.

* **[/release-info](https://api.cellosaur.us/v1/release-info)**
* **[/terminologies](https://api.cellosaur.us/v1/terminologies)**
* **[/references](https://api.cellosaur.us/v1/references)**
* **[/cell-lines](https://api.cellosaur.us/v1/cell-lines)**
* **[/cell-lines/{id}](https://api.cellosaur.us/v1/cell-lines/mcf7)**

## How to setup API locally

To run the API locally,

- Download and extract the [latest release](https://github.com/assefamaru/cellosaurus-api/releases/latest).
- Navigate to the [/scripts](scripts) directory and execute the script `db.sh`. This will setup the full database in your mysql instance.
- Create the following environment variables locally: `cellosaurus_user`, `cellosaurus_pass`, `cellosaurus_db`, `cellosaurus_host`. They represent your local mysql user, password, database name and host respectively.
- Now that the API is set up, you can either run one of the available [executables](build), or run `go run main.go`.

## Contributing

To contribute, you can:

* Fork this repository, make edits, and open a pull request.
* Offer suggestions, request new fetures, or report any errors by creating a [new](https://github.com/assefamaru/cellosaurus-api/issues/new) issue and assigning it an appropriate label.

## License

This project is under the MIT License - see [LICENSE](LICENSE) for details.
