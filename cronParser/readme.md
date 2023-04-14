# just a cron expression parser

* `*` - any
* `x-y` - range <x;y>
* `,` - value sepratarot
* `/` step, allowed grammar:
    * `*/s`
    * `x-y/s`
    * `1,2,3/s`
    * `5,20-30/5`
