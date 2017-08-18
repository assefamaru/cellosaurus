# Cellosaurus Restful API

[![Build Status](https://travis-ci.org/assefamaru/cellosaurus.svg?branch=master)](https://travis-ci.org/assefamaru/cellosaurus)
[![Build status](https://ci.appveyor.com/api/projects/status/os3kne9qkch7mils?svg=true)](https://ci.appveyor.com/project/assefamaru/cellosaurus)
[![GoDoc](https://godoc.org/github.com/assefamaru/cellosaurus?status.svg)](https://godoc.org/github.com/assefamaru/cellosaurus)

[Cellosaurus](http://web.expasy.org/cellosaurus/) is a knowledge resource on cell lines. It attempts to describe all cell lines used in biomedical research.

This API aims to make the data provided by Cellosaurus as integrable as possible, by enabling programmatic access to the full database.

## Accessing the API

All calls are made to the following URL, adding required parameters/endpoints for specific services.

```
https://cellosaurus.pharmacodb.com/v2/
```

Response is in `json` format.

## Endpoints

The following endpoints are currently supported.

* **[/cell_lines](docs/template.md)**
* **[/cell_lines/{id}](docs/template.md)**
* **[/search/{id}](docs/template.md)**

## Contributing

To contribute, fork this repository, make changes and submit a pull request.

## License

This project is under the MIT License - see [LICENSE](LICENSE) for details.
