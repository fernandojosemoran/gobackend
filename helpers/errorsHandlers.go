package helpers

import "log"

func PanicErrorHandler(err error) {
	if err != nil {
		panic(err)
	}
}

func ResolveErrorHandler(err error) {

}

func LogFatalHandler(msg string, err error) {
	if err != nil {
		log.Fatal("\n"+msg+"\n", err)
	} else {
		log.Printf("\nDATABASE CONNECTED SUCESSFULLY\n")
	}
}
