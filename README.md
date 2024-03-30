# hypervector

![hypie](./assets/hypie.png)

**Status: 🔨 Pre-first release**

Hypervector is a test data fixture engine intended for data-intensive production systems such as those reliant
on machine learning and data science components. It makes configuring & generating high dimensional feature 
vector fixtures for testing easy and repeatable.

## Use

Hypervector sits naturally alongside your integration and smoke test infrastructure
during continuous integration stages, and makes large high-dimensional feature vectors
available via endpoint.

<img src="./assets/flowchart.png" width=85%>

The service allows for benchmarks to be saved alongside their parent feature vectors. These can be 
asserted by test runners  via the Hypervector API, providing a sophisticated regression testing framework for
statistical properties of your fixtures.

<img src="./assets/flowchart-benchmark.png" width=86%>

## Build

You can run the test suite and build a local binary (useful for development) with:

```bash
make all
```

Building and running a Docker image is done with:

```bash
make docker
```
