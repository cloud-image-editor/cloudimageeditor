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
	"io/ioutil"
	"strings"
)

// OS types
var (
	OSCentos  = "centos"
	OSUbuntu  = "ubuntu"
	OSUnknown = "unknown"
)

// tryToGetOSFromOSRelease : on linux, try to get os type from /etc/os-release
func tryToGetOSFromOSRelease(mntPath string) (ostype string) {
	osrelease := mntPath + "/etc/os-release"
	/*	no need import other package ...
		_, err := os.Stat(osrelease)
		if err != nil || os.IsNotExist(err) {
			return ""
		}
		cfg, err := goconfig.LoadConfigFile(osrelease)
		if err != nil {
			return ""
		}

		ostype = cfg.MustValue("", "ID", "")
	*/
	osreleaseInfo, err := ioutil.ReadFile(osrelease)
	if err != nil {
		return ""
	}
	osreleaseInfoLines := strings.Split(string(osreleaseInfo), "\n")
	for _, line := range osreleaseInfoLines {
		kv := strings.Split(line, "=")
		if len(kv) < 2 {
			continue // maybe a BUG
		}
		if "ID" == kv[0] {
			ostype = kv[1] // bingo
			break
		}
	}

	// strip prefix and suffix
	if strings.HasPrefix(ostype, "\"") {
		ostype = ostype[1:len(ostype)]
	}
	if strings.HasSuffix(ostype, "\"") {
		ostype = ostype[0 : len(ostype)-1]
	}

	return ostype
}

// OSGuess : try to guess os type
func OSGuess(mntPath string) (ostype string) {
	ostype = tryToGetOSFromOSRelease(mntPath)
	if ostype != "" {
		return ostype
	}

	return OSUnknown
}
