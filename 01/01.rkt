#lang racket
(foldl + 0 (file->list "01.txt"))