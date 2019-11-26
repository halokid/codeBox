import setuptools

with open("README.md", "r") as fh:
  long_description = fh.read()

setuptools.setup(
  name = "xiaoshuaib",
  version = "0.0.1",
  author = "xiaoshuaib",
  long_description = long_description,
  long_description_content_type = "text/markdown",
  url = "http://xxxx",
  packages = setuptools.find_packages(),
  classifiers = [
    "python 3"
  ]
)

