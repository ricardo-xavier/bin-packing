//http://people.brunel.ac.uk/~mastjjb/jeb/orlib/binpackinfo.html
package main

import (
    "fmt"
    "os"
)

func main() {

    dataset := Load(os.Args[1])
    for _, instance := range(dataset.Instances) {

        fmt.Println(instance.ToString())
        bins := ff(instance)

        waste := 0
        for _, bin := range(bins) {
            waste += (instance.Capacity - bin.Total)
        }
        fmt.Printf("  ff %d +%d (%d)\n", len(bins), len(bins) - instance.BestSolution, waste)
    }
}
