# What is TiddlyBench CLI?

TiddlyBench CLI is a command line interface for TiddlyWiki.

TiddlyBench CLI is a companion to the (TiddlyBench browser extension)[https://github.com/destinmoulton/tiddlybench-extension] for Firefox and Chrome. Check out the [TiddlyBench website](https://tiddlybench.com) for more documentation on TiddlyBench CLI and the TiddlyBench browser extension.

# Features

-   Paste directly into a TiddlyWiki tiddler
-   Quickly add entries to your Journal or Inbox
-   Edit a tiddler using your favorite text editor
-   Pipe a file directly into a tiddler

# Warning

TiddlyBench is in early development so use it at your own risk.
Make frequent backups!

# Usage

## Tiddler Notes

If a Tiddler does not exist with a given title, a new tiddler will be created with that title.

## Basic

Set the title (`-t`) and add (`-a`) text to that tiddler.

```bash
$ tb -t "Title of Tiddler" -a "Text that will be added to the tiddler"
```

## Journal and Inbox

Add (`-a`) a new journal (`-j`) entry:

```bash
$ tb -j -a "This text will be added to the journal"
```

## Pasting

Paste (`-p`) the contents of your clipboard to your tiddler Inbox (`-i`):

```bash
$ tb -p -i
```

Paste a quote (`--quote`) from the clipboard into today's journal (`-j`):

```bash
$ tb -p -j --quote
```

## Piping

Pipe data into a tiddler (`-t`):

```bash
$ cat underground.txt | tb -t "Notes from the Underground"
```

## Blocks

Add (`-a`) a code block (`--code`) to the Inbox (`-i`):

```bash
$ tb -i --code -a "function falsify(){\nreturn false;\n}"
```

# CLI Options

## Command Line Flags and Options

| Flag | Command   | Description                                         |
| ---- | --------- | --------------------------------------------------- |
| -a   | --add     | Add or append text to the tiddler                   |
| -c   | --config  | Re-run the configuration questions                  |
| -e   | --edit    | Edit the new content with your favorite text editor |
| -i   | --inbox   | Add to the Inbox                                    |
| -j   | --journal | Add to the Journal                                  |
| -p   | --paste   | Paste form the [Clipboard](#clipboard-usage)        |
| -t   | --title   | Set the title of the tiddler you want to add/append |

Note: -t will override -i (inbox) or -j (journal)

## Block Options

You can set the block that will surround the content that you add.

| Command  | Description                 |
| -------- | --------------------------- |
| --code   | Monospaced code block       |
| --bullet | Bulleted list item          |
| --number | Numbered list item          |
| --quote  | Quote block                 |
| --h1     | H1 aka first level header   |
| --h2     | H2 aka second level header  |
| --h3     | H3 aka third level header   |
| --h4     | H4 aka fourth level header  |
| --h5     | H5 aka fifth level header   |
| --custom | Custom user definable block |

If you are editing the text using an external editor, using the `-e` flag, the block code will be visible and editable within the text editor.

See [Block Customization](#block-customization) for how to customize the blocks.

# Config File

The configuration file is named `config.json`.

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

## Block Customization

You can customize the blocks by configuring the `begin` and `end` string for a block.

Add `\n` to add a newline.

The `default` setting is the block that is used when no [block option](#block-options) are provided.

## Custom Text Editor

The config file includes an option to setup a custom text editor when you use the `-e` flag. The default is set to the `$EDITOR` environment variable. You can customize the editor and the arguments that the editor will receive.

# Authentication

tikli assumes that your TiddlyWiki installation is configured with basic authentication.

## Password Storage

During the configuration stage tikli will ask if you want to save your password.

Saving your password may not be safe if you are using an account where the administrator is not you.

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
