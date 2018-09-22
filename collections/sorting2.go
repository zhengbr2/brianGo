package main

import (
	"fmt"
	"os"
	"sort"
	"text/tabwriter"
	"time"
)

type Track struct {
	Title  string
	Artist string
	Album  string
	Year   int
	Length time.Duration
}

var tracks = []*Track{
	{"Go", "Delilah", "From the Roots Up", 2012, length("3m38s")},
	{"Go", "Moby", "Moby", 1992, length("3m37s")},
	{"Go Ahead", "Alicia Keys", "As I Am", 2007, length("4m36s")},
	{"Ready 2 Go", "Martin Solveig", "Smash", 2011, length("4m24s")},
}

func length(s string) time.Duration {
	d, err := time.ParseDuration(s)
	if err != nil {
		panic(s)
	}
	return d
}

func printTracks(tracks []*Track) {
	const format = "%v\t%v\t%v\t%v\t%v\t\n"
	tw := new(tabwriter.Writer).Init(os.Stdout, 0, 8, 2, ' ', 0)
	fmt.Fprintf(tw, format, "Title", "Artist", "Album", "Year", "Length")
	fmt.Fprintf(tw, format, "-----", "------", "-----", "----", "------")
	for _, t := range tracks {
		fmt.Fprintf(tw, format, t.Title, t.Artist, t.Album, t.Year, t.Length)
	}
	tw.Flush() // calculate column widths and print table
}



type defaultSort []*Track

func (x defaultSort) Len() int           { return len(x) }
func (x defaultSort) Swap(i, j int)      { x[i], x[j] = x[j], x[i] }
func (x defaultSort) Less(i, j int) bool { return x[i].Artist < x[j].Artist }


type byArtistReverse struct {
	defaultSort
}
func (x byArtistReverse) Less(i, j int) bool { return x.defaultSort[j].Artist < x.defaultSort[i].Artist }


type byYear struct {
	defaultSort
}
func (x byYear) Less(i, j int) bool { return x.defaultSort[i].Year < x.defaultSort[j].Year }


type customSort struct {
	defaultSort
	less func(x, y *Track) bool
}

func (x customSort) Less(i, j int) bool { return x.less(x.defaultSort[i], x.defaultSort[j]) }     ///  ATTENTION!!!!!!!!!!


func main() {
	fmt.Println("initial")
	printTracks(tracks)
	sort.Sort(defaultSort(tracks))
	fmt.Println("defaultSort")
	printTracks(tracks)
	//sort.Sort(sort.Reverse(defaultSort(tracks)))

	sort.Sort(byArtistReverse{tracks})
	fmt.Println("Reverse(defaultSort)")
	printTracks(tracks)
	sort.Sort(byYear{tracks})
	fmt.Println("byYear")
	printTracks(tracks)

	fmt.Println("customer Sort")
	sort.Sort(customSort{tracks, func(x, y *Track) bool {
		if x.Title != y.Title {
			return x.Title < y.Title
		}
		if x.Year != y.Year {
			return x.Year < y.Year
		}
		if x.Length != y.Length {
			return x.Length < y.Length
		}
		return false
	}})

	printTracks(tracks)
}
