package model

import (
    "fmt"
)

type Instance struct {
    Id string
    Capacity int
    NumItems int
    BestSolution int
    Items []Item
}

func (this Instance) Print() {
    fmt.Printf("instance %s %d %d %d\n", this.Id, this.Capacity, this.NumItems, this.BestSolution)
}
