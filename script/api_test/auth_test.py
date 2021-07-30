# https://www.python-httpx.org/
import os
import unittest

import httpx


class AuthTest(unittest.TestCase):

    def setUp(self) -> None:
        self.x = httpx.Client(base_url=os.getenv("API_BASE"))
        return super().setUp()

    def test_authLogin(self):
        r = self.x.post("auth/login", json={
            "email": "test@qq.com",
            "password": "wtfpasswd",
            "captcha": "000000"
        })
        print(r.json())


if __name__ == '__main__':
    unittest.main()