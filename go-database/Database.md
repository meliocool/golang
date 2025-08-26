# Go Database

## Introduction
- GoLang has a database package by default
- Package database is a package filled with standard interface to communicate to the database
- This makes the code nearly identical no matter the database
- The difference is only the SQL itself (MySQL, POSTGRESQL, SQLServer)

## Flow Database Package
- APP -> Database Interface -> Database Driver -> DBMS