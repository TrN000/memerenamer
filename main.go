package main

import (
	_ "embed"
	"html/template"
	"log"
	"net/http"
	"path/filepath"
	"regexp"
	"strings"

	// "net/url"
	"os"
)

//go:embed main.html
var main_html string

const imgPrefix string = "/img/"

const Port string = ":8080"

func isIn(s string, sl []string) bool {
	for _, l := range sl {
		if s == l {
			return true
		}
	}
	return false
}

func supportedFiletypes() []string {
	return []string{".jpg"}
}

func main() {
	// grab files
	dir, err := os.Getwd()
	if err != nil {
		log.Fatal(err)
	}

	files, err := os.ReadDir(dir)
	if err != nil {
		log.Fatal(err)
	}
	var images []string
	for _, s := range files {
		ext := filepath.Ext(s.Name())
		if s.Type().IsRegular() && isIn(ext, supportedFiletypes()) {
			images = append(images, s.Name())
		}
	}
	if len(images) == 0 {
		log.Fatal("no images in directory")
	}

	// handler for images
	http.Handle(imgPrefix, http.StripPrefix(imgPrefix, http.FileServer(http.Dir(dir))))

	// start server
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		// check if we're processing a post request
		if r.Method == http.MethodPost {
			log.Println("got a post")
			err = r.ParseForm()
			if err != nil {
				log.Fatalf("parse form failed with error: %v", err)
			}
			oldname := r.PostFormValue("oldname")
			ext := filepath.Ext(oldname)
			newname := r.PostFormValue("newname")
			if newname != "" {
				newname = strings.ToLower(newname)
				if !strings.HasSuffix(newname, ext) {
					newname += ext
				}
				re := regexp.MustCompile(`\s+`)
				newname = string(re.ReplaceAll([]byte(newname), []byte("_")))
				err = os.Rename(oldname, newname)
				if err != nil {
					log.Fatalf("failed to rename: %v", err)
				}
				log.Printf("successfully renamed %s to %s\n", oldname, newname)
			} else {
				log.Println("empty new name, skipping")
			}
		}

		// templating
		tmpl, err := template.New("main").Parse(main_html)
		if err != nil {
			log.Fatal(err)
		}

		currentImg := "oops, no more images"
		if len(images) > 0 {
			currentImg, images = images[0], images[1:]
		}

		// run it
		data := struct {
			Endpoint string
			Filename string
		}{
			Endpoint: imgPrefix + currentImg,
			Filename: currentImg,
		}
		err = tmpl.Execute(w, data)
		if err != nil {
			log.Fatal(err)
		}
	})

	// start server
	log.Printf("Starting server on  http://localhost%s\n", Port)
	log.Fatal(http.ListenAndServe(Port, nil))
}
