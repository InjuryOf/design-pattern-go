package struct_type

type server interface {
	handlerRequest(string, string) (int, string)
}

type nginx struct {
	application       *application
	maxAllowedRequest int
	rateLimiter       map[string]int
}

func newNginxServer() *nginx {
	return &nginx{
		application:       &application{},
		maxAllowedRequest: 2,
		rateLimiter:       make(map[string]int),
	}
}

func (n *nginx) handlerRequest(url, method string) (int, string) {
	if !n.checkRateLimiter(url) {
		return 403, "Not Allowed"
	}
	return n.application.handlerRequest(url, method)
}

func (n *nginx) checkRateLimiter(url string) bool {
	if n.rateLimiter[url] == 0 {
		n.rateLimiter[url] = 1
	}
	if n.rateLimiter[url] > n.maxAllowedRequest {
		return false
	}

	n.rateLimiter[url] = n.rateLimiter[url] + 1
	return true
}

type application struct {
}

func (a *application) handlerRequest(url, method string) (int, string) {
	if url == "create" && method == "post" {
		return 200, "OK"
	}
	if url == "search" && method == "get" {
		return 201, "search success"
	}
	return 404, "url not found"
}
