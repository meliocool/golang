# Repository Pattern

## Introduction
- In the book Domain Driven Design, Eric Evans said that Repository is a mechanism for encapsulating storage, retrieval, and search behavior, which emulates a collection of object
- Pattern Repository is used as a bridge between Business Logic and All the SQL Queries to the Database
- So All SQL queries is written in a Repository, and the business logic code just uses that repository

## Pattern
- Business Logic -> CALLS Repository -> USES Entity / Model
- Repository -> CALLS Repository Implementation -> CALLS Database

## Entity / Model
- Entity / Model in GoLang is represented as a Struct 
- For example if we query to a repository, it should return a converted data into struct Entity / Model 
- We can utilize the struct object itself