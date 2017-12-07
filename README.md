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

* **[/release-info](#Endpoints)**
* **[/terminologies](#Endpoints)**
* **[/references](#Endpoints)**
* **[/cell-lines](#Endpoints)**
* **[/cell-lines/{id}](#Endpoints)**

## Contributing

To contribute, you can:

* Fork this repository, make edits, and open a pull request.
* Offer suggestions, request new fetures, or report any errors by creating a [new](https://github.com/assefamaru/cellosaurus-api/issues/new) issue and assigning it an appropriate label.

## License

This project is under the MIT License - see [LICENSE](LICENSE) for details.
