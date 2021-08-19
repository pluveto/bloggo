from .model_provider import ModelProvider
from typing import List
from .model import Field, Mode, Model
from .api_builder import APIBuilder
from .api_group import API, APIGroup, Meta
from .ac_util import ac_util


class APIGroupBuilder:

    def __init__(self, grp: APIGroup, models_provider: ModelProvider) -> None:
        self.api_group = grp
        self.models_provider = models_provider

    def finish(self) -> APIGroup:
        return self.api_group

    def auto_complete(self) -> APIGroup:
        """
        本函数用于填充用户省略的信息
        """
        if(self.models_provider.models == None):
            raise RuntimeError("Missing models!")
        return self.build_api_group().finish()

    def build_api_group(self):
        if(self.api_group.route == "."):
            self.api_group.route = "/" + \
                ac_util.pascal_to_underscore(self.api_group.name)
        self.build_meta()

        for api in self.api_group.apis:
            self.build_api(api)
        return self

    def build_meta(self):
        meta: Meta = self.api_group.meta
        if(not(meta.json) or meta.json == "."):
            meta.json = ac_util.pascal_to_camel(self.api_group.name)
        if(not(meta.label) or meta.label == "."):
            meta.label = ac_util.en_to_ch(self.api_group.name)            
        return self

    def build_api(self, api: API):
        """
        自动补全 API 的相关信息
        """
        api.method = api.method.upper()
        builder = APIBuilder(self.api_group.name, api, self.models_provider)
        api = (builder.build_action()
               .build_route()
               .build_base_model()
               .build_action_on()
               .build_form()
               .build_paging()
               .build_permissions()
               .build_auto_assign()
               .build_req_fields()
               .build_ret_fields()
               .remove_useless_form_fields()
               .finish())
        return self
