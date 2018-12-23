package converter

import (
	"fmt"
	"testing"
)

func Test_StringSlice2Cvs(t *testing.T) {
	origin := []string{"first_name", "last_name", "username"}
	result, err := StringSlice2Cvs(origin)
	if err != nil {
		t.Error(err)
	}
	fmt.Printf("%s\n", result)
}

func Test_Cvs2StringSlice(t *testing.T) {
	in := `first_name,last_name,username
"Rob","Pike",rob
Ken,Thompson,ken
"Robert","Griesemer","gri"
`
	result, err := Cvs2StringSlice(in)
	if err != nil {
		t.Error(err)
	}
	fmt.Printf("%v\n", result)

}
