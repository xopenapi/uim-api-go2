package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"

	"github.com/xopenapi/uim-api-go2"
)

func main() {
	var (
		signingSecret string
	)

	flag.StringVar(&signingSecret, "secret", "YOUR_SIGNING_SECRET_HERE", "Your uim app's signing secret")
	flag.Parse()

	http.HandleFunc("/slash", func(w http.ResponseWriter, r *http.Request) {

		verifier, err := uim.NewSecretsVerifier(r.Header, signingSecret)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			return
		}

		r.Body = ioutil.NopCloser(io.TeeReader(r.Body, &verifier))
		s, err := uim.SlashCommandParse(r)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			return
		}

		if err = verifier.Ensure(); err != nil {
			w.WriteHeader(http.StatusUnauthorized)
			return
		}

		switch s.Command {
		case "/echo":
			params := &uim.Msg{Text: s.Text}
			b, err := json.Marshal(params)
			if err != nil {
				w.WriteHeader(http.StatusInternalServerError)
				return
			}
			w.Header().Set("Content-Type", "application/json")
			w.Write(b)
		default:
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
	})
	fmt.Println("[INFO] Server listening")
	http.ListenAndServe(":3000", nil)
}
