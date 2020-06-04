# Usage

If a Tiddler does not exist with that title, a new Tiddler will be created with that title.

## Basic

Set the title (`-t`) and add (`-a`) text to that tiddler.

```bash
$ tikli -t "Title of Tiddler" -a "Text that will be added to the tiddler"
```

## Journal and Inbox

Add (`-a`) a new journal (`-j`) entry:

```bash
$ tikli -j -a "This text will be added to the journal"
```

## Pasting

Paste (`-p`) the contents of your clipboard to your tiddler Inbox (`-i`):

```bash
$ tikli -p -i
```

Paste a quote (`--quote`) from the clipboard into today's journal (`-j`):

```bash
$ tikli -p -j --quote
```

## Piping

Pipe data into a tiddler (`-t`):

```bash
$ cat underground.txt | tikli -t "Notes from the Underground"
```

## Blocks

Add (`-a`) a code block (`--code`) to the Inbox (`-i`):

```bash
$ tikli -i --code -a "function falsify(){\nreturn false;\n}"
```

# CLI Options

## Commands

| Flag | Command   | Description                                         |
| ---- | --------- | --------------------------------------------------- |
| -a   | --add     | Add or append text to the tiddler                   |
| -c   | --config  | Re-run the configuration questions                  |
| -p   | --paste   | Paste form the [Clipboard](#clipboard-usage)        |
| -i   | --inbox   | Add to the Inbox                                    |
| -j   | --journal | Add to the Journal                                  |
| -t   | --title   | Set the title of the tiddler you want to add/append |

## Blocks

You can set how you want

# Config File

## Location of Config File

### Linux

If `$XDG_CONFIG_HOME` is defined:

```
$XDG_CONFIG_HOME/tikli/config.json
```

Otherwise it will default to:

```
$HOME/.config/tikli/config.json
```

### MacOS

```
$HOME/Library/Application Support/tikli/config.json
```

### Windows

```
%AppData%/tikli/config.json
```

## Block Configuration

You can customize the blocks by configuring the `begin` and `end` string for a block.

Add `\n` to add a newline.

# Authentication

tikli assumes that your TiddlyWiki installation is configured with basic authentication.

## Password Storage

tikli will ask if you want to save your password.

This may not be safe if you are using an account where the administrator is not you.

# Clipboard Usage

tikli can paste text directly to a tiddler.

Note: `xsel` or `xclip` are required for Linux usage

## Usage

Paste and ask for Tiddler title.

```bash
$ tikli -p
```

Paste and add to the Journal.

```bash
$ tikli -pf
```

# Blocks

The blocks can be configured by editing your configuration file

# Tiddler Titles

You can add a custom Tiddler title using the -t flag.

```
$ tikli -t "New Tiddler"
```

`tikli` is configured with two main tiddlers:

-   Inbox (`-i`)
-   Journal (`-j`)

You can repurpose either of them to fit your needs. General use case is to set the Inbox title to something static, like `Inbox`, and then set the Journal to a date generated title. See [Date and Time in Titles](#date-and-time-in-titles) for information on what date time string substitutions are available.

Note: Using either `-i` or `-j` flag overrides the `-t` title flag.

## Date and Time in Titles

You can configure the tiddlers (either Inbox or Journal) to use date and time string substitutions.

### List of Substitutions

The substitutions in the title are performed locally by `tikli`, not by TiddlyWiki. Therefore, not all of them have been implemented because I only have so many hours in a day.

| Token        | Substituted Value                          |
| ------------ | ------------------------------------------ |
| `DDD`        | Day of week in full (eg, "Monday")         |
| `ddd`        | Short day of week (eg, "Mon")              |
| `DD`         | Day of month                               |
| `0DD`        | Adds a leading zero                        |
| `DDth`       | _NOT AVAILABLE_                            |
| `WW`         | _NOT AVAILABLE_                            |
| `0WW`        | _NOT AVAILABLE_                            |
| `MMM`        | Month in full (eg, "July")                 |
| `mmm`        | Short month (eg, "Jul")                    |
| `MM`         | Month number                               |
| `0MM`        | Adds leading zero                          |
| `YYYY`       | Full year                                  |
| `YY`         | Two digit year                             |
| `wYYYY`      | Full year with respect to week number      |
| `wYY`        | Two digit year with respect to week number |
| `hh`         | Hours                                      |
| `0hh`        | _NOT AVAILABLE_                            |
| `hh12`       | Hours in 12 hour clock                     |
| `0hh12`      | Hours in 12 hour clock with leading zero   |
| `mm`         | Minutes                                    |
| `0mm`        | Minutes with leading zero                  |
| `ss`         | Seconds                                    |
| `0ss`        | Seconds with leading zero                  |
| `XXX`        | _NOT AVAILABLE_                            |
| `0XXX`       | _NOT AVAILABLE_                            |
| `am` or `pm` | Lower case AM/PM indicator                 |
| `AM` or `PM` | Upper case AM/PM indicator                 |
| `TZD`        | Timezone offset                            |
| `\x`         | _NOT AVAILABLE_                            |
| `[UTC]`      | _NOT AVAILABLE_                            |

See Also: [TiddlyWiki DateFormat Reference](https://tiddlywiki.com/static/DateFormat.html)

This can be configured directly when you initially run|
`tikli` or using the `-c` configuration flag, or by manually setting the title in the [config file](#config-file).
