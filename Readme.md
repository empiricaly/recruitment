# Empirica Recruitment (beta)

Empirica Recruitment is an application meant to simplify the operations needed
to recruit experiment participants using online participant platforms (currently
on MTurk supported).

It is currently in early beta and will require some effort from the user to
setup and use. It has been tested on limited experiments with up to 100
participants and multiple steps.

# Quick Usage

Download the latest pre-built binary release from the
[releases page](https://github.com/empiricaly/recruitment/releases)
(currently only macos supported, adding other platforms soon).

Then copy the config file `recruitment.example.yaml` to `recruitment.yaml` in
the folder where you will be running recruitment from. Modify the configuration
as needed.

Make the binary you downloaded executable with:

```sh
chmod u+x ./recruitment
```

Finally run with:

```sh
./recruitment
```

It will create the SQLLite database in the same folder (`recruiment.db`) and
start the HTTP server on port 8880. Head over to http://localhost:8880.
