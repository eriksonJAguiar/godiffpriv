# The package godiffpriv
### Godiffpriv is a golang module for help to work differential privacy in golang projects. The module has support to queries with numeric and symbolic values. Besides, we apply the probability Laplace distribution to generate a random numbers.


### How it works?

#### Data Type:
- Numeric data should be float64 array
- Symbolic data should be string array

### How to use?

#### Install:
`
go get github.com/eriksonJAguiar/godiffpriv
`

#### Import:

`
import "github.com/eriksonJAguiar/godiffpriv"
`

#### Hello world:

- **For Numeric data:**

```	
	package main

	import (
		"encoding/json"
		"fmt"
		"github.com/eriksonJAguiar/godiffpriv"
	)
	
	func main() {
		data := []string{"Male", "Female", "Male", "Female"}
		val := godiffpriv.PrivateDataFactory(data)
		res, _ := val.applyPrivacy(1)

		var response map[string]float64

		err := json.Unmarshal(res, &response)

		if err != nil {
			fmt.Println(err.Error())
		} else {
			fmt.Println(response)
		}
	}
```

- **For Symbolic data:**

```
  	package main

	import (
		"encoding/json"
		"fmt"
		"github.com/eriksonJAguiar/godiffpriv"
	)
	
	func main() {
		data := []float64{1.5, 2.3, 7.2, 9.1}
		val := godiffpriv.PrivateDataFactory(data)
		res, _ := val.applyPrivacy(1)

		var response map[string]float64

		err := json.Unmarshal(res, &response)

		if err != nil {
			fmt.Println(err.Error())
		} else {
			fmt.Println(response)
		}
	}
```


