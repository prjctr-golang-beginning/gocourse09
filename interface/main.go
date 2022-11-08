package main

import "fmt"

func main() {
	//var a fmt.Stringer
	//var b io.Writer
	//interface embedding
	//interface casting, switch/case
	//a := &Alligator{}
	//if ca, ok := any(a).(Crawler); ok {
	//	ca.Crawl()
	//} else {
	//	fmt.Println(`No!`)
	//}
	b := &SnakeFly{}
	CheckAnimal(b)
}

func CheckAnimal(a any) {
	switch a.(type) {
	case FlyerCrawler:
		fmt.Println(`I'm FlyerCrawler'`)
	case Flyer:
		fmt.Println(`I'm Flyer'`)
	case Crawler:
		fmt.Println(`I'm Crawler'`)
	}
}

type Crawler interface {
	Crawl()
}

type Flyer interface {
	Fly()
}

type FlyerCrawler interface {
	Crawler
	Flyer
}

type Alligator struct {
}

func (a *Alligator) Crawl() {
	fmt.Println(`I'm crawling'`)
}

func (a *Alligator) Color() string {
	return `green`
}

type Owl struct {
}

func (o *Owl) Fly() {

}

type SnakeFly struct {
}

func (o *SnakeFly) Fly() {
	fmt.Println(`I'm flying'`)
}

func (a *SnakeFly) Crawl() {
	fmt.Println(`I'm crawling'`)
}
