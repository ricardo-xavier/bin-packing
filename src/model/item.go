package model

import (
    "fmt"
)

type Item struct {
    Id int
    Size int
    Bin *Bin
}

func (this Item) Print(debug int) {
    if this.Bin != nil {
        fmt.Printf("item %d %d %d\n", this.Id, this.Size, this.Bin.Id)
    } else {
        fmt.Printf("item %d %d %d\n", this.Id, this.Size, -1)
    }
}
