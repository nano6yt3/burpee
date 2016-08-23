package main

import (
  	"flag"
	  "fmt"
	  "log"
	  "os"
	  "bufio"
    "net/http"
    "net/url"
    "crypto/tls"
	)

const usage = `
 Usage: burpee -file <file with URLs> -proxy <proxy server>
 Options:
  -file     Path to the file with URLs to visit.
  -proxy    Protocol://IP:Port (e.g. https://192.168.2.24:8081) - default http://localhost:8080.
`

func main(){
   flag.Usage = func(){fmt.Println(usage)}
   fmt.Println("Burpee started:")
    
   proxyPtr := flag.String("proxy","http://localhost:8080","The proxy server")
   pathPtr  := flag.String("file","","the path to the file with URLs - 1 per line")
   flag.Parse()

   fmt.Println("Proxy: ",*proxyPtr)
   fmt.Println("File : ",*pathPtr)

   proxyUrl, err := url.Parse(*proxyPtr)
   if err != nil {
        panic(err)
    }
   http.DefaultTransport = &http.Transport{Proxy: http.ProxyURL(proxyUrl),
   										   TLSClientConfig: &tls.Config{InsecureSkipVerify : true},
                            }

   // open a file
   if file, err := os.Open(*pathPtr); err == nil {
    	// make sure it gets closed
    	defer file.Close()
    	// create a new scanner and read the file line by line
    	scanner := bufio.NewScanner(file)
    	for scanner.Scan() {
      	  log.Println(scanner.Text())
	      resp, err := http.Get(scanner.Text())  
          if err != nil {
               fmt.Println("http.Get => %v", err.Error())
           } else {
                 defer resp.Body.Close()
                  if err != nil {
                        log.Fatal(err)
                   }
              }
   	      }

    	// check for errors
    	if err = scanner.Err(); err != nil {
    	   log.Fatal(err)
    	}

    } else { log.Fatal(err) }

}
