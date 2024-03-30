# hypervector

![hypie](./assets/hypie.png)

**Status: 🔨 Pre-first release**

Hypervector is a test data fixture engine intended for data-intensive production systems such as those reliant
on machine learning and data science components.

It makes configuring & generating high dimensional feature vectors easy, and provides data populations for testing
deterministically over HTTP. Output from your function under test can also be hashed and stored for future assertion to
detect specific regressions related to input data sub-populations.

## Use

<img src="./assets/flowchart.png" width=75%>

## Build

You can run the test suite and build a local binary (useful for development) with:

```bash
make all
```

Building and running a Docker image is done with:

```bash
make docker
```
