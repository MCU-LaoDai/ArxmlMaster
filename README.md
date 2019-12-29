# ArxmlMaster
All-round Arxml Parse Tool for Autosar v3 v4

With the help of powerful Golang XML operation, it is possible to map arxml to
Golang struct elegantly. This repository provide the mapping as a basic.

The mapping is not related to a certain field like communication or swc port
configuration but all field defined in autosar00xx.xsd in Autosar standard. In 
principle, all kind of arxml from tool definition to ecu extract can be handled
by Arxml Master.

Since the deviation between each version of Autosar, for every version, there is
a seperate file for mapping.

## Todo List
1. A simple command line to provide some convenient function to operate arxml
2. Extarcat more meaningful information from the map struct like ecu information
, port information and signal information.
3. Test for regression
4. Examples
