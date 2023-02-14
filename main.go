package golang

import (
    "sort"
)

func GroupAndSort(users []UserData) map[string][]UserData {
    grouped := make(map[string][]UserData, len(users))
    for _, user := range users {
        grouped[user.City] = append(grouped[user.City], user)
    }
    for _, citizens := range grouped {
        sort.Slice(citizens, func(i, j int) bool {
            return citizens[i].ID < citizens[j].ID
        })
    }
    return grouped
}
