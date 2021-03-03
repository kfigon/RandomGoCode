# Bad name
* just rename

# Duplications
* `extraction function`
* if not possible - `move instruction` -> `extract function`
* classes - `move up in hierarchy`
* ifs - `extract function` or `replace if with polymorphism`

# long function
* `extract function`
* too many temp vars and params - `change function to command pattern`
* hard ifs - `decompose conditional` or `replace if with polymorhism`
* hard loop - `split loop` -> `extract function`

# long param list
* `replace parameter with query`
* `wrap params with object`
* `remove flag parameter`

# globals
* `encapsulate var` (get/set for global)

# mutable data
* `encapsulate var`
* `split variable`
* move data away from algorithm `move instruction` & `extract function`
* `CQRS`
* `remove setter`
* `change reference to value`
* `replace variable with query`
* big area of mutation - `introduce class`, `introduce tranformation`

# divirgent change
* `split phases`
* `move function`
* `extract function` or `extract class`

# shotgun surgery
* `move function` and `move field`
* `change function to class`
* `introduce tranformation`
* `split phases`
* `inline function or class` -> `extract function or class`

# feature envy
* `move funciton`
* `extract function`

# data clump
* `wrap with a class`

# simple types
* `wrap with a class`
* `introduce polymorphism`

# lazy class
* `inline`
* `remove hierarchy`

# speculative generality
* `inline`
* `remove hierarchy`

# temp field
* `extract class`
* `introduce special case`

# message chain
* `hide delegate`
* `extract function`

# middle man
* `remove middle man`
* `inline`
* `exchange upclass/downclass with delegate`

# inapropiate intimacy
* `move field` or `move function`
* `hide delegate`

# big class
* `extract`
* `add subclass`

# rejected bequest
* `move method/field down`
* `introduce delegate`