package main

import (
    "model"
)

func ff(instance model.Instance) []model.Bin {

    var bins []model.Bin
    numBins := 0

    for _, item := range(instance.Items) {

        var bin *model.Bin
        for b:=0; b<numBins; b++ {
            if bins[b].Total + item.Size < instance.Capacity {
                bin = &bins[b]
                break
            }
        }

        if bin == nil {
            var newBin model.Bin
            bins = append(bins, newBin)
            bin = &bins[numBins]
            numBins++
        }
        bin.Total += item.Size
    }
    return bins
}
