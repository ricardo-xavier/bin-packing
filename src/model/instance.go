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

func (this Instance) ToString() string {
    return fmt.Sprintf("%s %d %d %d", this.Id, this.Capacity, this.NumItems, this.BestSolution)
}
