package main

import (
        "bytes"
        "flag"
        "fmt"
        "github.com/tormoder/fit"
        "io/ioutil"
        "log"
        "os"
)

var (
        inputFile string
        outputFile string
)

func fileExists(filename string) bool {
        info, err := os.Stat(filename)
        if err != nil {
                if os.IsNotExist(err) {
                        return false
                }
                return true
        }
        return !info.IsDir()
}

// func bail(err error) {
// 	if err != nil {
// 		log.Fatalf("%s", err.Error())
// 	}
// }

func main() {
        // remove all flags from log output
        log.SetFlags(0)

        flag.StringVar(&inputFile, "i", "", "a string var")
        flag.StringVar(&outputFile, "o", "", "a string var")

        flag.Parse()

        if !fileExists(inputFile) {
                log.Fatal("No input file found: " + inputFile)
        }

        if fileExists(outputFile) {
                log.Fatal("Output file already exists: " + outputFile)
        }

        fmt.Println("inputFile:", inputFile)
        fmt.Println("outputFile:", outputFile)



        // Read our FIT test file data
        testData, err := ioutil.ReadFile(inputFile)
        if err != nil {
                fmt.Println(err)
                return
        }

        // Decode the FIT file data
        fit, err := fit.Decode(bytes.NewReader(testData))
        if err != nil {
                fmt.Println(err)
                return
        }

        // Inspect the TimeCreated field in the FileId message
        fmt.Println(fit.FileId.TimeCreated)

        // // Inspect the dynamic Product field in the FileId message
        // fmt.Println(fit.FileId.GetProduct())

        // // Inspect the FIT file type
        // fmt.Println(fit.FileType())

        // // Get the actual activity
        // activity, err := fit.Activity()
        // if err != nil {
        //         fmt.Println(err)
        //         return
        // }

        // // Print the latitude and longitude of the first Record message
        // for _, record := range activity.Records {
        //         fmt.Println(record.PositionLat)
        //         fmt.Println(record.PositionLong)
        //         break
        // }

        // // Print the sport of the first Session message
        // for _, session := range activity.Sessions {
        //         fmt.Println(session.Sport)
        //         break
        // }

}
