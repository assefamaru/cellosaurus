#!/bin/bash

# Read user and password interactively.
read -p "Database name (database will be created using this name): " database
read -p "MySQL User: " user
read -s -p "MySQL Password: " password
echo

# Create database and tables.
# Then load tables with csv data.
mysql -u "$user" -p"$password" <<EOF
DROP DATABASE IF EXISTS $database;
CREATE DATABASE $database;
USE $database;
# ==============================================================================
CREATE TABLE releaseInfo(
    id INT AUTO_INCREMENT primary key NOT NULL,
    attribute VARCHAR(255) NOT NULL,
    content VARCHAR(255) NOT NULL
);
LOAD DATA LOCAL INFILE '../data/csv/releaseInfo.csv' INTO TABLE releaseInfo
FIELDS TERMINATED BY ',' ENCLOSED BY '"'
IGNORE 1 LINES;
# ==============================================================================
CREATE TABLE terminologies(
    id INT AUTO_INCREMENT primary key NOT NULL,
    abbreviation VARCHAR(255) NOT NULL,
    name VARCHAR(255) NOT NULL,
    server VARCHAR(255) NOT NULL,
    db_url VARCHAR(255) NOT NULL,
    cat VARCHAR(255) NOT NULL
);
LOAD DATA LOCAL INFILE '../data/csv/terminologies.csv' INTO TABLE terminologies
FIELDS TERMINATED BY ',' ENCLOSED BY '"'
IGNORE 1 LINES;
# ==============================================================================
CREATE TABLE cells(
    acp VARCHAR(20) primary key NOT NULL,
    id VARCHAR(255) NOT NULL,
    acs VARCHAR(500),
    sy VARCHAR(500),
    sx VARCHAR(255),
    ca VARCHAR(255),
    dt VARCHAR(255),
    INDEX id (id),
    INDEX acs (acs),
    INDEX sx (sx),
    INDEX ca (ca),
    INDEX sy (sy)
);
LOAD DATA LOCAL INFILE '../data/csv/cell_lines.csv' INTO TABLE cells
FIELDS TERMINATED BY ',' ENCLOSED BY '"'
IGNORE 1 LINES;
# ==============================================================================
CREATE TABLE attributes(
    id INT AUTO_INCREMENT primary key NOT NULL,
    accession VARCHAR(20) NOT NULL,
    attribute VARCHAR(60) NOT NULL,
    content VARCHAR(1000) NOT NULL,
    FOREIGN KEY (accession) REFERENCES cells(acp)
);
LOAD DATA LOCAL INFILE '../data/csv/attributes.csv' INTO TABLE attributes
FIELDS TERMINATED BY ',' ENCLOSED BY '"'
IGNORE 1 LINES;
# ==============================================================================
CREATE TABLE refs(
    id INT AUTO_INCREMENT primary key NOT NULL,
    rx VARCHAR(200) NOT NULL,
    ra VARCHAR(500),
    rg VARCHAR(200),
    rt VARCHAR(500) NOT NULL,
    rl VARCHAR(200) NOT NULL
);
LOAD DATA LOCAL INFILE '../data/csv/references.csv' INTO TABLE refs
FIELDS TERMINATED BY ',' ENCLOSED BY '"'
IGNORE 1 LINES;
# ==============================================================================
EOF

# Remove dump file, if it exists.
dump="../data/sql/cellosaurus_api.sql"
if [ -f $dump ] ; then
    rm $dump
fi

# Create a mysql dump of the cellosaurus api database.
mysqldump -u "$user" -p"$password" cellosaurus_api > $dump