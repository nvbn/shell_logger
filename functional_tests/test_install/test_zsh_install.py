image = ('shell_logger/zsh_install',
         '''FROM ubuntu:18.04
            RUN apt-get update
            RUN apt-get install -yy zsh socat curl sudo''',
         'sh')


setup = '''
cp /src/functional_tests/test_install/.zshrc /root/
'''


def test(spawnu, TIMEOUT):
    # Prepare container
    proc = spawnu(*image)
    proc.sendline(setup)
    proc.sendline('zsh')
    proc.sendline('sh -c "$(cat /src/install.sh)"')
    assert proc.expect([TIMEOUT, 'source ~/.zshrc'])
    proc.sendline('source ~/.zshrc')
    # Ensure that shell_logger is installed
    proc.sendline('test $SHELL_LOGGER_SOCKET && echo installed')
    assert proc.expect([TIMEOUT, 'echo installed'])
    assert proc.expect([TIMEOUT, 'installed'])
