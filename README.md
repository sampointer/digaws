# digaws

Look-up region and other information for any AWS-owned IP address:

```bash
digaws 52.94.76.5

digaws 2a05:d07a:a0ff:ffff:ffff:ffff:ffff:aaaa
```

## Installation

### Homebrew

### Nix

### Packages

## Usage

By default `digaws` fetches `ip-ranges.json` every time it is executed. If you
intend on performing batch processing or perhaps want to use a version of the
file at a specific point in time you can pass the `-i` flag with a path and
use a locally cached copy. See: `digaws -h` and [the AWS documentation][2] for
more information.

[1]: https://ip-ranges.amazonaws.com/ip-ranges.json
[2]: https://docs.aws.amazon.com/general/latest/gr/aws-ip-ranges.html
