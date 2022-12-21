# linstor-wait-until

![Latest release](https://img.shields.io/github/v/release/linbit/linstor-wait-until)

Waits until a specific component of LINSTOR is online and usable.

## Usage

The program will continuously loop until the condition is true. Currently implemented:

* `linstor-wait-until api-online` waits until the LINSTOR API is online, i.e. you can start sending client commands
* `linstor-wait-until satellite-online <satellite-name>` waits until a satellite's status is ONLINE.

## Configuration

`linstor-wait-until` uses the environment variables specified in the [`golinstor` library](https://pkg.go.dev/github.com/LINBIT/golinstor/client#NewClient)
for configuration.

| Variable               | Description                                                            |
|------------------------|------------------------------------------------------------------------|
| `LS_CONTROLLERS`       | A comma-separated list of LINSTOR controller URLs to connect to.       |
| `LS_USERNAME`          | Username to use for HTTP basic auth.                                   |
| `LS_PASSWORD`          | Password to use for HTTP basic auth.                                   |
| `LS_ROOT_CA`           | CA certificate to use for authenticating the server.                   |
| `LS_USER_KEY`          | TLS key to use for authenticating the client to the server.            |
| `LS_USER_CERTIFICATE`  | TLS certificate to use for authenticating the client to the server.    |
| `LS_BEARER_TOKEN_FILE` | Name of the file containing the token for Bearer Token Authentication. |
