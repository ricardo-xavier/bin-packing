//http://people.brunel.ac.uk/~mastjjb/jeb/orlib/binpackinfo.html
package main

import (
    "fmt"
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
            bins := ff(instance)

            waste := 0
            for _, bin := range(bins) {
                waste += (instance.Capacity - bin.Total)
            }
            fmt.Printf("  ff %d +%d (%d)\n", len(bins), len(bins) - instance.BestSolution, waste)
            if debug > 0 {
                for _, bin := range(bins) {
                    bin.Print(debug)
                }
                for _, item := range(instance.Items) {
                    item.Print(debug)
                }
            }
        }
    }
}
