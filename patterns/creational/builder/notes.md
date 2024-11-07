# Notes

builder seperates / abstracts the logic for creating an object from its representation
simplifies the creation of complex objects that could have several representations by encapsulating all of the code in one spot.

## advantages

from wiki
>Advantages of the builder pattern include:
>Allows you to vary a product's internal representation.
>Encapsulates code for construction and representation.
>Provides control over the steps of the construction process.

## general structure

- encapsulated logic for creating an object is seperate and not exportable
- builder has two structs
    - builder struct
    - complex object struct that is returned by builder
- builder provides constructor method that returns the newly created complex object
- Builder's setters often return themselves to enable method chaining.

## notes on example

The author in the builder-example with notifications uses a notification struct with lowercase fields. go doesn't support immutability so this was a suggested workaround to make something immutable. Since the fields aren't exported, and are only set through the builder it ensures that an importing package can't modify them. The fields are instead set through getters and setters
ie
`builder.SetSubtitle(message)`
