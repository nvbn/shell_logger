import pytest


def pytest_addoption(parser):
    group = parser.getgroup("resources_usage")
    group.addoption('--enable-resources-usage', action="store", default=0,
                    help="Enable resources usage tests")


@pytest.fixture(autouse=True)
def resources_usage(request):
    if request.node.get_marker('resources_usage') \
            and not request.config.getoption('enable_resources_usage'):
        pytest.skip('resources usage tests are disabled')


@pytest.fixture
def resources_usage_steps(request):
    n = request.config.getoption('enable_resources_usage')
    if n:
        return int(n)
    else:
        return 0
