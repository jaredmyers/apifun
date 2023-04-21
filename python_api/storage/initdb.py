from sqlmodel import create_engine, Session

DATABASE_URI = "postgresql://postgres:password@localhost/testdb2"
engine = create_engine(DATABASE_URI, connect_args={}, future=True, echo=True)


def get_session():
    with Session(engine) as session:
        yield session


"""
from sqlalchemy import create_engine
from sqlalchemy.ext.declarative import declarative_base
from sqlalchemy.orm import sessionmaker

DATABASE_URI = "postgresql://postgres:password@localhost/testdb2"

engine = create_engine(DATABASE_URI, connect_args={}, future=True)

SessionLocal = sessionmaker(autocommit=False, autoflush=False, bind=engine, future=True)

Base = declarative_base()


def get_session():
    db = SessionLocal()
    try:
        yield db
    finally:
        db.close()
"""
