#lang racket

;; Module for converting Cellosaurus' data text file to csv
;;   Text file can be found here:
;;   ftp://ftp.expasy.org/databases/cellosaurus/cellosaurus.txt

(struct cell-line (id ac as sy dr rx ww cc st di ox hi oi sx ca) #:mutable #:transparent)
(struct counter (x) #:mutable #:transparent)

(define cline (cell-line "" "" "" "" "" "" "" "" "" "" "" "" "" "" ""))
(define count (counter 1))

(define (print/cell-line cell-line)
  (printf "~a," (counter-x count))
  (set-counter-x! count (+ (counter-x count) 1))
  (if (equal? (cell-line-id cell-line) "")
      (printf "NA,")
      (printf "\"~a\"," (cell-line-id cell-line)))
  (if (equal? (cell-line-ac cell-line) "")
      (printf "NA,")
      (printf "\"~a\"," (cell-line-ac cell-line)))
  (if (equal? (cell-line-as cell-line) "")
      (printf "NA,")
      (printf "\"~a\"," (cell-line-as cell-line)))
  (if (equal? (cell-line-sy cell-line) "")
      (printf "NA,")
      (printf "\"~a\"," (cell-line-sy cell-line)))
  (if (equal? (cell-line-dr cell-line) "")
      (printf "NA,")
      (printf "\"~a\"," (cell-line-dr cell-line)))
  (if (equal? (cell-line-rx cell-line) "")
      (printf "NA,")
      (printf "\"~a\"," (cell-line-rx cell-line)))
  (if (equal? (cell-line-ww cell-line) "")
      (printf "NA,")
      (printf "\"~a\"," (cell-line-ww cell-line)))
  (if (equal? (cell-line-cc cell-line) "")
      (printf "NA,")
      (printf "\"~a\"," (cell-line-cc cell-line)))
  (if (equal? (cell-line-st cell-line) "")
      (printf "NA,")
      (printf "\"~a\"," (cell-line-st cell-line)))
  (if (equal? (cell-line-di cell-line) "")
      (printf "NA,")
      (printf "\"~a\"," (cell-line-di cell-line)))
  (if (equal? (cell-line-ox cell-line) "")
      (printf "NA,")
      (printf "\"~a\"," (cell-line-ox cell-line)))
  (if (equal? (cell-line-hi cell-line) "")
      (printf "NA,")
      (printf "\"~a\"," (cell-line-hi cell-line)))
  (if (equal? (cell-line-oi cell-line) "")
      (printf "NA,")
      (printf "\"~a\"," (cell-line-oi cell-line)))
  (if (equal? (cell-line-sx cell-line) "")
      (printf "NA,")
      (printf "\"~a\"," (cell-line-sx cell-line)))
  (if (equal? (cell-line-ca cell-line) "")
      (printf "NA\n")
      (printf "\"~a\"\n" (cell-line-ca cell-line)))
  (set-cell-line-id! cell-line "")
  (set-cell-line-ac! cell-line "")
  (set-cell-line-as! cell-line "")
  (set-cell-line-sy! cell-line "")
  (set-cell-line-dr! cell-line "")
  (set-cell-line-rx! cell-line "")
  (set-cell-line-ww! cell-line "")
  (set-cell-line-cc! cell-line "")
  (set-cell-line-st! cell-line "")
  (set-cell-line-di! cell-line "")
  (set-cell-line-ox! cell-line "")
  (set-cell-line-hi! cell-line "")
  (set-cell-line-oi! cell-line "")
  (set-cell-line-sx! cell-line "")
  (set-cell-line-ca! cell-line ""))

(define (iter in)
  (define line (read-line in))
  (unless (eof-object? line)
    (define lst (regexp-split #rx"   " line))
    (cond
      [(equal? (first lst) "ID")
       (set-cell-line-id! cline (second lst))]
      [(equal? (first lst) "AC")
       (set-cell-line-ac! cline (second lst))]
      [(equal? (first lst) "AS")
       (set-cell-line-as! cline (second lst))]
      [(equal? (first lst) "SY")
       (set-cell-line-sy! cline (second lst))]
      [(equal? (first lst) "DR")
       (cond
         [(equal? (cell-line-dr cline) "")
          (set-cell-line-dr! cline (second lst))]
         [else
          (set-cell-line-dr! cline (string-append
                                    (cell-line-dr cline) " | "
                                    (second lst)))])]
      [(equal? (first lst) "RX")
       (cond
         [(equal? (cell-line-rx cline) "")
          (set-cell-line-rx! cline (second lst))]
         [else
          (set-cell-line-rx! cline (string-append
                                    (cell-line-rx cline) " | "
                                    (second lst)))])]
      [(equal? (first lst) "WW")
       (cond
         [(equal? (cell-line-ww cline) "")
          (set-cell-line-ww! cline (second lst))]
         [else
          (set-cell-line-ww! cline (string-append
                                    (cell-line-ww cline) " | "
                                    (second lst)))])]
      [(equal? (first lst) "CC")
       (cond
         [(equal? (cell-line-cc cline) "")
          (set-cell-line-cc! cline (second lst))]
         [else
          (set-cell-line-cc! cline (string-append
                                    (cell-line-cc cline) " | "
                                    (second lst)))])]
      [(equal? (first lst) "ST")
       (cond
         [(equal? (cell-line-st cline) "")
          (set-cell-line-st! cline (second lst))]
         [else
          (set-cell-line-st! cline (string-append
                                    (cell-line-st cline) " | "
                                    (second lst)))])]
      [(equal? (first lst) "DI")
       (cond
         [(equal? (cell-line-di cline) "")
          (set-cell-line-di! cline (second lst))]
         [else
          (set-cell-line-di! cline (string-append
                                    (cell-line-di cline) " | "
                                    (second lst)))])]
      [(equal? (first lst) "OX")
       (cond
         [(equal? (cell-line-ox cline) "")
          (set-cell-line-ox! cline (second lst))]
         [else
          (set-cell-line-ox! cline (string-append
                                    (cell-line-ox cline) " | "
                                    (second lst)))])]
      [(equal? (first lst) "HI")
       (cond
         [(equal? (cell-line-hi cline) "")
          (set-cell-line-hi! cline (second lst))]
         [else
          (set-cell-line-hi! cline (string-append
                                    (cell-line-hi cline) " | "
                                    (second lst)))])]
      [(equal? (first lst) "OI")
       (cond
         [(equal? (cell-line-oi cline) "")
          (set-cell-line-oi! cline (second lst))]
         [else
          (set-cell-line-oi! cline (string-append
                                    (cell-line-oi cline) " | "
                                    (second lst)))])]
      [(equal? (first lst) "SX")
       (set-cell-line-sx! cline (second lst))]
      [(equal? (first lst) "CA")
       (set-cell-line-ca! cline (second lst))]
      [(equal? (first lst) "//")
       (print/cell-line cline)])
    (iter in)))
    
(define (convert infile outfile)
  (define in (open-input-file infile))
  (with-output-to-file outfile
    (lambda () (iter in)))
  (close-input-port in))

(convert "/path/to/text/file"
         "/path/to/csv/file")