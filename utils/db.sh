#!/bin/bash

# Read user and password interactively
read -p "User: " user
read -s -p "Password: " password
echo

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
    attribute VARCHAR(5) NOT NULL,
    content VARCHAR(1000) NOT NULL,
    FOREIGN KEY (accession) REFERENCES cells(accession)
);

CREATE TABLE refs(
    id INT AUTO_INCREMENT primary key NOT NULL,
    ref_identifier VARCHAR(200) NOT NULL,
    attribute VARCHAR(5) NOT NULL,
    content VARCHAR(1000) NOT NULL
);

EOF
