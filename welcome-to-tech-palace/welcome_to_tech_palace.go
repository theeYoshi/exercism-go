package techpalace

// WelcomeMessage returns a welcome message for the customer.
func WelcomeMessage(customer string) string {
	var answer string
	answer = "Welcome to the Tech Palace, " + customer
	return answer
}

// AddBorder adds a border to a welcome message.
func AddBorder(welcomeMsg string, numStarsPerLine int) string {
	var result string
	for i := 0; i < numStarsPerLine; i++ {
		result += "*"
	}
	result += "\n"
	result += welcomeMsg + "\n"
	for i := 0; i < numStarsPerLine; i++ {
		result += "*"
	}
	return result
}

// CleanupMessage cleans up an old marketing message.
func CleanupMessage(oldMsg string) string {
	panic("Please implement the CleanupMessage() function")
}
