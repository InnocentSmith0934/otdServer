 package main

  import (
      "os"
      "time"
      "math/rand"
      "path/filepath"
      "io/ioutil"
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

  func otdRand() ([]byte, error) {
      data, err := readRandomFile()
      if err != nil {
         return nil, err
      }

      today := otdEntry{}
      err = yaml.Unmarshal(data, &today)
      if err != nil {
         return nil, err
      }

      rendered, err := renderEntry(today)
      if err != nil {
         return nil, err
      }

      return rendered, nil
  }
