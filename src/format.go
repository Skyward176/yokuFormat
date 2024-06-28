package main
import "fmt"
import "os"
import "log"

func main() {
    var files []string; 
    findFiles(".", &files)
    for _, name := range files {
        fmt.Printf("FILE: %s \n", name)
    }
}
func findFiles(dirName string, files *[]string) {//finds all md files in the current 
    path := dirName
    entries, err := os.ReadDir(dirName)
    if err != nil {
        log.Fatal(err)
    }
 
    for _, e := range entries {
        if e.Name()[0:1] != "."{// ignore files hidden with .

            if e.IsDir() == true {
                    findFiles(path + "/" + e.Name(), files)
            } else{
                *files = append(*files, (path + "/" + e.Name()) )
            }

        }
    }
}
