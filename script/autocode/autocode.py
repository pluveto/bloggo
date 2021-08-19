"""
autocode.py
代码生成器
"""
from lib.autocode_master import AutocodeMaster
import sys
import os
import yaml
import pprint
import json

from datetime import datetime
from termcolor import colored, cprint
from lib import model_gen

# 配置文件，来自 config.yml
config = {}


def help():
    # todo: 实现帮助功能
    return


def about():
    """
    获取关于信息
    """
    name = colored('Autocode', 'magenta')
    versn = colored(' v2.0', 'green')
    return \
        f"+--------------------------------------+\n" \
        f"| {name}{versn}                        |\n" \
        f"| @author Pluveto                      |\n" \
        f"+--------------------------------------+\n" \



def load_config():
    """
    读取配置文件
    """

    config = yaml.load(open("config.yml", encoding="utf-8",
                       mode="r"), Loader=yaml.SafeLoader)
    return config


def gen_models():
    """
    读取 Go 模型结构体，并解析为标准格式，输出到文件
    """

    print('Generating models...')
    models = model_gen.gen_models(config['dir']['source']['model'])
    models_json = json.dumps({"models": models}, indent=2, ensure_ascii=False)
    models_path = config["dir"]["task"] + "/models.json"
    open(models_path, encoding="utf-8", mode="w").write(models_json)
    print("models output: " + models_path)


def get_time():
    return datetime.today().strftime('%Y-%m-%d %H:%M:%S')


def run_task(models_file: str, task_file: str):
    gen = AutocodeMaster(models_file, task_file)
    task_out_file = os.path.splitext(task_file)[0] + ".filled.json"
    open(task_out_file, encoding="utf-8",
         mode="w").write(json.dumps(gen.api_group.to_dict(), indent=2, ensure_ascii=False))
    print("task filled: %s" % task_out_file)

    gen.load_vars()

    task_basename = os.path.splitext(task_file)[0]

    gen.render_tpl('routes.jinja', task_basename + ".routes.txt")
    gen.render_tpl('api.jinja', task_basename + ".api.txt")
    gen.render_tpl('service.jinja', task_basename + ".service.txt")
    gen.render_tpl('repo.jinja', task_basename + ".repo.txt")    


if __name__ == "__main__":
    print(sys.argv)
    print(about())
    # todo: 性能计时
    print(get_time())
    config = load_config()
    noArg = len(sys.argv) == 1
    if(noArg):
        gen_models()
        exit(0)

    import argparse
    parser = argparse.ArgumentParser(description='Autocode parser')
    parser.add_argument('-t', '--task', dest='task', type=str,
                        help='Task filename', required=True)
    parser.add_argument('-m', '--models', dest='models',
                        type=str, help='Models filename', required=True)
    args = parser.parse_args()

    run_task(args.models, args.task)
