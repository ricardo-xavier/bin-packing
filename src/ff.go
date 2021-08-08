package main

import (
    "model"
)

func ff(instance model.Instance) []model.Bin {

    var bins []model.Bin
    numBins := 0

    for i, item := range(instance.Items) {

        instance.Items[i].Bin = nil
        for b:=0; b<numBins; b++ {
            if bins[b].Total + item.Size < instance.Capacity {
                instance.Items[i].Bin = &bins[b]
                break
            }
        }

        if instance.Items[i].Bin == nil {
            var newBin model.Bin
            newBin.Id = numBins
            bins = append(bins, newBin)
            instance.Items[i].Bin = &bins[numBins]
            numBins++
        }
        instance.Items[i].Bin.Total += item.Size
        instance.Items[i].Bin.Items = append(instance.Items[i].Bin.Items, &instance.Items[i])
    }
    return bins
}
