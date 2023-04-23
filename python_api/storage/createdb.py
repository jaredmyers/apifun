from sqlmodel import SQLModel
import models
from initdb import engine

SQLModel.metadata.create_all(engine)
