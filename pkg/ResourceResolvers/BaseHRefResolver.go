package ResourceResolvers

import (
	"strings"

	"github.com/maxence-charriere/go-app/v10/pkg/app"
)

type (
	BaseHRefResolverOptions struct {
		Version string
	}
)

func ResourceResolverWithBaseHRefResolverOptions(
	options BaseHRefResolverOptions,
) app.ResourceResolver {
	return versionedCacheBustingBaseHRefResourceResolver{
		version: options.Version,
	}
}

type versionedCacheBustingBaseHRefResourceResolver struct {
	app.ResourceResolver

	version string
}

func (r versionedCacheBustingBaseHRefResourceResolver) Resolve(path string) string {
	switch path {

	case
		"/",
		"/web",
		"/web/app.wasm",
		"/app.js",
		"/app-worker.js",
		"/wasm_exec.js",
		"/manifest.webmanifest",
		"/app.css",
		"/vi.css":

		path = strings.TrimPrefix(path, "/")
		path = path + "?v=" + r.version

	}

	return path
}
