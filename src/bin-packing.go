//http://people.brunel.ac.uk/~mastjjb/jeb/orlib/binpackinfo.html
package main

import (
    "fmt"
    "model"
    "os"
)

func main() {

    debug := 0
    var files []string
    for _, arg := range(os.Args) {
        if arg == "-d" {
            debug++
            continue
        }
        files = append(files, arg)
    }

    for _, file := range(files) {
        dataset := Load(file)
        for _, instance := range(dataset.Instances) {

            instance.Print()

            channel_ff := make(chan []model.Bin)
            channel_bf := make(chan []model.Bin)

            go ff(instance, channel_ff)
            go bf(instance, channel_bf)

            bins := <-channel_ff
            printResult(debug, "ff", instance, bins)

            bins = <-channel_bf
            printResult(debug, "bf", instance, bins)
        }
    }
}

func printResult(debug int, algorithm string, instance model.Instance, bins []model.Bin) {
    waste := 0
    for _, bin := range(bins) {
        waste += (instance.Capacity - bin.Total)
    }
    fmt.Printf("  %s %d +%d (%d)\n", algorithm, len(bins), len(bins) - instance.BestSolution, waste)
    if debug > 0 {
        for _, bin := range(bins) {
            bin.Print(debug)
        }
        if debug > 1 {
            for _, item := range(instance.Items) {
                item.Print(debug)
            }
        }
    }
}
