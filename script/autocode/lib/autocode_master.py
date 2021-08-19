"""

api_gen.py

本文件负责将 ApiGroup 对象分析后输出以下文件：
1. 路由
2. API 层
3. API 层的请求体、响应体模型
3. Service 层
4. Repo 层

本文件的工作原理：

task_file ---------------------> apiGroup ------------> template_vars
           api_group_from_dict              prepare()        | 
                                                             | renders.render()
                                                             v
                                                        result_files


template_vars 结构：
{
    groups [
        name: "Post"
        route: "/post"
        apis: [
            method: "GET"
            route: "/post/get"
            method_route: "/get"
            handler_name: "PostGet"
        ]
    ]
}

"""

import json
from .template_var_booster import TemplateVarBooster
from .api_group_builder import APIGroupBuilder
from .model_provider import ModelProvider
import re
from typing import Iterator

from jinja2.loaders import FileSystemLoader
from .api_group import APIGroup
from jinja2 import Environment, PackageLoader, select_autoescape


class AutocodeMaster:
    def __init__(self, models_file: str, task_file: str):
        task_json = open(task_file, encoding="utf-8", mode="r").read()
        api_group = APIGroup.from_json(task_json)
        self.model_provider = ModelProvider.from_file(models_file)
        builder = APIGroupBuilder(api_group, self.model_provider)
        self.api_group = builder.auto_complete()
        return

    def load_vars(self):
        booster = TemplateVarBooster(self.api_group, self.model_provider)
        booster.boost()
        self.template_vars = {"group": booster.get_vars()}
        self.env = Environment(
            loader=FileSystemLoader("template")
        )        

    def render_tpl(self, tpl_name, out_name = None):
        template = self.env.get_template(tpl_name)
        result = template.render(self.template_vars)
        if not out_name:
            print(result)
            return
        open(out_name, mode="w", encoding="utf-8").write(result)
        print("rendered %s" % out_name)
# for test
# print(pascalToUnderscore("HTTPRequest"))
