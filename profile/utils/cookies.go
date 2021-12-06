package utils

import "net/http"

func CopyCookies(f http.Request, t *http.Request) {
	cookies := f.Cookies()
	for _, cookie := range cookies {
		t.AddCookie(cookie)
	}
}
