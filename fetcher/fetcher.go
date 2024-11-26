package fetcher

type Fetcher struct {
}

func New() Fetcher {
	return Fetcher{}
}

func (f Fetcher) Users() {
	// TODO: Fetch users from vendor API based on the `vendorAPI.md`
}
