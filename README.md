# hypervector

![hypie](./hypie.png)

**Status: 🔨 Pre-first release**

Hypervector is a test data fixture engine intended for data-intensive production systems such as those reliant
on machine learning and data science components.

It makes configuring & generating high dimensional feature vectors easy, and provides data populations for testing
deterministically over HTTP. Output from your function under test can also be hashed and stored for future assertion to
detect specific regressions related to input data sub-populations.

## Use

![image](https://github.com/jsnctl/hypervector/assets/25587856/ce3a40cd-3cb2-411a-a4ae-c16bf72101a6)

## Build

You can run the test suite and build a local binary (useful for development) with:

```bash
make all
```

Building and running a Docker image is done with:

```bash
make docker
```
