package utils

import "strconv"

func Pagination(page string, itemPerPage string) (int, int) {
	cPage, _ := strconv.Atoi(page)
	limit, _ := strconv.Atoi(itemPerPage)
	offset := limit * (cPage - 1)
	return offset, limit
}
