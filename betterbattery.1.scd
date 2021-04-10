betterbattery(1) [""["betterbattery-VERSION"]]

# NAME

betterbattery - A battery printing utility.

# SYNOPSIS

*betterbattery [--help] [--symbols=symbols] [--config=target]*

# DESCRIPTION

*betterbattery* is a tool for monitoring battery status and running commands if
the percentage changed by a specified amount since the last time the tool was
run. For example you could run this command every 3 minutes and if the bettery
percentage dropped below 10% warn the user with a notification and lower the
screen brightness.

# OPTIONS

	*-c, --config* _path_
		Specifies an alternate path to read the config file from.
	*-h, --help*
		Print a help message.
	*-s, --symbols* _string_
		Two character to represent charging status. Increasing then decreasing.
		For example: +-

# AUTHORS

Maintained by Dakota Walsh <kota at nilsu.org>.
Up-to-date sources can be found at https://git.sr.ht/~kota/betterbattery/
