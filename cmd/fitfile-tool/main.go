package main

import (
        // "bufio"
        "bytes"
        "encoding/binary"
        "flag"
        "fmt"
        "github.com/tormoder/fit"
        "io/ioutil"
        "log"
        "os"
)

var (
        inputFile string
        zwiftFile string
        outputFile string
        // zwiftFlag bool = false
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

func getFitData(filename string) (*fit.File, error) {
        // Read our FIT test file data
        inData, err := ioutil.ReadFile(filename)
        if err != nil {
                return nil, err
        }

        // Decode the FIT file data
        fitData, err := fit.Decode(bytes.NewReader(inData))
        if err != nil {
                return nil, err
        }
        return fitData, nil
}

func main() {
        // remove all flags from log output
        log.SetFlags(0)

        flag.StringVar(&inputFile, "i", "", "a string var")
        flag.StringVar(&outputFile, "o", "", "a string var")
        flag.StringVar(&zwiftFile, "z", "", "a string var")
        // flag.BoolVar(&zwiftFlag, "z", true, "Zwift data")

        flag.Parse()

        if !fileExists(inputFile) {
                log.Fatal("No input file found: " + inputFile)
        }

        if zwiftFile != "" && !fileExists(zwiftFile) {
                log.Fatal("No Zwift file found: " + zwiftFile)
        }

        if fileExists(outputFile) {
                log.Fatal("Output file already exists: " + outputFile)
        }

        fmt.Println("inputFile:", inputFile)
        fmt.Println("outputFile:", outputFile)



        // Read our FIT test file data
        // inData, err := ioutil.ReadFile(inputFile)
        // if err != nil {
        //         fmt.Println(err)
        //         return
        // }

        // // Decode the FIT file data
        // fitData, err := fit.Decode(bytes.NewReader(inData))
        // if err != nil {
        //         fmt.Println(err)
        //         return
        // }
        fitData, err := getFitData(inputFile)
        if err != nil {
                fmt.Println(err)
                return
        }

        // Inspect the TimeCreated field in the FileId message
        fmt.Println(fitData.FileId.TimeCreated)

        // // Inspect the dynamic Product field in the FileId message
        // fmt.Println(fitData.FileId.GetProduct())

        // // Inspect the FIT file type
        // fmt.Println(fitData.FileType())

        // Get the actual activity
        activity, err := fitData.Activity()
        if err != nil {
                fmt.Println(err)
                return
        }

        if zwiftFile != "" {
                zwiftData, err := getFitData(zwiftFile)
                if err != nil {
                        fmt.Println(err)
                        return
                }
                fmt.Println(zwiftData.FileId.TimeCreated)

                zActivity, err := zwiftData.Activity()
                if err != nil {
                        fmt.Println(err)
                        return
                }

                if len(activity.Sessions) == len(zActivity.Sessions) && len(zActivity.Sessions) == 1 {
                        activity.Sessions[0].StartPositionLat = zActivity.Sessions[0].StartPositionLat
                        activity.Sessions[0].StartPositionLong = zActivity.Sessions[0].StartPositionLong
                        activity.Sessions[0].TotalDistance = zActivity.Sessions[0].TotalDistance
                        activity.Sessions[0].AvgSpeed = zActivity.Sessions[0].AvgSpeed
                        activity.Sessions[0].MaxSpeed = zActivity.Sessions[0].MaxSpeed
                        activity.Sessions[0].TotalAscent = zActivity.Sessions[0].TotalAscent
                        activity.Sessions[0].TotalDescent = zActivity.Sessions[0].TotalDescent
                        // activity.Sessions[0].AvgAltitude = zActivity.Sessions[0].AvgAltitude
                        // activity.Sessions[0].MaxAltitude = zActivity.Sessions[0].MaxAltitude
                        // activity.Sessions[0].GpsAccuracy = zActivity.Sessions[0].GpsAccuracy
                        // activity.Sessions[0].AvgPosVerticalSpeed = zActivity.Sessions[0].AvgPosVerticalSpeed
                        // activity.Sessions[0].AvgNegVerticalSpeed = zActivity.Sessions[0].AvgNegVerticalSpeed
                        // activity.Sessions[0].MaxPosVerticalSpeed = zActivity.Sessions[0].MaxPosVerticalSpeed
                        // activity.Sessions[0].MaxNegVerticalSpeed = zActivity.Sessions[0].MaxNegVerticalSpeed
                        // activity.Sessions[0].MinAltitude = zActivity.Sessions[0].MinAltitude
                        activity.Sessions[0].EnhancedAvgSpeed = zActivity.Sessions[0].EnhancedAvgSpeed
                        activity.Sessions[0].EnhancedMaxSpeed = zActivity.Sessions[0].EnhancedMaxSpeed
                        // activity.Sessions[0].EnhancedAvgAltitude = zActivity.Sessions[0].EnhancedAvgAltitude
                        // activity.Sessions[0].EnhancedMinAltitude = zActivity.Sessions[0].EnhancedMinAltitude
                        // activity.Sessions[0].EnhancedMaxAltitude = zActivity.Sessions[0].EnhancedMaxAltitude
                        // activity.Sessions[0].AvgVam = zActivity.Sessions[0].AvgVam
                        fmt.Printf("StartPositionLat: %s\n", zActivity.Sessions[0].StartPositionLat)
                        fmt.Printf("StartPositionLong: %s\n", zActivity.Sessions[0].StartPositionLong)
                        fmt.Printf("TotalDistance: %d\n", zActivity.Sessions[0].TotalDistance)
                        fmt.Printf("AvgSpeed: %d\n", zActivity.Sessions[0].AvgSpeed)
                        fmt.Printf("MaxSpeed: %d\n", zActivity.Sessions[0].MaxSpeed)
                        fmt.Printf("TotalAscent: %d\n", zActivity.Sessions[0].TotalAscent)
                        fmt.Printf("TotalDescent: %d\n", zActivity.Sessions[0].TotalDescent)
                        // fmt.Printf("AvgAltitude: %d\n", zActivity.Sessions[0].AvgAltitude)
                        // fmt.Printf("MaxAltitude: %d\n", zActivity.Sessions[0].MaxAltitude)
                        // fmt.Printf("GpsAccuracy: %d\n", zActivity.Sessions[0].GpsAccuracy)
                        // fmt.Printf("AvgPosVerticalSpeed: %d\n", zActivity.Sessions[0].AvgPosVerticalSpeed)
                        // fmt.Printf("AvgNegVerticalSpeed: %d\n", zActivity.Sessions[0].AvgNegVerticalSpeed)
                        // fmt.Printf("MaxPosVerticalSpeed: %d\n", zActivity.Sessions[0].MaxPosVerticalSpeed)
                        // fmt.Printf("MaxNegVerticalSpeed: %d\n", zActivity.Sessions[0].MaxNegVerticalSpeed)
                        // fmt.Printf("MinAltitude: %d\n", zActivity.Sessions[0].MinAltitude)
                        fmt.Printf("EnhancedAvgSpeed: %d\n", zActivity.Sessions[0].EnhancedAvgSpeed)
                        fmt.Printf("EnhancedMaxSpeed: %d\n", zActivity.Sessions[0].EnhancedMaxSpeed)
                        // fmt.Printf("EnhancedAvgAltitude: %d\n", zActivity.Sessions[0].EnhancedAvgAltitude)
                        // fmt.Printf("EnhancedMinAltitude: %d\n", zActivity.Sessions[0].EnhancedMinAltitude)
                        // fmt.Printf("EnhancedMaxAltitude: %d\n", zActivity.Sessions[0].EnhancedMaxAltitude)
                        // fmt.Printf("AvgVam: %d\n", zActivity.Sessions[0].AvgVam)
                // for _, session := range activity.Sessions {
                //         fmt.Println(session.Sport)
                //         break

                        // fmt.Print("Single session")

                        for _, record := range activity.Records {

                                for _, zrecord := range zActivity.Records {

                                        if zrecord.Timestamp == record.Timestamp {
                                                fmt.Print(".")
                                                // fmt.Printf(" Timestamp: %d\n", record.Timestamp)
                                                record.PositionLat = zrecord.PositionLat
                                                record.PositionLong = zrecord.PositionLong
                                                record.Speed = zrecord.Speed
                                                record.Altitude = zrecord.Altitude

                                                record.Speed1s = zrecord.Speed1s
                                                record.Resistance = zrecord.Resistance
                                                record.Distance = zrecord.Distance
                                                record.GpsAccuracy = zrecord.GpsAccuracy
                                                record.VerticalSpeed = zrecord.VerticalSpeed
                                                record.EnhancedSpeed = zrecord.EnhancedSpeed
                                                record.EnhancedAltitude = zrecord.EnhancedAltitude
                                                // fmt.Printf(" EnhancedSpeed: %d\n", record.PositionLong)
                                                break
                                        }
                                }
                        }
                }
        }


        // if zwiftFlag {
        //         // Print the latitude and longitude of the first Record message
        //         for _, record := range activity.Records {
        //                 // fmt.Println(record.PositionLat)
        //                 // fmt.Println(record.PositionLong)
        //                 record.Speed = 30
        //                 record.Altitude = 15
        //                 // break
        //         }
        // }

        // // Print the sport of the first Session message
        // for _, session := range activity.Sessions {
        //         fmt.Println(session.Sport)
        //         break
        // }

        // outBuf := bufio.NewWriter(outputFile)
        outBuf, err := os.Create(outputFile)
        if err != nil {
                panic(err)
        }

        err = fit.Encode(outBuf, fitData, binary.LittleEndian)
        if err != nil {
                log.Fatalf("encode: got error, want none; error is: %v", err)
        }

}
