package urlshort

import (
	"fmt"
	"net/http"
)

func mapHandler(paths map[string]string, fallback http.Handler) (http.HandlerFunc, error) {
	fmt.Println("mapHandler called")
	if len(paths) <= 0 && fallback == nil {
		fmt.Println(ErrEmptyMapNoDefault)
		return nil, ErrEmptyMapNoDefault
	}

	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		fmt.Printf("request: %v\n", r)
		fmt.Printf("paths: %v\n", paths)
		if r == nil {
			fmt.Println("No valid request, using fallback.")
			fallback.ServeHTTP(w, r)
			return
		}

		reqURLString := r.URL.String()
		fmt.Printf("request URL: %v\n", reqURLString)

		for url, redirect := range paths {
			fmt.Println("checking " + url + " " + redirect)
			if url == reqURLString {
				fmt.Printf("Redirecting %v to %v\n", url, redirect)
				w.WriteHeader(http.StatusMovedPermanently)
				w.Write([]byte(redirect))
				return
			}
		}

		fmt.Println("URL not found, using fallback")
		fallback.ServeHTTP(w, r)
		return
	}), nil
}
