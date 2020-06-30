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

func SpeedValue(value uint16) uint16 {
        return uint16(float32(value) / 0.0036)
}

func AltitudeValue(value uint16) uint16 {
        return ( value + 500 ) * 5
}

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

        // fmt.Println("inputFile:", inputFile)
        // fmt.Println("outputFile:", outputFile)



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
 
        for _, session := range activity.Sessions {
                session.SubSport = fit.SubSportVirtualActivity
                session.StartPositionLat = fit.NewLatitudeDegrees(-33.852222)
                session.StartPositionLong = fit.NewLongitudeDegrees(151.210556)
                // session.TotalDistance = 
                // session.MinSpeed = 1000
                session.AvgSpeed = SpeedValue(30)
                session.MaxSpeed = SpeedValue(30)
                session.MinAltitude = AltitudeValue(8848)
                session.AvgAltitude = AltitudeValue(8848)
                session.MaxAltitude = AltitudeValue(8848)
                session.TotalAscent = 0
                session.TotalDescent = 0
                // session.MinTemperature = 30
                session.AvgTemperature = 30
	        session.MaxTemperature = 30
                session.MinHeartRate = 160
                session.AvgHeartRate = 160
                session.MaxHeartRate = 160
                // session.MinCadence = 90
                session.AvgCadence = 90
                session.MaxCadence = 90
                // session.MinPower = 301
                session.AvgPower = 301
                session.MaxPower = 301
                // session.LeftRightBalance = 52
        }

        for _, lap := range activity.Laps {
                lap.StartPositionLat = fit.NewLatitudeDegrees(-33.852222)
                lap.StartPositionLong = fit.NewLongitudeDegrees(151.210556)
                lap.EndPositionLat = fit.NewLatitudeDegrees(-33.852222)
                lap.EndPositionLong = fit.NewLongitudeDegrees(151.210556)
                // lap.TotalDistance = 
                // lap.MinSpeed = 1000
                lap.AvgSpeed = SpeedValue(30)
                lap.MaxSpeed = SpeedValue(30)
                lap.MinAltitude = AltitudeValue(8848)
                lap.AvgAltitude = AltitudeValue(8848)
                lap.MaxAltitude = AltitudeValue(8848)
                lap.TotalAscent = 0
                lap.TotalDescent = 0
                // lap.MinTemperature = 30
                lap.AvgTemperature = 30
	        lap.MaxTemperature = 30
                lap.MinHeartRate = 160
                lap.AvgHeartRate = 160
                lap.MaxHeartRate = 160
                // lap.MinCadence = 90
                lap.AvgCadence = 90
                lap.MaxCadence = 90
                // lap.MinPower = 301
                lap.AvgPower = 301
                lap.MaxPower = 301
                lap.LeftRightBalance = 52
        }

        for _, record := range activity.Records {
                record.Power = 301
                record.Cadence = 90
                record.HeartRate = 160
                // record.LeftRightBalance = 50
                record.Speed = SpeedValue(30)
                record.Altitude = AltitudeValue(8848)
                record.PositionLat = fit.NewLatitudeDegrees(-33.852222)
                record.PositionLong = fit.NewLongitudeDegrees(151.210556)
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


        for _, record := range activity.Records {
                fmt.Printf("Timestamp: %d\n", record.Timestamp)
                fmt.Printf("PositionLat: %d\n", record.PositionLat)
                fmt.Printf("PositionLong: %d\n", record.PositionLong)
                fmt.Printf("Altitude: %d\n", record.Altitude)
                fmt.Printf("HeartRate: %d\n", record.HeartRate)
                fmt.Printf("Cadence: %d\n", record.Cadence)
                fmt.Printf("Distance: %d\n", record.Distance)
                fmt.Printf("Speed: %d\n", record.Speed)
                fmt.Printf("Power: %d\n", record.Power)
                // fmt.Printf("CompressedSpeedDistance: %d\n", record.CompressedSpeedDistance)
                // fmt.Printf("Grade: %d\n", record.Grade)
                // fmt.Printf("Resistance: %d\n", record.Resistance)
                // fmt.Printf("TimeFromCourse: %d\n", record.TimeFromCourse)
                // fmt.Printf("CycleLength: %d\n", record.CycleLength)
                // fmt.Printf("Temperature: %d\n", record.Temperature)
                // fmt.Printf("Speed1s: %d\n", record.Speed1s)
                // fmt.Printf("Cycles: %d\n", record.Cycles)
                // fmt.Printf("TotalCycles: %d\n", record.TotalCycles)
                // fmt.Printf("CompressedAccumulatedPower: %d\n", record.CompressedAccumulatedPower)
                // fmt.Printf("AccumulatedPower: %d\n", record.AccumulatedPower)
                fmt.Printf("LeftRightBalance: %d\n", record.LeftRightBalance)
                // fmt.Printf("GpsAccuracy: %d\n", record.GpsAccuracy)
                // fmt.Printf("VerticalSpeed: %d\n", record.VerticalSpeed)
                // fmt.Printf("Calories: %d\n", record.Calories)
                // fmt.Printf("VerticalOscillation: %d\n", record.VerticalOscillation)
                // fmt.Printf("StanceTimePercent: %d\n", record.StanceTimePercent)
                // fmt.Printf("StanceTime: %d\n", record.StanceTime)
                // fmt.Printf("ActivityType: %d\n", record.ActivityType)
                // fmt.Printf("LeftTorqueEffectiveness: %d\n", record.LeftTorqueEffectiveness)
                // fmt.Printf("RightTorqueEffectiveness: %d\n", record.RightTorqueEffectiveness)
                // fmt.Printf("LeftPedalSmoothness: %d\n", record.LeftPedalSmoothness)
                // fmt.Printf("RightPedalSmoothness: %d\n", record.RightPedalSmoothness)
                // fmt.Printf("CombinedPedalSmoothness: %d\n", record.CombinedPedalSmoothness)
                // fmt.Printf("Time128: %d\n", record.Time128)
                // fmt.Printf("StrokeType: %d\n", record.StrokeType)
                // fmt.Printf("Zone: %d\n", record.Zone)
                // fmt.Printf("BallSpeed: %d\n", record.BallSpeed)
                // fmt.Printf("Cadence256: %d\n", record.Cadence256)
                // fmt.Printf("FractionalCadence: %d\n", record.FractionalCadence)
                // fmt.Printf("TotalHemoglobinConc: %d\n", record.TotalHemoglobinConc)
                // fmt.Printf("TotalHemoglobinConcMin: %d\n", record.TotalHemoglobinConcMin)
                // fmt.Printf("TotalHemoglobinConcMax: %d\n", record.TotalHemoglobinConcMax)
                // fmt.Printf("SaturatedHemoglobinPercent: %d\n", record.SaturatedHemoglobinPercent)
                // fmt.Printf("SaturatedHemoglobinPercentMin: %d\n", record.SaturatedHemoglobinPercentMin)
                // fmt.Printf("SaturatedHemoglobinPercentMax: %d\n", record.SaturatedHemoglobinPercentMax)
                // fmt.Printf("DeviceIndex: %d\n", record.DeviceIndex)
                // fmt.Printf("EnhancedSpeed: %d\n", record.EnhancedSpeed)
                // fmt.Printf("EnhancedAltitude: %d\n", record.EnhancedAltitude)
                fmt.Println()
        }

        // outBuf := bufio.NewWriter(outputFile)
        outBuf, err := os.Create(outputFile)
        if err != nil {
                panic(err)
        }

        err = fit.Encode(outBuf, fitData, binary.LittleEndian)
        if err != nil {
                log.Fatalf("encode: got error, want none; error is: %v", err)
        }




        // test outputFile
        outData, err := getFitData(outputFile)
        if err != nil {
                fmt.Println(err)
                return
        }

        activity, err = outData.Activity()
        if err != nil {
                fmt.Println(err)
                return
        }

        fmt.Printf("StartPositionLat: %s\n", activity.Sessions[0].StartPositionLat)
        fmt.Printf("StartPositionLong: %s\n", activity.Sessions[0].StartPositionLong)
        fmt.Printf("TotalDistance: %d\n", activity.Sessions[0].TotalDistance)
        fmt.Printf("AvgSpeed: %d\n", activity.Sessions[0].AvgSpeed)
        fmt.Printf("MaxSpeed: %d\n", activity.Sessions[0].MaxSpeed)
        fmt.Printf("TotalAscent: %d\n", activity.Sessions[0].TotalAscent)
        fmt.Printf("TotalDescent: %d\n", activity.Sessions[0].TotalDescent)
        fmt.Printf("EnhancedAvgSpeed: %d\n", activity.Sessions[0].EnhancedAvgSpeed)
        fmt.Printf("EnhancedMaxSpeed: %d\n", activity.Sessions[0].EnhancedMaxSpeed)

        for _, record := range activity.Records {

                // fmt.Printf(" Timestamp: %d\n", record.Timestamp)
                // record.PositionLat = zrecord.PositionLat
                // record.PositionLong = zrecord.PositionLong
                // record.Speed = zrecord.Speed
                // record.Altitude = zrecord.Altitude

                // record.Speed1s = zrecord.Speed1s
                // record.Resistance = zrecord.Resistance
                // record.Distance = zrecord.Distance
                // record.GpsAccuracy = zrecord.GpsAccuracy
                // record.VerticalSpeed = zrecord.VerticalSpeed
                // record.EnhancedSpeed = zrecord.EnhancedSpeed
                // record.EnhancedAltitude = zrecord.EnhancedAltitude
                // fmt.Printf("PositionLat: %d\n", record.PositionLat)
                // fmt.Printf("PositionLong: %d\n", record.PositionLong)
                fmt.Printf("n Positions Long Lat: %d.%d\n", record.PositionLong, record.PositionLat)
                fmt.Printf("n Speed: %d\n", record.Speed)
                fmt.Printf("n Altitude: %d\n", record.Altitude)
                fmt.Printf("n Cadence: %d\n", record.Cadence)
                fmt.Printf("n HeartRate: %d\n", record.HeartRate)
                fmt.Printf("n Power: %d\n", record.Power)
                fmt.Printf("n LeftRightBalance: %d\n", record.LeftRightBalance)
        }

}
