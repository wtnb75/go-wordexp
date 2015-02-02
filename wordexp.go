package wordexp

/*
#include <stdlib.h>
#include <wordexp.h>
#include <fnmatch.h>

size_t get_wordoff(wordexp_t *we){
	return we->we_offs;
}

size_t get_wordc(wordexp_t *we){
	return we->we_wordc;
}

char *get_nth(wordexp_t *we, int n){
	if(n<we->we_wordc){
		return we->we_wordv[n];
	}
	return NULL;
}

wordexp_t *new_we(){
	return calloc(1, sizeof(wordexp_t));
}

void free_we(wordexp_t *we){
	wordfree(we);
	free(we);
}
*/
import "C"

import (
	"fmt"
)

const (
	WRDE_APPEND     = C.WRDE_APPEND
	WRDE_DOOFFS     = C.WRDE_DOOFFS
	WRDE_NOCMD      = C.WRDE_NOCMD
	WRDE_REUSE      = C.WRDE_REUSE
	WRDE_SHOWERR    = C.WRDE_SHOWERR
	WRDE_UNDEF      = C.WRDE_UNDEF
	WRDE_BADCHAR    = C.WRDE_BADCHAR
	WRDE_BADVAL     = C.WRDE_BADVAL
	WRDE_CMDSUB     = C.WRDE_CMDSUB
	WRDE_NOSPACE    = C.WRDE_NOSPACE
	WRDE_SYNTAX     = C.WRDE_SYNTAX
	FNM_NOESCAPE    = C.FNM_NOESCAPE
	FNM_PATHNAME    = C.FNM_PATHNAME
	FNM_PERIOD      = C.FNM_PERIOD
	FNM_LEADING_DIR = C.FNM_LEADING_DIR
	FNM_CASEFOLD    = C.FNM_CASEFOLD
	FNM_NOMATCH     = C.FNM_NOMATCH
)

type WordExpError struct {
	val int
}

func (w *WordExpError) Error() string {
	switch w.val {
	case WRDE_BADCHAR:
		return "Bad Character"
	case WRDE_BADVAL:
		return "Bad Variable"
	case WRDE_CMDSUB:
		return "Command execution not allowed"
	case WRDE_NOSPACE:
		return "Not enough memory to store the result"
	case WRDE_SYNTAX:
		return "Shell syntax error in words"
	default:
		return fmt.Sprintf("Unknown Error: %d", w.val)
	}
}

type FnMatchError struct {
	val int
}

func (f *FnMatchError) Error() string {
	switch f.val {
	case FNM_NOMATCH:
		return "String does not match"
	default:
		return fmt.Sprintf("Unknown Error: %d", f.val)
	}
}

func WordExp(pattern string, flags int) (res []string, err error) {
	res = []string{}
	we := C.new_we()
	defer C.free_we(we)
	retval := int(C.wordexp(C.CString(pattern), we, C.int(flags)))
	if retval != 0 {
		err = &WordExpError{val: retval}
		return
	}
	nret := int(C.get_wordc(we))
	for i := 0; i < nret; i++ {
		res = append(res, C.GoString(C.get_nth(we, C.int(i))))
	}
	return res, nil
}

func FnMatch(pattern, target string, flags int) error {
	res := int(C.fnmatch(C.CString(pattern), C.CString(target), C.int(flags)))
	if res != 0 {
		return &FnMatchError{val: res}
	}
	return nil
}
