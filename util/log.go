package cloudimageeditor

/*
 * cloudimageeditor package
 *
 * Authors:
 *     zhenwei pi <cloudimageeditor@126.com>
 *
 * This project is under Apache v2 License.
 *
 */

import (
	"fmt"
)

// LogD : log with debug level
func LogD(format string, a ...interface{}) {
	fmt.Printf(format, a)
}
