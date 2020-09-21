# Specification to Translate to BDD

Given a data store
initially
    it has no items

when a new Environment is added
    there are no Android nodes
    there are no iOS nodes
    getting the Android nodes returns a string 'UNKNOWN'
    getting the iOS nodes returns a string 'UNKNOWN'
    converting the Environment to string creates a table format string

when a node is added for a nonexistent environment
    an error is returned

when as Android node is added for an existing environment with no nodes
    that Environment has the added node
    the Node shows up when the Environments is converted to a table

when an Android node is added for an existing Environment
    that Environment has the added node and the original node
    both nodes show up when the Environment is converted to a table
    
when an Environment had an Android and and 2 iOS nodes
    the Environment table is laid our correctly when converted to a table

when a Node is removed for a nonexistent Environment
    an error is returned

when a Environment has 2 nodes and the 1st is removed
    the first is gone
    the second is there
    the table correct when the Environment is formatted as a table

when a Environment has 2 nodes and the 2nd is removed
    the first is there
    the second is gone
    the table correct when the Environment is formatted as a table

when a Environment has 2 nodes and both are removed
    both nodes are gone
    attempting to retrieve a node returns a string 'UNKNOWN'
    the table correct when the Environment is formatted as a table
