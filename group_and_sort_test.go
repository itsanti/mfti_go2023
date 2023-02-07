package golang

import (
    "reflect"
    "testing"

    "github.com/stretchr/testify/assert"
)

func TestGroupAndSort(t *testing.T) {
    type testcase struct {
        name  string
        input []UserData
        want  map[string][]UserData
    }
    cases := [...]testcase{
        {
            name:  "Nil case",
            input: nil,
            want:  make(map[string][]UserData),
        },
        {
            name: "Single city",
            input: []UserData{
                {
                    ID:      "3",
                    Name:    "Rob",
                    Surname: "Pike",
                    City:    "Las Vegas",
                    Contacts: &Contacts{
                        Email: "rob.pike@golang.com",
                        Phone: "",
                    },
                },
                {
                    ID:      "1",
                    Name:    "Dmitry",
                    Surname: "Vyukov",
                    City:    "Las Vegas",
                    Contacts: &Contacts{
                        Email: "",
                        Phone: "6161616161",
                    },
                },
                {
                    ID:       "2",
                    Name:     "Unknown",
                    Surname:  "",
                    City:     "Las Vegas",
                    Contacts: nil,
                },
            },
            want: map[string][]UserData{
                "Las Vegas": {
                    {
                        ID:      "1",
                        Name:    "Dmitry",
                        Surname: "Vyukov",
                        City:    "Las Vegas",
                        Contacts: &Contacts{
                            Email: "",
                            Phone: "6161616161",
                        },
                    },
                    {
                        ID:       "2",
                        Name:     "Unknown",
                        Surname:  "",
                        City:     "Las Vegas",
                        Contacts: nil,
                    },
                    {
                        ID:      "3",
                        Name:    "Rob",
                        Surname: "Pike",
                        City:    "Las Vegas",
                        Contacts: &Contacts{
                            Email: "rob.pike@golang.com",
                            Phone: "",
                        },
                    },
                },
            },
        },
        {
            name: "Several cities",
            input: []UserData{
                {
                    ID:      "3",
                    Name:    "Rob",
                    Surname: "Pike",
                    City:    "Las Vegas",
                    Contacts: &Contacts{
                        Email: "rob.pike@golang.com",
                        Phone: "",
                    },
                },
                {
                    ID:      "1",
                    Name:    "Dmitry",
                    Surname: "Vyukov",
                    City:    "New York",
                    Contacts: &Contacts{
                        Email: "",
                        Phone: "6161616161",
                    },
                },
                {
                    ID:       "2",
                    Name:     "Unknown",
                    Surname:  "",
                    City:     "Moscow",
                    Contacts: nil,
                },
                {
                    ID:       "4",
                    Name:     "Unknown",
                    Surname:  "",
                    City:     "Moscow",
                    Contacts: nil,
                },
            },
            want: map[string][]UserData{
                "Las Vegas": {{
                    ID:      "3",
                    Name:    "Rob",
                    Surname: "Pike",
                    City:    "Las Vegas",
                    Contacts: &Contacts{
                        Email: "rob.pike@golang.com",
                        Phone: "",
                    },
                }},
                "New York": {{
                    ID:      "1",
                    Name:    "Dmitry",
                    Surname: "Vyukov",
                    City:    "New York",
                    Contacts: &Contacts{
                        Email: "",
                        Phone: "6161616161",
                    },
                }},
                "Moscow": {
                    {
                        ID:       "2",
                        Name:     "Unknown",
                        Surname:  "",
                        City:     "Moscow",
                        Contacts: nil,
                    },
                    {
                        ID:       "4",
                        Name:     "Unknown",
                        Surname:  "",
                        City:     "Moscow",
                        Contacts: nil,
                    },
                },
            },
        },
    }
    for _, tc := range cases {
        t.Run(tc.name, func(t *testing.T) {
            got := GroupAndSort(tc.input)
            assert.True(t, reflect.DeepEqual(got, tc.want))
        })
    }
}
