# Passing Data To Templates

The dot (.) in the template is the current value of the data at that point in execution.

You get to pass in one value - that's it!

Fortunately, we have many different types which that value can be including composite types which compose together values. (These are also known as aggregate data types - they aggregate together many different values).

## Slice
Use this for passing in a bunch of values of the same type. We could have a []int or a []string or a slice of any type.

## Map 
Use this for passing in key-value data.

## Struct
This is probably the most commonly used data type when passing data to templates. A struct allows you to compose together values of different types.

From https://github.com/GoesToEleven/golang-web-dev/tree/master/005_data
