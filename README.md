# Cellosaurus API

The [Cellosaurus](http://web.expasy.org/cellosaurus/) is a knowledge resource on cell lines. 
A detailed content description of Cellosaurus can be found [here](http://web.expasy.org/cellosaurus/description.html). Cellosaurus currently provides text/obo files to make their data publicly available, and can be found [here](ftp://ftp.expasy.org/databases/cellosaurus).

**cellosaurus-api** was built to make this data more accessible, making it more convenient to build applications or integrate
Cellosaurus data. All outputs are currently in `JSON` format.

To access the API, make a call to the following URL, adding the required endpoints for a given query.
```
https://cellosaurus.herokuapp.com/api/v1/
```

## Endpoints

* /cell_lines/all
* /cell_lines/{accession}

## License

This project is under [MIT](LICENSE) license.
