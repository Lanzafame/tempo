package tempo

import (
	"io/ioutil"
	"log"
	"os"
)

// MustGetwd returns the current working directory.
// It panics if err is not nil.
func MustGetwd() string {
	dir, err := os.Getwd()
	if err != nil {
		log.Panicf("mustgetwd: couldn't get current working directory err: %s", err)
	}
	return dir
}

// LoadTemplates takes a directory path as a string and
// returns a slice of template files to be loaded into
// the global var `templates`.
func LoadTemplates(tmpldir string) []string {
	files, err := ioutil.ReadDir(tmpldir)
	if err != nil {
		log.Panicf("loadtemplates: unable to read dir %s err: %s", tmpldir, err)
	}
	tmpls := []string{}
	for _, f := range files {
		str := tmpldir + f.Name()
		tmpls = append(tmpls, str)
	}
	return tmpls
}
