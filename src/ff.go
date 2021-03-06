package main

import (
    "model"
)

func ff(instance model.Instance, channel chan []model.Bin) {

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
            newBin.Id = numBins
            bins = append(bins, newBin)
            bin = &bins[numBins]
            numBins++
        }
        bin.Total += item.Size
        bin.Items = append(bin.Items, item)
    }
    channel <- bins
}
