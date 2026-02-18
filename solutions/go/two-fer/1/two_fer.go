package twofer

// ShareWith takes a name and return a sentence
func ShareWith(name string) string {
	if name == ""{
        name = "you"
    }
	return "One for " + name +", one for me."
}
