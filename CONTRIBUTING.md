## Contributing in general

Our project welcomes external contributions. A good way to familiarize yourself with the codebase and the contribution process is to look for and address issues in the [issue tracker](https://github.com/sysflow-telemetry/sysflow/issues).

To contribute code or documentation, please submit a [pull request](https://github.com/sysflow-telemetry/sf-processor/pulls); and please quickly [get in touch](#communication) with us before embarking on a more ambitious contribution.

> **Note:** We appreciate your effort, and want to avoid a situation where a contribution
requires extensive rework (by you or by us), sits in backlog for a long time, or
cannot be accepted at all!

### Proposing new features

If you would like to implement a new feature, please [raise an issue](https://github.com/sysflow-telemetry/sysflow/issues)
before sending a pull request so that the proposed feature can be discussed first. This is to avoid
putting an effort on a feature that the project developers would not be able to accept into the code base.

### Fixing bugs

If you would like to fix a bug, please [raise an issue](https://github.com/sysflow-telemetry/sysflow/issues) before sending a
pull request so that the bug fix can be tracked properly.

### Merge approval

The project maintainers use LGTM (Looks Good To Me) in comments on the code
review to indicate acceptance. A change requires LGTMs from two of the
maintainers of each component affected.

For a list of the maintainers, please see the [maintainers page](MAINTAINERS.md).

## Legal

Each source file must include a license header for the Apache
Software License 2.0. Using the SPDX format is the simplest approach.
For example,

```
/*
Copyright <holder> All Rights Reserved.

SPDX-License-Identifier: Apache-2.0
*/
```

We have tried to make it as easy as possible to make contributions. This
applies to how we handle the legal aspects of contribution. We use the
same approach - the [Developer's Certificate of Origin 1.1 (DCO)](https://github.com/hyperledger/fabric/blob/master/docs/source/DCO1.1.txt) - that the Linux® Kernel [community](https://elinux.org/Developer_Certificate_Of_Origin)
uses to manage code contributions.

We simply ask that when submitting a patch for review, the developer
must include a sign-off statement in the commit message.

Here is an example Signed-off-by line, which indicates that the
submitter accepts the DCO:

```
Signed-off-by: John Doe <john.doe@example.com>
```

You can include this automatically when you commit a change to your
local git repository using the following command:

```
git commit -s
```

## Communication

Please feel free to connect with us on our [Slack channel](https://join.slack.com/t/sysflow-telemetry/shared_invite/enQtODA5OTA3NjE0MTAzLTlkMGJlZDQzYTc3MzhjMzUwNDExNmYyNWY0NWIwODNjYmRhYWEwNGU0ZmFkNGQ2NzVmYjYxMWFjYTM1MzA5YWQ) or
via [email](mailto:sysflow@us.ibm.com). Note that the projects in this repository are not formal products. As a result, the communication channels are to the maintainers and not to a support staff.

## Setup

The documentation is a work in progress but should provide a good overview on how to get started with the project. The Dockerfile also provides a treasure trove of information
on how to build the application, dependencies, and how to test the collector. 

## Testing

TBD

## Coding style guidelines
We follow the [Golang coding standards](https://golang.org/doc/effective_go.html) in this project. You can use the go compiler or your IDE of choice to automatically lint your code.
