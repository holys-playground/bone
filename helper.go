/********************************
*** Multiplexer for Go        ***
*** Bone is under MIT license ***
*** Code by CodingFerret      ***
*** github.com/go-zoo         ***
*********************************/

package bone

import (
	"fmt"
	"net/http"
	"net/url"
)

// Handle when a request does not match a registered handler.
func (m *Mux) HandleNotFound(rw http.ResponseWriter, req *http.Request) {
	if m.notFound != nil {
		rw.WriteHeader(http.StatusNotFound)
		m.notFound(rw, req)
	} else {
		http.NotFound(rw, req)
	}
}

// Check if the path don't end with a /
func (m *Mux) valid(path string) bool {
	plen := len(path)
	if plen > 1 && path[plen-1:] == "/" {
		return false
	}
	return true
}

// Check if the request path is for Static route
func (m *Mux) isStatic(p string) (string, bool) {
	for k, s := range m.Static {
		if len(p) >= s.Size && p[:s.Size] == s.Path {
			return k, true
		}
		continue
	}
	return "", false
}

// Debugging function
func (m *Mux) inspect(meth string) {
	for i, r := range m.Routes[meth] {
		fmt.Printf("#%d => %s\n", i+1, r.Path)
	}
}

// Insert the url value into the vars stack
func (r *Route) insert(req *http.Request, uv url.Values) {
	for k := range uv {
		if vars[req] == nil {
			vars[req] = make(map[string]string)
		}
		vars[req][k] = uv.Get(k)
	}

}

// GetValue Return the key value, of the current *http.Request
func GetValue(req *http.Request, key string) string {
	return vars[req][key]
}
