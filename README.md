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

* [api](#api)
* [applications](#applications)
* [applicationpackgekeys](#applicationpackgekeys)
* [area](#area)
* [members](#members)
* [memberapplications](#memberapplications)
* [roles](#roles)
* [services](#services)
* [endpoints](#endpoints)
* [packages](#packages)
* [packagekeys](#packagekeys)
* [plans](#plans)
* [planservices](#planservices)
* [plandesigner](#plandesigner)
* [config](#config)
* [help](#help)


### api

* [import](#api_import)

Usage:

```json
mashcli api import [command options] filename PublicDomain
```

Options:

#### **applications**

* [delete](#applications-delete)
* [export](#applications-export)
* [import](#applications-import)
* [show](#applications-show)
* [showall](#applications-showall)

##### <a id="applications-delete">*delete*</a>
Delete an Application in Mashery.

Usage:

```json
mashcli applications delete [command options] ApplicationID
```

Command Options:

* --area value, -a value    Area Configuration Name. If not provided defaults to mashcli.config

##### <a id="applications-export">*export*</a>
Export an Application specification from Mashery.

Usage:

```json
mashcli applications export [command options] ApplicationID
```

Command Options:

* --area value, -a value      Area Configuration Name. If not provided defaults to mashcli.config
* --filename value, -f value  The export Filename for the Application Definition

*Omission of the --filename option will output the definition to stdout.*

Examples:

```json
$ mashcli applications export d985067d-1568-4152-94f7-82acdbf4537d
{
 "id": "d985067d-1568-4152-94f7-82acdbf4537d",
 "created": "2017-12-12T14:06:45Z",
 "updated": "2017-12-14T10:03:43Z",
 "username": "mashclitest",
 "name": "mashclitest_application",
 "description": "mashcli test",
 "type": "WEB_WIDGET",
 "commercial": true,
 "ads": true,
 "adsSystem": "mashclitest",
 "usageModel": "OPEN_VERY_WIDE",
 "tags": "mashcli",
 "notes": "mashcli",
 "howDidYouHear": "mashcli",
 "preferredProtocol": "REST",
 "preferredOutput": "JSON",
 "externalId": "mashcli",
 "uri": "http://mashcli",
 "status": "draft",
 "isPackaged": true,
 "oauthRedirectUri": "http://mashcli"
}
```

```json
$ mashcli applications export -f out.json d985067d-1568-4152-94f7-82acdbf4537d
```

##### <a id="applications-import">*import*</a>
Import an Application specification in to Mashery. Should be used to only update an existing Application.
If you need to create a new Application from specification please use [memberapplications import] command.

Usage:

```json
mashcli applications import [command options]
```

Command Options:

* --area value, -a value      Area Configuration Name. If not provided defaults to mashcli.config
* --filename value, -f value  The import Filename for the Application Definition


##### <a id="applications_show">*show*</a>
Show a specific Application in Mashery.

Usage:

```json
mashcli applications show [command options]
```

Options:

* --area value, -a value    Area Configuration Name. If not provided defaults to mashcli.config
* --output value, -o value  Output format value 'table' or 'json'. Default is 'table'

Examples:

```json
$ mashcli applications show bd096f15-dad4-4318-b2ba-09ba6b1cf831
+--------------------------------------+-------------+---------------------------+--------------+-------------+--------+-------------+------------+--------------+----------------+----------+--------+---------------------+---------------------+
|                  ID                  | EXTERNAL ID |           NAME            | DESCRIPTION  |  USERNAME   | STATUS |    TYPE     | COMMERCIAL | RUNS ADVERTS |     USAGE      | PROTOCOL | OUTPUT |       CREATED       |       UPDATED       |
+--------------------------------------+-------------+---------------------------+--------------+-------------+--------+-------------+------------+--------------+----------------+----------+--------+---------------------+---------------------+
| bd096f15-dad4-4318-b2ba-09ba6b1cf831 |             | Weather Application       |              | mmussett    | draft  |             | OFF        | OFF          |                |          |        | 2017-09-15T08:44:01 | 2017-09-15T08:44:01 |
+--------------------------------------+-------------+---------------------------+--------------+-------------+--------+-------------+------------+--------------+----------------+----------+--------+---------------------+---------------------+
```


##### <a id="applications_showall">*showall*</a>
Show a list of all known Applications in Mashery.

Usage:

```json
mashcli applications showall [command options]
```

Options:

* --area value, -a value    Area Configuration Name. If not provided defaults to mashcli.config
* --filter value, -f value  Filter expression as colon-separated name/value pair i.e -filter 'name:Basic'
* --output value, -o value  Output format value 'table' or 'json'. Default is 'table'

Examples:

```json
$ mashcli applications showall
+--------------------------------------+-------------+---------------------------+--------------+-------------+--------+-------------+------------+--------------+----------------+----------+--------+---------------------+---------------------+
|                  ID                  | EXTERNAL ID |           NAME            | DESCRIPTION  |  USERNAME   | STATUS |    TYPE     | COMMERCIAL | RUNS ADVERTS |     USAGE      | PROTOCOL | OUTPUT |       CREATED       |       UPDATED       |
+--------------------------------------+-------------+---------------------------+--------------+-------------+--------+-------------+------------+--------------+----------------+----------+--------+---------------------+---------------------+
| bd096f15-dad4-4318-b2ba-09ba6b1cf831 |             | Weather Application       |              | mmussett    | draft  |             | OFF        | OFF          |                |          |        | 2017-09-15T08:44:01 | 2017-09-15T08:44:01 |
| ff6e6a14-1e42-4b25-a474-b013704a86ca |             | ACME Application          |              | mmussett    | draft  |             | OFF        | OFF          |                |          |        | 2017-12-08T10:10:21 | 2017-12-08T10:10:21 |
+--------------------------------------+-------------+---------------------------+--------------+-------------+--------+-------------+------------+--------------+----------------+----------+--------+---------------------+---------------------+
```

```json
$ mashcli applications showall -output json
[
    {
        "id": "bd096f15-dad4-4318-b2ba-09ba6b1cf831",
        "created": "2017-09-15T08:44:01Z",
        "updated": "2017-09-15T08:44:01Z",
        "username": "mmussett",
        "name": "Weather Application",
        "status": "draft",
        "isPackaged": true
    },
    {
        "id": "ff6e6a14-1e42-4b25-a474-b013704a86ca",
        "created": "2017-12-08T10:10:21Z",
        "updated": "2017-12-08T10:10:21Z",
        "username": "mmussett",
        "name": "ACME Application",
        "status": "draft",
        "isPackaged": true
    }
]
```
```json
$ mashcli applications showall -filter 'username:mmussett'
+--------------------------------------+-------------+---------------------------+-------------+----------+--------+------+------------+--------------+-------+----------+--------+---------------------+---------------------+
|                  ID                  | EXTERNAL ID |           NAME            | DESCRIPTION | USERNAME | STATUS | TYPE | COMMERCIAL | RUNS ADVERTS | USAGE | PROTOCOL | OUTPUT |       CREATED       |       UPDATED       |
+--------------------------------------+-------------+---------------------------+-------------+----------+--------+------+------------+--------------+-------+----------+--------+---------------------+---------------------+
| bd096f15-dad4-4318-b2ba-09ba6b1cf831 |             | Weather Application       |             | mmussett | draft  |      | OFF        | OFF          |       |          |        | 2017-09-15T08:44:01 | 2017-09-15T08:44:01 |
| ff6e6a14-1e42-4b25-a474-b013704a86ca |             | ACME Application          |             | mmussett | draft  |      | OFF        | OFF          |       |          |        | 2017-12-08T10:10:21 | 2017-12-08T10:10:21 |
+--------------------------------------+-------------+---------------------------+-------------+----------+--------+------+------------+--------------+-------+----------+--------+---------------------+---------------------+
```

### applicationpackagekeys

* showall

### area

* backup
* restore


### members

* add
* delete
* export
* import
* setstatus
* show
* showall

### memberapplications

* export
* import
* showall

### *services*

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


### endpoints

* add
* clone
* delete
* export
* import
* show
* showall

### packages

* add
* clone
* delete
* export
* import
* show
* showall

### packagekeys

* delete
* export
* import
* setrates
* setstatus
* show
* showall

### plans

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

### planservices

* show
* showall

### plandesigner

* add
* delete

### config

* show
* add
