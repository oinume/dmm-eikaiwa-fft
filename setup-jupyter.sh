#!/bin/sh

if [ ! -d .jupyter ]; then
    virtualenv --distribute --python /usr/local/bin/python3 .jupyter
    . .jupyter/bin/activate
    pip install -r requirements-jupyter.txt
fi
