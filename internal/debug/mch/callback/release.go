// +build !wechat_debug

package callback

import (
	"io"
	"net/http"

	"github.com/chunbaise/goutil/utxml"
)

func DebugPrintRequest(r *http.Request) {}

func DebugPrintRequestMessage(msg []byte) {}

func EncodeXMLResponseMessage(w io.Writer, msg map[string]string) (err error) {
	return utxml.EncodeXMLFromMap(w, msg, "xml")
}
