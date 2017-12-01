#!/bin/bash

# Read user and password interactively.
read -p "User: " user
read -s -p "Password: " password
echo

# Create database and tables.
# Then load tables with csv data.
mysql -u "$user" -p"$password" <<EOF
DROP DATABASE IF EXISTS cellosaurus_api;
CREATE DATABASE cellosaurus_api;

USE cellosaurus_api;

CREATE TABLE cells(
    accession VARCHAR(20) primary key NOT NULL,
    identifier VARCHAR(255) NOT NULL,
    INDEX identifier (identifier)
);

CREATE TABLE attributes(
    id INT AUTO_INCREMENT primary key NOT NULL,
    accession VARCHAR(20) NOT NULL,
    attribute VARCHAR(60) NOT NULL,
    content VARCHAR(1000) NOT NULL,
    FOREIGN KEY (accession) REFERENCES cells(accession)
);

CREATE TABLE refs(
    id INT AUTO_INCREMENT primary key NOT NULL,
    ref_identifier VARCHAR(200) NOT NULL,
    attribute VARCHAR(60) NOT NULL,
    content VARCHAR(1000) NOT NULL
);

CREATE TABLE stats(
    id INT AUTO_INCREMENT primary key NOT NULL,
    attribute VARCHAR(100) NOT NULL,
    content VARCHAR(100) NOT NULL
);

LOAD DATA LOCAL INFILE '../data/csv/cells.csv' INTO TABLE cells
FIELDS TERMINATED BY ',' ENCLOSED BY '"'
IGNORE 1 LINES;

LOAD DATA LOCAL INFILE '../data/csv/attributes.csv' INTO TABLE attributes
FIELDS TERMINATED BY ',' ENCLOSED BY '"'
IGNORE 1 LINES;

LOAD DATA LOCAL INFILE '../data/csv/refs.csv' INTO TABLE refs
FIELDS TERMINATED BY ',' ENCLOSED BY '"'
IGNORE 1 LINES;

LOAD DATA LOCAL INFILE '../data/csv/stats.csv' INTO TABLE stats
FIELDS TERMINATED BY ',' ENCLOSED BY '"'
IGNORE 1 LINES;
EOF

# Remove dump file, if it exists.
dump="../data/sql/cellosaurus_api.sql"
if [ -f $dump ] ; then
    rm $dump
fi

# Create a mysql dump of the cellosaurus api database.
mysqldump -u "$user" -p"$password" cellosaurus_api > $dump
