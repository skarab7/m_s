import unittest
import os
from multiprocessing import Pool
import fcntl
import subprocess


class TestLsLock(unittest.TestCase):
    """
    Integration-test of ls_lock
    """
    TEST_TMP_DIR = "/tmp"
    LOCK_FILE_PREFIX = "lslock-test"
    LS_LOCK_EXEC="bash ls_lock.sh"

    def test_when_twoLockFileExist_then_printLocks(self):
        """
        given: two processes are holding lock
        """
        filename = self._get_lock_file_name("_t1")
        self._create_file_if_not_exist(filename)
        my_pid = os.getpid()

        output = None
        with open(filename, "r+") as f:
            fcntl.flock(f.fileno(), fcntl.LOCK_EX)
            output = self._get_output_ls_lock(TestLsLock.TEST_TMP_DIR)
        self._delete_file_if_exists(filename)
        lines = output.split()[2:]
        for l in lines:
            print l
            pid = l.split(" ")[0]
            fl = l.split(" ")[1] 
            print pid
            print fl          

    def _create_file_if_not_exist(self, filename):
        if not os.path.exists(filename):
            with open(filename, 'w+') as f:
                f.write("0")

    def _get_lock_file_name(self, postfix):
        return "{0}/{1}_{2}".format(TestLsLock.TEST_TMP_DIR, TestLsLock.LOCK_FILE_PREFIX, postfix)

    def _get_output_ls_lock(self, target_dir):
        cmd = TestLsLock.LS_LOCK_EXEC.split()
        cmd.append(target_dir)
        process = subprocess.Popen(cmd, stdout=subprocess.PIPE)
        output = process.communicate()[0] 
        return output

    def _delete_file_if_exists(self, filename):
        if(os.path.isfile(filename)):
            os.remove(filename)

