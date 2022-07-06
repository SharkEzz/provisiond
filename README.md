<div align="center">

# provisiond

<img src="assets/provisiond.png" height="86">

<p>An easy to use and straightforward deployment automation software, written in Go</p>

</div>

## How to use ?

### Executing a deployment

#### CLI

This is the most simple way to use provisiond, simply launch the executable with the flag `-file` followed by the relative or absolute path to the deployment file.

```
./provisiond -file deployment.yaml
```

#### REST API

The REST API allow you to trigger deployments remotely.

You have to provide the flags `-api` and `-apiPassword`. The `apiPassword` must be followed by a string that will allow accessing the API.

```
./provisiond -api -apiPassword password
```

Aditionally, you can provide the flag `-apiPort` followed by the port you want to listen (must be an integer between 0 and 65536).

By default, the API will be available on `0.0.0.0:7655`.

There is currently 2 routes available:
  - `/v1/healthcheck`: will return a 200 response code and a JSON body with `{ 'result': 'ok' }`
  - `/v1/deploy`: this route is POST only, you must provide a header named `password` with the choosen password as the value.
  The body of the request must be a deployment file content in plain text.

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

You can mark a job as allowed to fail, it means that even if the job return an error, it will be ignored and the rest of the deployment will be executed. By default, all jobs are not allowed to fail unless stated explicitly.

To allow a job to fail, add this key to the job:
```yaml
    allow_failure: true
```

##### Shell
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

##### File

The file plugin allow you to interact with files easily.

- Create
  ```yaml
    file:
      action: create
      path: path_to_the_file.txt
      content: |
        content of the file (can be single of multiline)
  ```

- Exist
  
  The exist action assert that a file exist, if it does not, throw an error
  ```yaml
    file:
      action: exist
      path: path_to_the_file.txt
  ```

- Delete

  As is name say, this action allow to delete a specific file (multiple file support is coming!)
  ```yaml
    file:
      delete: delete
      path: path_to_the_file.txt
  ```

### Configuration

> The configuration is currently only supported when executing provisiond from the CLI, not from the REST API

provisiond allow you to define global configuration for your deployments.

As of now, a very little options are configurable globally:
- Job timeout
- Deployment timeout
- Allow job failure

This options are written in a configuration file named `config.yaml`, this file must be in the same folder as the provisiond executable.

Example:

Examples can be found in the `/examples/config` folder.

```yaml
job_timeout: 60
deployment_timeout: 120
allow_failure: false
```

Here, we set a job timeout of 60 seconds, a deployment timeout of 120 seconds and disallow job failure.

Please note that all the timeout numbers are in seconds, and if the `config.yaml` file is present, you must define all the options yourself.
By default, the job timeout is set to one hour, the deployment timeout to 1 day and the jobs are not allowed to fail.

### Examples

You can check the example files in the current repo (in examples folders).

## Contribute

### Writing plugins

provisiond use the [plugin system from Go](https://pkg.go.dev/plugin) to load external `.so` files and register them for use in deployments files.

> A plugin **MUST** implement the `Plugin` interface defined in this repo.

> A plugin **MUST** export a function with this signature : `GetPlugin() (p any)`, which will return an instance of newly created the plugin

The `Plugin` interface only define one method: `Execute`, which take 2 parameters:
- data : type `any`
    - This is the data received from the deployment file
- ctx : type `JobContext`
    - It contains the methods to interact with the targeted system

The builded plugins must be in a `plugins` folder, in the same directory as the provisiond executable.

Check the example in the `pkg/plugin/example` folder.

## Roadmap

- Add tests
- Launch jobs in goroutines (allow for supervision, timeout, cancels, etc...)
- Clean TODOS
- Dashboard :eyes:
- Online documentation
