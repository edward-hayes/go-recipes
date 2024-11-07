# notes

provides the ability for a subject (object that is being observed) to notify other objects that are interested in the subject that a change has occured

The other objects that are subscribed to changes are called observers

this design is also called "publish and subscribe" or "event driven system"

provides loose coupling between different parts of an architecture that have to interact with each other

Very commonly used

## advantages

- observers and subject don't need to know how the other elements are implemented. observers just register themselves with the subject and receive updates when a change occurs
- used when changes in state are not predictable and the list of observers is not known in advance

## basic structure

- subject object that maintains a list of interested observers
- subject has API for registering and unregistering observers
- subject has one (or more) API for receiving updates from the subject (ie Notifyall(), StatusChanged() etc)
- observer implements an API that is typically interface that is called by the subject (OnUpdate() etc )

## notes on example

- observer-main is the example of the observer (the object that is receiving elements)
