package helpers

import (
	"math/rand"
	"nand2tetris/vm-translator/constants"
	"runtime"
	"strings"
)

func RemoveInlineComments(s string) string {
	for strings.ContainsAny(s, constants.CommentMarker) {
		i := strings.IndexAny(s, constants.CommentMarker)
		s = s[:i]
	}

	return s
}

// Takes the full input VM file path and returns the name of the VM. Useful when
// handling static memory segment
// e.g. /home/test/Add.vm -> Add
func GetStaticLabel(fileName string) string {
	sep := "/"
	if runtime.GOOS == "windows" {
		sep = "\\"
	}

	i := strings.LastIndex(fileName, sep)
	j := strings.LastIndex(fileName, ".vm")

	return fileName[i+1 : j]
}

// Based on Paul Hankin's solution from this stack overflow question
// https://stackoverflow.com/questions/22892120/how-to-generate-a-random-string-of-a-fixed-length-in-go/22892986#22892986
func GetRandomString(length int) string {
	alphaNumeric := []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789")

	out := make([]rune, length)
	for i := range out {
		out[i] = alphaNumeric[rand.Intn(len(alphaNumeric))]
	}

	return string(out)
}
