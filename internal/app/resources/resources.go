//go:generate fyne bundle --package resources --output resources_generated.go --name padLockOpenImageResource images/lock-open.svg
//go:generate fyne bundle --package resources --output resources_generated.go --name padLockClosedImageResource --append images/lock-closed.svg

package resources

var (
	PadLockOpenImageResource   = padLockOpenImageResource
	PadLockClosedImageResource = padLockClosedImageResource
)
