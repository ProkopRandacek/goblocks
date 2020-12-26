# goblocks
Status bar manager for dwm

## What can it do?
- custom separators
- independent update interval
- sleeps most of the time
- json configuration

## What can't it do?
- colors
- start a thermo-nuclear war

## configuration
All config files live in `~/.config/goblocks`. Most of the config is done in `blocks.json`.

`
[
  { 
    "Command": "time",
    "Interval": 1,
    "Separator": "][" 
   },
   {
    "Command": "network",
    "Interval": 20,
    "Separator": "]["
   }
]
`
The command option rund a command in the config folder.
