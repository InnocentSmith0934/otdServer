 package main

  import (
      "fmt"
      "os"
      "time"
      "math/rand"
      "path/filepath"
      "io/ioutil"
      "strconv"
      "bytes"
      "html/template"
      "gitlab.com/golang-commonmark/markdown"
      "gopkg.in/yaml.v2"
  )

  var dirname string

  func init() {
      rand.Seed(time.Now().UnixNano())

      if value, ok := os.LookupEnv("CONTENT_DIR"); ok {
         dirname = value
      } else {
         dirname = "./content/otds/"
      }
  }

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

  func readRandomFile() ([]byte, error) {
      var files []string

      // make a slice containing names of all regular files with .yaml extension
      err := filepath.Walk(dirname, func(path string, info os.FileInfo, err error) error {
          if err != nil {
              return err
          }
          if info.Mode().IsRegular() {
              if filepath.Ext(info.Name()) == ".yaml" {
                  files = append(files, info.Name())
              }
          }
          return nil
      })

      if err != nil {
          return nil, err
      }

      i := rand.Intn(len(files))
      data, err := ioutil.ReadFile(dirname + files[i])
      return data, err
  }

  func renderEntry(entry otdEntry) ([]byte, error) {
      tmpl := template.Must(template.ParseFiles("otdentry.html"))

      var output bytes.Buffer

      err := tmpl.Execute(&output, entry)

      if err != nil {
         return nil, err
      }

      return output.Bytes(), err
  }

  func main() {

      data, err := readRandomFile()
      if err != nil {
          fmt.Println(err)
          os.Exit(1)
      }

      today := otdEntry{}
      err = yaml.Unmarshal(data, &today)
      if err != nil {
          fmt.Println(err)
          os.Exit(1)
      }

      rendered, err := renderEntry(today)
      if err != nil {
          fmt.Println(err)
          os.Exit(1)
      }

      fmt.Println(string(rendered))
  }
