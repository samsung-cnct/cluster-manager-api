package website

import (
	"mime"
	"net/http"

	"github.com/philips/go-bindata-assetfs"
	"github.com/samsung-cnct/cluster-manager-api/pkg/ui/data/homepage"
	"github.com/samsung-cnct/cluster-manager-api/pkg/ui/data/protobuf"
	"github.com/samsung-cnct/cluster-manager-api/pkg/ui/data/swagger"
	"github.com/samsung-cnct/cluster-manager-api/pkg/ui/data/swaggerjson"
)

func AddWebsiteHandles(mux *http.ServeMux) {
	serveSwagger(mux)
	serveSwaggerJSON(mux)
	serveProtoBuf(mux)
	serveHomepage(mux)
}

func serveSwagger(mux *http.ServeMux) {
	mime.AddExtensionType(".svg", "image/svg+xml")

	// Expose files in third_party/swagger-ui/ on <host>/swagger-ui
	fileServer := http.FileServer(&assetfs.AssetFS{
		Asset:    swagger.Asset,
		AssetDir: swagger.AssetDir,
		Prefix:   "swagger-ui",
	})
	prefix := "/swagger-ui/"
	mux.Handle(prefix, http.StripPrefix(prefix, fileServer))
}

func serveSwaggerJSON(mux *http.ServeMux) {
	mime.AddExtensionType(".json", "application/json")

	fileServer := http.FileServer(&assetfs.AssetFS{
		Asset:    swaggerjson.Asset,
		AssetDir: swaggerjson.AssetDir,
	})
	prefix := "/swagger/"
	mux.Handle(prefix, http.StripPrefix(prefix, fileServer))
}

func serveProtoBuf(mux *http.ServeMux) {

	fileServer := http.FileServer(&assetfs.AssetFS{
		Asset:    protobuf.Asset,
		AssetDir: protobuf.AssetDir,
	})
	prefix := "/protobuf/"
	mux.Handle(prefix, http.StripPrefix(prefix, fileServer))
}

func serveHomepage(mux *http.ServeMux) {

	fileServer := http.FileServer(&assetfs.AssetFS{
		Asset:    homepage.Asset,
		AssetDir: homepage.AssetDir,
	})
	prefix := "/"
	mux.Handle(prefix, http.StripPrefix("/", fileServer))
}
