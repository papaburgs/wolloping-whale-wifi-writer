package main

import (
    "flag"
    "time"
)

var manMap map[string]string
var d = flag.Int("duration", 2, "number of seconds to run for")
var w = flag.Int("workers", 1, "number of workers")
var u = flag.Int("units", 100, "number of units to mock")
var units []string
var c = make(chan string)
var s, id string
var t <-chan time.Time
var j int

var hexdigits = []string{"1", "2", "3", "4", "5", "6", "7", "8", "9", "a", "b", "c", "d", "e", "f"}
var mans = []string{}
var reasons = []string{"start", "stop"}

func init() {
    manMap = make(map[string]string)
    manMap["fj-01-01"] = "Samsung S6"
    manMap["fj-01-02"] = "Samsung Note"
    manMap["fj-01-03"] = "Samsung S5"
    manMap["fj-01-04"] = "Samsung S2"
    manMap["fj-01-05"] = "Samsung J3"
    manMap["fj-01-06"] = "Samsung Tab"
    manMap["fj-02-01"] = "HTC OneM9"
    manMap["fj-02-02"] = "HTC OneM8"
    manMap["fj-03-01"] = "Blackberry Priv"
    manMap["fj-03-02"] = "Blackberry Storm"
    manMap["fj-03-03"] = "Blackberry Classic"
    manMap["fj-04-01"] = "Apple IPhone7"
    manMap["fj-04-02"] = "Apple IPhone6S"
    manMap["fj-04-03"] = "Apple IPhone5"
    manMap["fj-04-04"] = "Apple IPhone6SPlus"
    manMap["fj-04-05"] = "Apple IPhone7Plus"
    manMap["fj-04-06"] = "Apple IPad"
    manMap["fj-05-01"] = "LG Nexus"
    manMap["fj-05-02"] = "LG G5"
    manMap["fj-06-01"] = "Motorola XPlay"
    manMap["fj-06-02"] = "Motorola Force"
    manMap["fj-06-01"] = "Novatel Mifi"

    for x, _ := range manMap {
        mans = append(mans, x)
    }

}
