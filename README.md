# loom
Loom - A repository for Weave Files

Used by [GoWeave](https://github.com/deferpanic/goweave)

A weave file contains aspects to weave into projects.

The weave file is subject to dramatic change right now. There is no
formal specification as of yet.

[Hello World](https://github.com/deferpanic/loom#examples)
[Examples](https://github.com/deferpanic/loom#examples)


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

### Examples:

[Loom](https://github.com/deferpanic/loom) has a listing of samples and
user-contributed weaves you might wish to look at or utilize.

Essentially we support aspect files w/in a project. If code exists in
your project we should be able to provide aspect coverage for it.

There are a few design decisions that need to be made to support across
projects && into stdlib. Stdlib probably won't come until we move to IR.

To try things out first try running `go build`. Then try running `goweave`.

#### Before Function Execution
```go
aspect {
  pointcut: execute(beforeBob())
  advice: {
    before: {
      fmt.Println("before bob")
    }
  }
}
```

#### After Function Execution
```
aspect {
  pointcut: execute(afterAnnie())
  advice: {
    after: {
      fmt.Println("after annie")
    }
  }
}
```

#### Before Main Execution
```go
aspect {
  pointcut: execute(main())
  advice: {
    before: {
      fmt.Println("hello world")
    }
  }
}
```

#### Global Advice
```go
aspect {
  pointcut: execute(*)
  advice: {
    before: {
      var globalCntr int = 0
    }
  }
}
```

#### Before Function Call
```go
aspect {
  pointcut: call(beforeBob())
  advice: {
    before: {
      fmt.Println("before bob")
    }
  }
}
```

#### After Function Call
```
aspect {
  pointcut: call(afterAnnie())
  advice: {
    after: {
      fmt.Println("after annie")
    }
  }
}
```

### Around Function Call
```
aspect {
  pointcut: call(http.HandleFunc(d, s))
  advice: {
    around: {
      http.HandleFunc(d, dps.HTTPHandlerFunc(s))
    }
  }
}
```


