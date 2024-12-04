numRating, err := strconv.ParseFloat(strings.TrimSpace(input), 64)

	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println("added 1 to your rating:", numRating+1)
	}

###

initiallly the type of input was string..in order to add 1 to it we need to convert the string to number and then we've to remove an trailing spaces presrent after the input.