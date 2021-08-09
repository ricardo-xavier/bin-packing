package model

import (
    "fmt"
)

type Item struct {
    Id int
    Size int
}

func (this Item) Print(debug int) {
    fmt.Printf("item %d %d\n", this.Id, this.Size)
}
