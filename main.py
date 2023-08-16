from config import init_config
from action.login import check_login
from args import get_action
from action.reset import reset
from action.login import login
from action.logout import logout
from action.clone import clone
from action.list import list
from action.develop import develop
from action.show import show
from action.idea import idea
from action.vscode import vscode
from action.clean import clean
from action.start import start
from action.package import package
from action.docker import docker
from action.deploy import deploy
from action.website import website
from action.branch import branch
from action.git import git
from action.pr import pr
# from action.init import init
# from action.frontend import frontend
# from action.backend import backend


if __name__ == '__main__':
    init_config()
    switch_action = {
        "reset": reset,
        "login": login,
        "logout": logout,
        "clone": clone,
        "list": list,
        "develop": develop,
        "show": show,
        "idea": idea,
        "vscode": vscode,
        "clean": clean,
        "start": start,
        "package": package,
        "docker": docker,
        "deploy": deploy,
        "website": website,
        "branch": branch,
        "git": git,
        "pr": pr,
        # "init": init,
        # "frontend": frontend,
        # "backend": backend,
    }
    if get_action() != 'login':
        check_login()
    do_action = switch_action.get(get_action())
    do_action()
