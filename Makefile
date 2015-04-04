.PHONY: clean-pyc clean-build docs clean

_VIRT_ENV_NAME=ls_lock


virtualenv_create_env:
	bash -c ". $$(which virtualenvwrapper.sh); mkvirtualenv $(_VIRT_ENV_NAME);"

virtualenv_install_packages:
	bash -c ". $$(which virtualenvwrapper.sh) ; \
	workon $(_VIRT_ENV_NAME) ; \
	pip install -U -r requirements.txt ;" \
	echo "Use workon $(_VIRT_ENV_NAME)" ;

virtualenv_install_test_packages:
	bash -c ". $$(which virtualenvwrapper.sh) ; \
	workon $(_VIRT_ENV_NAME) ; \
	pip install -U -r test-requirements.txt ;" \
	echo "Use workon $(_VIRT_ENV_NAME)" ;


# How to use it:
# make infinitest NOSETEST=test-integration/minion/test_cpu_freq_governor.py
infinitest:
	while true; do nosetests --nocapture $(NOSETEST); sleep 1; done

help:
	@echo "clean - remove all build, test, coverage and Python artifacts"
	@echo "clean-build - remove build artifacts"
	@echo "clean-pyc - remove Python file artifacts"
	@echo "clean-test - remove test and coverage artifacts"
	@echo "lint - check style with flake8"
	@echo "test - run tests quickly with the default Python"
	@echo "test-all - run tests on every Python version with tox"
	@echo "coverage - check code coverage quickly with the default Python"
	@echo "docs - generate Sphinx HTML documentation, including API docs"
	@echo "release - package and upload a release"
	@echo "dist - package"

clean: clean-build clean-pyc clean-test

clean-build:
	rm -fr build/
	rm -fr dist/
	rm -fr *.egg-info

clean-pyc:
	find . -name '*.pyc' -exec rm -f {} +
	find . -name '*.pyo' -exec rm -f {} +
	find . -name '*~' -exec rm -f {} +
	find . -name '__pycache__' -exec rm -fr {} +

clean-test:
	rm -fr .tox/
	rm -f .coverage
	rm -fr htmlcov/

lint:
	flake8 ch-heat-controller tests

# test:
#	python setup.py nosetests

# test-all:
#	tox

coverage:
	coverage run --source ch-heat-controller setup.py test
	coverage report -m
	coverage html
	open htmlcov/index.html

docs:
	rm -f docs/ch-heat-controller.rst
	rm -f docs/modules.rst
	sphinx-apidoc -o docs/ ch-heat-controller
	$(MAKE) -C docs clean
	$(MAKE) -C docs html
	open docs/_build/html/index.html

release: clean
	python setup.py sdist upload
	python setup.py bdist_wheel upload

dist: clean
	python setup.py sdist
	python setup.py bdist_wheel
	ls -l dist
.PHONY: test_arduinod integration_test_hw_arduinod
