# Cellosaurus API

[![Build Status](https://travis-ci.org/assefamaru/cellosaurus-api.svg?branch=master)](https://travis-ci.org/assefamaru/cellosaurus-api)
[![Build status](https://ci.appveyor.com/api/projects/status/ssw9ljftsj3pbom5?svg=true)](https://ci.appveyor.com/project/assefamaru/cellosaurus-api)
[![GoDoc](https://godoc.org/github.com/assefamaru/cellosaurus-api?status.svg)](https://godoc.org/github.com/assefamaru/cellosaurus-api)

[Cellosaurus](http://web.expasy.org/cellosaurus/) is a knowledge resource on cell lines. It attempts to describe all cell lines used in biomedical research.

This API aims to make the data provided by Cellosaurus as integrable as possible, by providing programmatic access to the full database.

## Accessing the API

All calls are made to the following URL, adding required endpoints for specific services.

```
https://cellosaur.us/api/v1/
```

All response is in `json` format.

## Endpoints

The following endpoints are currently supported.

* GET **[/cell_lines](#)**
* GET **[/cell_lines/{id}](#)**
* GET **[/search/{id}](#)**

| Endpoints | Parameters | Examples |
| :--- | :--- | :--- |
| `/cell_lines` | `page`, `per_page`, `all`, `indent`, `include` | https://cellosaur.us/api/v1/cell_lines?page=2&per_page=10 |
| `/cell_lines/{id}` | `indent`, `type` | https://cellosaur.us/api/v1/cell_lines/mcf-7?type=identifier |
| `/search/{id}` | `indent` | https://cellosaur.us/api/v1/search/mcf-7?indent=false |

A cell line can be searched using its `identifier`, `accession id`, or its `synonym names`. For example, the cell line **mcf-7** has the following attributes:

| Identifier | Accession ID | Synonyms |
| :--- | :------------------- | :--- |
| MCF-7 | CVCL_0031 | "MCF 7", "MCF.7", "MCF7", "Michigan Cancer Foundation-7", "ssMCF-7", "ssMCF7", "MCF7/WT", "IBMF-7", "MCF7-CTRL" |

So, in order to look up this cell line, request options look like the following:

```
GET /cell_lines/CVCL_0031
GET /cell_lines/mcf-7?type=identifier
GET /cell_lines/ssMCF7?type=synonym

GET /search/CVCL_0031
GET /search/mcf-7
GET /search/ssMCF7
```

## Responses

All responses are in `json` format.

| Status Code | Status Message | Description |
| :--- | :--- | :--- |
| 200 | Status OK | Normal response with no errors |
| 400 | Bad Request | The request URL is not supported by the API |
| 404 | Not Found | The requested resource was not found |
| 500 | Internal Server Error | Something bad happened internally |

## Contributing

You can offer suggestions, request new features, or report any erros by creating a [new](https://github.com/assefamaru/cellosaurus-api/issues/new) issue and assigning it an appropriate label.

## License

This project is licensed under the MIT License - see [LICENSE](LICENSE) for details.
