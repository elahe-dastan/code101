package main

import "github.com/uber/h3-go/v4"

func main() {
	// H3 index: 613280304353771519 (resolution 8)
	//snappH3Index := h3.Cell(613280304353771519)
	//snappChildren := snappH3Index.Children(12)
	//for c := 0; c < len(snappChildren); c++ {
	//	print("'")
	//	print(snappChildren[c].String())
	//	print("',")
	//}

	//gobolH3Index := h3.Cell(613280305205215231)
	//gobolChildren := gobolH3Index.Children(12)
	//for c := 0; c < len(gobolChildren); c++ {
	//	print("'")
	//	print(gobolChildren[c].String())
	//	print("',")
	//}

	//sideGobolH3Index := h3.Cell(613280304703995903)
	//sideGobolChildren := sideGobolH3Index.Children(12)
	//for c := 0; c < len(sideGobolChildren); c++ {
	//	print("'")
	//	print(sideGobolChildren[c].String())
	//	print("',")
	//}

	//H3 index: 613280304703995903 (resolution 7)
	snappH3Index := h3.Cell(608776704730595327)
	snappChildren := snappH3Index.Children(12)
	for c := 0; c < len(snappChildren); c++ {
		print("'")
		print(snappChildren[c].String())
		print("',")
	}

	//gobolH3Index := h3.Cell(608776705586233343)
	//gobolChildren := gobolH3Index.Children(12)
	//for c := 0; c < len(gobolChildren); c++ {
	//	print("'")
	//	print(gobolChildren[c].String())
	//	print("',")
	//}

	//sideGobolH3Index := h3.Cell(608776705082916863)
	//sideGobolChildren := sideGobolH3Index.Children(12)
	//for c := 0; c < len(sideGobolChildren); c++ {
	//	print("'")
	//	print(sideGobolChildren[c].String())
	//	print("',")
	//}
}
