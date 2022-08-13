import uvicorn
import time
from fastapi import FastAPI
from fastapi import Response
from pydantic import BaseModel
import uuid


def get_current_time() -> str:
    return time.strftime("%Y-%M-%d %H:%m:%S")


class Human(BaseModel):
    id: str
    name: str
    age: int
    created_at: str
    updated_at: str

class CreateHuman(BaseModel):
    name: str
    age: int

class UpdateHuman(BaseModel):
    name: str|None
    age: int|None


human_dict: dict[str, Human] = dict()


app = FastAPI(
    docs_url="/swagger/docs",
    redoc_url="/swagger/redoc",
)

@app.get("/")
def read_all():
    return list(human_dict.values())

@app.get("/{human_id}")
def read(human_id: str):
    if human_id not in human_dict:
        return {"message": "not found"}
    return human_dict[human_id]

@app.post("/")
def create(human: CreateHuman):
    new_id: str = str(uuid.uuid4())
    current_time: str = get_current_time()
    new_human = Human(
        id=new_id,
        name=human.name,
        age=human.age,
        created_at=current_time,
        updated_at=current_time,
    )
    human_dict[new_id] = new_human
    return new_human

@app.put("/{human_id}")
def update(human_id: str, human: UpdateHuman):
    if human_id not in human_dict:
        return {"message": "not found"}
    old_human = human_dict[human_id]
    if human.name is not None:
        old_human.name = human.name
        old_human.updated_at = get_current_time()
    if human.age is not None:
        old_human.age = human.age
        old_human.updated_at = get_current_time()
    return old_human

@app.delete("/{human_id}")
def delete(human_id: str):
    if human_id not in human_dict:
        return {"message": "not found"}
    del human_dict[human_id]
    return {"message": "OK"}

uvicorn.run(
    app,
    host="0.0.0.0",
    port=5000,
)
