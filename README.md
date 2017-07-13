# Cellosaurus API

[![Build Status](https://travis-ci.org/assefamaru/cellosaurus-api.svg?branch=master)](https://travis-ci.org/assefamaru/cellosaurus-api)
[![MIT licensed](https://img.shields.io/badge/license-MIT-blue.svg)](https://github.com/assefamaru/cellosaurus-api/blob/master/LICENSE)

[Cellosaurus](http://web.expasy.org/cellosaurus/) is a knowledge resource on cell lines. It attempts to describe all cell lines used in biomedical research.

This API aims to make the data provided by Cellosaurus as integrable as possible.

## Accessing the API

All calls are made to the following URL, adding required parameters/endpoints for specific services.

```
https://cellosaurus.pharmacodb.com/v1/
```

Returned data is in `json` format.

## Endpoints

* [/cell_lines](docs/cells.md)
* [/cell_lines/{id}](docs/cell.md)

## Contributing

To contribute, fork this repository, make changes and submit a pull request.

## License

This project is under the MIT License - see [LICENSE](LICENSE) for details.
