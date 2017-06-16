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

* **/cell_lines/{id}**
* **/cell_lines/{id}/get_accession**
* **/cell_lines/{id}/get_synonyms**

`/cell_lines/{id}` returns the cell line whose `identifier = id` or `accession = id` or `synonym = id`. The returned data will have the following structure:

```json
{
	"category": "value",
	"sex": "value",
	"identifier": "value",
	"accession": {
		"primary": "value",
		"secondary": "value"
	},
	"synonyms": [
		"values"
	],
	"species-of-origin": [
		{
			"terminology": "value",
			"accession": "value",
			"species": "value"
		},
	],
	"cross-references": [
		{
			"database": "value",
			"accession": "value"
		},
	],
	"reference-identifier": [
		"values"
	],
	"web-pages": [
		"values"
	],
	"comments": [
		{
			"category": "value",
			"comment": "value"
		},
	],
	"str-profile-data": {
		"sources": [
			"values"
		],
		"markers": [
			{
				"id": "value",
				"alleles": "value"
			},
		]
	},
	"diseases": [
		{
			"terminology": "value",
			"accession": "value",
			"disease": "value"
		},
	],
	"hierarchy": [
		{
			"terminology": "value",
			"accession": "value",
			"derived-from": "value"
		},
	],
	"same-origin-as": [
		{
			"terminology": "value",
			"accession": "value",
			"identifier": "value"
		}
	]
}
```

`value` is either a `string` or `null`. See [Cellosaurus](http://web.expasy.org/cellosaurus/) for more information on each field.

`/cell_lines/{id}/get_accession` returns the accession id associated with the queried cell line (if cell line exists in db).

`/cell_lines/{id}/get_synonyms` returns an array of synonyms listed under the cell line being queried.

## Contributing

To contribute, fork this repository, make updates/edits and submit a pull request.

## License

This project is licensed under the MIT License - see [LICENSE](LICENSE) for details.
