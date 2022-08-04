zero value of a variable is the empty or default value for a variable type. When you declare a variable without a value it is by default assigned a zero value.

Go does this for only core types

| **Type** | Zero value |
| ----------- | ----------- |
| Numbers (integers and floats) | 0 |
| String | "" |
| Pointers, functions, interfaces, slices, channels, and maps | nil |


We're using fmt.Printf to help us in this exercise as we can get more
detail about a value's type. fmt.Printf uses a template language that allows us to
transform passed values. The substitution we're using is %#v. This transformation is a
useful tool for showing a variable's value and type. Some other common substitutions
you can try are as follows:

| **Substitution** | **Formatting** |
| ----------- | ----------- |
| %v | Any value regardless of the type |
| %+v | Values with extra information such as struct field names |
| %#v | Go syntax such as %+v with the addition of the name of the type of the variable |
| %T | Print the variable's type |
| %d | Decimal (base 10) |
| %s | String |