from config import get_current_project_path
import os


def remove():
    command = 'cd ' + get_current_project_path() + ' && ./gradlew web'
    os.system(command)