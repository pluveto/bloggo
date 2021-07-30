# https://www.python-httpx.org/
import os
import unittest

import httpx


class TestStringMethods(unittest.TestCase):

    def setUp(self) -> None:
        self.x = httpx.Client(base_url=os.getenv("API_BASE"))
        return super().setUp()

    def test_ping(self):
        r = self.x.get("/ping")

        assert(r.status_code == 200)
        print(r.json()['data']['message'])

    def test_adminCreate(self):
        r = self.x.post("admin/create", json={
            "email": "test@qq.com",
            "password": "wtfpasswd"
        })
        print(r.json())


if __name__ == '__main__':
    unittest.main()