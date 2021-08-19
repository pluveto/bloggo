"""
APIBuilder
负责的是 APIGroup 中 API 的生成。
本类由 APIGroupBuilder 类调用，生成其 apis 字段
"""
from lib.model import Field
from typing import List
from .action import Action
from .api_form_builder import APIFormBuilder
from .model_provider import ModelProvider
from .api_group import API
from .ac_util import ac_util


class APIBuilder:
    def __init__(self, model_name: str, api: API, models_provider: ModelProvider) -> None:
        self.model_name = model_name
        self.api = api
        self.models_provider = models_provider

        self.build_base_model()
        self.build_action()
        self.model = self.models_provider.get_model(self.api.base_model)
        action = self.api.action
        self.form_builder = APIFormBuilder(self.model, action)
        pass

    def finish(self):
        return self.api

    def build_action(self):
        if(not(self.api.action) or self.api.action == "."):
            self.api.action = ac_util.http_method_to_action(self.api.method)
            self.action: Action = Action(
                ac_util.pascal_to_underscore(self.api.action))
        return self

    def build_route(self):
        if(not(self.api.route) or self.api.route == "."):
            self.api.route = "/" + \
                ac_util.pascal_to_underscore(self.api.action + self.api.action_on)
        return self

    def build_base_model(self):
        if(not(self.api.base_model) or self.api.base_model == "."):
            self.api.base_model = self.model_name
        return self

    def build_action_on(self):
        """
        API 操作的资源对象（如果和模型类相同则可以省略名称）
        例：GetPost，则 action_on 为 ""
        例：GetRecentPost，则 action_on 为 "Recent"
        """
        if(not(self.api.action_on) or self.api.action_on == "."):
            self.api.action_on = ""
        return self

    def build_form(self):
        """
        请求表单
        """
        if(not bool(self.api.form)):
            self.api.form = self.form_builder.auto_complete()
        return self

    def build_paging(self):
        """
        分页
        """
        if(not bool(self.api.paging)):
            self.api.paging = self.__complete_api_paging()
        return self

    def build_permissions(self):
        """
        权限
        """
        if(not bool(self.api.permissions)):
            # Module.Action
            self.api.permissions = [self.model_name + "." + self.api.action]
        return self

    def build_auto_assign(self):
        """
        自动赋值
        """
        self.__complete_api_auto_assign()
        return self

    def build_req_fields(self):
        """
        格式：
        {
            "Username": {
                "name": "Username",
                "json": "username",
                "column": "username",
                "column_meta": {
                    "column": "username"
                },
                "kind": "string",
                "weight": 0,
                "searchable": true,
                "sortable": true,
                "mode": [
                    "r",
                    "w"
                ],
                "label": "用户名",
                "type": "text",
                "required": True
            }
        }
        """
        if(bool(self.api.req_fields)):
            return self
        """
        需要根据 Action 的类型，来决定请求需要提交的字段有哪些。
        """        
        self.api.req_fields = self.form_builder.select_fields_for_request()        
        return self

    def build_ret_fields(self):
        if(bool(self.api.ret_fields)):
            return self
        self.api.ret_fields = self.form_builder.select_fields_for_response()
        return self

    def __complete_api_auto_assign(self):
        model = self.models_provider.get_model(self.api.base_model)
        auto_assign_is_empty = not bool(self.api.auto_assign)
        if auto_assign_is_empty:
            print("auto_assign_is_empty")
            self.api.auto_assign = {}
            if(model.get_field("PublishedAt")):
                self.api.auto_assign["PublishedAt"] = "${created_at}"
            if(model.get_field("CreatedAt")):
                self.api.auto_assign["CreatedAt"] = "${created_at}"
            if(model.get_field("ModifiedAt")):
                self.api.auto_assign["ModifiedAt"] = "${updated_at}"
            if(model.get_field("UpdatedAt")):
                self.api.auto_assign["UpdatedAt"] = "${updated_at}"
        else:
            for fieldName, autoValue in self.api.auto_assign:
                print("auto assign: " + field.json)
                field = model.get_field(fieldName)
                # 如果是自动创建的字段，则不应该在请求表单出现
                del self.api.form["properties"][field.json]
                if autoValue != ".":
                    continue
                if fieldName in ["PublishedAt", "CreatedAt"]:
                    self.api.auto_assign[fieldName] = "${created_at}"
                elif fieldName in ["ModifiedAt", "UpdatedAt"]:
                    self.api.auto_assign[fieldName] = "${updated_at}"

    def remove_useless_form_fields(self):
        useless_fields = ["publishedAt", "createdAt", "modifiedAt", "updatedAt"]

        props: dict = self.api.form["properties"]
        if(bool(props)):
            for k in list(props.keys()):
                if (str(k)) in useless_fields:
                    del props[k]

        props: List[Field] = self.api.req_fields
        props = [x for x in props if (not x.json in useless_fields)]
        self.api.req_fields = props
        return self

    def __complete_api_paging(self):
        """
        自动判断是否需要分页返回
        """
        act = ac_util.pascal_to_underscore(self.api.action)
        return {
            "create": False,
            "get": False,
            "list": True,
            "update": False,
            "search": True,
            "delete": False,
            "enable": False,
            "disable": False,
        }[act]
