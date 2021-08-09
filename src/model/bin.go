package model;

import (
    "fmt"
)

type Bin struct {
    Id int
    Total int
    Items []Item
}

func (this Bin) Print(debug int) {
    fmt.Printf("bin %d %d %d\n", this.Id, this.Total, len(this.Items))
    if debug > 1 {
        fmt.Print("  ")
        for _, item := range(this.Items) {
            fmt.Printf("%d ", item.Size)
        }
        fmt.Println()
    }
}
