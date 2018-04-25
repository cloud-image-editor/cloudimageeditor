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
	"testing"

	cloudimageeditorcore "github.com/cloud-image-editor/cloudimageeditor/core"
)

func TestImageImport(t *testing.T) {
	srcImagePath := "http://uec-images.ubuntu.com/bionic/current/bionic-server-cloudimg-amd64.img"
	destImagePath := "ubuntu.qcow2"
	err := cloudimageeditorcore.ImageImport(srcImagePath, "auto", destImagePath, "auto")
	if err != nil {
		t.Fatalf("cloudimageeditorcore.ImageImport : %s\n", err.Error())
	}
}
