<div align="center">

# provisiond

<img src="assets/provisiond.png" height="86">

<p>An easy to use and straightforward deployment automation software, written in Go</p>

</div>

## How to use ?

### Writing a deployment file

A deployment file consist in 4 parts:
- The name of the deployment
- The configuration for the hosts
- The environment variables
- The jobs

Let's see each of them:

#### Name
```yaml
name: The name of your deployment
```

#### Hosts

You can define a list of hosts to use in your jobs.
If you want to run a job locally, you do not need to add a `localhost` host, provisiond is smart enought to see when you will need to execute a job on the local machine.

```yaml
config:
  ssh:
    host1:
      ...
    host2:
      ...
```

##### Username / Password

```yaml
  host1:
    host: 127.0.0.1
    port: 22
    type: password
    username: your_username
    password: your_password
```

##### Private key

```yaml
  host: 127.0.0.1
  port: 22
  type: key
  username: your_username
  keyPath: ./path/to/the/private_key
  keyPass: ""
```

#### Variables

The variables are loaded as environment variables, you can use them in your jobs.

```yaml
variables:
  VARIABLE_1: hello
  VARIABLE_2: 1.525
  VARIABLE_3: |
    a multiline
    variable
```

#### Jobs

The jobs are where you define the commands to run.

```yaml
jobs:
  - name: Job 1
    hosts: [host1, host2]
    ...
```

The name and the hosts are the 2 basic required components.
provisiond use a system of plugins, each of them is identified by their name.

Currently, there is only a `shell` plugin (expect more to come).
It only take a string which is the command to run.

```yaml
    shell: echo Hello from deployment > hello.txt
```

```yaml
    shell: |
        echo multiline command
        cat /dev/zero
```

If the command output something, you will see it in the terminal where provisiond is ran.

### Examples

You can check the example files in the current repo (in examples folders).

## Contribute

### Writing plugins

provisiond use the [plugin system from Go](https://pkg.go.dev/plugin) to load external `.so` files and register them for use in deployments files.

> A plugin **MUST** implement the `Plugin` interface defined in this repo.

> A plugin **MUST** export a function with this signature : `GetPlugin() (p any)`, which will return an instance of newly created the plugin

The `Plugin` plugin interface only define one method: `Execute`, which take 2 parameters:
- data : type `any`
    - This is the data received from the deployment file
- ctx : type `JobContext`
    - It contains the methods to interact with the targeted system

The builded plugins must be in a `plugins` folder, in the same directory as the provisiond executable.

Check the example in the `pkg/plugin/example` folder.

## Roadmap