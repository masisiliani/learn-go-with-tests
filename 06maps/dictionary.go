package main

//Search method return the content of a word filter is it is in dictionary map
func Search(dictionary map[string]string, word string) string {
	return dictionary[word]
}
