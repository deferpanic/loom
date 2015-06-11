# loom
Loom - A repository for Weave Files

Used by [GoWeave](https://github.com/deferpanic/goweave)

A weave file contains aspects to weave into projects.

The weave file is subject to dramatic change right now. There is no
formal specification as of yet.

##Hello World Weave File:##
```
aspect {
  imports (
    "fmt"
  )
  pointcut: execute(*)
  advice: {
    before: {
      fmt.Println("hi")
    }
  }
}
```
