package main

import (
    "flag"
    "fmt"
    "math/rand"
    "strconv"
    "time"
)

func main() {
    // load command line args
    flag.Parse()
    var newunit, reason string
    var randIn, randOut int
    for i := 0; i < *u; i++ {
        newunit = mans[rand.Intn(len(mans))] + "-"
        newunit += hexdigits[rand.Intn(len(hexdigits))] + hexdigits[rand.Intn(len(hexdigits))] + "-"
        newunit += hexdigits[rand.Intn(len(hexdigits))] + hexdigits[rand.Intn(len(hexdigits))] + "-"
        newunit += hexdigits[rand.Intn(len(hexdigits))] + hexdigits[rand.Intn(len(hexdigits))]

        units = append(units, newunit)
    }

    fmt.Println("Setting up %v workers to run for %v seconds\n", *w, *d)
    for i := 0; i < *w; i++ {
        id = strconv.Itoa(i + 1)
        go func(id string, units []string) {
            fmt.Printf("Worker %v checking in\n", id)
            for {
                reason = reasons[rand.Intn(len(reasons))]
                if reason == "start" {
                    randIn = 0
                    randOut = 0
                } else {
                    randIn = rand.Intn(10000)
                    randOut = rand.Intn(10000)
                }

                c <- fmt.Sprintf("%s|%v|%v|%s", units[rand.Intn(len(units))], randIn, randOut, reason)
            }
        }(id, units)
    }
    t = time.After(time.Duration(*d) * time.Second)
Forever:
    for {
        select {
        case <-t:
            fmt.Println("time is up")
            break Forever
        case s = <-c:
            j++
            fmt.Println(s)
        }
    }
    fmt.Printf("got %v ticks", j)
}
