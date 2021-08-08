package main

import (
    "bufio"
    "model"
    "os"
    "strconv"
    "strings"
)

func Load(fileName string) model.Dataset {

    var dataset model.Dataset
    file, err := os.Open(fileName)
    if err != nil {
        panic(err)
    }

    scanner := bufio.NewScanner(file)

    scanner.Scan()
    text := scanner.Text()
    numInstances, _ := strconv.Atoi(text)

    for i:=0; i<numInstances; i++ {

        var instance model.Instance

        scanner.Scan()
        instance.Id = strings.TrimSpace(scanner.Text())

        scanner.Scan()
        text = strings.TrimSpace(scanner.Text())
        fields := strings.Split(text, " ")

        instance.Capacity, _ = strconv.Atoi(fields[0])
        instance.NumItems, _ = strconv.Atoi(fields[1])
        instance.BestSolution, _ = strconv.Atoi(fields[2])

        for j:=0; j<instance.NumItems; j++ {
            scanner.Scan()
            text = scanner.Text()
            size, _ := strconv.Atoi(text)
            var item model.Item
            item.Size = size
            instance.Items = append(instance.Items, item)
        }

        dataset.Instances = append(dataset.Instances, instance)
    }

    file.Close()
    return dataset
}
