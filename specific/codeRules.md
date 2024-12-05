# My Coding differences
So, I'm green in GOlang and go-projects architecture.

### Naming
For me was very unexpectable to recognize, package's 
files can be named by any words. 
My [XXX] idea for trying to write this project:
divide any package on 2 parts:
 - Common structure/interface ```Model```
 - Implementation ```Impl```

And any file in project (except the ```main.go```) has 
_Model_ or _Impl_ suffix. It helps me to understand and memorize
the content-table of any package. It seems like headers in C/++ or
Java project structure, when the complex of entities must be implemented and
not thrashed, (as far as I know...).

### Model insights
Model-files inside project contains 3 type of entities.
 - Interfaces. Describe and represent API of package,
 - Structures. Main enitities for with the functions includes,
 - Constructor. If package represents the service or background-worker,
it also provides the Contructor function that returns structure and service's API.

### Impl insights
Impl-files inside project contains only functions implification.
Every Impl-file represents only once interface implification 
from Modeling file.
 - postgresModel.go (```DataQueryable```) => ```postgresQueryableImpl.go``` 
 - postgresModel.go (```DataController```) => ```postgresManipulativeImpl.go```

and etc...

### Package's definition

