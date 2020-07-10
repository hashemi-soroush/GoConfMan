# GoConfMan
GoConfman's idea is that:

There are multiple sources for configuration (e.g. env vars, files, defaults, etc.). No matter how many sources
GoConfMan supports, there are always sources that some developers need but it doesn't support. So, adding a new source
must be easy, feels natural, and fits beautifully along with the existing sources.

Besides, control over the order of which we read the sources and fill the configurations must be given to the developer
completely, but sensible defaults are nice.

Given these ideas, GoConfMan has a simple structure.

## Codebase structure
For each source, an interface, named `ConfigWithX`, is defined that has a method called `BindX()`. Beside it, there is
a function called `LoadFromX(interface{})`. Something like this:

```go
package goconfman

type ConfigWithDefaults interface {
    BindDefaults()
}

func LoadFromDefaults(config interface{}) {
    ...
}
```

And that's it. That's the whole codebase of GoConfMan. A bunch of files like this. You wanna add a new source? Just add
a new file like this to your project or be kind and put it in GoConfMan codebase and send a merge request.

You might think that `LoadFromX(interface{})` functions are some huge functions that do a lot of stuff and are hard and
time-consuming to write. Actually, the only thing they do is calling their corresponding `BindX` method for all the
fields in the given config struct recursively. Of course, there are some differences here and there, but basically
that's the whole point of these functions: to load the config recursively.

And finally, there's a `LoadFromAll(interface{})` function that calls all the `LoadFromX` functions in a sensible order,
so you can simply use it, or copy and paste it in your own project and change the order.

Now that we know how the whole thing works, let's actually use it.

## How to use it
Just like this

```go
package my_package
import "github.com/Sayed-Soroush-Hashemi/GoConfMan/pkg/goconfman"

type MyConfig struct {
    Port int
    Workers int
    DB DBConfig
}

type DBConfig struct {
    IP string
    Port int
    Username string
    Password string
}

func (m *MyConfig) BindDefaults() {
    m.Port = 8000
    m.Workers = 4
}

func (d *DBConfig) BindDefaults() {
    d.IP = "localhost"
    d.Port = 5678
    d.Username = "admin"
}

func myFunc() {
    myConfig := MyConfig{}
    goconfman.LoadFromAll(&myConfig)
}
``` 
You have a number of config structs which can be composed into each other. You load the whole config with the last two
lines in the above code snippet.

As you can see, you write normal code to determine the default values. "Not making things complicated" is one of the
goals of GoConfMan, so the same normal code routine happens in almost all the `BindX` methods. For example, here's
`BindAliases` methods for the above config structs.

```go
func (d *DBConfig) BindAliases() {
    d.Password = d.Username
}
```

You want more examples? Check out the `tests` folder. There are a few `X_config.go` files that show you how to use
GoConfMan. You can also check out the `load_from_X_test.go` files to see how to use the `LoadFromX` functions in your
code.

## Hazards
1. If you have a config struct inside another config struct (like the `DBConfig` in the above code snippet), it must be
exposed, otherwise the `LoadFromX` functions can't see it and can't call the `BindX()` methods on it.
 
2. You can do whatever you want in `BindX` methods. You can set defaults in the `BindAliases()` methods and vice versa.
But, for the sake of clarity and separation, don't do that.

## Migrating to GoConfMan
GoConfman enables you to gradually migrate from your old configuration manager to GoConfMan. You can call `LoadFromX`
functions on any struct and it recursively looks for GoConfMan-compatible structs and load them. So, you can start small
and just migrate one of your configuration structs and yet you can call `LoadFromX` on your top-most config struct and
it works. For an example, check out the `non_goconfman_config.go` and `load_from_defaults.go` files in the `tests`
folder. 