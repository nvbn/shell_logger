import json


image = ('shell_logger/bash_usage',
         '''FROM ubuntu:18.04
            RUN apt-get update
            RUN apt-get install -yy bash socat curl
            RUN curl https://raw.githubusercontent.com/rcaloras/bash-preexec/master/bash-preexec.sh -o /root/.bash-preexec.sh''',
         'sh')


setup = '''
cp /src/functional_tests/shell_logger /usr/bin/
chmod +x /usr/bin/shell_logger
cp -a /src/functional_tests/common /root/
cp /src/functional_tests/test_happy_path/.bashrc /root/
'''


def test(spawnu, TIMEOUT):
    # Prepare container
    proc = spawnu(*image)
    proc.sendline(setup)
    proc.sendline('bash')
    # Ensure that shell_logger is running
    proc.sendline('test $SHELL_LOGGER_SOCKET && echo ready')
    assert proc.expect([TIMEOUT, 'echo ready'])
    assert proc.expect([TIMEOUT, 'ready'])
    # Execute command
    proc.sendline('ls')
    # Ensure that shell_logger recorded command
    proc.sendline('socat - UNIX-CONNECT:$SHELL_LOGGER_SOCKET')
    proc.sendline(json.dumps({
        "type": "list",
        "count": 5,
    }))
    assert proc.expect([TIMEOUT, '"command":"ls"'])
    assert proc.expect([TIMEOUT, '"returnCode":0'])
