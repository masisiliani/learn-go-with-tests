package main

import "errors"

//Dictionary type
type Dictionary map[string]string

//ErrNotFound is an error when the word does not exists in dictionary
var ErrNotFound = errors.New("could not find the word you were looking for")

//Search method return the content of a word filter is it is in dictionary map
func (d Dictionary) Search(word string) (string, error) {
	definition, ok := d[word]

	if !ok {
		return "", ErrNotFound
	}

	return definition, nil

}
