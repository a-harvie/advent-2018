#lang racket
(define input (file->list "01.txt"))
(define testInput (file->list "01a_test.txt"))
(define freq 0)
(define seen (mutable-set 0))
(define duplicate #f)

(for ([i (in-naturals)]) ;loop forever
  (for ([f input])
    (set! freq (+ freq f))
    (when (set-member? seen freq) (set! duplicate freq))
    #:break (set-member? seen freq)
    (set-add! seen freq)
  )
  #:break duplicate
  (printf "Outer loop ~a: ~a ~a\n" i freq duplicate)
)
(printf "Found duplicate: ~a\n" duplicate)
