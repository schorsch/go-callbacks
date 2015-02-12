# Go Callbacks

Building extensible software always requires a system to add loosely coupled parts. There are many names for such
be it plugins, components, extensions or modules.

go-callbacks provides a system to register and run methods for a given callback name.

Golang callbacks was created with a Wordpress-like hook-based system in mind, where modules register to callbacks
and optionally provide own callback hooks. Don't worry this library is nowhere to close to wordpress' fucked-up
spaghetti-code.

Golang with it's static nature makes it somewhat difficult to build a completely dynamic plugin-system. There are
approaches to work with RPC-based micro-services. This is about the only way to extend a GO-based system at runtime but
while the idea is not bad, it opposes lots of hurdles. GO-callbacks assume a simpler approach by defining components in
the sources and compile them with the base system.

See the tiny web-framework example, where modules are manually registered in the main framework file. A core module
orchestrates the application and provides basic hooks. The /modules folder is an example of where 3rd party plugins
could reside and it might even be possible to use git submodules for external extensions.

## Usage

see callback_test and /examples

## Notes

Not battle tested yet! This is my first golang lib and i hope to use it in one of my production apps soon.

Feedback is welcome!


MIT License | Copyright 2015 Georg Leciejewski