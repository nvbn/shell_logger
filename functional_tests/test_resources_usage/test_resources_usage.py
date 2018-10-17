import sys
import pytest


image = ('shell_logger/bash_memory_usage',
         '''FROM ubuntu:18.04
            RUN apt-get update
            RUN apt-get install -yy bash socat curl
            RUN curl https://raw.githubusercontent.com/rcaloras/bash-preexec/master/bash-preexec.sh -o /root/.bash-preexec.sh''',
         'sh')


setup = '''
cp /src/functional_tests/shell_logger /usr/bin/
chmod +x /usr/bin/shell_logger
cp -a /src/functional_tests/common /root/
cp /src/functional_tests/test_memory_usage/.bashrc /root/
'''


def _prepare_line(*items):
    return ''.join(part.ljust(16) for part in items)


def _print_stats(stage, proc):
    stats = proc.docker_stats()
    print(_prepare_line(stage,
                        stats['CPUPerc'],
                        stats['MemUsage'].split('/')[0],
                        stats['MemPerc']))


@pytest.mark.resources_usage
def test(spawnu, TIMEOUT, resources_usage_steps):
    # Prepare container
    proc = spawnu(*image)
    proc.sendline(setup)
    proc.sendline('bash')

    proc.sendline('test $SHELL_LOGGER_SOCKET && echo ready')
    assert proc.expect([TIMEOUT, 'echo ready'])
    assert proc.expect([TIMEOUT, 'ready'])

    sys.stdout.write(
        '\n\n\033[1m'
        + _prepare_line('Stage', 'CPUPerc', 'MemUsage', 'MemPerc')
        + '\033[0m\n')
    _print_stats('ready', proc)
    try:
        for n in range(resources_usage_steps):
            proc.sendline("cat /dev/urandom "
                          "| tr -dc 'a-zA-Z0-9~!@#$%^&*_-' "
                          "| fold -w 200 "
                          "| head -n 10000 "
                          "&& clear")

            proc.sendline('echo check-{}'.format(n))
            _print_stats('handling-{}'.format(n), proc)
            assert proc.expect([TIMEOUT, 'echo check-{}'.format(n)])
            assert proc.expect([TIMEOUT, 'check-{}'.format(n)])

            _print_stats('after-{}'.format(n), proc)
    finally:
        proc.sendline('clear')
