#!/bin/bash

if [ $# -ne 2 ]; then
  echo "Usage: ${0} <py_source> <testdata>" 1>&2
  exit 1
fi

cat ${2} | python ${1}
