package binpath

func Base() string {
	_, base := Path()
	return base
}
