	var symbol string

	// CALL OF THE SYMBOL
	n, err := fmt.Scanf("%v", &symbol)

	// RETURN ERROR STATUS FAILED TO SCAN ANY VALUES
	if err != nil {
		fmt.Println("failed to scan any correct symbol", n, err, symbol)
	}

	// RETURN ERROR UNKNOWN ELEMENT
	if _, ok := elements[symbol]; !ok {
		fmt.Println("unknown element symbol")
	} else {
		fmt.Println("Here is the the element you request for:", elements[symbol])
		// RETURN THE RIGHT ANSWER
	}
