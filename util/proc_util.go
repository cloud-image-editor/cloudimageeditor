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
	"io/ioutil"
	"os"
	"strconv"
	"strings"
)

// GetProcInt64 : generic method for get int64 from /proc/xxx
func GetProcInt64(path string) (result int64, err error) {
	resultString, err := ioutil.ReadFile(path)
	if err != nil {
		return 0, err
	}

	if strings.HasSuffix(string(resultString), "\n") {
		resultString = resultString[0 : len(resultString)-1]
	}

	return strconv.ParseInt(string(resultString), 0, 64)
}

// SetProcInt64 : generic method for set int64 to /proc/xxx
func SetProcInt64(path string, value int64) (err error) {
	f, err := os.OpenFile(path, os.O_RDWR, 0644)
	if err != nil {
		return err
	}
	defer f.Close()

	valueStr := fmt.Sprintf("%v", value)
	f.WriteString(valueStr)

	return nil
}
