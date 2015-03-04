# go-jellyfish

[![Build Status](https://travis-ci.org/jamesturk/go-jellyfish.svg)](https://travis-ci.org/jamesturk/go-jellyfish)

go-jellyfish is a Go library for approximate and phonetic matches of strings.

go-jellyfish is based on [the C/Python version of jellyfish](https://github.com/sunlightlabs/jellyfish).

Written by James Turk <james.p.turk@gmail.com> and released under a BSD-style license.  (See LICENSE for details.)

Porter Stemmer implementation based upon [Alex Gonopolskiy's go-stem](https://github.com/agonopol/go-stem), with permission.

## Included Algorithms

String comparison:

* Levenshtein Distance
* Damerau-Levenshtein Distance
* Jaro Distance
* Jaro-Winkler Distance
* Match Rating Approach Comparison
* Hamming Distance

Phonetic encoding:

* American Soundex
* Metaphone
* NYSIIS (New York State Identification and Intelligence System)
* Match Rating Codex

