from dataclasses import field
from .action import Action
from typing import List
from .model import Field, Mode, Model
from .ac_util import ac_util


class APIFormBuilder:

    def __init__(self, model: Model, action: str) -> None:
        print("builder with action: " + action)
        self.model = model
        self.action = action
        self.form = {}
        pass

    def finish(self):
        return self.form

    def auto_complete(self) -> dict:
        """
        自动补全请求表单（用于客户端请求和查询）
        """
        form = {}
        fields = self.select_fields_for_request()
        for fi in fields:
            print(fi)
            form[fi.json] = {
                "type": ac_util.go_kind_to_form_type(fi.kind)
            }
        self.form = {
            "type": "object",
            "properties": form
        }
        return self.finish()



    def merge_fields(self, l1: List[Field], l2: List[Field]) -> List[Field]:
        def fields_has(l: List[Field], name:str)->bool:
            for fi in l:
                if fi.name == name:
                    return True
            return False

        ret  :List[Field] = []
        if(bool(l1)):
            for fi in l1:
                if(not(fields_has(ret, fi.name))):
                    ret.append(fi)
        if(bool(l2)):
            for fi in l2:                
                if not bool(fi):
                    continue
                if(not(fields_has(ret, fi.name))):
                    ret.append(fi)
        return ret

    def select_fields_for_request(self) -> List[Field]:
        """
        筛选各种行为（action）提交表单所需要的字段
        """
        model: Model = self.model
        action: str = self.action
        act = Action(ac_util.pascal_to_underscore(action))
        if(act == Action.create):
            fields = model.get_fields_by_mode(Mode.W)  # 寻找所有可写字段
            exclude = []  # 要排除的字段
            include = []  # 要强制添加的字段
            fields = list(filter(lambda fi: (not fi.name in exclude), fields))
            # fields = [*fields, *include] 此处不能简单合并

        elif(act == Action.get):
            # 仅限 primaryKey, unique 字段
            fields = model.get_fields_unique()

        elif(act == Action.list):
            fields = []

        elif(act == Action.update):
            fields = model.get_fields_by_mode(Mode.R)
            exclude = []  # 要排除的字段
            include = [model.get_pkey_field()]  # 需要指定主键才能更新
            fields = list(filter(lambda fi: (not fi.name in exclude), fields))
            fields = self.merge_fields(fields, include)

        elif(act == Action.search):
            fields = model.get_fields_searchable()

        elif(act == Action.delete):
            fields = model.get_fields_unique()

        elif(act == Action.enable):
            fields = model.get_fields_unique()

        elif(act == Action.disable):
            fields = model.get_fields_unique()
        else:
            raise(RuntimeError("unknown action: %s" % act))

        return fields

    def select_fields_for_response(self) -> List[Field]:
        """
        筛选各种行为（action）提交表单所需要的字段
        """
        model: Model = self.model
        action: str = self.action
        act = Action(ac_util.pascal_to_underscore(action))
        fields = []
        """
        创建后返回 ID
        """
        if(act == Action.create):
            pkey = model.get_pkey_field()  # 寻找所有可写字段
            if(bool(pkey)):
                fields = model.get_fields_unique()
        elif(act == Action.get):
            """
            GET 后返回所有公开字段
            """
            # 仅限 primaryKey, unique 字段
            fields = model.get_fields_by_mode(Mode.R)

        elif(act == Action.list):
            fields = model.get_fields_by_mode(Mode.R)

        elif(act == Action.update):
            fields = []

        elif(act == Action.search):
            fields = model.get_fields_by_mode(Mode.R)

        elif(act == Action.delete):
            fields = []

        elif(act == Action.enable):
            fields = model.get_fields_by_mode(Mode.R)

        elif(act == Action.disable):
            fields = model.get_fields_by_mode(Mode.R)
        else:
            raise(RuntimeError("unknown action: %s" % act))

        return fields
