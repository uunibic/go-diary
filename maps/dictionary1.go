package main

import "errors"

var ErrNotFound = errors.New("could not find the word you are looking for")

type Dictionary map[string]string

func (d Dictionary) Search(word string) (string, error) {
	definition, ok := d[word]

	if !ok {
		return "", ErrNotFound
	}
	return definition, nil
}

// func Search(dictionary map[string]string, word string) string {
// 	return dictionary[word]
// }