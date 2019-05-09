 package main

  import (
      "fmt"
      "os"
      "time"
      "math/rand"
      "path/filepath"
      "io/ioutil"
      "strconv"
      "gitlab.com/golang-commonmark/markdown"
      "gopkg.in/yaml.v2"
  )

  type otdEntry struct {
    Year int `yaml:"year"`
    Title string `yaml:"title"`
    Intro string `yaml:"intro"`
    Document string `yaml:"document"`
  }

  func (o otdEntry) date() string {
      day := time.Now().Format("January 2")
      year := strconv.Itoa(o.Year)

      return day + ", " + year
  }

  var dirname string = "./content/otds/"

  func init() {
      rand.Seed(time.Now().UnixNano())
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

  func renderEntry(entry otdEntry) []byte {
      md := markdown.New(markdown.HTML(true))

      day := entry.date()
      title := entry.Title
      intro := md.RenderToString([]byte(entry.Intro))
      document := md.RenderToString([]byte(entry.Document))

      output := "<h2>" + day + "</h2><h3>" + title +  "</h3>" + intro + "<blockquote>" + document + "</blockquote>"

      return []byte(output)
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

      rendered := renderEntry(today)
      fmt.Println(string(rendered))
  }
