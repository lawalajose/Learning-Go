//--Summary:
//  Create a system monitoring dashboard using the existing dashboard
//  component structures. Each array element in the components represent
//  a 1-second sampling.
//
//--Requirements:
//* Create functions to calculate averages for each dashboard component
//* Using struct embedding, create a Dashboard structure that contains
//  each dashboard component
//* Print out a 5-second average from each component using promoted
//  methods on the Dashboard

package main

import (
	"fmt"
)

type Bytes int
type Celcius float32

type BandwidthUsage struct {
	amount []Bytes
}

type CpuTemp struct {
	temp []Celcius
}

type MemoryUsage struct {
	amount []Bytes
}
type Dashboard struct {
	BandwidthUsage
	CpuTemp
	MemoryUsage
}

func (c *CpuTemp) CpuAverage() int {
	sum := 0
	for _, b := range c.temp {
		sum += int(b)
	}
	average := sum / len(c.temp)
	return average
}
func (b *BandwidthUsage) BandAverage() int {
	sum := 0
	for _, r := range b.amount {
		sum += int(r)
	}
	average := sum / len(b.amount)
	return average
}

func (m *MemoryUsage) MemoryAverage() int {
	sum := 0
	for _, r := range m.amount {
		sum += int(r)
	}
	average := sum / len(m.amount)
	return average
}

func (d Dashboard) String() string {
	return fmt.Sprintf("BandwidthUsage: %v \nCpuTemp: %v \n MemoryUsage: %v", d.BandwidthUsage, d.CpuTemp, d.MemoryUsage)
}

func main() {
	bandwidth := BandwidthUsage{[]Bytes{50000, 100000, 130000, 80000, 90000}}
	temp := CpuTemp{[]Celcius{50, 51, 53, 51, 52}}
	memory := MemoryUsage{[]Bytes{800000, 800000, 810000, 820000, 800000}}
	dashboard := Dashboard{
		BandwidthUsage: bandwidth,
		CpuTemp:        temp,
		MemoryUsage:    memory,
	}

	averageBand := dashboard.BandAverage()
	fmt.Println(averageBand)

	fmt.Println(dashboard)

}
