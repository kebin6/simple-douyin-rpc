package share

type ShareIdHandle interface {
	GetShareId() (shareIdStr string, err error)
}
