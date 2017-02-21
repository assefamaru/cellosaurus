# Cellosaurus API

[Cellosaurus](http://web.expasy.org/cellosaurus/) is a knowledge resource on cell lines. It attempts to describe all cell lines used in biomedical research. 

This API aims to make the data provided by Cellosaurus as integrable as possible. 

## Accessing the API

All calls are made to the following URL, adding required parameters/endpoints for specific services.

```
http://cellosaurus.pharmacodb.com/v1/
```

Returned data is in `json` format.

## Endpoints

* **/cell_lines/{id}**

`/cell_lines/{id}` returns the cell line with either its `identifier = id` or `accession = id` or `synonyms include id`. The returned data will have the following structure:

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
	"species-of-origin": {
		"terminology": "value",
		"accession": "value",
		"species": "value"
	},
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

`value` is either a `string` or `null`.

## Contributing

Fork repo, make updates and submit a pull request.

## License

This project is licensed under the MIT License - see the [LICENSE](LICENSE) file for details.