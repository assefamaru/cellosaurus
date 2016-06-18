# Cellosaurus API

The [Cellosaurus](http://web.expasy.org/cellosaurus/) is a knowledge resource on cell lines. 
A detailed content description of Cellosaurus can be found [here](http://web.expasy.org/cellosaurus/description.html). Cellosaurus currently provides text/obo files to make their data publicly available, and can be found [here](http://web.expasy.org/cellosaurus/).

**cellosaurus-api** was built to make this data more accessible, making it more convenient to build applications or integrate
Cellosaurus data. All outputs are currently in `JSON` format.

To access the API, make a call to the following URL, adding the required endpoints for a given query.
```
https://cellosaurus.herokuapp.com/api/v1/
```

## Endpoints

* /cell_lines/all
* /cell_lines/{accession}

For example, using `accession number = CVCL_8264`, you can make a `GET` request to:
`https://cellosaurus.herokuapp.com/api/v1/cell_lines/CVCL_8264`

which will output the following:

```
{
  "Identifier (cell line name)": "CVCL_8264",
  "Accession (CVCL_xxxx)": "CVCL_8264",
  "Synonyms": "Minnesota-EE; Minn. EE",
  "Cross-references": [
    "ATCC; CCL-4",
    "BioSample; SAMN03151744"
  ],
  "References identifier": [
    "PubMed=566722;",
    "PubMed=1246601;",
    "PubMed=13472685;",
    "PubMed=20143388;"
  ],
  "Web pages": [
    "http://iclac.org/wp-content/uploads/Cross-Contaminations-v7_2.pdf"
  ],
  "Comments": [
    "Problematic cell line: Contaminated. Shown to be a HeLa derivative (PubMed 566722, PubMed 1246601, PubMed 20143388).",
    "Discontinued: ATCC; CCL-4; true."
  ],
  "STR data": [
    "NA"
  ],
  "Diseases": [
    "NCIt; C4029; Cervical adenocarcinoma"
  ],
  "Species of origin": [
    "NCBI_TaxID=9606; ! Homo sapiens"
  ],
  "Hierarchy": [
    "CVCL_0030 ! HeLa"
  ],
  "Originate from same individual": [
    "NA"
  ],
  "Sex (gender) of cell": "Female",
  "Category": "Cancer cell line"
}

```

## License

This project is under [MIT](LICENSE) license.
