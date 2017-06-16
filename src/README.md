# Cellosaurus API

[![Build Status](https://travis-ci.org/assefamaru/cellosaurus-api.svg?branch=master)](https://travis-ci.org/assefamaru/cellosaurus-api) [![MIT License](https://badges.frapsoft.com/os/mit/mit.png?v=103)](https://opensource.org/licenses/mit-license.php) [![GitHub version](https://badge.fury.io/gh/assefamaru%2Fcellosaurus-api.svg)](https://badge.fury.io/gh/assefamaru%2Fcellosaurus-api) [![Issue Count](https://codeclimate.com/github/assefamaru/cellosaurus-api/badges/issue_count.svg)](https://codeclimate.com/github/assefamaru/cellosaurus-api)

[Cellosaurus](http://web.expasy.org/cellosaurus/) is a knowledge resource on cell lines. It attempts to describe all cell lines used in biomedical research.

This API aims to make the data provided by Cellosaurus as integrable as possible.

## Accessing the API

All calls are made to the following URL, adding required parameters/endpoints for specific services.

```
https://cellosaurus-api.herokuapp.com/v1/cell_lines/
```

Returned data is in `json` format.

## Endpoints

* **`/cell_lines/{id}`**

## Contributing

To contribute, fork this repository, make updates/edits and submit a pull request.

## License

This project is licensed under the MIT License - see [LICENSE](LICENSE) for details.
