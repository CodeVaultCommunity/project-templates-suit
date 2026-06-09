import typer
from typing import Annotated
import json
from pathlib import Path
from dataclasses import dataclass
from enum import Enum
import requests

@dataclass
class _DepedencieModel:
    path: str
    need_name: str
    
    @property
    def name(self) -> Path:
        return Path(self.need_name)
    
    @property
    def full_name(self) -> Path:
        return Path(self.path) / self.need_name
    
    @property
    def url(self) -> str:
        return str(self.full_name).replace('\\', '/')
    
    @staticmethod
    def from_dict(d: dict) -> _DepedencieModel:
        path = d.get('path')
        need_name = d.get('need_name')
        if path is None or need_name is None:
            raise ValueError(f"{d} is not well formated; path ({path}) and/or name ({need_name}) not exists.")
        
        return _DepedencieModel(path=path, need_name=need_name)

class _Templates(str, Enum):
    go_api = 'golang_api'
    github_web = 'github_web'
    node_ts_lib = 'node_ts-lib'
    
    @property
    def path(self) -> Path:
        splited = self.value.split('_')
        path = Path()
        for item in splited:
            path /= item
        return path
    
    @property
    def url(self) -> str:
        return str(self.path).replace('\\', '/')

def _load_dependencies(path: Path) -> list[_DepedencieModel]:
    with open(path, 'r') as f:
        dependencies: list = json.load(f)
    
    result: list[_DepedencieModel] = []
    for item in dependencies:
        if type(item) is not dict:
            raise ValueError(f"{path} is not well formated; {item} is not of type dict")
        _item = _DepedencieModel.from_dict(item)
        result.append(_item)
    
    return result


_GITHUB_URL = f"https://api.github.com/repos/CodeVaultCommunity/project-templates-suit/contents"
_GITHUB_RAW_URL = "https://raw.githubusercontent.com/CodeVaultCommunity/project-templates-suit"

def _download_file(destination: Path, url: str) -> None:
    content = requests.get(url)
    content.raise_for_status()
                
    destination.write_bytes(content.content)
    
def _download_content(destination: Path, content_name: str, branch: str) -> None:
    url = f"{_GITHUB_RAW_URL}/{branch}/${content_name}"
    _download_file(destination, url)

def _download_folder(folder: str, branch: str, trg: Path) -> None:    
    def _walk(path: str):
        url = f"{_GITHUB_URL}/{path}?ref={branch}"
        response = requests.get(url)
        response.raise_for_status()
        for item in response.json():
            print(f'downloading {item.get('path')}')
            if item.get('type') == 'file':
                destination: Path = trg / item.get('path', '')    
                destination.parent.mkdir(parents=True, exist_ok=True)
                download_url = item.get('download_url', '')
                _download_file(destination, download_url)
                
            elif item.get('type') == "dir":
                _walk(item["path"])
    _walk(folder)
    
app = typer.Typer()

@app.command()
def project(
    name: Annotated[str, typer.Option("--name", '-n')],
    template: Annotated[_Templates, typer.Option("--template", '-t')],
    branch: Annotated[str, typer.Option("--branch", '-b')] = 'main',
    ):
    current_path = Path(__file__).parent.resolve()
    project_path = current_path / name
    project_path.mkdir(parents=True, exist_ok=True)
    
    _download_folder(template.url, branch, project_path)
    needs_json = project_path / '.needs.json'
    if needs_json.exists():
        for dependencie in _load_dependencies(needs_json):
            _download_content(project_path, dependencie.url, branch)
    
    
def main():
    app()

if __name__ == "__main__":
    main()
