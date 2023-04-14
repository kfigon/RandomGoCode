# just a cron expression parser

* `*` - any
* `x-y` - range <x;y>
* `,` - value sepratator
* `/` step, used with wildcards and ranges. Allowed grammar:
    * `*/s`
    * `x-y/s`
