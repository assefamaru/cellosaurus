# cellosaurus-api

[![Build Status](https://travis-ci.org/assefamaru/cellosaurus-api.svg?branch=master)](https://travis-ci.org/assefamaru/cellosaurus-api)
[![Build status](https://ci.appveyor.com/api/projects/status/ssw9ljftsj3pbom5?svg=true)](https://ci.appveyor.com/project/assefamaru/cellosaurus-api)
[![GoDoc](https://godoc.org/github.com/assefamaru/cellosaurus-api/src?status.svg)](https://godoc.org/github.com/assefamaru/cellosaurus-api/src)

The [Cellosaurus](https://web.expasy.org/cellosaurus/) is a knowledge resource on cell lines. It attempts to describe all cell lines used in biomedical research. This API aims to make the data provided by Cellosaurus as integrable as possible, by providing programmatic access to the full database.

NOTE: The live version of the API is hosted on a free tier plan on heroku, and [sleeps after 30 minutes of inactivity](https://devcenter.heroku.com/articles/dynos#dyno-idling). As a result, you might experience a lag on your first request.

## Accessing the API

All calls are made to the following base URL, adding required endpoints for specific services.

```
https://api.cellosaur.us/v1/
```

Responses are in `json` format.

## Endpoints

The following endpoints are currently supported:

* **[/cell-lines](https://api.cellosaur.us/v1/cell-lines)**
* **[/cell-lines/{id}](https://api.cellosaur.us/v1/cell-lines/mcf7)**
* **[/references](https://api.cellosaur.us/v1/references)**
* **[/terminologies](https://api.cellosaur.us/v1/terminologies)**
* **[/release-info](https://api.cellosaur.us/v1/release-info)**

### Requests

All endpoints accept `GET` HTTP method, and each endpoint has a set of parameters/options that allow a user to format responses. See table below for a summary of how to structure URLs with parameters.

| Endpoints | Parameters | Examples |
| :--- | :--- | :--- |
| **/cell-lines** | `page`, `per_page`, `indent` | https://api.cellosaur.us/v1/cell-lines?page=1&per_page=20 |
| **/cell-lines/{id}** | `indent` | https://api.cellosaur.us/v1/cell-lines/mcf7?indent=true |
| **/references** | `page`, `per_page`, `indent` | https://api.cellosaur.us/v1/references?page=1&per_page=10&indent=true |
| **/terminologies** | `indent` | https://api.cellosaur.us/v1/terminologies?indent=true |
| **/release-info** | `indent` | https://api.cellosaur.us/v1/release-info |

Parameters need not be present in request URLs. When parameters are not included in request, they are set to their default values. The following are their default values:

| Parameters | Default Values |
| :--- | :---: |
| `page` | `1` |
| `per_page` | `10` |
| `indent` | `false` |

### Responses

All responses are in `json` format. Endpoints that contain `page` and `per_page` parameters will have a `meta` field in their response containing pagination information, as well as the total number of records under the requested resource type. Such a response will look as follows:

```json
{
    "meta": {
        "page": __,
        "per-page": __,
        "last-page": __,
        "total": __
    },
    "data": [ ... ]
}
```

## Setting up the API locally

To set up the API locally,

1. Download and unzip the [latest release](https://github.com/assefamaru/cellosaurus-api/releases/latest) of cellosaurus-api on your local machine.
2. Navigate to the [scripts](scripts) directory and execute the script `db.sh`. This will setup the full database in your local mysql instance.
3. Create the following environment variables locally: `cellosaurus_user`, `cellosaurus_pass`, `cellosaurus_db`, `cellosaurus_host`. They represent your local mysql user, password, database name and host respectively. Set them to their appropriate values accordingly.

With the above steps, the API should now be fully setup on your local machine. You can now run the API using one of two options:

1. Run one of the pre-built [executables](build). Depending on your operating system, you should pick the right binary. For example, if you are on a 64 bit Ubuntu OS, you can use `linux-amd64`, or use one of the `.exe` builds for windows etc.
2. You can also simply run `go run main.go` in the main directory.

## Contributing

You can contribute in various ways:

- Send PRs (pull requests) with new feature implementations, fixed bugs etc. Or,
- Offer suggestions, request new features, or report any errors by creating [new](https://github.com/assefamaru/cellosaurus-api/issues/new) issues and assigning them appropriate labels.

## License

This project is under the MIT License - see [LICENSE](LICENSE) for details.
