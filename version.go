package version

import (
    "fmt"
    "time"
    "log"
)

func V(version string) {


 fmt.Printf("Version: %s  Univrs.io\n", version)
 fmt.Printf("Current time: %s\n", time.Now())

}

func ErrNil ( err error, mesg string) {
		if err != nil {
		log.Fatalln(mesg)
	}
}