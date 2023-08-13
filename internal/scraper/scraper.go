package scraper

type IScraper interface {
    // run daily scraper method: get all mutations within that month
    // run initial scraper method: get all mutation within 6 months or so, depends on the banks config for how far behind we can get mutations
}

type Scraper struct {
    // account repo
    // mutation repo
    // selenium repo
}
