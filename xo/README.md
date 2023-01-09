# xo

- [When should I use a pointer to an interface?](https://www.reddit.com/r/golang/comments/kit3da/whats_the_meaning_of_a_pointer_to_an_interface/)

- 
i) args: a object we pass around and collect data.
ii) loaderImp: a single implementation of loader
  - for each sql driver we create different object.

- for loaderImp I would have preferred separate implementation for each driver  

dir/internal:
args: pass arguments each for whole process 
 args stores data like DTO for the process

loader: a generic loader object and single implementation

dir/loaders:
specific loaders like mysql
