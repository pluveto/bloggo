import re
import requests
import json
from typing import List

class ac_util:

    @staticmethod
    def en_to_ch(query):
        url = 'http://fanyi.youdao.com/translate'
        data = {
            "i": query,
            "from": "en",
            "to": "zh-CHS",
            "smartresult": "dict",
            "client": "fanyideskweb",
            "salt": "16081210430989",
            "doctype": "json",
            "version": "2.1",
            "keyfrom": "fanyi.web",
            "action": "FY_BY_CLICKBUTTION"
        }
        headers = {
            'User-Agent': 'Mozilla/5.0 (Windows NT 10.0; Win64; x64)'
            + ' AppleWebKit/537.36 (KHTML, like Gecko) Ch'
            + 'rome/90.0.4430.72 Safari/537.36 Edg/90.0.818.39'
        }
        response = requests.post(url, data=data, headers=headers)
        html = response.content.decode('utf-8')
        result = json.loads(html)['translateResult'][0][0]['tgt']
        return result

    @staticmethod
    def any_in_str(strs: List[str], str) ->bool:
        for s in strs:
            if s in str:
                return True
        return False

    @staticmethod
    def pascal_to_underscore(inp: str) -> str:
        regex = r"([A-Z][A-Z0-9]*(?=$|[A-Z][a-z0-9])|[A-Za-z][a-z0-9]+)"
        words = [x.lower() for x in re.findall(regex, inp)]
        return '_'.join(words)


    @staticmethod
    def pascal_to_camel(inp: str) -> str:
        if(inp.upper() == inp):
            """全为大写的转换成全为小写（如 ID）"""
            return inp.lower()
        if(len(inp) == 0):
            return inp
        return inp[0].lower() + inp[1:]


    @staticmethod
    def http_method_to_action(inp: str) -> str:
        d = {
            "GET": "Get",
            "POST": "Create",
            "PUT": "Update",
            "DELETE": "Delete",
        }
        if inp in d:
            return d[inp]
        raise ValueError('%s is not a valid http method.' % inp)


    @staticmethod
    def go_kind_to_form_type(inp: str) -> str:
        d = {
            "bool": "integer",
            "int": "integer",
            "int8": "integer",
            "int16": "integer",
            "int32": "integer",
            "int64": "integer",
            "uint": "integer",
            "uint8": "integer",
            "uint16": "integer",
            "uint32": "integer",
            "uint64": "integer",
            "uintptr": "integer",
            "float32": "number",
            "float64": "number",
            "complex64": "string",
            "complex128": "string",
            "array": "array",
            "interface": "object",
            "map": "object",
            "slice": "array",
            "string": "string",
        }
        if inp in d:
            return d[inp]
        raise ValueError('%s is not a valid kind, or not support yet.' % inp)
