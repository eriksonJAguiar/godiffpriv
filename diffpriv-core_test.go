package godiffpriv

import (
	"encoding/json"
	"fmt"
	"testing"
)

func TestDiffPrivSymbolic(t *testing.T) {
	data := []string{"Unknown", "Male", "Female", "Male", "Female"}
	val := PrivateDataFactory(data)
	res, _ := val.ApplyPrivacy(1)

	var response map[string]float64

	err := json.Unmarshal(res, &response)

	if err != nil {
		fmt.Println(err.Error())
	} else {
		fmt.Println(response)
	}

}

func TestDiffPrivNumeric(t *testing.T) {
	data := []float64{1.5, 2.3, 7.2, 9.1}
	val := PrivateDataFactory(data)
	res, _ := val.ApplyPrivacy(1)

	var response map[string]float64

	err := json.Unmarshal(res, &response)

	if err != nil {
		fmt.Println(err.Error())
	} else {
		fmt.Println(response)
	}

}
