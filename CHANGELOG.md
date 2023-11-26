# Changelog

All notable changes to this project will be documented in this file.

The format is based on [Keep a Changelog](https://keepachangelog.com/en/1.0.0/),
and this project adheres to [Semantic Versioning](https://semver.org/spec/v2.0.0.html).

## [0.0.4] - 2023-11-26

* Adjust naming of interfaces and types to better represent their purpose.
  * `options.Options` is now collection of `option.Option` setters.
  * `options.Resolver` is what previously `options.Options` was. 

## [0.0.3] - 2023-11-17

* Simplify `Resolve` by removing necessity of instancing `options.Options` at first.
* Allow for interface option to be `nil` if default value is `nil`.

## [0.0.2] - 2023-10-12

* Fix providing option with interface implementation.

## [0.0.1] - 2023-10-09

* Release to GitHub

[0.0.4]: https://github.com/SevenOfSpades/go-just-options/releases/tag/v0.0.4
[0.0.3]: https://github.com/SevenOfSpades/go-just-options/releases/tag/v0.0.3
[0.0.2]: https://github.com/SevenOfSpades/go-just-options/releases/tag/v0.0.2
[0.0.1]: https://github.com/SevenOfSpades/go-just-options/releases/tag/v0.0.1