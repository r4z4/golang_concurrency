#!/bin/sh

# container post_create_script

cd /tmp
git clone --branch v0.4.4 https://github.com/pgvector/pgvector.git
cd pgvector
make
make install 