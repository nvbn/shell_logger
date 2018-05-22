image = ('shell_logger/bash_install',
         '''FROM ubuntu:18.04
            RUN apt-get update
            RUN apt-get install -yy bash curl sudo''',
         'sh')


setup = '''
cp /src/functional_tests/test_install/.bashrc /root/
'''


def test(spawnu, TIMEOUT):
    # Prepare container
    proc = spawnu(*image)
    proc.sendline(setup)
    proc.sendline('bash')
    proc.sendline('sh -c "$(cat /src/install.sh)"')
    assert proc.expect([TIMEOUT, 'source ~/.bashrc'])
    proc.sendline('source ~/.bashrc')
    # Ensure that shell_logger is installed
    proc.sendline('test $SHELL_LOGGER_SOCKET && echo installed')
    assert proc.expect([TIMEOUT, 'echo installed'])
    assert proc.expect([TIMEOUT, 'installed'])
