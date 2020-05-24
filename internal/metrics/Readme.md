# Metrics package

The `metrics` package contains a pattern to add Prometheus type metrics to a
MantisNet program. It has useful tools to integrate with cobra and viper command
line framework to configure flags, and a wrapper around the Prometheus client
package to normalize it's utilization within MantisNet.

## Current assigned ports

- bflow: 9901
- bftpassembly: 9902
- bs3archive: 9903
- btransform: 9904
- ppe: 9905
- rfpng_exporter: 9906
