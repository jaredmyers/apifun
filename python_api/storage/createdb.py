from sqlmodel import SQLModel
from schemas import User, Team
from initdb import engine

SQLModel.metadata.create_all(engine)
