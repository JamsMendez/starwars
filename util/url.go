package util

import (
	"net/url"
	"strconv"
	"strings"
)

func ParseURLToID(sUrl string) uint64 {
	rUrl, err := url.Parse(sUrl)
	if err != nil {
		return uint64(0)
	}

	paths := strings.Split(rUrl.Path, "/")

	size := len(paths)
	index := size - 2
	if index >= 0 {
		value := paths[index]

    ID, err := strconv.Atoi(value)
    if err != nil {
      return uint64(ID)
    } 

    return uint64(ID)
	}

	return uint64(0) 
}
