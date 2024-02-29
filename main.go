package main

import (
	"fmt"
	"slices"
)

type BrowserHistory struct {
	History     []string
	CurrentPage int
}

func newBrowserHistory(homepage string) *BrowserHistory {
	fmt.Printf("Setting home page to %s\n", homepage)
	var browserHistory BrowserHistory
	browserHistory.History = append(browserHistory.History, homepage)
	browserHistory.CurrentPage++
	return &browserHistory
}

func (h *BrowserHistory) Visit(site string) {
	fmt.Printf("Now going to visit site %s\n", site)
	if !slices.Contains(h.History, site) {
		h.History = append(h.History, site)
		fmt.Printf("Setting current page index to %d\n", h.CurrentPage)
		h.CurrentPage++

	} else {
		fmt.Printf("%s already in slice. Staying on page %s\n", site, h.History[h.CurrentPage-1])
	}
}

// Back func allows users to go back to pages they have already visited.
// if the provided number is too large, we just go back to our homepage.
func (h *BrowserHistory) Back(size int) {
	if size > len(h.History) {
		fmt.Printf("Cannopt move %d pages back. Going to homescreen.")
		h.CurrentPage = 0
	} else {
		fmt.Printf("Going back to page %s\n", h.History[len(h.History)-size])
		h.CurrentPage = len(h.History) - size
	}
}

// Forward func allows users to move to pages they have already visited in this session
// if too high of a number is provided we default to the latest, last know page.
func (h *BrowserHistory) Forward(size int) {
	if size+h.CurrentPage >= len(h.History) {
		fmt.Printf("Cannot move %d pages forward. Going to last known page: %s\n", size, h.History[len(h.History)-1])
		h.CurrentPage = len(h.History)
	} else {
		fmt.Printf("Going forwards to page %s", h.History[size+h.CurrentPage])
		h.CurrentPage = len(h.History) + size
	}

}

func main() {

	historyAccumulator := newBrowserHistory("Dashboard")
	fmt.Printf("Current Page index")
	historyAccumulator.Visit("Google")
	historyAccumulator.Visit("Facebook")
	historyAccumulator.Visit("Twitter")
	historyAccumulator.Visit("Youtube")
	historyAccumulator.Visit("BoardGameGeek")
	historyAccumulator.Visit("lorcana")
	historyAccumulator.Back(4)
	historyAccumulator.Forward(1)
	//historyAccumulator.Visit("Google")

}
