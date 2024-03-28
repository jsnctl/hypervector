# hypervector

Hypervector is a test data fixture engine intended for data-intensive production systems such as those reliant
on machine learning and data science components.

It makes configuring & generating high dimensional feature vectors easy, and provides data populations for testing
deterministically over HTTP. Output from your function under test can also be hashed and stored for future assertion to
detect specific regressions related to input data sub-populations.
