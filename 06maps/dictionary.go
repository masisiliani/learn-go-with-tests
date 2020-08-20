package main

//Dictionary type
type Dictionary map[string]string

const (
	//ErrNotFound is an error when the word does not exists in dictionary
	ErrNotFound = DictionaryErr("could not find the word you were looking for")

	//ErrWordExists occurs when word already exixts
	ErrWordExists = DictionaryErr("word already exixts")
)

//DictionaryErr struct represents error on this package
type DictionaryErr string

func (e DictionaryErr) Error() string {
	return string(e)
}

//Search method return the content of a word filter is it is in dictionary map
func (d Dictionary) Search(word string) (string, error) {
	definition, ok := d[word]

	if !ok {
		return "", ErrNotFound
	}

	return definition, nil
}

//Add insert new word in dictionary
func (d Dictionary) Add(word, definition string) error {
	_, err := d.Search(word)

	switch err {
	case ErrNotFound:
		d[word] = definition
	case nil:
		return ErrWordExists
	default:
		return err
	}

	return nil
}