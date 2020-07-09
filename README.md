# The package godiffpriv
### It is golang module for help developer using differential privacy in your golang projects


### How it works?



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


