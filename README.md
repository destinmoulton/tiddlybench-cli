# CLI Options

## Commands

| Flag | Command   | Description                                                  |
| ---- | --------- | ------------------------------------------------------------ |
| -p   | --paste   | Paste the clipboard. See [Clipboard Usage](#clipboard-usage) |
| -i   | --inbox   | Add to the Inbox                                             |
| -j   | --journal | Add to the Journal                                           |

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

You can add info to two types of tiddlers:

-   Inbox (`-i`)
-   Journal (`-j`)

You can repurpose either of them to fit your needs. General use case is to set the Inbox title to something static, like `Inbox`, and then set the Journal to a date generated title. See [Date and Time in Titles](#date-and-time-in-titles) for information on what date time string substitutions are available.

## Date and Time in Titles

You can configure the tiddlers (either Inbox or Journal) to use date and time string substitutions.

| Token | Substituted Value                  |
| ----- | ---------------------------------- |
| `DDD` | Day of week in full (eg, "Monday") |
| `ddd` | Short day of week (eg, "Mon")      |

`DD` Day of month `0DD` Adds a leading zero `DDth` Adds a suffix `WW` ISO\-8601 week number of year `0WW` Adds a leading zero `MMM` Month in full (eg, "July") `mmm` Short month (eg, "Jul") `MM` Month number `0MM` Adds leading zero `YYYY` Full year `YY` Two digit year `wYYYY` Full year with respect to week number `wYY` Two digit year with respect to week number `hh` Hours `0hh` Adds a leading zero `hh12` Hours in 12 hour clock `0hh12` Hours in 12 hour clock with leading zero `mm` Minutes `0mm` Minutes with leading zero `ss` Seconds `0ss` Seconds with leading zero `XXX` Milliseconds `0XXX` Milliseconds with leading zero `am` or `pm` Lower case AM/PM indicator `AM` or `PM` Upper case AM/PM indicator `TZD` Timezone offset `\x` Used to escape a character that would otherwise have special meaning `[UTC]` Time\-shift the represented date to UTC. Must be at very start of format string
[TiddlyWiki DateFormat Reference](https://tiddlywiki.com/static/DateFormat.html)

This can be configured directly when you initially run `tikli` or using the `-c` configuration flag, or by manually setting the title in the [config file](#config-file).
