 package main

  import (
      "time"
      "strconv"
      "bytes"
      "html/template"
      "gitlab.com/golang-commonmark/markdown"
  )

  type otdEntry struct {
    Year int `yaml:"year"`
    Title string `yaml:"title"`
    Intro string `yaml:"intro"`
    Document string `yaml:"document"`
  }

  func (o otdEntry) Date() string {
      day := time.Now().Format("January 2")
      year := strconv.Itoa(o.Year)

      return day + ", " + year
  }

  func (o otdEntry) IntroHTML() template.HTML {
      md := markdown.New(markdown.HTML(true))

      i := md.RenderToString([]byte(o.Intro))

      return template.HTML(i)
  }

  func (o otdEntry) DocHTML() template.HTML {
      md := markdown.New(markdown.HTML(true))

      d := md.RenderToString([]byte(o.Document))

      return template.HTML(d)
  }

  func renderEntry(entry otdEntry) ([]byte, error) {
      tmpl := template.Must(template.ParseFiles("otdEntry.html"))

      var output bytes.Buffer

      err := tmpl.Execute(&output, entry)

      if err != nil {
         return nil, err
      }

      return output.Bytes(), err
  }

