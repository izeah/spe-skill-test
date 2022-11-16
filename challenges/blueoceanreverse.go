package challenges

func BlueOcean(datas []int, filter []int) []int {
	newData := datas

	for _, f := range filter {
		newData = cleanData(newData, f)
	}

	return newData
}

func cleanData(datas []int, d int) []int {
	newData := []int{}
	for _, v := range datas {
		if v != d {
			newData = append(newData, v)
		}
	}

	return newData
}
