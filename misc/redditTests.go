package main

import (
	"fmt"
	"math/rand"
	"net/http"
  "bytes"
  "io/ioutil"
  "encoding/json"
)

type firstMain []firstData

type firstData struct {
	Data    childrenData    `json:"data"`
  // Kind    string    `json:"kind"`
}

type childrenData []postList

type postList struct {
  Hash    string    `json:"hash"`
  Title   string    `json:"title"`
}


func main() {
  link, title := getLinks()
  fmt.Println(link)
  fmt.Println(title)
}

func getLinks() (string, string) {
  var redditLink firstData

  subDomain := [3]string {"shorthairedhotties", "PrettyGirls", "beautifulwomen"}
  Domain := "http://imgur.com/"

  randNum := rand.Intn(len(subDomain))
  // randNum = 0
  // fmt.Println(randNum)
  getDomain := Domain + "r/" + subDomain[randNum] + "/new.json"

  response, err := http.Get(getDomain)
  if err != nil {
    fmt.Print("40: ")
    fmt.Println(err)
  }
  defer response.Body.Close()
  body, err := ioutil.ReadAll(response.Body)
  if err != nil {
    fmt.Print("46: ")
    fmt.Println(err)
  }
  data := bytes.TrimSpace(body)
  data = bytes.TrimPrefix(data, []byte("// "))
  err = json.Unmarshal(data, &redditLink)
  if err != nil {
    fmt.Print("53: ")
    fmt.Println(err)
  }

  linkHash := redditLink.Data[0].Hash
  linkTitle := redditLink.Data[0].Title

  finalLink := Domain + linkHash + ".jpg"

  return finalLink, linkTitle

}
