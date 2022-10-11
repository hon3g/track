### Usage
```
$ track help

Usage:
  track [command]

Available Commands:
  completion  Generate the autocompletion script for the specified shell
  help        Help about any command
  lasership   Track a LaserShip package given a tracking number

Flags:
  -h, --help   help for track

Use "track [command] --help" for more information about a command.
```
```
$ track lasership LS12345678

Estimated Delivery Date:
   2022-10-10

Event(s):
   2022-10-10T19:34:47
   NY BROOKLYN 11204
   Your package has been delivered
   Left at: Front Door

```
```
$ track lasership LS12345678 --all

Estimated Delivery Date:
   2022-10-10

Event(s):
   ...

   2022-10-10T16:32:45
   NY MASPETH 11378
   Out for Delivery

   2022-10-10T19:34:47
   NY BROOKLYN 11204
   Your package has been delivered
   Left at: Front Door

```