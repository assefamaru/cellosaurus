;; Module converts Cellosaurus' text file data to csv format.
;; The text file(s) can be found at one of:
;; * ftp://ftp.expasy.org/databases/cellosaurus
;; * https://github.com/calipho-sib/cellosaurus
;; The text files are cellosaurus.txt and cellosaurus_refs.txt

#lang racket

;; Headers defining the list of column names for the csv files.
(define header-cells "\"acp\",\"id\",\"acs\",\"sy\",\"sx\",\"ca\"")
(define header-attrs "\"\",\"accession\",\"attribute\",\"content\"")
(define header-refs "\"\",\"ref_identifier\",\"attribute\",\"content\"")

;; write-cells-to-csv iterates over each line of cellosaurus.txt,
;; parsing the following attributes for each cell line entry:
;;  ID : Identifier (cell line name)
;;  AC : Accession (CVCL_xxxx)
;;  AS : Secondary accession number(s)
;;  SY : Synonyms
;;  SX : Sex (gender) of cell
;;  CA : Category
(define (write-cells-to-csv in)
  (define line (read-line in))
  (unless (eof-object? line)
    (let loop ([line line]
               [id ""]
               [ac ""]
               [as ""]
               [sy ""]
               [sx ""]
               [ca ""])
      (define lst (string-split line "   "))
      (cond
        [(equal? (car lst) "//")
         (printf "\"~a\",\"~a\",\"~a\",\"~a\",\"~a\",\"~a\"\n" ac id as sy sx ca)
         (write-cells-to-csv in)]
        [(equal? (car lst) "ID")
         (set! id (cadr lst))
         (loop (read-line in) id ac as sy sx ca)]
        [(equal? (car lst) "AC")
         (set! ac (cadr lst))
         (loop (read-line in) id ac as sy sx ca)]
        [(equal? (car lst) "AS")
         (set! as (cadr lst))
         (loop (read-line in) id ac as sy sx ca)]
        [(equal? (car lst) "SY")
         (set! sy (cadr lst))
         (loop (read-line in) id ac as sy sx ca)]
        [(equal? (car lst) "SX")
         (set! sx (cadr lst))
         (loop (read-line in) id ac as sy sx ca)]
        [(equal? (car lst) "CA")
         (set! ca (cadr lst))
         (loop (read-line in) id ac as sy sx ca)]
        [else
         (loop (read-line in) id ac as sy sx ca)]))))

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

;; write-attrs-to-csv iterates over each line of cellosaurus_refs.txt,
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

;; Convert Cellosaurus' data from txt to csv format.
(define (convert txt csv table)
  (define in (open-input-file txt))
  (with-output-to-file csv
    (lambda ()
      (cond
        [(equal? table "cells")
         (printf "~a\n" header-cells)
         (write-cells-to-csv in)]
        [(equal? table "attributes")
         (printf "~a\n" header-attrs)
         (write-attrs-to-csv in)]
        [(equal? table "refs")
         (printf "~a\n" header-refs)
         (write-refs-to-csv in)])))
  (close-input-port in))

;; --- "cells.csv" output will contain unique cell lines,
;;      each with their identifier and accession attributes only.
;; --- "attributes.csv" ouput will contain each attribute type and attribute content
;;      for a cell line (referenced using accession).
;; --- "refs.csv" output will contain reference data for each reference identifier
;; (convert "../data/txt/cellosaurus.txt" "../data/csv/cells.csv" "cells")
;; (convert "../data/txt/cellosaurus.txt" "../data/csv/attributes.csv" "attributes")
;; (convert "../data/txt/cellosaurus_refs.txt" "../data/csv/refs.csv" "refs")
