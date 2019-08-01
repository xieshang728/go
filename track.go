package main

import (
	"fmt"
	"os"
	"sort"
	"text/tabwriter"
	"time"
)

type Track struct {
	Title string
	Artist string
	Album string
	Year int
	Length time.Duration
}

type byArtist []*Track

var tracks = [] *Track{
	{"GO","Delilah","From the Roots",2012,length("3m8s")},
	{"Go","Moby","Moby",1992,length("3m8s")},
	{"Go Ahead","Alicia Keys","As I Am",2007,length("4m36s")},
	{"Read 2 Go","Martin Solving","Smash",2011,length("4m24s")},
}

func length(s string) time.Duration{
	d,err := time.ParseDuration(s)
	if err != nil{
		panic(err)
	}
	return d
}

func printTracks(tracks [] *Track){
	const format  = "%v\t%v\t%v\t%v\t%v\t\n"
	tw := new(tabwriter.Writer).Init(os.Stdout,0,8,2,' ',0)
	fmt.Fprintf(tw,format,"Title","Artist","Album","Year","Length")
	fmt.Fprintf(tw,format,"------","------","------","------","------")
	for _,t := range tracks{
		fmt.Fprintf(tw,format,t.Title,t.Artist,t.Album,t.Year,t.Length)
	}
	tw.Flush()
}

func (b byArtist) Len() int{
	return len(b)
}

func(b byArtist) Swap(i,j int){
	b[i],b[j] = b[j],b[i]
}

func (b byArtist) Less(i,j int) bool{
	return b[i].Artist < b[j].Artist
}

func main()  {

	sort.Sort(byArtist(tracks))
	printTracks(tracks)
}