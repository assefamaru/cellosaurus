;; Module converts Cellosaurus' .txt files to csv format.
;; The text file(s) can be found at either of:
;; * ftp://ftp.expasy.org/databases/cellosaurus
;; * https://github.com/calipho-sib/cellosaurus
;; The text files are cellosaurus.txt and cellosaurus_refs.txt

#lang racket

;; Headers defining the list of column names for the csv files.
(define header-cells "\"\",\"identifier\",\"accession\"")
(define header-attrs "\"\",\"accession\",\"attribute\",\"content\"")
(define header-refs "\"\",\"ref_identifier\",\"attribute\",\"content\"")

;; write-cells-to-csv iterates over each line of cellosaurus.txt,
;; parsing identifier and accession attributes, and writing
;; content to csv (via printf).
;; For example, lines such as:
;;                            "ID   #15310-LN"
;;                            "AC   CVCL_E548"
;;              gets stored in csv as: n,"#15310-LN","CVCL_E548"
;;              whereby n is some row number (integer) > 0
(define (write-cells-to-csv in [row 1])
  (define line (read-line in))
  (unless (eof-object? line)
    (define id (cadr (string-split line "   ")))
    (define ac (cadr (string-split (read-line in) "   ")))
    (printf "~a,\"~a\",\"~a\"\n" row id ac)
    (let loop ()
      (if (equal? (read-line in) "//")
          (write-cells-to-csv in (add1 row))
          (loop)))))

;; write-attrs-to-csv iterates over each line of cellosaurus.txt,
;; parsing line and writing content to csv (via printf).
;; For example, a line such as: "SX   Female"
;;              gets stored in csv as: n,"CVCL_xxxx","SX","Female"
;;              whereby n is some row number (integer) > 0
(define (write-attrs-to-csv in [row 1])
  (define line (read-line in))
  (unless (eof-object? line)
    (define ac (cadr (string-split (read-line in) "   ")))
    (let loop ([row row])
      (define line (read-line in))
      (cond
        [(equal? line "//")
         (write-attrs-to-csv in row)]
        [else
         (define lst (string-split line "   "))
         (printf "~a,\"~a\",\"~a\",\"~a\"\n" row ac (car lst) (cadr lst))
         (loop (add1 row))]))))

;; write-attrs-to-csv iterates over each line of cellosaurus.txt,
;; parsing line and writing content to csv (via printf).
(define (write-refs-to-csv in [row 1])
  (define line (read-line in))
  (unless (eof-object? line)
    (define rx (cadr (string-split line "   ")))
    (let loop ([row row])
      (define line (read-line in))
      (cond
        [(equal? line "//")
         (write-refs-to-csv in row)]
        [else
         (define lst (string-split line "   "))
         (printf "~a,\"~a\",\"~a\",\"~a\"\n" row rx (car lst) (cadr lst))
         (loop (add1 row))]))))

;; Convert cellosaurus.txt to csv format.
(define (convert txt csv table)
  (define in (open-input-file txt))
  (with-output-to-file csv
    (lambda ()
      (cond
        [(equal? table "cell_lines")
         (printf "~a\n" header-cells)
         (write-cells-to-csv in)]
        [(equal? table "attributes")
         (printf "~a\n" header-attrs)
         (write-attrs-to-csv in)]
        [(equal? table "references")
         (printf "~a\n" header-refs)
         (write-refs-to-csv in)])))
  (close-input-port in))

;; --- Uncomment below to convert .txt to .csv
;; --- "cell_lines.csv" output will contain unique cell lines,
;;      each with their identifier and accession attributes only.
;; --- "attributes.csv" ouput will contain each attribute type and attribute content
;;      for a cell line (referenced using accession).
;; (convert "cellosaurus.txt" "cell_lines.csv" "cell_lines")
;; (convert "cellosaurus.txt" "attributes.csv" "attributes")
;; (convert "cellosaurus_refs.txt" "references.csv" "references")
