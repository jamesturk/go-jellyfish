#!/usr/bin/env python3
import csv
import sys
import os

if __name__ == '__main__':
    with open('testdata/{}.csv'.format(sys.argv[1])) as f:
        reader = csv.reader(f)

        for n, line in enumerate(reader):
            with open('fuzz/corpus/{}{}'.format(sys.argv[1], n), 'w') as out:
                out.write(line[0])
    os.system('go-fuzz-build github.com/jamesturk/go-jellyfish/fuzz/{}'.format(sys.argv[1]))
    os.system('go-fuzz -bin=./{0}-fuzz -corpus=./fuzz/corpus -workdir=./fuzz/{0}-workdir'.format(sys.argv[1]))
