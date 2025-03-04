#!/usr/bin/python3
import sys
import os
import argparse
import re
import datetime
import subprocess


default_path = "./src/" # Where to spawn the files
default_template = "day.template" # name of the template-file
file_extension = "go" # The file extension to use (instead of ".template") - dont include a period
test_day_command = "go run ./src/{}/day{}.go" # The command used to test a specific day
test_all_command = "UNIMPLEMENTED_FOR_GO" # The command used to test all days at once

current_year = datetime.datetime.now().year
year_format = "aoc{year}"
day_format = "{day:02d}"
filename_format = "day{}.{}".format(day_format, file_extension)
filename_pattern = r"day(\d+)." + re.escape(file_extension) # The pattern of the created files.

possible_days = list(map(lambda n: str(n),range(1,25+1)))


def new(args):
    print("Creating new file...")
    package = year_format.format(year=args.year)
    path = os.path.join(args.dir, package)
    if not args.day:
        # Find the relevant day
        day = find_day(path) + 1
    elif args.day in possible_days:
        # Use specified day
        day = int(args.day)
    else:
        # undefined
        print("Undefined behaviour! args.day=" + args.day)
        day = int(args.day)

    day_padded = day_format.format(day=day)

    # Get Timestamp
    starttime = datetime.datetime.now().strftime("%H:%M:%S") #.time()

    # Get template
    with open(args.template) as f:
        template = f.read()
        output = format_template(args, template, SafeDict(package=package, year=args.year, day=day_padded, starttime=starttime))

    # Create today's file
    newfile = os.path.join(path, filename_format.format(day=day))
    print("Created {} at time {}!".format(newfile, starttime))
    with open(newfile, "w", newline='') as f:
        f.write(output)

def test(args):
    print("Testing aoc: year={} - day={} ...".format(args.year,args.day or "latest"))
    if "all" in args.day:
        cmd = test_all_command
        print(cmd)
        # os.system(cmd)
        subprocess.run(cmd, shell=True)
    else:
        if args.day == "":
            path = os.path.join(args.dir, year_format.format(year=args.year))
            day = find_day(path)
        elif int(args.day):
            day = int(args.day)
        day_padded = day_format.format(day=day)
        testonly_format = test_day_command.format(year_format, day_format)
        cmd = testonly_format.format(year=args.year, day=int(day_padded))
        print(cmd+"\n")
        os.environ["PROJECT_ROOT"] = os.path.realpath(os.path.dirname(__file__))
        subprocess.run(cmd, shell=True) # works


class SafeDict(dict):
     def __missing__(self, key):
         return '{' + key + '}'


def find_day(path):
    last_day = 0
    for name in os.listdir(path):
        if re.search(filename_pattern, name) is not None:
            num = int(re.search(filename_pattern, name).group(1))
            last_day = num if num > last_day else last_day
    return last_day


def format_template(args, template, dict, placeholder_pattern=r"%{{(.+?)}}"):
    escaped = re.sub(r"([{}])", r"\1\1", template) # Temporarily escape curly-brackets
    placeholder = re.sub(placeholder_pattern, r"{\1}", escaped) # Convert placeholder to regular python-format-placeholders
    formatted = placeholder.format_map(dict)
    unescaped = re.sub(r"([{}]){2}", r"\1", formatted) # unescape curly-brackets
    return unescaped


def parse_args():
    parser = argparse.ArgumentParser()
    parser.add_argument('-d', '--dir', nargs=1, default=default_path,
                        help="The directory of the src (Default: {}".format(default_path))
    parser.add_argument('-y', '--year', nargs="?", default=current_year,
                        help="The year (Default: current ({}))".format(current_year))

    subparsers = parser.add_subparsers(dest='command', help='Which command to run. (Default: test)\n'
                                                            'Use `aoc.py new` to create today\'s file.')
    subparsers.required = True

    # parser for the "test" command
    parser_test = subparsers.add_parser("test", formatter_class=argparse.RawTextHelpFormatter,
            help="Test solution(s), runs '{}'".format(test_day_command.format(year_format, day_format)))
    parser_test.add_argument('day', nargs="?", choices=["","all"]+possible_days,
                             default="", help="Which day(s) to test (Default: ""):\n"
                                              "    \"\" - Only latest day.\n"
                                              "    \\d - Test only the specified day.\n"
                                              "    \"all\" - Test everything. Runs '{}' instead.""".format(test_all_command))
    parser_test.set_defaults(func=test)

    # parser for the "new" command
    parser_new = subparsers.add_parser('new', formatter_class=argparse.RawTextHelpFormatter,
                                       help='Create a new file.')
    parser_new.add_argument('-t', '--template', nargs=1, default=default_template,
                            help="The path to the template (Default: '{}')".format(default_template))
    parser_new.add_argument('day', nargs="?", choices=[""]+possible_days,
                            default="", help='Which day to create (Default: ""):\n'
                                             '    "" - Current highest + 1\n'
                                             '    \\d - Create the specified day.\n')
    parser_new.set_defaults(func=new)

    return parser.parse_args()


def main():
    args = parse_args()
    args.func(args)

if __name__ == "__main__":
    main()
