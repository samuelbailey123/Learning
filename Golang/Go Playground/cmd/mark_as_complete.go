package cmd

func MarkAsComplete(complete string) {
	//TODO Mark the item as complete
	remove(ToDo, complete)
}

func remove(s []string, r string) []string {
	for i, v := range s {
		if v == r {
			return append(s[:i], s[i+1:]...)
		}
	}
	return s
}
