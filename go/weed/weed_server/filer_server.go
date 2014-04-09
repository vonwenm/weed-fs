package weed_server

import (
	"code.google.com/p/weed-fs/go/filer"
	"code.google.com/p/weed-fs/go/glog"
	"net/http"
	"strconv"
)

type FilerServer struct {
	port       string
	master     string
	collection string
	filer      filer.Filer
}

func NewFilerServer(r *http.ServeMux, port int, master string, dir string, collection string) (fs *FilerServer, err error) {
	fs = &FilerServer{
		master:     master,
		collection: collection,
		port:       ":" + strconv.Itoa(port),
	}

	if fs.filer, err = filer.NewFilerEmbedded(dir); err != nil {
		glog.Fatal("Can not start filer in dir:", dir)
		return
	}

	r.HandleFunc("/", fs.filerHandler)

	return fs, nil
}