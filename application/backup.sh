#!/bin/bash

cd $1
sqlite3 test.db <<EOF
.output backup.sql
.dump
.exit
EOF

