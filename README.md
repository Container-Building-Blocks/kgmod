![kgmod](kgmod.png)

*kgmod* is a simple command line utility aimed to help improve the developer experience for creating [GoLang](https://golang.org/) services to deploy in [Kubernetes](https://kubernetes.io/) via the package manager [Helm](https://helm.sh).

## Table of Contents

- [Overview](#overview)
- [Installing](#installing)
- [Getting Started](#getting-started)
  * [Init](#init)
  * [Custom config](#custom-config)
- [Contributing](CONTRIBUTING.md)
- [License](#license)

# Overview
*kgmod* is a cli utility to do a lot of boiler plate process. kgmod doesn't create or eliminate code, rather it eliminates the boiler plate processes amongst the developers.

This command-line tools looks for a `.kgmod.yaml` file in the local directory where the command is executed. Sharing a simple `.kgmod.yaml` file is nearly equivalent enough to share the workspace settings. As the tool refers to the specified configurations for the modules, flags that indicate whether to create or not the Dockerfile and Helm charts.

# Installing

Installing *kgmod* is very simple. Go to the [releases](https://github.com/Container-Building-Blocks/kgmod/releases) page. Choose the version and platform of the binary you want to download. 

After downloading the binary, rename it to simply `kgmod`, chmod to executable and move it to the search path.

# Getting Started

Typically *kgmod* follows the below cinfiguration file format.

```yaml
modules:
  - k8s.io/klog/v2
replace:
  - github.com/gkarthiks/k8s-discovery=/Users/gkarthiks/k8s-discovery
docker: true
chart-helm: true
```

Using the above sample configuration will create a module in the current working directory. Fetches the modules listed. Replaces the module to the mentioned local path. Creates a skeleton docker file for containerizing and also creates the helm v3 initial chart for the developer to fill in. Thus minimizing all the boiler plate process works.

## Init

Initializing the kgmod is pretty simple. There are multiple ways to do it. If there is no `.kgmod.yaml` file in the current working directory but running the init command will not throw any errors. This will just create a `go module` and doesn't do the `go get` as there will not be any pre-configured modules. But creates the skeleton docker file and helm chart.

To execute the init command, 

```sh
kgmod init --pkg-name my-fancy-module --docker true --chart-helm true
```

Alternatively the *kgmod* comes with a pull command that is handy to pull from a specific location or pulls the default configuration [here](https://raw.githubusercontent.com/gkarthiks/kgmod/master/kgmod.yaml).


To pull the default configuration,

```sh
kgmod pull
```

## Custom config

Alternative to the above default configuration, *kgmod* can be used with a custom configuration as well. This is very helpful with sharing the config across the team. For example, a team's basic development structure/configuration can be stored/exposed from an endpoint.

To pull the config from an external endpoint

```sh
kgmod pull --location=https://example.com/kgmod.yaml
```

The configuration of this can be also stored once and used multiple times. No need to have the `.kgmod.yaml` in all the project directory structure. To do so, have the config file stored in one place somewhere local and execute by passing the file path with the help of config option as 

```sh
kgmod init --pkg-name my-fancy-module --config /Users/.kgmod.yaml
```


# License

Cobra is released under the Apache 2.0 license. See [LICENSE](https://github.com/Container-Building-Blocks/kgmod/blob/master/LICENSE)