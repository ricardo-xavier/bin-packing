package main

import (
    "model"
)

func bf(instance model.Instance, channel chan []model.Bin) {

    var bins []model.Bin
    numBins := 0

    for _, item := range(instance.Items) {

        var bin *model.Bin
        minWaste := instance.Capacity
        for b:=0; b<numBins; b++ {
            if bins[b].Total + item.Size < instance.Capacity {
                waste := instance.Capacity - bins[b].Total
                if waste < minWaste { 
                    bin = &bins[b]
                    minWaste = waste
                }
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
