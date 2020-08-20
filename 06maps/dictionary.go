package main

//Dictionary type
type Dictionary map[string]string

//Search method return the content of a word filter is it is in dictionary map
func (d Dictionary) Search(dictionary map[string]string, word string) string {
	return d[word]
}
