mashcli
=======

This repository contains the source code for TIBCO Mashery CLI tool (mashcli).

Written in golang v1.9

## Installation and Build

To download mashcli, simply run:
```
$go get github.com/mmussett/mashcli
```

To build and install mashcli, simply run:
```
go install ./...
```

### Supported platforms
mashcli was built and tested on Go v1.9 on  OS X.


### Getting Started
```
$ mashcli
NAME:
   mashcli - TIBCO Mashery - Command Line Interface

USAGE:
   mashcli [global options] command [command options] [arguments...]

VERSION:
   0.1.0

AUTHOR:
   Mark Mussett (mmussett@me.com)

COMMANDS:
     api                         Manage application package keys-related operations. For additional help, use 'mashcli applicationpackagekeys --help'
     applications, ap            Manage application-related operations for the current user. For additional help, use 'mashcli applications --help'
     applicationpackagekeys, ak  Manage application package keys-related operations. For additional help, use 'mashcli applicationpackagekeys --help'
     area, ar                    Manage application package keys-related operations. For additional help, use 'mashcli area --help'
     members, m                  Manage member-related operations. For additional help, use 'mashcli members --help'
     memberapplications, ma      Manage application-related operations for the current user. For additional help, use 'mashcli memberapplications --help'
     services, s                 Manage service-related operations for the current user. For additional help, use 'mashcli services --help'
     endpoints, e                Manage endpoint-related operations for the current user. For additional help, use 'mashcli endpoints --help'
     packages, pa                Manage package-related operations for the current user. For additional help, use 'mashcli packages --help'
     packagekeys, k              Manage package key-related operations. For additional help, use 'mashcli packagekeys --help'
     plans, pl                   Manage plan-related operations for the current user. For additional help, use 'mashcli plans --help'
     planservices, ps            Manage planservice-related operations for the current user. For additional help, use 'mashcli planservices --help'
     plandesigner, pd            Manage plan designer operations for the current user. For additional help, use 'mashcli plandesigner --help'
     config, c                   Configuration for the current user. For additional help, use 'mashcli config --help'
     help, h                     Shows a list of commands or help for one command

GLOBAL OPTIONS:
   --help, -h     show help
   --version, -v  print the version

COPYRIGHT:
   2017 Mark Mussett.
```

### Setup

mashcli must be provided with your Mashery Area configuration before being able to work. 
Once provide mashcli will be able to authenticate and use the Mashery platform APIs.

mashcli stores Mashery Area configuration files in $HOME/.mashcli. 

Each configuration files stores the parameters settings for a single Mashery Area.

By default mashcli will read Mashery Area configuration from `mashcli.config` unless overridden from the command line using the --area flag

Each configuration file contains a json object that follows the following format:
```json
{
  "userid": "",
  "password": "",
  "apikey": "",
  "apikeysecret": "",
  "name": "",
  "area": "",
  "tm": "",
  "ccurl": ""
}
```

Example:
```json
{
  "userid": "mmussett",
  "password": "pa55w0rd",
  "apikey": "s3st2bzysqcv1qcu2t5vz977",
  "apikeysecret": "8pYXg7HQXa",
  "name": "demo",
  "area": "c7e8e2d5-ff91-42eb-9885-10f2aa2cc3f5",
  "tm": "demo.api.mashery.com",
  "ccurl": "https://demo.admin.mashery.com/control-center"
}
```

#### Adding your Mashery Area configuration

To create your area configuration, simply run:
```
$mashcli config add
```
Follow on-screen prompts and provide the following following:
```
Configuration Name?
Enter a value (Default is mashcli): 

Area ID?
Enter a value:

Traffic Manager?
Enter a value:

Control Centre URL?
Enter a value (Default is https://<<area>>.admin.mashery.com/control-center): 

API Key?
Enter a value: 

API Key Secret?
Enter a value: 

User ID?
Enter a value: 
 
User Password?
Enter a value: 
```


### Commands

#### api

Usage:

```json
mashcli api import [command options] filename PublicDomain
```

Options:

#### applications

* delete
* export
* import
* show
* showall

#### applicationpackagekeys

* showall

#### area

* backup
* restore


#### members

* add
* delete
* export
* import
* setstatus
* show
* showall

#### memberapplications

* export
* showall

#### *services*

* add
* clone
* delete
* export
* import
* show
* show all

##### _show all services_
Show a list of all know services
Usage:
```json
mashcli services showall [command options] [arguments...]
```
Options:
* --area value, -a value    Area Configuration Name
* --output value, -o value  Output format table or json

Example
```json
$ mashcli services showall
+--------------------------+--------------------------------------------------+--------------------------------+----------+------------------+---------------------+---------------------+
|        SERVICE ID        |                       NAME                       |          DESCRIPTION           | AGG  QPS |     VERSION      |       CREATED       |       UPDATED       |
+--------------------------+--------------------------------------------------+--------------------------------+----------+------------------+---------------------+---------------------+
| 9dyupurdkktfwstmrw3z45vg | Swagger Petstore                                 |                                |        0 | 1.0.0            | 2018-01-05T14:19:14 | 2018-01-05T14:19:14 |
+--------------------------+--------------------------------------------------+--------------------------------+----------+------------------+---------------------+---------------------+
```

```json
$ mashcli services showall -output json
[
    {
        "id": "9dyupurdkktfwstmrw3z45vg",
        "created": "2018-01-05T14:19:14.000+0000",
        "crossdomainPolicy": "\u003c?xml version=\"1.0\"?\u003e\n\u003c!DOCTYPE cross-domain-policy SYSTEM \"http://www.macromedia.com/xml/dtds/cross-domain-policy.dtd\"\u003e\n\u003ccross-domain-policy\u003e\n   \u003callow-access-from domain=\"*\"/\u003e\n\u003c/cross-domain-policy\u003e",
        "editorHandle": "mmussett",
        "name": "Swagger Petstore",
        "revisionNumber": 2,
        "updated": "2018-01-05T14:19:14.000+0000",
        "version": "1.0.0"
    }
]
```
##### _show service_
Show service
Usage:
```json
mashcli services show [command options] ID
```
Options:
* --area value, -a value    Area Configuration Name
* --output value, -o value  Output format table or json

Example
```json
$ mashcli services show 9dyupurdkktfwstmrw3z45vg
+--------------------------+------------------+-------------+----------+---------+---------------------+---------------------+
|        SERVICE ID        |       NAME       | DESCRIPTION | AGG  QPS | VERSION |       CREATED       |       UPDATED       |
+--------------------------+------------------+-------------+----------+---------+---------------------+---------------------+
| 9dyupurdkktfwstmrw3z45vg | Swagger Petstore |             |        0 | 1.0.0   | 2018-01-05T14:19:14 | 2018-01-05T14:19:14 |
+--------------------------+------------------+-------------+----------+---------+---------------------+---------------------+
```

```json
$ mashcli services show 9dyupurdkktfwstmrw3z45vg -output json
{
    "id": "9dyupurdkktfwstmrw3z45vg",
    "created": "2018-01-05T14:19:14.000+0000",
    "crossdomainPolicy": "\u003c?xml version=\"1.0\"?\u003e\n\u003c!DOCTYPE cross-domain-policy SYSTEM \"http://www.macromedia.com/xml/dtds/cross-domain-policy.dtd\"\u003e\n\u003ccross-domain-policy\u003e\n   \u003callow-access-from domain=\"*\"/\u003e\n\u003c/cross-domain-policy\u003e",
    "editorHandle": "mmussett",
    "name": "Swagger Petstore",
    "revisionNumber": 2,
    "updated": "2018-01-05T14:19:14.000+0000",
    "version": "1.0.0"
}
```


#### endpoints

* add
* clone
* delete
* export
* import
* show
* showall

#### packages

* add
* clone
* delete
* export
* import
* show
* showall

#### packagekeys

* delete
* export
* import
* setrates
* setstatus
* show
* showall

#### plans

* add
* clone
* delete
* export
* import
* setkeyprops
* setratelimits
* setstatus
* show
* showall

#### planservices

* show
* showall

#### plandesigner

* add
* delete

#### config

* show
* add
