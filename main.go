package main

import "sort"

// how to sort in go https://stackoverflow.com/a/40932847

func GroupAndSort(users []UserData) map[string][]UserData {
	var res = make(map[string][]UserData, len(users))
	for _, user := range users {
		res[user.City] = append(res[user.City], user)
	}
	for _, userData := range res {
		userDataSort(userData)
	}
	return res
}

func userDataSort(userData []UserData) {
	sort.Slice(userData, func(i, j int) bool {
		return userData[i].ID < userData[j].ID
	})
}
